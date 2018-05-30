package main

import (
	"net/http"
	"testing"
	"net/http/httptest"
	"bytes"
)


var a App

func TestItSavesParticipant(t *testing.T) {

	a = App{}
	a.Initialize()

	payload := []byte(`{"name":"Michael Horsley","email":"Michael.Horsley@sorted.com"}`)

	req, _ := http.NewRequest("POST", "/participants", bytes.NewBuffer(payload))
	response := executeRequest(req)

	checkResponseCode(t, http.StatusCreated, response.Code)
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}