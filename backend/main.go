package main

import (
	"flight/configuration"
	"flight/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic("Failed to load env file")
	}
	httpServer := gin.Default()

	db := configuration.ConnectDB()

	httpServer.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200,
			gin.H{
				"working": "everything well",
			},
		)
	})

	httpServer.Use(cors.Default())

	routes.FlightRoutes(httpServer, db)

	httpServer.Run(":8080")
}
