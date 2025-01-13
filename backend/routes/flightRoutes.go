package routes

import (
	"flight/controllers"
	"flight/repository"
	"flight/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func FlightRoutes(router *gin.Engine, db *gorm.DB) {
	flightRepository := repository.FlightRepository{DB: db}

	flightService := services.FlightServices{Repo: flightRepository}

	flightController := controllers.FlightController{FlightService: flightService}

	router.GET("/station", flightController.GetStationController)
	router.GET("/flights", flightController.SearchFlightController)
	router.GET("/flights/two-way", flightController.TwoWayFlightSearchController)
}
