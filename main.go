package main

import (
	"github.com/Imran-Sarkar-Sabbir/gin-jwt-auth/controllers"
	"github.com/Imran-Sarkar-Sabbir/gin-jwt-auth/initializers"
	"github.com/Imran-Sarkar-Sabbir/gin-jwt-auth/middlewares"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/signup", controllers.SignUp)
	r.POST("/signin", controllers.Login)
	r.GET("/validate", middlewares.AuthMiddleware, controllers.Validate)

	r.Run() // listen and serve on 0.0.0.0:8080

}
