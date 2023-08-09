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

type Vote struct {
    VoterID string `json:"voter_id"`
    CandidateID string `json:"candidate_id"`
}

type Status struct {
    Code int        `json:"code"`
    Message string `json:"message"`
}

func TestServeRoot(t *testing.T) {
    // Test case 1: Successful GET request
    req, err := http.NewRequest(http.MethodGet, "/", bytes.NewBufferString("{}"))
    if err!= nil {
        t.Error(err)
        return
    }
    resp := &http.Response{
        Header: http.Header{"Access-Control-Allow-Origin": {"*"}},
    }
    serveRoot(resp, req)
    if resp.StatusCode!= http.StatusOK {
        t.Errorf("Expected HTTP status code 200, got %d", resp.StatusCode)
    }

    // Test case 2: Invalid POST request
    req, err := http.NewRequest(http.MethodPost, "/", bytes.NewBufferString("{\"candidate_id\": \"\"}"))
    if err!= nil {
        t.Error(err)
        return
    }
    resp := &http.Response{
        Header: http.Header{"Access-Control-Allow-Origin": {"*"}},
    }
    serveRoot(resp, req)
    if resp.StatusCode!= http.StatusBadRequest {
        t.Errorf("Expected HTTP status code 400, got %d", resp.StatusCode)
    }
}
