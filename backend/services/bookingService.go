package services

import (
	"errors"
	"flight/models"
	"flight/repository"
	"fmt"
	"html/template"
	"math/rand"

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
		FlightID:           bookingDTO.FlightID,
		Price:              bookingDTO.DefaultPrice,
		TicketType:         bookingDTO.TicketType, // Assuming default ticket type
		PNRnumber:          pnrNumber,
		BookingReferenceNo: generateBookingReferenceNo(),
	}

	result, err := service.BookingRepo.Create(booking)
	if err != nil {
		return nil, err
	}

	for _, passengerDetail := range bookingDTO.PassengerDetails {
		passenger := models.Passenger{
			BookingID:  int(result.ID),
			Title:      passengerDetail.Title,
			FirstName:  passengerDetail.FirstName,
			LastName:   passengerDetail.LastName,
			SeatNumber: passengerDetail.SeatNumber,
			Email:      passengerDetail.Email,
		}

		_, err := service.PassengerRepo.Create(passenger)
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

func (service *BookingService) FindBookingByReferenceNo(referenceNo string) (*models.Booking, error) {
	return service.BookingRepo.FindByReferenceNumber(referenceNo)
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

func generateHtml(bookingDetails models.Booking) string {
	html := `
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <style>
        body {
            font-family: Arial, sans-serif;
            background-color: #f9fafb;
            margin: 0;
            padding: 20px;
        }
        .container {
            max-width: 800px;
        }
        .ticket {
            background: white;
            border-radius: 12px;
            box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
            margin-bottom: 24px;
            overflow: hidden;
        }
        .ticket-header {
            background: linear-gradient(to right, #2563eb, #1d4ed8);
            padding: 24px;
            color: white;
            display: flex;
            justify-content: space-between;
            align-items: center;
        }
        .ticket-title {
            font-size: 20px;
            font-weight: bold;
            margin: 0;
        }
        .pnr {
		color:black;
            font-size: 14px;
            opacity: 0.9;
        }
        .ticket-body {
            padding: 24px;
        }
        .flight-info {
            display: flex;
            justify-content: space-between;
            align-items: center;
            margin-bottom: 16px;
        }
        .flight-id {
            color: #6b7280;
            font-size: 14px;
        }
        .route{
		width: 100%;
		 border-collapse: collapse;
		  text-align: center;
		  position:relative;
        }
        .location {
		display:flex;
            text-align: center;
            flex: 1;
        }
        .location-label {
            color: #6b7280;
            font-size: 14px;
            margin-bottom: 4px;
        }
        .location-city {
            font-size: 18px;
            font-weight: bold;
            color: #111827;
            margin: 4px 0;
        }
        .location-time {
            font-size: 14px;
            color: #374151;
        }
        .route-line {
            flex: 1;
            height: 2px;
            border-top: 2px dashed #d1d5db;
            margin: 0 20px;
            position: relative;
        }
        .route-arrow {
            position: absolute;
            top: 50%;
            left: 50%;
            transform: translate(-50%, -50%);
            color: #2563eb;
            font-size: 24px;
        }
        .passenger-details {
            margin-top: 20px;
            padding-top: 20px;
            border-top: 1px solid #e5e7eb;
        }
        .passenger-box {
            background-color: #f9fafb;
            padding: 16px;
            border-radius: 8px;
        }
        .passenger-label {
            color: #6b7280;
            font-size: 14px;
            margin-bottom: 8px;
        }
        .passenger-name {
            font-size: 18px;
            font-weight: bold;
            color: #111827;
            margin: 0;
        }
    </style>
</head>
<body>
    <div class="container">
    `

	for _, passenger := range bookingDetails.Passengers {
		html += fmt.Sprintf(`
        <div class="ticket">
            <div class="ticket-header">
                <h2 class="ticket-title">Flight Ticket</h2>
                <span class="pnr">PNR: %s</span>
				<h2 class="ticket-title">Flight Ticket</h2>
				<span class="pnr">Seat Number: %d</span>
            </div>




            <div class="ticket-body">
                <div class="flight-info">
                    <span class="flight-id">Flight ID: %d</span>
                </div>

                <table class="route">
				<tr>
                            <td style="padding: 8px;">

                        <div class="location-label">From</div>
                        <div class="location-city">%s</div>
                        <div class="location-time">%s</div>
         </td>


                        <td class="route-arrow">→</td>


                    <td class="padding: 8px;">
                        <div class="location-label">To</div>
                        <div class="location-city">%s</div>
                        <div class="location-time">%s</div>
                    </td>
					</tr>
                </table>

                <div class="passenger-details">
                    <div class="passenger-box">
                        <div class="passenger-label">Passenger Details</div>
                        <p class="passenger-name">%s %s</p>
                    </div>
                </div>


            </div>
        </div>`,
			bookingDetails.PNRnumber,
			passenger.SeatNumber,
			bookingDetails.FlightID,
			bookingDetails.Flight.Source,
			formatTime(bookingDetails.Flight.DepartureTime),
			bookingDetails.Flight.Destination,
			formatTime(bookingDetails.Flight.ArrivalTime),
			passenger.FirstName,
			passenger.LastName,
		)
	}

	html += `
    </div>
</body>
</html>`

	return html
}

func formatTime(timestamp time.Time) string {
	return timestamp.Format("03:04 PM")
}

func (service *BookingService) GeneretePdfToDownload(referenceNo string) (*template.Template, error) {
	result, err := service.BookingRepo.FindByReferenceNumber(referenceNo)
	if err != nil {
		return nil, err
	}
	tmpl := template.Must(template.New("page").Parse(generateHtml(*result)))
	return tmpl, nil
}

func generateBookingReferenceNo() string {
	rand.Seed(time.Now().UnixNano())
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, 9)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
