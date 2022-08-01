package api

import (
	"log"
	"win/fake-cards/internal/database"
)

type ApiConfig struct {
	Port            string
	InfoLog, ErrLog *log.Logger
	DB              database.DBFaker
}
