package tests

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

type ResultBoard struct {
    TotalVotes int `json:"total_votes"`
    Votes      []struct {
        CandidateID int `json:"candidate_id"`
        VoterID     int `json:"voter_id"`
    } `json:"votes"`
}

type Vote struct {
    CandidateID int `json:"candidate_id"`
    VoterID     int `json:"voter_id"`
}

type Status struct {
    Code int `json:"code"`
    Msg  string `json:"message"`
}

func TestBallot(t *testing.T) {
    // Test case 1: Successful GET request
    testGetBallotCountSuccessfully(t)

    // Test case 2: Successful POST request
    testPostBallotSuccessfully(t)
}

func testGetBallotCountSuccessfully(t *testing.T) {
    // Arrange
    port := rand.Intn(10) + 8080
    url := net.JoinHostPort("", port)
    expectedResponse := `{
        "total_votes": 5
    }`

    // Act
    resp, err := httpClientRequest(http.MethodGet, url, "/", nil)
    if err!= nil {
        t.Errorf("Failed to make GET request: %v", err)
        return
    }

    // Assert
    actualResponse := string(resp)
    if actualResponse!= expectedResponse {
        t.Errorf("Expected response: %#v\nActual response: %#v", expectedResponse, actualResponse)
    }
}

func testPostBallotSuccessfully(t *testing.T) {
    // Arrange
    port := rand.Intn(10) + 8080
    url := net.JoinHostPort("", port)
    voteReq := Vote{
        CandidateID: rand.Intn(10),
        VoterID:     rand.Intn(10),
    }
    reqBuf, err := json.Marshal(voteReq)
    if err!= nil {
        t.Errorf("Failed to marshal vote request: %v", err)
        return
    }

    // Act
    resp, err := httpClientRequest(http.MethodPost, url, "/", bytes.NewReader(reqBuf))
    if err!= nil {
        t.Errorf("Failed to make POST request: %v", err)
        return
    }

    // Assert
    actualResponse := string(resp)
    expectedResponse := `{
        "status": "success"
    }`
    if actualResponse!= expectedResponse {
        t.Errorf("Expected response: %#v\nActual response: %#v", expectedResponse, actualResponse)
    }
}

func httpClientRequest(method, url string, path string, body interface{}) (*http.Response, error) {
    client := &http.Client{}
    req, err := http.NewRequest(method, url+path, body)
    if err!= nil {
        return nil, err
    }
    res, err := client.Do(req)
    if err!= nil {
        return nil, err
    }
    defer res.Body.Close()
    return res, nil
}
