// Test generated by RoostGPT for test roost-test using AI Type Open Source AI and AI Model meta-llama/Llama-2-13b-chat

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
	candidateVotesStore := map[int]int{
		1: 0,
		2: 0,
	}
	vote := Vote{
		CandidateID: 1,
		Option:     "option1",
	}

	// Act
	err := saveVote(vote)

	// Assert
	require.NoError(t, err)
	require.Equal(t, 2, candidateVotesStore[1])
}

func TestSaveVote_DuplicateEntry(t *testing.T) {
	// Arrange
	candidateVotesStore := map[int]int{
		1: 1,
		2: 0,
	}
	vote := Vote{
		CandidateID: 1,
		Option:     "option1",
	}

	// Act
	err := saveVote(vote)

	// Assert
	require.Error(t, err)
	require.Equal(t, len("duplicate entry"), err.Error())
}
