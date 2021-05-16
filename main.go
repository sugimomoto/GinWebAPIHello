package main

import (
	"GinWebAPIHello/data"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello")

	r := gin.Default()

	r.GET("/tasks", GetTasks)
	r.GET("/tasks/:taskid", GetTask)
	r.POST("/tasks", AddTask)

	r.Run()
}

func GetTasks(c *gin.Context) {

	tasks, err := data.ReadTasks()

	if err != nil {
		c.JSON(500, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(200, tasks)
}

func GetTask(c *gin.Context) {

	taskid, err := strconv.Atoi(c.Param("taskid"))

	if err != nil {
		c.JSON(500, gin.H{
			"message": "TaskId mast be integer type.",
		})
		return
	}

	task, err := data.ReadTask(taskid)

	if err != nil {
		c.JSON(500, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(200, task)
}

func AddTask(c *gin.Context) {
	subject := c.PostForm("subject")
	priority := c.PostForm("priority")

	task := data.Task{
		Subject:  subject,
		Priority: priority,
	}

	insertid, err := task.CreateTask()

	if err != nil {
		c.JSON(500, gin.H{
			"message": err,
		})
		return
	}

	result, err := data.ReadTask(int(insertid))

	if err != nil {
		c.JSON(500, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(200, result)
}
