package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net"
	"net/http"
	"sort"
	"strings"
	"sync"
)

func TestGetCandidatesVote(t *testing.T) {
	testCases := []struct {
		name     string
		expected map[string]int
	}{
		{
			name: "empty store",
			expected: make(map[string]int),
		},
		{
			name: "single vote",
			expected: map["candidate1": 1],
		},
		{
			name: "multiple votes",
			expected: map{"candidate1": 1, "candidate2": 2},
		},
		{
			name: "non existent candidate",
			expected: nil,
		},
	}

	for _, tc := range testCases {
		tc.name := fmt.Sprintf("Test%s", tc.name)
		defer func() {
			if r := recover(); r!= nil {
				t.Errorf("Panic %v\n", r)
			}
		}()
		actual := getCandidatesVote()
		if!reflect.DeepEqual(actual, tc.expected) {
			t.Errorf("%s: expected %#v, got %#v\n", tc.name, tc.expected, actual)
		}
	}
}

func TestGetCandidatesVote_OnceDo(t *testing.T) {
	var once sync.Once
	store := make(map[string]int)
	getCandidatesVote() <- struct{}{}
	once.Do(func() {
		store["candidate1"] = 1
	})
	actual := getCandidatesVote()
	if len(actual)!= 1 || actual["candidate1"]!= 1 {
		t.Errorf("Expected candidate1 to have vote 1 but got %#v\n", actual)
	}
}
