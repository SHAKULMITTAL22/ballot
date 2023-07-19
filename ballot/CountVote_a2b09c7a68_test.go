package main

import (
	"testing"
	"sort"
)

type CandidateVotes struct {
	CandidateID string
	Votes       int
}

type ResultBoard struct {
	Results    []CandidateVotes
	TotalVotes int
}

func getCandidatesVote() map[string]int {
	return map[string]int{
		"John":  5,
		"Alice": 7,
		"Bob":   10,
	}
}

func countVote() (res ResultBoard, err error) {
	votes := getCandidatesVote()
	for candidateID, voteCount := range votes {
		res.Results = append(res.Results, CandidateVotes{candidateID, voteCount})
		res.TotalVotes += voteCount
	}

	sort.Slice(res.Results, func(i, j int) bool {
		return res.Results[i].Votes > res.Results[j].Votes
	})
	return res, err
}

func TestCountVote_a2b09c7a68(t *testing.T) {
	res, err := countVote()
	if err != nil {
		t.Errorf("countVote() error = %v", err)
		return
	}

	if res.TotalVotes != 22 {
		t.Errorf("countVote() TotalVotes = %v, want %v", res.TotalVotes, 22)
	}

	if len(res.Results) != 3 {
		t.Errorf("countVote() len(Results) = %v, want %v", len(res.Results), 3)
	}

	if res.Results[0].CandidateID != "Bob" || res.Results[0].Votes != 10 {
		t.Errorf("countVote() Results[0] = %v, want %v", res.Results[0], CandidateVotes{"Bob", 10})
	}
}
