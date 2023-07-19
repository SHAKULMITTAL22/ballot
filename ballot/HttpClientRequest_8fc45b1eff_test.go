package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHttpClientRequest_8fc45b1eff(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte(`OK`))
	}))
	defer server.Close()

	tests := []struct {
		name       string
		operation  string
		hostAddr   string
		command    string
		params     io.Reader
		wantStatus int
		wantBody   []byte
		wantErr    error
	}{
		{
			name:       "Test case 1: Valid host address and operation",
			operation:  "GET",
			hostAddr:   server.URL,
			command:    "/",
			params:     nil,
			wantStatus: 200,
			wantBody:   []byte("OK"),
			wantErr:    nil,
		},
		{
			name:       "Test case 2: Invalid host address",
			operation:  "GET",
			hostAddr:   "invalid",
			command:    "/",
			params:     nil,
			wantStatus: 400,
			wantBody:   nil,
			wantErr:    errors.New("Failed to create HTTP request.Get \"http://invalid/\": dial tcp: lookup invalid: no such host"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotStatus, gotBody, gotErr := httpClientRequest(tt.operation, tt.hostAddr, tt.command, tt.params)
			if gotStatus != tt.wantStatus {
				t.Errorf("httpClientRequest() gotStatus = %v, want %v", gotStatus, tt.wantStatus)
			}
			if !bytes.Equal(gotBody, tt.wantBody) {
				t.Errorf("httpClientRequest() gotBody = %v, want %v", string(gotBody), string(tt.wantBody))
			}
			if (gotErr != nil && tt.wantErr == nil) || (gotErr == nil && tt.wantErr != nil) || (gotErr != nil && tt.wantErr != nil && gotErr.Error() != tt.wantErr.Error()) {
				t.Errorf("httpClientRequest() gotErr = %v, want %v", gotErr, tt.wantErr)
			}
		})
	}
}

func httpClientRequest(operation, hostAddr, command string, params io.Reader) (int, []byte, error) {

	url := "http://" + hostAddr + command
	if strings.Contains(hostAddr, "http://") {
		url = hostAddr + command
	}

	req, err := http.NewRequest(operation, url, params)
	if err != nil {
		return http.StatusBadRequest, nil, errors.New("Failed to create HTTP request." + err.Error())
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	httpClient := http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}
	defer resp.Body.Close()

	body, ioErr := ioutil.ReadAll(resp.Body)
	if hBit := resp.StatusCode / 100; hBit != 2 && hBit != 3 {
		if ioErr != nil {
			ioErr = fmt.Errorf("status code error %d", resp.StatusCode)
		}
	}
	return resp.StatusCode, body, ioErr
}
