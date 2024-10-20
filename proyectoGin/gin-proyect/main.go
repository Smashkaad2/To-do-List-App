package main

import (
    "github.com/gin-gonic/gin"
)

func main() {
    // Crear un enrutador Gin
    router := gin.Default() // En Gin siempre parece haber un enrutador que es el que manda las peticiones

	router.GET("/", func(c *gin.Context) { // Si no quieres especificar una ruta para la URL
		c.String(200, "Funny")
	})

	router.GET("/bye", func(c *gin.Context) { // Si quiero especificar una ruta para la URL (Sin JSON)
		c.String(200, "GoodBye World")
	})

	// router.POST("/create", func(c *gin.Context) { // Ejemplo de POST antes de ver la documentacion
	// 	c.JSON(200, gin.H{
    //         "nombre": "Javier",
	// 		"edad": 23,
	// 		"lenguaje" : "Golang",
    //     })
	// })

    router.GET("/ping", func(c *gin.Context) { // Si quiero especificar una ruta para la URL (Con JSON)
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })

    router.Run(":8081") // Puerto en el que corre el enrutador
	// "go run main.go" --> Comando para correr el servidor

	//Probando Branch
}
