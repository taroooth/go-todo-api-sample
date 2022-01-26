package main

import (
	"github.com/gin-gonic/gin"
	"todo-api-sample/controller"
	"todo-api-sample/middleware"
)

func main() {
	engine := gin.Default()
	engine.Use(middleware.RecordUaAndTime)
	todoEngine := engine.Group("/todo")
	{
		v1 := todoEngine.Group("/v1")
		{
			v1.POST("/add", controller.TodoAdd)
			v1.GET("/list", controller.TodoList)
			v1.PUT("/update/:id", controller.TodoUpdate)
			v1.DELETE("/delete/:id", controller.TodoDelete)
		}
	}
	engine.Run(":3000")
}
