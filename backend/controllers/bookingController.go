package controllers

import (
	"bytes"
	"flight/models"
	"flight/services"
	"log"
	"net/http"

	"github.com/andrewcharlton/wkhtmltopdf-go"
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

func (ctrl *BookingController) FindBooking(ctx *gin.Context) {
	referenceNo := ctx.Param("referenceNo")

	result, err := ctrl.BookingService.FindBookingByReferenceNo(referenceNo)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}

func (ctrl *BookingController) DownloadPdf(ctx *gin.Context) {
	referenceNo := ctx.Param("referenceNo")

	temp, err := ctrl.BookingService.GeneretePdfToDownload(referenceNo)
	if err != nil {

		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	buf := &bytes.Buffer{}
	temp.Execute(buf, ctx.Request.URL.String())
	doc := wkhtmltopdf.NewDocument()
	pg, err := wkhtmltopdf.NewPageReader(buf)
	if err != nil {
		log.Fatal("Error reading page buffer")
	}
	doc.AddPages(pg)

	// Set headers and serve the PDF
	ctx.Header("Content-Type", "application/pdf")
	ctx.Header("Content-Disposition", `attachment; filename="tickets.pdf"`)
	err = doc.Write(ctx.Writer)
	if err != nil {
		log.Fatal("Error serving pdf")
	}
}

func (ctrl *BookingController) SendingEmail(ctx *gin.Context) {
	var emailDto models.EmailSendingDTO
	if err := ctx.ShouldBindJSON(&emailDto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	result, err := ctrl.BookingService.SendMail(emailDto.Email, emailDto.ReferencNo)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

	}

	ctx.JSON(http.StatusOK, gin.H{
		"result":  result,
		"message": "email has been sent successfull",
	})
}
