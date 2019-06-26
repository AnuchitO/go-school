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

type Handler struct {
	ts TodoServicer
}

func NewHandler() *Handler {
	db := database.NewPostgres()
	return &Handler{
		ts: &TodoService{
			db: db,
		},
	}
}

// GetHandler ...
func (h *Handler) GetHandler(c *gin.Context) {
	todos, err := h.ts.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, todos)
}

// GetByIdHandler ...
func (h *Handler) GetByIdHandler(c *gin.Context) {
	pID := c.Param("id")
	id, err := strconv.Atoi(pID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo, err := h.ts.FindByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, todo)
}

func (h *Handler) CreateHandler(c *gin.Context) {}
func (h *Handler) UpdateHandler(c *gin.Context) {}
func (h *Handler) DeleteHandler(c *gin.Context) {}
