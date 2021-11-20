package postgresql

import (
	"fmt"
	"strconv"

	"github.com/seefmitrais/go-rest-api-practice/internal/config"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

// create connection to postgresql database
func CreateConnection() {
	var err error
	dbHost := config.Get("DB_HOST")
	p := config.Get("DB_PORT")
	dbUser := config.Get("DB_USER")
	dbPassword := config.Get("DB_PASSWORD")
	dbName := config.Get("DB_NAME")
	dbPort, err := strconv.ParseUint(p, 10, 32)
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)
	DB, err = gorm.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to PostgreSQL Database!")
}
