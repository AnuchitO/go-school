package main

import (
	"github.com/AnuchitO/school/todo"
	"github.com/gin-gonic/gin"
)

func newRoute() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")

	td := todo.NewHandler()

	api.GET("/todos", td.GetHandler)
	api.GET("/todos/:id", td.GetByIdHandler)
	api.POST("/todos", td.UpdateHandler)
	api.DELETE("/todos/:id", td.DeleteHandler)

	return r
}

func main() {
	r := newRoute()

	r.Run(":1234")
}
