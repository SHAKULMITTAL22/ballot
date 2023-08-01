package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountVote(t *testing.T) {
	type args struct {
		votes map[string]int
	}
	tests := []struct {
		name string
		args args
		want ResultBoard
	}{
		{"empty vote", {}, ResultBoard{}},
		{"single candidate vote", {"candidate1": 50}, ResultBoard{{"candidate1", 50}}},
		{"multi candidate vote", {"candidate1": 30, "candidate2": 40}, ResultBoard{{"candidate1", 30}, {"candidate2", 40}}},
		{"vote with invalid candidate id", {"invalid": 10}, Error("Invalid candidate ID")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := countVote(tt.args.votes)
			if!assert.NoError(t, err) {
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestCountVote_Sort(t *testing.T) {
	type args struct {
		votes map[string]int
	}
	tests := []struct {
		name string
		args args
		want ResultBoard
	}{
		{"sorted results", {"candidate1": 50, "candidate2": 40}, ResultBoard{{"candidate1", 50}, {"candidate2", 40}}},
		{"unsorted results", {"candidate2": 40, "candidate1": 50}, ResultBoard{{"candidate2", 40}, {"candidate1", 50}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := countVote(tt.args.votes)
			if!assert.NoError(t, err) {
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
