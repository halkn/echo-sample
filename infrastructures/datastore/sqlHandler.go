package datastore

import (
	"database/sql"

	"github.com/halkn/echo-sample/interfaces/gateway"

	// "lib/pq" is a package fot postgresql driver and is not used directly.
	_ "github.com/lib/pq"
)

// SQLHandler is a handler responsible for connection to DB.
type SQLHandler struct {
	Conn *sql.DB
}

// NewSQLHandler create a new sqlhandler.
func NewSQLHandler(conn *sql.DB) *SQLHandler {
	return &SQLHandler{Conn: conn}
}

// Query exetutes a query that returns rows.
func (sqlHandler *SQLHandler) Query(statement string, args ...interface{}) (gateway.Row, error) {

	rows, err := sqlHandler.Conn.Query(statement)
	if err != nil {
		return new(SQLRow), err
	}

	row := new(SQLRow)
	row.Rows = rows
	return row, nil

}

// SQLRow is an implementation of presenters.Row.
type SQLRow struct {
	Rows *sql.Rows
}

// Scan ...
func (r SQLRow) Scan(dest ...interface{}) error {
	return r.Rows.Scan(dest...)
}

// Next ...
func (r SQLRow) Next() bool {
	return r.Rows.Next()
}

// Close ...
func (r SQLRow) Close() error {
	return r.Rows.Close()
}
