package router

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"win/fake-cards/cmd/api"
)

// testing routes
var app api.ApiConfig

func TestRoutes(t *testing.T) {
	fakeData := []byte(`"is_test": true`)

	req, err := http.NewRequest("POST",
		"/api/txintent?card=1111222233332900&cv=157&amount=100",
		bytes.NewBuffer(fakeData),
	)
	if err != nil {
		t.Fatal(err)
	}

	defer req.Body.Close()

	recorder := httptest.NewRecorder()

	handler := http.HandlerFunc(app.TxIntentHandler)

	handler.ServeHTTP(recorder, req)
	// change this json message to test  POST/PATCH or DELETE method not implemented
	wantError := `{"error":true,"message":"POST is not implemented"}`

	if recorder.Code == http.StatusNotImplemented && recorder.Body.String() == wantError {
		t.Logf("expeced error %s menthod not implemented", req.Method)
	}
}
