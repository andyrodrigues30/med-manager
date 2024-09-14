package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"andyrodrigues30/account-service/config"
	"andyrodrigues30/account-service/controllers"
)

func main() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Content-Type", "Authorization"},
	}))

	// ==ROUTES==

	// users
	router.POST("/register", controllers.Register)
	// router.POST("/login", controllers.Login)
	// router.POST("/verify", controllers.Verify)

	PORT := config.GetEnvVariable("PORT")
	URL := "localhost:" + PORT

	router.Run(URL)
}
