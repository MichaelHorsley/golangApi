package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"poolgolang/src/participant"
	"strconv"
	"testing"
)

var a App

func Test_GivenValidParticipant_WhenPostToParticipants_ThenReturns201Code(t *testing.T) {

	a = App{}
	a.Initialize()

	response := AddParticipant()

	checkResponseCode(t, http.StatusCreated, response.Code)
}

func Test_GivenValidParticipant_WhenPutToParticipants_ThenReturns200Code(t *testing.T) {
	a = App{}
	a.Initialize()

	postResponse := AddParticipant()
	participant := participant.Participant{}
	json.Unmarshal(postResponse.Body.Bytes(), &participant)

	payload := []byte(`{"name":"Michael Horsley","email":"Michael.Horsley@sorted.com"}`)

	req, _ := http.NewRequest("PUT", "/participants/", bytes.NewBuffer(payload))
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)
}

func AddParticipant() *httptest.ResponseRecorder {
	payload := []byte(`{"name":"Michael Horsley","email":"Michael.Horsley@sorted.com"}`)

	req, _ := http.NewRequest("POST", "/participants", bytes.NewBuffer(payload))
	response := executeRequest(req)
	return response
}

func Benchmark_Hello(b *testing.B) {
	a = App{}
	a.Initialize()

	for i := 0; i < b.N; i++ {
		payload := []byte(`{"name":"Michael Horsley` + strconv.Itoa(i) + `","email":"Michael.Horsley` + strconv.Itoa(i) + `@sorted.com"}`)
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
