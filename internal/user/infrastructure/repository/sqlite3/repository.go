package userreposqlite3

import "ifoodish-store/pkg/sqlite3"

type UserSQLite3Repository struct {
	db *sqlite3.Connection
}

func New(db *sqlite3.Connection) (repo *UserSQLite3Repository) {
	return &UserSQLite3Repository{
		db: db,
	}
}
