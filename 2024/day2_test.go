package main

import (
	"testing"

	"golang.org/x/exp/slices"
)

func TestDay2Part1(t *testing.T) {

	testCases := []struct {
		data         []int
		expectResult bool
	}{
		{
			data:         []int{7, 6, 4, 2, 1},
			expectResult: true,
		},
		{
			data:         []int{1, 2, 7, 8, 9},
			expectResult: false,
		},
		{
			data:         []int{9, 7, 6, 2, 1},
			expectResult: false,
		},
		{
			data:         []int{1, 3, 2, 4, 5},
			expectResult: false,
		},
		{
			data:         []int{8, 6, 4, 4, 1},
			expectResult: false,
		},
		{
			data:         []int{1, 3, 6, 7, 9},
			expectResult: true,
		},
	}

	for i, tc := range testCases {
		if reportIsSafe(tc.data) != tc.expectResult {
			t.Errorf("case %d, got result %v", i, tc.expectResult)
		}
	}
}

func TestDay2Part2WithDamper(t *testing.T) {

	testCases := []struct {
		data         []int
		expectResult bool
	}{
		{
			data:         []int{7, 6, 4, 2, 1},
			expectResult: true,
		},
		{
			data:         []int{1, 2, 7, 8, 9},
			expectResult: false,
		},
		{
			data:         []int{9, 7, 6, 2, 1},
			expectResult: false,
		},
		{
			data:         []int{1, 3, 2, 4, 5},
			expectResult: true,
		},
		{
			data:         []int{8, 6, 4, 4, 1},
			expectResult: true,
		},
		{
			data:         []int{1, 3, 6, 7, 9},
			expectResult: true,
		},
		// from input data, fixing bug which missed removing last int then
		// checking again with damper
		{
			data:         []int{48, 51, 54, 56, 60},
			expectResult: true,
		},
	}

	for i, tc := range testCases {
		if reportIsSafeWithDamper(tc.data) != tc.expectResult {
			t.Errorf("expected safe %v, (%d) %v", tc.expectResult, i, tc.data)
		}
	}
}

func TestDay2Part2(t *testing.T) {

	data := [][]int{
		[]int{7, 6, 4, 2, 1},
		[]int{1, 2, 7, 8, 9},
		[]int{9, 7, 6, 2, 1},
		[]int{1, 3, 2, 4, 5},
		[]int{8, 6, 4, 4, 1},
		[]int{1, 3, 6, 7, 9},
	}

	result := day2Part2(data)

	if 4 != result {
		t.Errorf("expected 4, got result %v", result)
	}

}

func TestRemoveIndex(t *testing.T) {

	testCases := []struct {
		input  []int
		idx    int
		result []int
	}{
		{
			input:  []int{1, 2, 3, 4, 5},
			idx:    0,
			result: []int{2, 3, 4, 5},
		},
		{
			input:  []int{1, 2, 3, 4, 5},
			idx:    4,
			result: []int{1, 2, 3, 4},
		},
		{
			input:  []int{1, 2, 3, 4, 5},
			idx:    2,
			result: []int{1, 2, 4, 5},
		},
	}

	for _, tc := range testCases {

		result := removeIndex(tc.input, tc.idx)

		if !slices.Equal(result, tc.result) {
			t.Errorf("not different %v", tc.input)
		}
	}
}
