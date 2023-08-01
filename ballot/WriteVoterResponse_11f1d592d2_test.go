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
	Code    int    `json:"code"`
}

func TestWriteVoterResponse(t *testing.T) {
	// Success scenario: valid status object
	status := Status{
		Message: "Successful voter response",
		Code:    200,
	}
	writer := &bytes.Buffer{}
	writeVoterResponse(writer, status)
	response := writer.String()
	expected := `[{"message":"Successful voter response","code":200}]`
	if response!= expected {
		t.Errorf("Invalid response. Got %v, Expected %v", response, expected)
	}

	// Failure scenario: invalid status object
	status = Status{
		Message: "",
		Code:    -1,
	}
	writer = &bytes.Buffer{}
	writeVoterResponse(writer, status)
	response = writer.String()
	expected = `[{"message":"Failed to generate voter response. Code:-1"},]`
	if response!= expected {
		t.Errorf("Invalid response. Got %v, Expected %v", response, expected)
	}
}

func writeVoterResponse(w io.Writer, s Status) {
	enc := json.NewEncoder(w)
	err := enc.Encode(s)
	if err!= nil {
		panic(err)
	}
}
