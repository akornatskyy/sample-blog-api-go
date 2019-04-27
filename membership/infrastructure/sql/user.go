package sql

import (
	"database/sql"
	"log"

	"github.com/akornatskyy/sample-blog-api-go/membership/domain/user"
	"github.com/akornatskyy/sample-blog-api-go/shared/sqlx"
)

type userRepository struct {
	SelectAuthInfo   *sql.Stmt `SELECT id, password_hash, is_locked FROM users WHERE username = ?`
	SelectUserByID   *sql.Stmt `SELECT username, first_name, last_name FROM users where id = ?`
	ExistsByUsername *sql.Stmt `SELECT '' FROM users WHERE username = ?`
	InsertAccount    *sql.Stmt `INSERT INTO users (id, username, password_hash, email) VALUES (?, ?, ?, ?)`
}

func NewUserRepository(db *sql.DB) user.Repository {
	r := &userRepository{}
	if err := sqlx.Prepare(db, r); err != nil {
		log.Fatal(err)
	}
	return r
}

func (r *userRepository) FindAuthInfo(username string) (*user.AuthInfo, error) {
	m := &user.AuthInfo{}
	return m, r.SelectAuthInfo.QueryRow(username).Scan(
		&m.UserID, &m.PasswordHash, &m.IsLocked,
	)
}

func (r *userRepository) FindUserByID(id string) (*user.User, error) {
	m := &user.User{}
	return m, r.SelectUserByID.QueryRow(id).Scan(
		&m.Username, &m.FirstName, &m.LastName,
	)
}

func (r *userRepository) HasAccount(username string) (bool, error) {
	var s string
	err := r.ExistsByUsername.QueryRow(username).Scan(&s)
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
	_, err := r.InsertAccount.Exec(
		reg.UserID, reg.Username, reg.PasswordHash, reg.Email,
	)
	if err != nil {
		return false, err
	}
	return true, nil
}
