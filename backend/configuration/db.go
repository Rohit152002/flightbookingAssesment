package configuration

import (
	"flight/models"
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DatabaseDsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DATABASE"),
	)
}

func ConnectDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open(DatabaseDsn()), &gorm.Config{})
	if err != nil {
		panic("Failed to connect DB")
	}
	fmt.Print("database connected")
	db.AutoMigrate(&models.Flight{})
	db.AutoMigrate(&models.Booking{})
	db.AutoMigrate(&models.Passenger{})
	return db
}
