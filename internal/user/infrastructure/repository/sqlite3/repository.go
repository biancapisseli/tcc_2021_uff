package userreposqlite3

type UserSQLite3Repository struct{}

func New() (repo *UserSQLite3Repository) {
	return &UserSQLite3Repository{}
}
