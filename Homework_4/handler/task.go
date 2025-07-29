package handler

import (
	"net/http"
	"strconv"
	"todo/model"
	"todo/store"

	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	store *store.MemoryStore
}

func NewTaskHandler(s *store.MemoryStore) *TaskHandler {
	return &TaskHandler{store: s}
}

func (h *TaskHandler) GetAll(c *gin.Context) {
	c.JSON(http.StatusOK, h.store.GetAll())
}

func (h *TaskHandler) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	task, err := h.store.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
		return
	}
	c.JSON(http.StatusOK, task)
}

func (h *TaskHandler) Create(c *gin.Context) {
	var task model.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}
	newTask := h.store.Create(task)
	c.JSON(http.StatusCreated, newTask)
}

func (h *TaskHandler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	var task model.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}
	if err := h.store.Update(id, task); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
		return
	}
	c.Status(http.StatusNoContent)
}

func (h *TaskHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	if err := h.store.Delete(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
		return
	}
	c.Status(http.StatusNoContent)
}
