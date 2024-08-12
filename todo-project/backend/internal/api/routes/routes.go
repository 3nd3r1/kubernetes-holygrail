package routes

import (
	"net/http"
	"todo-project-backend/internal/api/handlers"
	"todo-project-backend/internal/repositories"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) error {
	todoRepository, err := repositories.NewTodoRepository()
	if err != nil {
		return err
	}
	todoHandler := handlers.NewTodoHandler(todoRepository)

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})

	apiGroup := router.Group("/api")
	{
		apiGroup.GET("/todos", todoHandler.GetAllTodos)
		apiGroup.POST("/todos", todoHandler.CreateTodo)
	}

	return nil
}
