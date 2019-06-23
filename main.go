package main

import (
	"github.com/AnuchitO/school/todo"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	r := gin.Default()

	r.GET("/api/todos", todo.GetHandler)
	r.GET("/api/todos/:id", todo.GetByIdHandler)
	r.POST("/api/todos", todo.UpdateHandler)
	r.DELETE("/api/todos/:id", todo.DeleteHandler)

	// r.GET("/api/students", student.GetHandler)
	// r.GET("/api/students/:id", student.GetByIdHandler)
	// r.POST("/api/students", student.CreateHandler)
	// r.DELETE("/api/students/:id", student.DeleteHandler)

	r.Run(":1234")
}
