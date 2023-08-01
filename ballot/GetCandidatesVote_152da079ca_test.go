package mytest

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetCandidatesVote_Success(t *testing.T) {
	// Arrange
	expected := map[string]int{"candidate1": 50, "candidate2": 30, "candidate3": 20}
	mockStore := &MockCandidateVotesStore{
		Map: expected,
	}
	getCandidatesVoteFunc := func() map[string]int { return mockStore.Map }

	// Act
	actual := getCandidatesVoteFunc()

	// Assert
	require.Equal(t, expected, actual)
}

func TestGetCandidatesVote_EdgeCase(t *testing.T) {
	// Arrange
	expected := map[string]int{}
	mockStore := &MockCandidateVotesStore{
		Map: expected,
	}
	getCandidatesVoteFunc := func() map[string]int { return mockStore.Map }

	// Act
	actual := getCandidatesVoteFunc()

	// Assert
	require.Equal(t, expected, actual)
}

func TestGetCandidatesVote_Failure(t *testing.T) {
	// Arrange
	mockStore := &MockCandidateVotesStore{
		Err: errors.New("failed to retrieve candidates votes"),
	}
	getCandidatesVoteFunc := func() map[string]int { return mockStore.Map }

	// Act
	actual := getCandidatesVoteFunc()

	// Assert
	require.NotNil(t, actual)
	require.True(t, actual == nil || len(actual) == 0)
	require.Equal(t, mockStore.Err, actual)
}

type MockCandidateVotesStore struct {
	Map     map[string]int
	Err     error
	mutex   sync.Mutex
	called  bool
}{
	func (m *MockCandidateVotesStore) Get(key string) int {
		m.mutex.Lock()
		defer m.mutex.Unlock()
		if m.called {
			return 0
		}
		m.called = true
		return m.Map[key]
	},
	func (m *MockCandidateVotesStore) Set(key string, vote int) {
		m.mutex.Lock()
		defer m.mutex.Unlock()
		m.Map[key] = vote
	},
}
