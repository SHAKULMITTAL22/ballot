package main

import (
	"sync"
	"testing"
)

var once sync.Once
var candidateVotesStore map[string]int

func getCandidatesVote() map[string]int {
	once.Do(func() {
		candidateVotesStore = make(map[string]int)
	})
	return candidateVotesStore
}

func TestGetCandidatesVote_152da079ca(t *testing.T) {
	// Test case 1: Check if the function returns a map
	result := getCandidatesVote()
	if result == nil {
		t.Error("Expected a map but got nil")
	}

	// Test case 2: Check if the function returns an empty map on first call
	if len(result) != 0 {
		t.Errorf("Expected an empty map but got a map with length %d", len(result))
	}

	// Test case 3: Check if the function returns the same map on subsequent calls
	candidateVotesStore["John Doe"] = 100
	result = getCandidatesVote()
	if len(result) != 1 {
		t.Errorf("Expected a map with length 1 but got a map with length %d", len(result))
	}
	if result["John Doe"] != 100 {
		t.Errorf("Expected 100 votes for John Doe but got %d", result["John Doe"])
	}
}
