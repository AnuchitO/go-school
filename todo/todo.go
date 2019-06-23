package todo

import (
	"net/http"

	"github.com/AnuchitO/school/database"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

func GetHandler(c *gin.Context) {
	conn := database.Connect()
	rows := database.SelectAllTodo(conn)

	todos := []Todo{}
	for rows.Next() {
		t := Todo{}
		err := rows.Scan(&t.ID, &t.Title, &t.Status)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		todos = append(todos, t)
	}

	c.JSON(http.StatusOK, todos)
}

func CreateHandler(c *gin.Context)  {}
func GetByIdHandler(c *gin.Context) {}
func UpdateHandler(c *gin.Context)  {}
func DeleteHandler(c *gin.Context)  {}
