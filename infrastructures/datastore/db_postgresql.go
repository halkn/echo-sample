package datastore

import (
	"database/sql"
	"fmt"
)

func NewDB() *sql.DB {

	conn, err := sql.Open("postgres", "host=127.0.0.1 port=5555 user=postgres password=password dbname=tododb sslmode=disable")
	if err != nil {
		fmt.Println(err)
	}

	return conn

}
