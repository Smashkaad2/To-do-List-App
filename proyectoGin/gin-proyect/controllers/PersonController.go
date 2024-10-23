package controllers

import (
	"gin-proyect/configs"
	"gin-proyect/models"

	"github.com/gin-gonic/gin"
)

type PersonRequestBody struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Phone   uint   `json:"phone"`
}

func PersonCreate(c *gin.Context) {

	body := PersonRequestBody{}

	c.BindJSON(&body)

	person := &models.Person{Name: body.Name, Address: body.Address, Phone: body.Phone}

	result := configs.DB.Create(&person)

	if result.Error != nil {
		c.JSON(500, gin.H{"Error": "Failed to insert"})
		return
	}

	c.JSON(200, &person)
}

func PersonGet(c *gin.Context) {
	var persons []models.Person

	if err := configs.DB.Preload("Tasks").Find(&persons).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch persons"})
		return
	}

	c.JSON(200, &persons)
}

func PersonGetById(c *gin.Context) {
	id := c.Param("id")
	var person models.Person
	if err := configs.DB.Preload("Tasks").Find(&person,id).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch persons"})
		return
	}
	c.JSON(200, &person)
	
}

func PersonUpdate(c *gin.Context) {

	id := c.Param("id")
	var person models.Person
	configs.DB.First(&person, id)

	body := PersonRequestBody{}
	c.BindJSON(&body)
	data := &models.Person{Name: body.Name, Address: body.Address, Phone: body.Phone}

	result := configs.DB.Model(&person).Updates(data)

	if result.Error != nil {
		c.JSON(500, gin.H{"Error": true, "message": "Failed to update"})
		return
	}

	c.JSON(200, &person)
}

func PersonDelete(c *gin.Context) {
	id := c.Param("id")
	var person models.Person
	configs.DB.Delete(&person, id)
	c.JSON(200, gin.H{"deleted": true})
	return
}
