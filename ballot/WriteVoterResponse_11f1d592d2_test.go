// Test generated by RoostGPT for test roost-test using AI Type Open Source AI and AI Model meta-llama/Llama-2-13b-chat

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
	"testing"
)

type Status struct {
	Message string `json:"message"`
}

func TestWriteVoterResponse(t *testing.T) {
	// Successful response
	status := Status{
		Message: "Your vote has been recorded.",
	}
	w := &mockResponseWriter{}
	writeVoterResponse(w, status)
	response := w.GetMockResponse()
	expected := `[{"message":"Your vote has been recorded."}]`
	if response!= expected {
		t.Errorf("Expected %q, got %q", expected, response)
	}

	// Error response
	status = Status{
		Message: "An error occurred while recording your vote.",
	}
	err := errors.New(" voting system unavailable.")
	w = &mockResponseWriter{}
	writeVoterResponse(w, status)
	response = w.GetMockResponse()
	expected = `[{"message":"An error occurred while recording your vote.","details":["voting system unavailable"]}]`
	if response!= expected {
		t.Errorf("Expected %q, got %q", expected, response)
	}
}

type mockResponseWriter struct {
	responses []string
}

func (m *mockResponseWriter) Write(p []byte) (int, error) {
	m.responses = append(m.responses, string(p))
	return len(p), nil
}

func (m *mockResponseWriter) GetMockResponse() string {
	return strings.Join(m.responses, "")
}
