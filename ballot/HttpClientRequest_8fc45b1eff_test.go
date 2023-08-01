package test

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"strings"
	"sync"
)

func TestHTTPClientRequest(t *testing.T) {
	testCases := map[string]struct{}{
		"Successful GET request": {
			operation: "GET",
			hostAddr: "example.com",
			command: "/path",
			params: bytes.NewBufferString(`{"key":"value"}`),
			expectedStatusCode: http.StatusOK,
			expectedBody:       []byte(`{"result": "success"}),
		},
		"Unsuccessful POST request": {
			operation: "POST",
			hostAddr: "example.com",
			command: "/path",
			params: bytes.NewBufferString(`{"key":"value"}`),
			expectedStatusCode: http.StatusCreated,
			expectedBody:       []byte(`{"result": "success"}),
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			// Set up mock server
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				switch r.Method {
				case "GET":
					w.WriteHeader(testCase.expectedStatusCode)
					w.Write(testCase.expectedBody)
				case "POST":
					w.WriteHeader(testCase.expectedStatusCode)
					w.Write(testCase.expectedBody)
				}
			}))
			defer server.Close()

			// Call httpClientRequest with test parameters
			statusCode, body, err := httpClientRequest(testCase.operation, testCase.hostAddr, testCase.command, testCase.params)
			if statusCode!= testCase.expectedStatusCode {
				t.Errorf("Expected status code %d, got %d", testCase.expectedStatusCode, statusCode)
			}
			if!reflect.DeepEqual(body, testCase.expectedBody) {
				t.Errorf("Expected body %v, got %v", testCase.expectedBody, body)
			}
			if err!= testCase.expectedError {
				t.Errorf("Expected error %v, got %v", testCase.expectedError, err)
			}
		})
	}
}

func TestHTTPClientRequest_errorHandling(t *testing.T) {
	testCases := map[string]struct{}{
		"Invalid URL": {"example.com"},
		"Nil request": {nil},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			_, _, err := httpClientRequest(testCase.operation, testCase.hostAddr, testCase.command, testCase.params)
			if err == nil {
				t.Error("Expected error, but none was returned.")
			}
		})
	}
}
