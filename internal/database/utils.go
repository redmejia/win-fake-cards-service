package database

import (
	"context"
	"log"
	"time"
	"win/fake-cards/internal/data"
)

// GenerateFakeCards generate testing cards
func (p *Postgresql) GenerateFakeCards(twelveNum string, amountInCent int, statusCode int, proceed bool) (fakeCardPool []data.Card) {
	// cent 100000 = $1000
	// success 2 status code
	// decline 4 status code

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	query := `
 		insert into cards (card_issuer, card_number, card_cv_number, status_code, amount, proceed)
 		values ($1, $2, $3, $4, $5, $6)
 	`
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

		_, err := p.Db.ExecContext(
			ctx,
			query,
			cards.FullName,
			cards.CardNumber,
			cards.CvNumber,
			statusCode,
			cards.Amount,
			proceed,
		)

		if err != nil {
			log.Println("err ", err)
			return
		}

		fakeCardPool = append(fakeCardPool, cards)

	}

	return
}

// GeyInfo get ths card information
func (p *Postgresql) GetInfo(cardNum, cardCv string) data.Card {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `
		select 
			card_issuer, 
			card_number, 
			card_cv_number, 
			status_code, 
			amount, 
			proceed 
		from 
			cards

		where
			card_number = $1 and card_cv_number = $2
			`

	row := p.Db.QueryRowContext(ctx, query, cardNum, cardCv)

	var card data.Card

	err := row.Scan(
		&card.FullName,
		&card.CardNumber,
		&card.CvNumber,
		&card.StatusCode,
		&card.Amount,
		&card.Proceed,
	)

	if err != nil {
		p.ErrorLog.Fatal(err)
	}

	return card
}
