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

type Status struct {
	Message string `json:"message"`
}

func TestWriteVoterResponse(t *testing.T) {
	// Success scenario
	status := Status{
		Message: "Your vote has been recorded.",
	}
	writer := &mockWriter{}
	writeVoterResponse(writer, status)
	expected := `[{"message":"Your vote has been recorded."}]`
	actual := writer.written
	if actual!= expected {
		t.Errorf("Expected %q written, got %q", expected, actual)
	}

	// Failure scenario
	status = Status{
		Message: "",
	}
	writer = &mockWriter{}
	writeVoterResponse(writer, status)
	expected = `[{"message":null}]`
	actual = writer.written
	if actual!= expected {
		t.Errorf("Expected %q written, got %q", expected, actual)
	}
}

type mockWriter struct {
	written []byte
}

func (m *mockWriter) Write(p []byte) (int, error) {
	m.written = append(m.written, p...)
	return len(p), nil
}
