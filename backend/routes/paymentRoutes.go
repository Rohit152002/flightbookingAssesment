package routes

import (
	"flight/controllers"

	"github.com/gin-gonic/gin"
)

func PaymentRoutes(router *gin.Engine) {
	paymentController := controllers.PaymentController{}
	router.POST("/verification", paymentController.PaymentVerification)
}
