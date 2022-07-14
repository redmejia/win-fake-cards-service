package data

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// Card Fake cards
type Card struct {
	FullName   string `json:"card_issuer"`
	CardNumber string `json:"card_number"`
	CvNumber   string `json:"cv_number"`
	StatusCode int    `json:"status_code"`
	Amount     int    `json:"amount"`
	Proceed    bool   `json:"proceed"`
}

const (
	// last four number on the card
	minNum = 1000
	maxNum = 3000
	// cv number card
	minCv = 100
	maxCv = 200
)

func GenFakeCards(firstTwelveNum string) (card string) {
	rand.Seed(time.Now().UnixNano())

	genNum := rand.Intn(maxNum-minNum) + minNum
	lastFour := strconv.Itoa(genNum)

	card = fmt.Sprintf("%s%s", firstTwelveNum, lastFour)
	return
}

func GenFakeCv() (cv string) {
	rand.Seed(time.Now().UnixNano())

	genCvNum := rand.Intn(maxCv-minCv) + minCv
	cvNum := strconv.Itoa(genCvNum)

	cv = fmt.Sprintf("%s", cvNum)

	return

}
