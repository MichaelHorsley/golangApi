package main

import (
	"net/http"
	"testing"
	"net/http/httptest"
	"bytes"
)


var a App

func Test_GivenValidParticipant_WhenPostToParticipants_ThenReturns201Code(t *testing.T) {

	a = App{}
	a.Initialize()

	payload := []byte(`{"name":"Michael Horsley","email":"Michael.Horsley@sorted.com"}`)

	req, _ := http.NewRequest("POST", "/participants", bytes.NewBuffer(payload))
	response := executeRequest(req)

	checkResponseCode(t, http.StatusCreated, response.Code)
}

func BenchmarkHello(b *testing.B) {

	a = App{}
	a.Initialize()

	for i := 0; i < b.N; i++ {

		payload := []byte(`{"name":"Michael Horsley","email":"Michael.Horsley@sorted.com"}`)

		req, _ := http.NewRequest("POST", "/participants", bytes.NewBuffer(payload))
		executeRequest(req)
	}
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