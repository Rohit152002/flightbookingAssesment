package controllers

import (
	"flight/models"
	"flight/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PaymentController struct {
}

func (pc *PaymentController) PaymentVerification(ctx *gin.Context) {
	var payment models.PaymentModel
	if err := ctx.ShouldBindJSON(&payment); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	valid, err := services.CheckPaymentValidation(&payment)
	if !valid {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Payment details are valid"})
}
