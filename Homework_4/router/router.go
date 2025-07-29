package router

import (
	"todo/handler"
	"todo/store"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	memStore := store.NewMemoryStore()
	taskHandler := handler.NewTaskHandler(memStore)

	r.GET("/tasks", taskHandler.GetAll)
	r.GET("/tasks/:id", taskHandler.GetByID)
	r.POST("/tasks", taskHandler.Create)
	r.PUT("/tasks/:id", taskHandler.Update)
	r.DELETE("/tasks/:id", taskHandler.Delete)

	return r
}
