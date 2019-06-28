package infrastructure

import (
	"database/sql"
	"fmt"

	"github.com/halkn/echo-sample/interfaces/presenters"

	// "lib/pq" is a package fot postgresql driver and is not used directly.
	_ "github.com/lib/pq"
)

// SQLHandler is a handler responsible for connection to DB.
type SQLHandler struct {
	Conn *sql.DB
}

// NewSQLHandler create a new sqlhandler.
func NewSQLHandler() *SQLHandler {

	conn, err := sql.Open("postgres", "host=127.0.0.1 port=5555 user=postgres password=password dbname=tododb sslmode=disable")
	if err != nil {
		fmt.Println(err)
	}

	sqlHandler := new(SQLHandler)
	sqlHandler.Conn = conn

	return sqlHandler

}

// Query exetutes a query that returns rows.
func (sqlHandler *SQLHandler) Query(statement string, args ...interface{}) (presenters.Row, error) {

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
