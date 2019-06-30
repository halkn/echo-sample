package gateway

// SQLHandler is interface of SQLHandler
type SQLHandler interface {
	Query(string, ...interface{}) (Row, error)
	Exec(string, ...interface{}) (Result, error)
}

// Row is a wrapper interface of sql.Row
type Row interface {
	Scan(...interface{}) error
	Next() bool
	Close() error
}

// Result is a wraper interface og sql.Result.
type Result interface {
	LastInsertId() (int64, error)
	RowsAffected() (int64, error)
}
