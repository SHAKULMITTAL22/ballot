package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServeRoot_e6109c0b6f(t *testing.T) {
	// Test case 1: GET method
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(serveRoot)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Test case 2: POST method with valid vote data
	vote := `{"VoterID": "voter1", "CandidateID": "candidate1"}`
	req, err = http.NewRequest("POST", "/", bytes.NewBufferString(vote))
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}

	// Test case 3: POST method with invalid vote data
	vote = `{"VoterID": "", "CandidateID": ""}`
	req, err = http.NewRequest("POST", "/", bytes.NewBufferString(vote))
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}

	// Test case 4: Unsupported method
	req, err = http.NewRequest("PUT", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusMethodNotAllowed)
	}
}
