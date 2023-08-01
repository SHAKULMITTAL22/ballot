package test

import (
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/roost/RoostGPT/ballot-go-test/bbcd45fe-4d51-4d38-9a3b-da324aec6c73/ballot"
)

func TestRunTest(t *testing.T) {
	type status struct {
		Message string `json:"message"`
		Code   int    `json:"code"`
	}

	var (
		w  http.ResponseWriter
		r  *http.Request
		status status
		err error
	)

	// Test successful execution
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	log.Println("ballot endpoint tests running")
	status.Message = "Test Cases passed"
	status.Code = http.StatusOK
	writeVoterResponse(w, status)

	// Test failed execution with error
	err = errors.New("test error")
	status.Message = fmt.Sprintf("Test Cases Failed with error : %v", err)
	status.Code = http.StatusBadRequest
	writeVoterResponse(w, status)
}

func writeVoterResponse(w http.ResponseWriter, status status) {
	resp := map[string]interface{}{
		"message": status.Message,
		"code":    status.Code,
	}
	json.NewEncoder(w).Encode(resp)
}
