package main

import (
	"testing"

	"golang.org/x/exp/slices"
)

func TestGetType2(t *testing.T) {

	testCases := []struct {
		hand   string
		result int
	}{
		{"AAAAA", 7},
		{"AAAAJ", 7},
	}

	for _, tc := range testCases {

		result := getType2(tc.hand)
		if tc.result != result {
			t.Errorf("failed: %s type %d", tc.hand, result)
		}
	}

}
func TestFirstBeatsSecondString2(t *testing.T) {

	testCases := []struct {
		first  string
		second string
		result bool
	}{
		// five of a kind
		{"33332", "2AAAA", true},
		{"77888", "77788", true},
		{"KK677", "KTJJT", true},
		{"KTJJT", "KK677", false},
	}

	for _, tc := range testCases {

		result := firstBeatsSecond2(tc.first, tc.second)
		if tc.result != result {
			t.Errorf("failed: %s beats %s", tc.first, tc.second)
		}
	}

}

func TestFirstBeatsSecondType2(t *testing.T) {

	testCases := []struct {
		first  string
		second string
		result bool
	}{
		// five of a kind
		{"AAAAA", "AAAAB", true},
		{"AAAAB", "AAAAA", false},

		// four of a kind
		{"AAAAB", "AAABB", true},
		{"AAABB", "AAAAB", false},

		// fullhouse
		{"AAABB", "AABBC", true},
		{"AABBC", "AAABB", false},

		// two pair
		{"AABBC", "AABCD", true},
		{"AABCD", "AABBC", false},

		// one pair
		{"AABCD", "ABCDE", true},
		{"ABCDE", "AABCD", false},
	}

	for _, tc := range testCases {

		result := firstBeatsSecond2(tc.first, tc.second)
		if tc.result != result {
			t.Errorf("failed: %s beats %s", tc.first, tc.second)
		}
	}

}

func TestFirstBeatsSecondType(t *testing.T) {

	testCases := []struct {
		first  string
		second string
		result bool
	}{
		// five of a kind
		{"AAAAA", "AAAAB", true},
		{"AAAAB", "AAAAA", false},

		// four of a kind
		{"AAAAB", "AAABB", true},
		{"AAABB", "AAAAB", false},

		// fullhouse
		{"AAABB", "AABBC", true},
		{"AABBC", "AAABB", false},

		// two pair
		{"AABBC", "AABCD", true},
		{"AABCD", "AABBC", false},

		// one pair
		{"AABCD", "ABCDE", true},
		{"ABCDE", "AABCD", false},
	}

	for _, tc := range testCases {

		result := firstBeatsSecond(tc.first, tc.second)
		if tc.result != result {
			t.Errorf("failed: %s beats %s", tc.first, tc.second)
		}
	}

}

func TestFirstBeatsSecondString(t *testing.T) {

	testCases := []struct {
		first  string
		second string
		result bool
	}{
		// five of a kind
		{"33332", "2AAAA", true},
		{"77888", "77788", true},
		{"KK677", "KTJJT", true},
		{"KTJJT", "KK677", false},
	}

	for _, tc := range testCases {

		result := firstBeatsSecond(tc.first, tc.second)
		if tc.result != result {
			t.Errorf("failed: %s beats %s", tc.first, tc.second)
		}
	}

}

func TestAddHandToHands(t *testing.T) {

	testCases := []struct {
		hand   string
		hands  []string
		result []string
	}{
		// five of a kind
		{"66666", []string{}, []string{"66666"}},
		{"55555", []string{"66666"}, []string{"55555", "66666"}},
		{"77777", []string{"66666"}, []string{"66666", "77777"}},
		{"99999", []string{"66666", "77777"}, []string{"66666", "77777", "99999"}},
		{"88888", []string{"66666", "77777", "99999"}, []string{"66666", "77777", "88888", "99999"}},
		{"32T3K", []string{}, []string{"32T3K"}},
		{"T55J5", []string{"32T3K"}, []string{"32T3K", "T55J5"}},
		{"KK677", []string{"32T3K", "T55J5"}, []string{"32T3K", "KK677", "T55J5"}},
		{"KTJJT", []string{"32T3K", "KK677", "T55J5"}, []string{"32T3K", "KTJJT", "KK677", "T55J5"}},
	}

	for _, tc := range testCases {

		result := addHandToHands(tc.hand, tc.hands)
		if !slices.Equal(tc.result, result) {
			t.Errorf("failed: %s added to %s gave %s", tc.hand, tc.hands, result)
		}
	}

}
