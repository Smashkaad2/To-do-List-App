package main

import (
    "github.com/gin-gonic/gin"
	"log"
	"time"
)

//* Esta función solamente se encarga de medir el tiempo de respuesta que toma una de las varias peticiones que tengo aqui en mi main, devuelve esto y el estado de dicha petición (200, 400, etc).
//* Es una función de Middleware de GIN
func LoggerMiddleware() gin.HandlerFunc { // Middleware, funciones que interceptan peticiones y respuestas HTTP pueden, pueden ejecutarse antes o despues de que la petición alcance su ruta designada
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		duration := time.Since(start)
		log.Printf("Request - Method: %s | Status: %d | Duration: %v", c.Request.Method, c.Writer.Status(), duration)
	}
}

//Middleware personalizado
func AuthMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		apiKey := c.GetHeader("X-API-Key")
		if apiKey == "" {
			c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
			return
		}
		c.Next()
	}
}

func main() {
    // Crear un enrutador Gin
    router := gin.Default() // En Gin siempre parece haber un enrutador que es el que manda las peticiones

	authGroup := router.Group("/api") //Para crear un enrutador especifico para ciertas peticiones
	authGroup.Use(AuthMiddleware())
	{
		authGroup.GET("/data", func(c *gin.Context) { // las peticiones con las rutas en cuestión
			c.JSON(200, gin.H{"message": "Authenticated and authorized!"})
		})
	}

	router.Use(LoggerMiddleware())// asi se llama en el main un middleware

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

	router.GET("/favicon.ico", func(c *gin.Context) {
		c.Status(204) // Retorna un estado 204 No Content para evitar el error 404 con el Favicon
	})

    router.Run(":8081") // Puerto en el que corre el enrutador
	// "go run main.go" --> Comando para correr el servidor

	//Probando Branch
}
