package gateway

// SQLHandler is interface of SQLHandler
type SQLHandler interface {
	Query(string, ...interface{}) (Row, error)
}

// Row is a wrapper interface of sql.Row
type Row interface {
	Scan(...interface{}) error
	Next() bool
	Close() error
}
