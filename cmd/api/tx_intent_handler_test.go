package api

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"win/fake-cards/internal/data"
)

type dbMockPost struct {
	DB                []data.Card
	InfoLog, ErrorLog *log.Logger
}

func (d *dbMockPost) GenerateFakeCards(twelveNum string, amountInCent int, statusCode int, proceed bool) (fakeCardPool []data.Card) {
	fakeCardPool = []data.Card{}
	return
}

func (d *dbMockPost) GetInfo(cardNum, cardCv string) (data.Card, error) {
	for _, c := range d.DB {
		if c.CardNumber == cardNum && c.CvNumber == cardCv {
			return c, nil
		}
	}
	return data.Card{}, nil
}

func TestTxHandler(t *testing.T) {

	// http:/localhost:8080/api/txintent?card=1222&cv=123&amount=1000
	req, err := http.NewRequest("GET", "/api/txintent?card=1111222233334444&cv=123&amount=100", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(app.TxIntentHandler)

	handler.ServeHTTP(recorder, req)

	// this response can change on every request
	// card can have enough amoun or not
	expected := `{"proceed":true,"tx_amount_intent":100,"tx_status_code":2,"tx_message":"Transanction Accepted"}`
	if expected != recorder.Body.String() {
		t.Errorf("expected %s but %s was recived", expected, recorder.Body.String())
	}
}
