package models

type BookingDetailsDTO struct {
	FlightID         int               `json:"flightId"`
	DefaultPrice     float64           `json:"defaultPrice"`
	TicketType       string            `json:"ticketType"`
	PassengerDetails []PassengerDetail `json:"passengerDetails"`
}

type PassengerDetail struct {
	Title      string `json:"title"`
	FirstName  string `json:"firstname"`
	LastName   string `json:"lastname"`
	Email      string `json:"email"`
	SeatNumber string `json:"seatNumber"`
}
