package routers

import (
	"github.com/WTBacon/todo-app-gin/controllers"
	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.GET("todo", controllers.FetchAllTodo)
		v1.GET("todo/:id", controllers.FetchSingleTodo)
		v1.POST("todo", controllers.CreateTodo)
		v1.PUT("todo/:id", controllers.UpdateTodo)
		v1.DELETE("todo/:id", controllers.DeleteTodo)
	}
	return r
}
