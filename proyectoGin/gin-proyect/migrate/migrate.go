package main

import (
	"gin-proyect/configs"
	"gin-proyect/models"
)

func init() {
	configs.ConnectToDB()
}

func main() {
	configs.DB.AutoMigrate(&models.Person{})
	configs.DB.AutoMigrate(&models.Task{})
}