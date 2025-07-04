package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID       string `json:"id"`
	Item     string `json:"item"`
	Complete bool   `json:"completed"`
}

var todos = []todo{
	{ID: "1", Item: "Clean Room", Complete: false},
	{ID: "2", Item: "Read Book", Complete: false},
	{ID: "3", Item: "Record Video", Complete: false},
}

func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

func addTodo(context *gin.Context) {
	var newTodo todo

	if err := context.BindJSON(&newTodo); err != nil {
		return
	}
	todos = append(todos, newTodo)
	context.IndentedJSON(http.StatusCreated, newTodo)
}

func getTodo(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodoById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo Not Found"})
		return
	}

	context.IndentedJSON(http.StatusOK, todo)
}


func toggleTodoStatus(context *gin.Context){
	id:=context.Param("id")
	todo,err:=getTodoById(id)
	if err!=nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo Not Found"})
		return
	}
	todo.Complete=!todo.Complete
	context.IndentedJSON(http.StatusOK,todo)
}


func getTodoById(id string) (*todo, error) {
	for i, t := range todos {
		if t.ID == id {
			return &todos[i], nil
		}
	}
	return nil, errors.New("todo not found")
}

func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.GET("/todos/:id", getTodo)
	router.PATCH("/todos/:id", toggleTodoStatus)
	router.POST("/todos", addTodo)
	router.Run("localhost:9090")
}
