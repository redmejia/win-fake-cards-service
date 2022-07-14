package router

import (
	"net/http"
	"win/fake-cards/cmd/api"
)

func Router(apiConfig *api.ApiConfig) http.Handler {

	mux := http.NewServeMux()

	mux.HandleFunc("/api/txintent", apiConfig.TxIntentHandler)

	return mux

}
