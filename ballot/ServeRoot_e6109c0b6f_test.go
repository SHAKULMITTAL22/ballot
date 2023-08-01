package test_serve_root

import (
    "testing"

    "github.com/stretchr/testify/assert"
)

func TestServeRoot(t *testing.T) {
    // Test case 1: Successful GET request
    req, err := http.NewRequest(http.MethodGet, "/", bytes.NewBufferString(""))
    assert.NoError(t, err)
    resp := &http.Response{
        Header: make(http.Header),
    }
    serveRoot(resp, req)
    assert.Equal(t, resp.StatusCode, http.StatusOK)
    assert.Equal(t, resp.Header.Get("Content-Type"), "application/json")
    assert.Equal(t, resp.Header.Get("Access-Control-Allow-Origin"), "*")

    // Test case 2: Invalid POST request
    req, err = http.NewRequest(http.MethodPost, "/", bytes.NewBufferString(""))
    assert.NoError(t, err)
    resp = &http.Response{
        Header: make(http.Header),
    }
    serveRoot(resp, req)
    assert.Equal(t, resp.StatusCode, http.StatusBadRequest)
    assert.Equal(t, resp.Header.Get("Content-Type"), "application/json")
    assert.Equal(t, resp.Header.Get("Access-Control-Allow-Origin"), *)
}
