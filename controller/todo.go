package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"todo-api-sample/model"
	"todo-api-sample/service"

	"github.com/gin-gonic/gin"
)

func TodoAdd(c *gin.Context) {
     todo := model.Todo{}
     err := c.Bind(&todo)
     if err != nil{
         c.String(http.StatusBadRequest, "Bad request")
         return
     }
    todoService :=service.TodoService{}
    err = todoService.SetTodo(&todo)
    if err != nil{
        fmt.Printf("eeeeeee: %v\n", err)
        c.String(http.StatusInternalServerError, err.Error())
        return
    }
    c.JSON(http.StatusCreated, gin.H{
        "status": "ok",
    })
}

func TodoList(c *gin.Context){
    todoService :=service.TodoService{}
    TodoLists := todoService.GetTodoList()
    c.JSONP(http.StatusOK, gin.H{
        "message": "created",
        "data": TodoLists,
    })
}

func TodoUpdate(c *gin.Context){
    todo := model.Todo{}
    err := c.Bind(&todo)
    if err != nil{
        c.String(http.StatusBadRequest, "Bad request")
        return
    }
    todoService :=service.TodoService{}
    id := c.Param("id")
    intId, parseErr := strconv.ParseInt(id, 10, 0)
    if parseErr != nil{
        c.String(http.StatusBadRequest, "Bad request")
        return
    }
    err = todoService.UpdateTodo(int(intId), &todo)
    if err != nil{
        c.String(http.StatusInternalServerError, "Server Error")
        return
    }
    c.JSON(http.StatusCreated, gin.H{
        "status": "ok",
    })
}

func TodoDelete(c *gin.Context){
    id := c.Param("id")
    intId, err := strconv.ParseInt(id, 10, 0)
    if err != nil{
        c.String(http.StatusBadRequest, "Bad request")
        return
    }
    todoService :=service.TodoService{}
    err = todoService.DeleteTodo(int(intId))
    if err != nil{
        c.String(http.StatusInternalServerError, "Server Error")
        return
    }
    c.JSON(http.StatusCreated, gin.H{
        "status": "ok",
    })
}