package main

import (
	"GinWebAPIHello/data"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello")

	r := gin.Default()

	r.GET("/tasks", GetTasks)

	r.Run()
}

func GetTasks(c *gin.Context) {

	tasks, err := data.ReadTasks()

	if err != nil {
		c.JSON(503, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(200, tasks)
}
