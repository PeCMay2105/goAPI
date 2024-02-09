package main

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

type task struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var tasks = []task{
	{ID: 1, Title: "Task 1", Completed: false},
	{ID: 2, Title: "Task 2", Completed: true},
	{ID: 3, Title: "Task 3", Completed: false},
}

func getTasks(context *gin.Context) {
	context.IndentedJSON(200, tasks)
}

func getOneTask(id int) (*task, error) {
	for item, taskList := range tasks {
		if taskList.ID == id {
			return &tasks[item], nil
		}
	}
	return nil, errors.New("Task not found")
}

func getTask(context *gin.Context) {
	id := context.Param("id")
	idInt, errSTR := strconv.Atoi(id)
	if errSTR != nil {
		fmt.Printf("An invalid value for 'ID' was passed previously. Raw error: %v", errSTR)
	}
	task, err := getOneTask(idInt)
	if err != nil {
		context.IndentedJSON(404, gin.H{"message": "Task not found"})
		return
	} else {
		context.IndentedJSON(200, task)

	}

}

func postTask(context *gin.Context) {
	var newTask task

	if err := context.BindJSON(&newTask); err != nil {
		fmt.Printf("An error occurred while binding the data to a json format: %v", err)
		return
	}

	tasks = append(tasks, newTask)
	context.IndentedJSON(201, newTask)

}

func completeTask(context *gin.Context) {
	id := context.Param("id")
	idInt, errSTR := strconv.Atoi(id)
	if errSTR != nil {
		fmt.Printf("An invalid value for 'ID' was passed previously. Raw error: %v", errSTR)
	}
	task, err := getOneTask(idInt)
	if err != nil {
		context.IndentedJSON(404, gin.H{"message": "Task not found"})
		return
	}
	task.Completed = !task.Completed
	context.IndentedJSON(202, task)

}

func main() {
	router := gin.Default()
	router.GET("/taskManager", getTasks)
	router.GET("/taskManager/:id", getTask)
	router.POST("/taskManager", postTask)
	router.PATCH("/taskManager/:id", completeTask)
	router.Run("localhost:3314")

}
