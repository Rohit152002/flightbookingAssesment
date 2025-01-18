package controllers

import (
	"flight/models"
	"flight/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BookingController struct {
	BookingService *services.BookingService
}

func (ctrl *BookingController) BookPassengersAsync(ctx *gin.Context) {
	var books models.BookingDetailsDTO
	if err := ctx.ShouldBindJSON(&books); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	result, err := ctrl.BookingService.CreateBooking(books)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"results": result,
	})

}
