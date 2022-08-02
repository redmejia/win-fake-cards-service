package database

import "testing"

func TestPostgresqlConnection(t *testing.T) {

	db, err := Connection()
	if err != nil {
		t.Log(err)
	}
	defer db.Close()

	connected, err := ConnectionPing(db)
	if err != nil {
		t.Fatal(err)
	}

	if !connected {
		t.Log("unable to connect")
	}
	t.Log("Connected")

}
