package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"win/fake-cards/cmd/api"
	"win/fake-cards/cmd/router"
	"win/fake-cards/internal/database"
)

func main() {

	db, err := database.Connection()
	if err != nil {
		log.Println("ERROR ", err)
	}
	defer db.Close()

	fmt.Println("TESTING CARDS")

	// test database
	var valid = database.GenFakesTestCardsToInsert("111122223333", 100000, 2, true)
	fmt.Println(" =========== valid cards")
	fmt.Printf("%s\n", database.CardForTest(valid))
	fmt.Println(" =========== invalid cards")
	var invalid = database.GenFakesTestCardsToInsert("222233334444", 0, 4, false)
	fmt.Printf("%s\n", database.CardForTest(invalid))

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime)

	app := api.ApiConfig{
		Port:    ":8083",
		InfoLog: infoLog,
		ErrLog:  errLog,
		DB: &database.Postgresql{
			Db:       db,
			InfoLog:  infoLog,
			ErrorLog: errLog,
		},
	}

	srv := &http.Server{
		Addr:    app.Port,
		Handler: router.Router(&app),
	}
	log.Println()
	log.Println("Server is running at http://localhost:8083/")
	log.Fatal(srv.ListenAndServe())

}
