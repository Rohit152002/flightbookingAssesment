package models

type BookingDetailsDTO struct {
	FlightID         int               `json:"flightId"`
	DefaultPrice     float64           `json:"defaultPrice"`
	PassengerDetails []PassengerDetail `json:"passengerDetails"`
}

type PassengerDetail struct {
	Title      string `json:"title"`
	Name       string `json:"name"`
	Age        int    `json:"age"`
	Gender     string `json:"gender"`
	SeatNumber string `json:"seatNumber"`
}
