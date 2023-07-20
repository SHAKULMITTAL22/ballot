// Test generated by RoostGPT for test roost-test using AI Type Vertex AI and AI Model code-bison

package main

import (
  "bytes"
  "errors"
  "fmt"
  "io/ioutil"
  "net/http"
  "strings"
  "testing"
)

func TestHttpClientRequest(t *testing.T) {
  expectedStatus := http.StatusOK

  // Arrange
  operation := "GET"
  hostAddr := "example.com"
  command := "/api/v1/users"
  jsonParams := `{"name": "John"}`
  var buf bytes.Buffer
  _, _ = buf.WriteString(jsonParams)
  reqBody := &buf

  actualStatus, resBody, err := httpClientRequest(operation, hostAddr, command, reqBody)

  if err!= nil || actualStatus!= expectedStatus {
    t.Fatalf("Expected status code '%d', got '%d' with error: %s\nResponse Body:\n%s",
      expectedStatus, actualStatus, err, resBody)
  }
}
