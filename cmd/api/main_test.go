package api

import (
	"log"
	"os"
	"testing"
	"win/fake-cards/internal/data"
)

var app ApiConfig

func TestMain(m *testing.M) {

	var data = []data.Card{
		{
			FullName:   "Elon Musk",
			CardNumber: "1111222233334444",
			CvNumber:   "123",
			StatusCode: 2,
			Amount:     100000,
			Proceed:    true,
		},
	}

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime)

	var dbMock dbMockPost
	dbMock.DB = data
	dbMock.InfoLog = infoLog
	dbMock.ErrorLog = errLog

	// dbMock.InfoLog = infoLog
	// dbMock.ErrorLog = errLog

	app.DB = &dbMock
	app.InfoLog = dbMock.InfoLog
	app.ErrLog = dbMock.ErrorLog

	code := m.Run()

	os.Exit(code)
}
