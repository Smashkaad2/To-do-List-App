package controllers

import (
	"fmt"
	"gin-proyect/configs"
	"gin-proyect/models"

	"github.com/gin-gonic/gin"
)

type TaskRequestBody struct {
	TaskName string        `json:"taskname"`
	Status   models.Status `json:"taskstatus"`
}

func TaskCreate(c *gin.Context) {
	body := TaskRequestBody{}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(400, gin.H{"Error": "Invalid request body"})
		return
	}

	if !models.StatusValido(body.Status) {
		c.JSON(400, gin.H{"Error": fmt.Sprintf("Invalid status: %s. Allowed values are: COMPLETED, PENDING, FREE", body.Status)})
		return
	}

	task := &models.Task{TaskName: body.TaskName, Status: body.Status}

	result := configs.DB.Create(&task)

	if result.Error != nil {
		errorMessage := fmt.Sprintf("Failed to insert task: %s with status: %s", body.TaskName, body.Status)
		c.JSON(500, gin.H{"Error": errorMessage})
		return
	}
	c.JSON(200, &task)
}

func TaskGet(c *gin.Context) {
	var tasks []models.Task
	configs.DB.Find(&tasks)
	c.JSON(200, &tasks)
	return
}

func TaskGetById(c *gin.Context) {
	id := c.Param("id")
	var task models.Task
	configs.DB.First(&task, id)
	c.JSON(200, &task)
	return
}

func TaskUpdate(c *gin.Context) {

	id := c.Param("id")
	var task models.Task
	configs.DB.First(&task, id)

	body := TaskRequestBody{}
	c.BindJSON(&body)
	data := &models.Task{TaskName: body.TaskName, Status: body.Status}

	result := configs.DB.Model(&task).Updates(data)

	if result.Error != nil {
		c.JSON(500, gin.H{"Error": true, "message": "Failed to update"})
		return
	}

	c.JSON(200, &task)
}

func TaskDelete(c *gin.Context) {
	id := c.Param("id")
	var task models.Task
	configs.DB.Delete(&task, id)
	c.JSON(200, gin.H{"deleted": true})
	return
}
