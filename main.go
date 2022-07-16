package main

import (
	"log"
	"net/http"
	"os"
	"win/fake-cards/cmd/api"
	"win/fake-cards/cmd/router"
	"win/fake-cards/internal/database"
)

func main() {

	db, err := database.Connection()
	defer db.Close()
	if err != nil {
		log.Println("ERROR ", err)
	}

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime)

	app := api.ApiConfig{
		Port:    ":8083",
		InfoLog: infoLog,
		ErrLog:  errLog,
		DB: database.Postgresql{
			Db:       db,
			InfoLog:  infoLog,
			ErrorLog: errLog,
		},
	}

	srv := &http.Server{
		Addr:    app.Port,
		Handler: router.Router(&app),
	}

	log.Println("Server is running at http://localhost:8080/")
	log.Fatal(srv.ListenAndServe())

}
