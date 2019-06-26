package todo

import (
	"net/http"
	"strconv"

	"github.com/AnuchitO/school/database"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

type TodoServicer interface {
	FindAll() ([]Todo, error)
	FindById(id string) (Todo, error)
}

type TodoService struct {
}

func (ts *TodoService) FindAll() ([]Todo, error) {
	conn := database.Connect()
	rows := database.SelectAllTodo(conn)

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
	conn := database.Connect()
	row := database.SelectByID(conn, id)
	t := Todo{}
	err := row.Scan(&t.ID, &t.Title, &t.Status)
	if err != nil {
		return Todo{}, err
	}
	return t, nil
}

// GetHandler ...
func GetHandler(c *gin.Context) {
	ts := &TodoService{}

	todos, err := ts.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, todos)
}

// GetByIdHandler ...
func GetByIdHandler(c *gin.Context) {
	ts := &TodoService{}
	pID := c.Param("id")
	id, err := strconv.Atoi(pID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo, err := ts.FindByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, todo)
}

func CreateHandler(c *gin.Context) {}
func UpdateHandler(c *gin.Context) {}
func DeleteHandler(c *gin.Context) {}
