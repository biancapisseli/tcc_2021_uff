package sqlxtx

type SQLResult interface {
	LastInsertId() (int64, error)
	RowsAffected() (int64, error)
}

type Transaction interface {
	NamedExec(query string, arg interface{}) (SQLResult, error)
	Select(dest interface{}, query string, args ...interface{}) error
	Get(dest interface{}, query string, args ...interface{}) error
	Exec(query string, args ...interface{}) (SQLResult, error)
}

type TransactionFinisher interface {
	Transaction
	Commit() (err error)
	Rollback() (err error)
}

type Transactioner interface {
	BeginTransaction() (tx TransactionFinisher, err error)
}
