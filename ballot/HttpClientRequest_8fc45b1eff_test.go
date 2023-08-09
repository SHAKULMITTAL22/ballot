package test

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net"
	"net/http"
	"sort"
	"strings"
	"sync"
)

func TestHttpClientRequest(t *testing.T) {
	// Test successful request
	testCase1 := func() {
		operation := "GET"
		hostAddr := "example.com"
		command := "/path"
		params := bytes.NewBufferString(`{"key": "value"}`)
		expectedStatusCode := http.StatusOK
		expectedBody := `{
			"message": "success"
		}`
		expectedError := ""

		statusCode, body, err := httpClientRequest(operation, hostAddr, command, params)
		if statusCode!= expectedStatusCode {
			t.Errorf("Expected status code %d, got %d", expectedStatusCode, statusCode)
		}
		if string(body)!= expectedBody {
			t.Errorf("Expected body %q, got %q", expectedBody, body)
		}
		if err!= nil {
			t.Errorf("Expected no error, got %v", err)
		}
	}()

	// Test failed request with invalid URL
	testCase2 := func() {
		operation := "GET"
		hostAddr := "invalid"
		command := "/path"
		params := bytes.NewBufferString(`{"key": "value"}`)
		expectedStatusCode := http.StatusBadRequest
		expectedBody := `{"error": "invalid URL"}`
		expectedError := "Failed to create HTTP request."

		statusCode, body, err := httpClientRequest(operation, hostAddr, command, params)
		if statusCode!= expectedStatusCode {
			t.Errorf("Expected status code %d, got %d", expectedStatusCode, statusCode)
		}
		if string(body)!= expectedBody {
			t.Errorf("Expected body %q, got %q", expectedBody, body)
		}
		if err == nil ||!strings.HasPrefix(err.Error(), expectedError) {
			t.Errorf("Expected error %q, got %v", expectedError, err)
		}
	}()

	// Test failed request with invalid JSON payload
	testCase3 := func() {
		operation := "POST"
		hostAddr := "example.com"
		command := "/path"
		params := bytes.NewBufferString(`{"key": "value"})
		expectedStatusCode := http.StatusBadRequest
		expectedBody := `{"error": "invalid JSON payload"}`
		expectedError := "Failed to parse JSON payload."

		statusCode, body, err := httpClientRequest(operation, hostAddr, command, params)
		if statusCode!= expectedStatusCode {
			t.Errorf("Expected status code %d, got %d", expectedStatusCode, statusCode)
		}
		if string(body)!= expectedBody {
			t.Errorf("Expected body %q, got %q", expectedBody, body)
		}
		if err == nil ||!strings.HasPrefix(err.Error(), expectedError) {
			t.Errorf("Expected error %q, got %v", expectedError, err)
		}
	}()
}
