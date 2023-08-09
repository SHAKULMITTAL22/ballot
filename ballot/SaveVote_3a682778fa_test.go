package test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type Vote struct {
	CandidateID int
	Option      string
}

func TestSaveVote_Success(t *testing.T) {
	// Arrange
	candidateVotesStore := make(map[int]int)
	vote := Vote{
		CandidateID: 1,
		Option:     "option1",
	}

	// Act
	err := saveVote(vote, candidateVotesStore)

	// Assert
	require.NoError(t, err)
	require.Equal(t, 2, candidateVotesStore[1])
}

func TestSaveVote_Failure(t *testing.T) {
	// Arrange
	candidateVotesStore := make(map[int]int)
	vote := Vote{
		CandidateID: 2,
		Option:     "option1",
	}

	// Act
	err := saveVote(vote, candidateVotesStore)

	// Assert
	require.Error(t, err)
	require.Equal(t, 1, candidateVotesStore[2])
}

func saveVote(vote Vote, candidateVotesStore map[int]int) error {
    // Your implementation here
    return nil
}
