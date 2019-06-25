package database

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

func Connect() *sql.DB {
	conn, _ := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	return conn
}

func SelectAllTodo(conn *sql.DB) *sql.Rows {
	stmt, _ := conn.Prepare("SELECT id, title, status FROM todos")
	rows, _ := stmt.Query()

	return rows
}
