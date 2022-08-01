package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"win/fake-cards/internal/data"

	_ "github.com/jackc/pgx/v4/stdlib"
)

type Postgresql struct {
	Db                *sql.DB
	DNS               string
	InfoLog, ErrorLog *log.Logger
}

// DBFaker for new cards check founding and create cards methods
type DBFaker interface {
	GenerateFakeCards(twelveNum string, amountInCent int, statusCode int, proceed bool) (fakeCardPool []data.Card)
	GetInfo(cardNum, cardCv string) (data.Card, error)
}

// Ping connection to postgresql database
func ConnectionPing(db *sql.DB) (bool, error) {
	err := db.Ping()
	if err != nil {
		return false, err
	}
	return true, nil
}

// Connection
func Connection() (*sql.DB, error) {

	port, _ := strconv.Atoi(os.Getenv("DBPORT"))
	connDB := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s",
		os.Getenv("HOSTNAME"), port, os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
	)

	db, err := sql.Open("pgx", connDB)
	if err != nil {
		return nil, err
	}
	return db, err

}

// DnsConnection
func DnsConnection(dns string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dns)
	if err != nil {
		return nil, err
	}

	return db, nil
}
