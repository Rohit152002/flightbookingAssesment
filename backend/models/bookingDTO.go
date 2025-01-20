package models

type BookingDetailsDTO struct {
	FlightID         int               `json:"flightId"`
	DefaultPrice     float64           `json:"defaultPrice"`
	TicketType       string            `json:"ticketType"`
	PassengerDetails []PassengerDetail `json:"passengerDetails"`
}

type EmailSendingDTO struct {
	Email      string `json:"email"`
	ReferencNo string `json:"referenceNo"`
}

type PassengerDetail struct {
	Title      string `json:"title"`
	FirstName  string `json:"firstname"`
	LastName   string `json:"lastname"`
	Email      string `json:"email"`
	SeatNumber int    `json:"seatNumber"`
}
