package services

import (
	"bytes"
	"errors"
	"flight/models"
	"flight/repository"
	"fmt"
	"html/template"
	"io"
	"log"
	"math/rand"

	"time"

	"github.com/andrewcharlton/wkhtmltopdf-go"
	"gopkg.in/gomail.v2"
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
	basehtml := `
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>Flight Booking Report</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            background-color: #f9fafb;
            margin: 0;
            padding: 20px;
        }
        table {
            border-collapse: collapse;
            width: 100%%;
        }
        .main-container {
            width: 800px;
            margin: 0 auto;
        }
        .report-header {
            background: white;
            padding: 24px;
            border-radius: 12px;
            margin-bottom: 24px;
            box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
        }
        .section-title {
            color: #111827;
            font-size: 24px;
            margin: 0 0 24px 0;
            padding: 0;
        }
        .sub-title {
            color: #374151;
            font-size: 18px;
            margin: 0 0 16px 0;
            padding: 0;
        }
        .info-cell {
            background: #f9fafb;
            padding: 16px;
            border-radius: 8px;
            vertical-align: top;
        }
        .info-label {
            color: #6b7280;
            font-size: 14px;
            margin-bottom: 4px;
        }
        .info-value {
            font-size: 16px;
            font-weight: bold;
            color: #111827;
            margin-bottom: 12px;
        }
        .spacer-row {
            height: 24px;
        }
        .divider {
            border-top: 1px solid #e5e7eb;
            margin: 24px 0;
        }
    </style>
</head>
<body>
    <table class="main-container">
        <tr>
            <td>
                <div class="report-header">
                    <h1 class="section-title">Flight Booking Report</h1>

                    <!-- Booking Info -->
                    <table>
                        <tr>
                            <td width="50%%" class="info-cell">
                                <div class="info-label">Booking Reference Number</div>
                                <div class="info-value">%s</div>
                            </td>
                            <td width="50%%" class="info-cell" style="padding-left: 16px;">
                                <div class="info-label">Flight ID</div>
                                <div class="info-value">%d</div>
                            </td>
                        </tr>
                    </table>

                    <div class="divider"></div>

                    <!-- Flight Details -->
                    <h2 class="sub-title">Flight Details</h2>
                    <table>
                        <tr>
                            <td width="50%%" class="info-cell">
                                <div class="info-label">From</div>
                                <div class="info-value">%s</div>
                                <div class="info-label">Departure Time</div>
                                <div class="info-value">%s</div>
                            </td>
                            <td width="50%%" class="info-cell" style="padding-left: 16px;">
                                <div class="info-label">To</div>
                                <div class="info-value">%s</div>
                                <div class="info-label">Arrival Time</div>
                                <div class="info-value">%s</div>
                            </td>
                            <td width="50%%" class="info-cell" style="padding-left: 16px;">
                                <div class="info-label">Price</div>
                                <div class="info-value">Rs %.2v</div>
                            </td>
                        </tr>
                    </table>

                    <div class="divider"></div>

                    <!-- Passenger Details -->
                    <h2 class="sub-title">Passenger Details</h2>`

	html := fmt.Sprintf(basehtml,
		bookingDetails.PNRnumber,
		bookingDetails.FlightID,
		bookingDetails.Flight.Source,
		formatTime(bookingDetails.Flight.DepartureTime),
		bookingDetails.Flight.Destination,
		formatTime(bookingDetails.Flight.ArrivalTime),
		int(bookingDetails.Price)*len(bookingDetails.Passengers))

	// Add passenger details
	for _, passenger := range bookingDetails.Passengers {
		passengerHtml := fmt.Sprintf(`
                    <table>
                        <tr>
                            <td class="info-cell" style="margin-bottom: 12px;">
                                <table width="">
                                    <tr>
                                        <td width="50%%">
                                            <div class="info-label">Passenger Name</div>
                                            <div class="info-value">%s %s</div>
                                        </td>
                                        <td width="50%%">
                                            <div class="info-label">Seat Number</div>
                                            <div class="info-value">%d</div>
                                        </td>
                                    </tr>
                                </table>
                            </td>
                        </tr>
                        <tr class="spacer-row"><td></td></tr>`,
			passenger.FirstName,
			passenger.LastName,
			passenger.SeatNumber)
		html += passengerHtml
	}

	// Close all tags
	html += `
                </div>
            </td>
        </tr>
    </table>
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
func (service *BookingService) SendMail(email string, referenceNo string) (bool, error) {
	result, err := service.BookingRepo.FindByReferenceNumber(referenceNo)
	if err != nil {
		return false, err
	}
	htmlContent := generateHtml(*result)
	pdfData, err := generatePDF(htmlContent)
	if err != nil {
		return false, err
	}

	msg := gomail.NewMessage()
	msg.SetHeader("From", "laishramrohit15@gmail.com")
	msg.SetHeader("To", email)
	msg.SetHeader("Subject", "Flight Ticket Report")
	msg.SetBody("text/html", "<b>Here is your flight ticket</b>")
	msg.Attach("booking_report.pdf", gomail.SetCopyFunc(func(w io.Writer) error {
		_, err := w.Write(pdfData.Bytes())
		return err
	}))
	n := gomail.NewDialer("smtp.gmail.com", 587, "laishramrohit15@gmail.com", "bnrk tpou pjxx gdyu")

	// Send the email
	if err := n.DialAndSend(msg); err != nil {
		panic(err)
	}
	return true, nil
}

func generatePDF(htmlContent string) (*bytes.Buffer, error) {
	doc := wkhtmltopdf.NewDocument()
	buf := bytes.NewBufferString(htmlContent)
	pg, err := wkhtmltopdf.NewPageReader(buf)
	if err != nil {
		log.Fatal("Error reading from reader.")
	}
	doc.AddPages(pg)

	output := &bytes.Buffer{}
	errs := doc.Write(output)
	if errs != nil {
		log.Fatal("Error writing to writer.")
	}
	return output, nil
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
