package controllers

import (
	"fmt"
	"gin-proyect/configs"
	"gin-proyect/models"
	"github.com/gin-gonic/gin"
)

type TaskRequestBody struct {
	TaskName string `json:"taskname"`
	Status   string `json:"taskstatus"`
}

func TaskCreate(c *gin.Context) {
	body := TaskRequestBody{}

	c.BindJSON(&body)

	task := &models.Task{TaskName: body.TaskName, Status: body.Status}

	result := configs.DB.Create(&task)

	if result.Error != nil {
		errorMessage := fmt.Sprintf("Failed to insert task: %s with status: %s", body.TaskName, body.Status)
		c.JSON(500, gin.H{"Error": errorMessage})
		return
	}
	c.JSON(200, &task)
}