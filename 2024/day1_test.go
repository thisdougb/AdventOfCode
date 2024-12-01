package main

import "testing"

func TestDay1Part1(t *testing.T) {

	testCases := []struct {
		list1        []int
		list2        []int
		expectResult int
	}{
		{
			list1:        []int{3, 4, 2, 1, 3, 3},
			list2:        []int{4, 3, 5, 3, 9, 3},
			expectResult: 11,
		},
	}

	for _, tc := range testCases {
		result := day1Part1(tc.list1, tc.list2)
		if result != tc.expectResult {
			t.Errorf("expected %v, got result %v", tc.expectResult, result)
		}
	}
}

func TestDay1Part2(t *testing.T) {

	testCases := []struct {
		list1        []int
		list2        []int
		expectResult int
	}{
		{
			list1:        []int{3, 4, 2, 1, 3, 3},
			list2:        []int{4, 3, 5, 3, 9, 3},
			expectResult: 31,
		},
	}

	for _, tc := range testCases {
		result := day1Part2(tc.list1, tc.list2)
		if result != tc.expectResult {
			t.Errorf("expected %v, got result %v", tc.expectResult, result)
		}
	}
}
