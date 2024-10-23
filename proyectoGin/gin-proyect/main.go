package main

import (
	"log"
	"net/http"

	"gin-proyect/configs"
	"gin-proyect/routes"
	"gin-proyect/models"

	"github.com/gin-gonic/gin"
)

func init() {
	configs.ConnectToDB()
}

func main() {

	r := gin.Default()

	routes.PersonRouter(r)
	routes.TaskRouter(r)

	if err := configs.DB.AutoMigrate(&models.Person{}); err != nil {
		log.Fatalf("Failed to migrate Person: %v", err)
	}

	if err := configs.DB.AutoMigrate(&models.Task{}); err != nil {
		log.Fatalf("Failed to migrate Task: %v", err)
	}

	configs.DB.AutoMigrate(&models.Person{}, &models.Task{})

	log.Println("Database migration completed!")

	configs.SeedDB()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello world from server Go.",
		})
	})

	r.Run(":8081")
	log.Println("Server Running....")
	// 	// "go run main.go" --> Comando para correr el servidor
	// 	//Probando Branch
}
