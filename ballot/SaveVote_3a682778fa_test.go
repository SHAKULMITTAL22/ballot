package vote_test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type Vote struct {
	CandidateID int
	Option      string
}

func getCandidatesVote() map[int]int {
	return map[int]int{
		1: 0,
		2: 0,
		3: 0,
	}
}

func saveVote(vote *Vote) error {
	candidateVotesStore := getCandidatesVote()
	candidateVotesStore[vote.CandidateID]++
	return nil
}

func TestSaveVoteSuccess(t *testing.T) {
	vote := &Vote{
		CandidateID: 1,
		Option:      "option1",
	}
	err := saveVote(vote)
	require.NoError(t, err)
	require.Equal(t, 1, candidateVotesStore[1])
}

func TestSaveVoteFailure(t *testing.T) {
	vote := &Vote{
		CandidateID: 42,
		Option:      "option1",
	}
	err := saveVote(vote)
	require.Error(t, err)
	require.Equal(t, 0, candidateVotesStore[42])
}
