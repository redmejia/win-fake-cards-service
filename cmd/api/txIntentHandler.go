package api

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type Transaction struct {
	Proceed        bool `json:"proceed"`
	TxAmountIntent int  `json:"tx_amount_intent"`
	TxStatusCode   int  `json:"tx_status_code"`
}

func (a *ApiConfig) TxIntentHandler(w http.ResponseWriter, r *http.Request) {
	// valid card 1111222233332369 | 103
	// http:/localhost:8080/api/txintent?card=1222&cv=123&amount=1000
	// transaction intent
	txIntent := r.URL.Query()

	card := a.DB.GetInfo(txIntent.Get("card"), txIntent.Get("cv"))

	txAmount, _ := strconv.Atoi(txIntent.Get("amount"))

	var tx Transaction

	if txAmount < card.Amount {
		// sent to new micro service to make the envoice
		tx.Proceed = card.Proceed
		tx.TxAmountIntent = txAmount
		tx.TxStatusCode = card.StatusCode
	} else {
		tx.Proceed = card.Proceed
		tx.TxAmountIntent = txAmount
		tx.TxStatusCode = card.StatusCode
	}

	txByte, err := json.Marshal(tx)
	if err != nil {
		a.ErrLog.Fatal(err)
	}

	a.InfoLog.Println(string(txByte))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(txByte)

}
