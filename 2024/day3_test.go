package main

import (
	"testing"
)

func TestDay3Part1(t *testing.T) {

	testCases := []struct {
		data         []string
		expectResult int32
	}{
		{
			data:         []string{"xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"},
			expectResult: 161,
		},
	}

	for i, tc := range testCases {
		if result := day3Part1(tc.data); result != tc.expectResult {
			t.Errorf("case %d, got result %v", i, result)
		}
	}
}

func TestDay3Part2(t *testing.T) {

	testCases := []struct {
		data         []string
		expectResult int32
	}{
		{
			data:         []string{"xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"},
			expectResult: 48,
		},
	}

	for i, tc := range testCases {
		if result := day3Part2(tc.data); result != tc.expectResult {
			t.Errorf("case %d, got result %v", i, result)
		}
	}
}
