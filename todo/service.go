package todo

import (
	"github.com/AnuchitO/school/database"
)

type TodoServicer interface {
	FindAll() ([]Todo, error)
	FindByID(id int) (Todo, error)
}

type TodoService struct {
	db *database.Postgres // TODO: change to interface
}

func (ts *TodoService) FindAll() ([]Todo, error) {
	rows := ts.db.SelectAllTodo()

	todos := []Todo{}
	for rows.Next() {
		t := Todo{}
		err := rows.Scan(&t.ID, &t.Title, &t.Status)
		if err != nil {
			return []Todo{}, err
		}
		todos = append(todos, t)
	}

	return todos, nil
}
func (ts *TodoService) FindByID(id int) (Todo, error) {
	row := ts.db.SelectByID(id)
	t := Todo{}
	err := row.Scan(&t.ID, &t.Title, &t.Status)
	if err != nil {
		return Todo{}, err
	}
	return t, nil
}
