package main

import (
	"github.com/AnuchitO/school/todo"
	"github.com/gin-gonic/gin"
)

func newRoute() *gin.Engine {
	r := gin.Default()

	r.GET("/api/todos", todo.GetHandler)
	r.GET("/api/todos/:id", todo.GetByIdHandler)
	r.POST("/api/todos", todo.UpdateHandler)
	r.DELETE("/api/todos/:id", todo.DeleteHandler)
	return r
}

func main() {
	r := newRoute()

	r.Run(":1234")
}
