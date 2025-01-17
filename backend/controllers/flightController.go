package controllers

import (
	"flight/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FlightController struct {
	FlightService services.FlightServices
}

func (ctrl *FlightController) GetStationController(c *gin.Context) {
	station, err := ctrl.FlightService.GetDistinctSources()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not get Stations"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": station,
	})
}

func (ctrl *FlightController) SearchFlightController(c *gin.Context) {
	source := c.Query("source")
	destination := c.Query("destination")
	date := c.Query("date")

	flights, err := ctrl.FlightService.GetFlights(source, destination, date)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error in parsing date"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": flights,
	})
}

func (ctrl *FlightController) TwoWayFlightSearchController(c *gin.Context) {
	source := c.Query("source")
	destination := c.Query("destination")
	departureDate := c.Query("departure-date")
	returnDate := c.Query("return-date")

	departureFlights, returnFlights, err := ctrl.FlightService.GetFlightsWithReturn(source, destination, departureDate, returnDate)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error in parsing date"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"departureFlights": departureFlights,
		"returnFlights":    returnFlights,
	})
}
