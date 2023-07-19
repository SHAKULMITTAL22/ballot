package main

import (
	"testing"
)

type Vote struct {
	CandidateID string
}

var candidateVotesStore = make(map[string]int)

func saveVote(vote Vote) error {
	candidateVotesStore[vote.CandidateID] += 2
	return nil
}

func getCandidatesVote() map[string]int {
	return candidateVotesStore
}

func TestSaveVote_3a682778fa(t *testing.T) {
	candidateID := "candidate1"
	vote := Vote{CandidateID: candidateID}

	// Test case 1: Successful scenario
	err := saveVote(vote)
	if err != nil {
		t.Error("Expected no error, got ", err)
	}
	if candidateVotesStore[candidateID] != 2 {
		t.Error("Expected 2 votes, got ", candidateVotesStore[candidateID])
	}

	// Test case 2: Failure scenario
	delete(candidateVotesStore, candidateID)
	err = saveVote(vote)
	if err != nil {
		t.Error("Expected no error, got ", err)
	}
	if candidateVotesStore[candidateID] != 2 {
		t.Error("Expected 2 votes, got ", candidateVotesStore[candidateID])
	}
}
