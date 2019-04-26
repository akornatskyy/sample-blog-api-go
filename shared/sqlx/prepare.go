package sqlx

import (
	"database/sql"
	"errors"
	"fmt"
	"reflect"
)

func Prepare(db *sql.DB, r interface{}) error {
	v := reflect.ValueOf(r)
	if v.Kind() != reflect.Ptr || v.IsNil() {
		return errors.New("expects not nil ptr")
	}
	v = v.Elem()
	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		if !f.CanSet() {
			continue
		}
		switch f.Interface().(type) {
		case *sql.Stmt:
			query := string(v.Type().Field(i).Tag)
			if query == "" {
				continue
			}
			stmt, err := db.Prepare(query)
			if err != nil {
				return fmt.Errorf(
					"sql.DB: Prepare(query) for %s.%s: %s",
					v.Type().Name(),
					v.Type().Field(i).Name,
					err.Error())
			}
			f.Set(reflect.ValueOf(stmt))
		}
	}
	return nil
}
