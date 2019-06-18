package sql

import (
	"database/sql"

	"github.com/akornatskyy/sample-blog-api-go/membership/domain/user"
	"github.com/akornatskyy/sample-blog-api-go/shared/sqlx"
)

type userRepository struct {
	selectAuthInfo   *sql.Stmt
	selectUserByID   *sql.Stmt
	existsByUsername *sql.Stmt
	insertAccount    *sql.Stmt
}

func NewUserRepository(db *sql.DB) user.Repository {
	return &userRepository{
		selectAuthInfo: sqlx.MustPrepare(db, `
			SELECT id, password_hash, is_locked FROM users WHERE username = ?`),
		selectUserByID: sqlx.MustPrepare(db, `
			SELECT username, first_name, last_name FROM users where id = ?`),
		existsByUsername: sqlx.MustPrepare(db, `
			SELECT '' FROM users WHERE username = ?`),
		insertAccount: sqlx.MustPrepare(db, `
			INSERT INTO users (id, username, password_hash, email) VALUES (?, ?, ?, ?)`),
	}
}

func (r *userRepository) FindAuthInfo(username string) (*user.AuthInfo, error) {
	m := &user.AuthInfo{}
	return m, r.selectAuthInfo.QueryRow(username).Scan(
		&m.UserID, &m.PasswordHash, &m.IsLocked,
	)
}

func (r *userRepository) FindUserByID(id string) (*user.User, error) {
	m := &user.User{}
	return m, r.selectUserByID.QueryRow(id).Scan(
		&m.Username, &m.FirstName, &m.LastName,
	)
}

func (r *userRepository) HasAccount(username string) (bool, error) {
	var s string
	err := r.existsByUsername.QueryRow(username).Scan(&s)
	switch {
	case err == sql.ErrNoRows:
		return false, nil
	case err != nil:
		return true, err
	default:
		return true, nil
	}
}

func (r *userRepository) CreateAccount(reg *user.Registration) (bool, error) {
	_, err := r.insertAccount.Exec(
		reg.UserID, reg.Username, reg.PasswordHash, reg.Email,
	)
	if err != nil {
		return false, err
	}
	return true, nil
}
