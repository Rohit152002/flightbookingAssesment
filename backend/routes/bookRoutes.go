package routes

import (
	"flight/controllers"
	"flight/repository"
	"flight/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func BookRoutes(router *gin.Engine, db *gorm.DB) {
	passengerRepo := repository.PassengerRepository{DB: db}
	bookingRepo := repository.BookRepository{DB: db}

	bookingSer := services.BookingService{BookingRepo: &bookingRepo, PassengerRepo: &passengerRepo}

	bookingController := controllers.BookingController{BookingService: &bookingSer}

	router.POST("/books", bookingController.BookPassengersAsync)
}
