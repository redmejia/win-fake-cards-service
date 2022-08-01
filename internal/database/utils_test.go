package database

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"testing"
	"win/fake-cards/internal/data"
)

// dbCard
type dbCard []data.Card

// mockDBPost
type mockDBPost struct {
	Db               dbCard
	DNS              string
	InfoLog, ErroLog *log.Logger
}

// newMockDBPost
func newMockDBPost() *mockDBPost {
	db := dbCard{
		data.Card{
			FullName:   "Elok Musk",
			CardNumber: "1111222233331234",
			CvNumber:   "123",
			Amount:     100000000,
			StatusCode: 2, // has enough
			Proceed:    true,
		},
		data.Card{
			FullName:   "Elok Musk",
			CardNumber: "1111222233334444",
			CvNumber:   "123",
			Amount:     0,
			StatusCode: 4, // not conplete no enough founding to proceed
			Proceed:    false,
		},
	}

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime)

	return &mockDBPost{
		Db:      db,
		DNS:     "",
		InfoLog: infoLog,
		ErroLog: errLog,
	}

}

func (m *mockDBPost) GenerateFakeCards(twelveNum string, amountInCent int, statusCode int, proceed bool) (fakeCardPool []data.Card) {

	for i := 0; i < 5; i++ {

		fakeCard := data.GenFakeCards(twelveNum)

		fakeCvNum := data.GenFakeCv()

		cards := data.Card{
			FullName:   "Elon Musk",
			CardNumber: fakeCard,
			CvNumber:   fakeCvNum,
			StatusCode: statusCode,
			Amount:     amountInCent,
			Proceed:    proceed,
		}
		// insert or append data card to mock database
		m.Db = append(m.Db, cards)

		fakeCardPool = append(fakeCardPool, cards)

	}
	return
}

func (m *mockDBPost) GetInfo(cardNum, cardCV string) (data.Card, error) {
	for _, v := range m.Db {
		if v.CardNumber == cardNum && v.CvNumber == cardCV {
			return v, nil
		}
	}
	return data.Card{}, NoRowInResultSet
}

func TestGetInfo(t *testing.T) {
	// test card and cv found must return card type
	// "1111222233331234"
	// 123
	cardNumer := "1111222233331234"
	cardCv := "123"

	dbase := newMockDBPost()
	card, err := dbase.GetInfo(cardNumer, cardCv)
	if err != nil {
		if errors.Is(err, NoRowInResultSet) {
			var notFoundRecord = struct {
				IsError bool   `json:"is_error"`
				Message string `json:"message"`
			}{
				IsError: true,
				Message: "record not found",
			}
			b, _ := json.Marshal(&notFoundRecord)
			expectError := `{"is_error":true,"message":"record not found"}`
			if expectError != string(b) {
				t.Errorf("not matching expect %s, but  %s was recived\n", expectError, string(b))
			}
		}
	}

	if card.FullName != "" {
		t.Log(card)
	}

}

func TestGenFakeCards(t *testing.T) {

	dbase := newMockDBPost()
	_ = dbase.GenerateFakeCards("111122223333", 1000, 2, true)

	if len(dbase.Db) > 2 {
		t.Log("Records were inserted")
	}

}
