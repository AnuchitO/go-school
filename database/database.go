package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type Postgres struct {
	conn *sql.DB
}

func NewPostgres() *Postgres {
	conn := connect()
	return &Postgres{
		conn: conn,
	}
}

func connect() *sql.DB {
	conn, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("con't connect to database ", err.Error())
	}
	return conn
}

func (p *Postgres) SelectAllTodo() *sql.Rows {
	stmt, _ := p.conn.Prepare("SELECT id, title, status FROM todos")
	rows, _ := stmt.Query()
	return rows
}

func (p *Postgres) SelectByID(id int) *sql.Row {
	stmt, _ := p.conn.Prepare("SELECT id, title, status FROM todos WHERE id = $1")
	row := stmt.QueryRow(id)
	return row
}
