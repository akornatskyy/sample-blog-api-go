package sqlx

import (
	"database/sql"
	"log"
)

func MustPrepare(db *sql.DB, query string) *sql.Stmt {
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatalf(
			"sql.DB: Prepare(query) %s: %s",
			err.Error(),
			query)
	}
	return stmt
}
