package sqlite3

import (
	"database/sql"
	"errors"

	sqlite3 "github.com/mattn/go-sqlite3"
)

func isConstraintErr(err error, extendedCode sqlite3.ErrNoExtended) bool {
	if sqlite3Err, ok := err.(sqlite3.Error); ok {
		if sqlite3Err.Code == sqlite3.ErrConstraint &&
			sqlite3Err.ExtendedCode == extendedCode {
			return true
		}
	}
	return false
}

func IsErrNoRows(err error) bool {
	return errors.Is(err, sql.ErrNoRows)
}

func IsUniqueErr(err error) bool {
	return isConstraintErr(err, 2067)
}

func IsPrimaryKeyErr(err error) bool {
	return isConstraintErr(err, 1555)
}

func IsForeignKeyErr(err error) bool {
	return isConstraintErr(err, 787)
}
