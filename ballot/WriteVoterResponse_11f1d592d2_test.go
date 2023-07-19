package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

type Status struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func writeVoterResponse(w http.ResponseWriter, status Status) {
	w.Header().Set("Content-Type", "application/json")
	resp, err := json.Marshal(status)
	if err != nil {
		log.Println("error marshaling response to vote request. error: ", err)
	}
	w.Write(resp)
}

func TestWriteVoterResponse_11f1d592d2(t *testing.T) {
	status := Status{Code: 200, Message: "Success"}

	// create a response recorder
	w := httptest.NewRecorder()

	// call the function
	writeVoterResponse(w, status)

	// check the status code
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	// check the response body
	body, _ := ioutil.ReadAll(w.Body)
	var resp Status
	json.Unmarshal(body, &resp)

	if resp.Code != status.Code || resp.Message != status.Message {
		t.Errorf("Expected response %v, got %v", status, resp)
	}

	// check the content type
	if w.Header().Get("Content-Type") != "application/json" {
		t.Errorf("Expected content type %s, got %s", "application/json", w.Header().Get("Content-Type"))
	}
}

func TestWriteVoterResponse_11f1d592d2_Error(t *testing.T) {
	status := Status{Code: 200, Message: string(make([]byte, 999999))} // this will cause an error

	// create a response recorder
	w := httptest.NewRecorder()

	// call the function
	writeVoterResponse(w, status)

	// check the status code
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	// check the response body
	body, _ := ioutil.ReadAll(w.Body)
	var resp Status
	err := json.Unmarshal(body, &resp)

	if err == nil {
		t.Error("Expected error, got nil")
	}
}
