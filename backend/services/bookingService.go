package services

import (
	"errors"
	"flight/models"
	"flight/repository"
	"math/rand"
	"strconv"
	"time"
)

type BookingService struct {
	BookingRepo   *repository.BookRepository
	PassengerRepo *repository.PassengerRepository
}

func (service *BookingService) CreateBooking(bookingDTO models.BookingDetailsDTO) (*models.Booking, error) {
	if len(bookingDTO.PassengerDetails) == 0 {
		return nil, errors.New("passenger details cannot be empty")
	}

	// Generate PNR number
	pnrNumber := generatePNRNumber()

	booking := models.Booking{
		FlightID:   bookingDTO.FlightID,
		Price:      bookingDTO.DefaultPrice,
		TicketType: bookingDTO.TicketType, // Assuming default ticket type
		PNRnumber:  pnrNumber,
	}

	result, err := service.BookingRepo.Create(booking)
	if err != nil {
		return nil, err
	}

	for _, passengerDetail := range bookingDTO.PassengerDetails {
		passenger := models.Passenger{
			BookingId:  int(result.ID),
			Title:      passengerDetail.Title,
			FirstName:  passengerDetail.FirstName,
			LastName:   passengerDetail.LastName,
			SeatNumber: convertSeatNumber(passengerDetail.SeatNumber),
			Email:      passengerDetail.Email,
		}

		_, err := service.PassengerRepo.Create(passenger)
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

func convertSeatNumber(seatNumber string) int {
	seatNum, err := strconv.Atoi(seatNumber)
	if err != nil {
		return 0 // or handle the error as needed
	}
	return seatNum
}

func generatePNRNumber() string {
	rand.Seed(time.Now().UnixNano())
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, 6)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
