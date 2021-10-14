package sqlite

import (
	"errors"
	"fmt"
	"sort"

	"github.com/jmoiron/sqlx"
)

var (
	ErrMigrationAlreadyExecuted    = errors.New("migration already executed")
	ErrEmptyMigrationList          = errors.New("empty migration list")
	ErrReachEndOfFirstVersionAfter = errors.New("reach the end of 'firstVersionAfter' function")
)

type migrator struct {
	db        *sqlx.DB
	attempted bool
	list      map[string][]string
}

func newMigrator(db *sqlx.DB, list map[string][]string) *migrator {
	return &migrator{
		db:        db,
		attempted: false,
		list:      list,
	}
}

func (m *migrator) run() (err error) {
	defer func() {
		if r := recover(); r != nil {
			switch v := r.(type) {
			case error:
				err = v
			default:
				err = fmt.Errorf("%+v", v)
			}
		}
	}()

	if m.attempted {
		return ErrMigrationAlreadyExecuted
	}
	m.attempted = true

	if m.list == nil || len(m.list) == 0 {
		return ErrEmptyMigrationList
	}

	sortedKeys := m.sortKeys()
	maxVersion := len(sortedKeys)
	var versionToBeApplied = 0

	for versionToBeApplied != maxVersion {
		if err := m.transact(func(tx *sqlx.Tx) error {

			if err := tx.QueryRow("PRAGMA user_version").Scan(&versionToBeApplied); err != nil {
				return err
			}

			if versionToBeApplied == maxVersion {
				return nil
			}

			if versionToBeApplied > maxVersion {
				return fmt.Errorf("current version %d is greater then latest version %d", versionToBeApplied, maxVersion)
			}

			for _, instruction := range m.list[sortedKeys[versionToBeApplied]] {
				if _, err := tx.Exec(instruction); err != nil {
					return err
				}
			}

			if _, err := tx.Exec(fmt.Sprintf("PRAGMA user_version=%d", versionToBeApplied+1)); err != nil {
				return err
			}

			return nil

		}); err != nil {
			return err
		}
	}

	return nil
}

func (m *migrator) transact(handler func(*sqlx.Tx) error) (err error) {
	if m.db == nil {
		return errors.New("nil db")
	}
	transaction, err := m.db.Beginx()
	if err != nil {
		return err
	}
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%+v", r)
			transaction.Rollback()
		} else if err != nil {
			transaction.Rollback()
		} else {
			err = transaction.Commit()
		}
	}()
	if err = handler(transaction); err != nil {
		return err
	}

	return nil
}

func (m *migrator) sortKeys() []string {
	keys := make([]string, len(m.list))
	i := 0
	for k := range m.list {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	return keys
}
