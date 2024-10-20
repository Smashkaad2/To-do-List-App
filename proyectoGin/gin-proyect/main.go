package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gin-proyect/configs"
	"gin-proyect/routes"
)

func init() {
	configs.ConnectToDB()
}

func main() {

	r := gin.Default()

	routes.PersonRouter(r)

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello world from server Go.",
		})
	})
	r.Run(":8081")
	// 	// "go run main.go" --> Comando para correr el servidor
	// 	//Probando Branch
}
