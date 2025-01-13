package repository

import (
	"fmt"
	"log"
	"testing"
	

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func mockDatabase() (*gorm.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		log.Fatalf("an error '%s'", err)
	}

	gormDb, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})

	if err != nil {
		log.Fatalf("gormDatabase %s", err)
	}
	return gormDb, mock
}

// func TestFlight(t *testing.T) {

// 	mockDb, mock := mockDatabase()
// 	flightRepository := FlightRepository{DB: mockDb}
// 	t.Run("searchFlight", func(t *testing.T) {

// 		columns := []string{"id", "source", "destination", "date", "departureTime", "arrivalTime", "priceSaver", "priceFlexi", "Airline"}

// 		rows := sqlmock.NewRows(columns).AddRow(1, "delhi", "mumbai", "12/03/2025", "12/03/2025", "12/03/2025", 2000, 3000, "indigo")

// 		mock.ExpectQuery("SELECT * FROM `flights` WHERE (source= ? and destination= ? and date = ? ) AND `flights`.`deleted_at` IS NULL").
// 			WithArgs("delhi", "mumbai", "2025-01-08 14:51:57").
// 			WillReturnRows(rows)

// 		// var flights []models.Flight
// 		result, err := flightRepository.GetFlights("delhi", "mumbai", time.Now())

// 		assert.NoError(t, err)
// 		for _, value := range *result {
// 			assert.Equal(t, value.Source, "delhi")
// 			assert.Equal(t, value.Destination, "mumbai")
// 		}

// 	})
// }

func TestFlight(t *testing.T) {
	mockDb, mock := mockDatabase()
	flightRepository := FlightRepository{DB: mockDb}

	t.Run("searchFlight", func(t *testing.T) {
		columns := []string{"id", "source", "destination", "date", "departureTime", "arrivalTime", "saverPrice", "flexiPrice", "airline"}

		rows := sqlmock.NewRows(columns).
			AddRow(1, "delhi", "mumbai", "2025-01-08", "2025-01-08", "2025-01-08", 2000, 3000, "indigo")

		// Modify the regex to allow spaces between SQL parts
		mock.ExpectQuery("SELECT * FROM `flights` WHERE (source = ? and destination = ? and date = ?) AND `flights`.`deleted_at` IS NULL").
			WithArgs("delhi", "mumbai", "2025-01-08").
			WillReturnRows(rows)

		result, err := flightRepository.GetFlights("delhi", "mumbai", "2025-01-10")
		fmt.Println("Actual SQL Query: ", mock.ExpectationsWereMet())

		assert.NoError(t, err)
		for _, value := range *result {
			assert.Equal(t, value.Source, "delhi")
			assert.Equal(t, value.Destination, "mumbai")
		}
	})
}
