package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/jackc/pgx/v4/stdlib"
)

type Postgresql struct {
	Db                *sql.DB
	DNS               string
	InfoLog, ErrorLog *log.Logger
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
	connDB := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("HOSTNAME"), port, os.Getenv("DBUSER"), os.Getenv("DBPASSWORD"),
		os.Getenv("DBNAME"), os.Getenv("DBSSLMODE"),
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
