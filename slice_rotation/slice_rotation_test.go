package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRotateRight(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5, 6}

	type testCase struct {
		name           string
		step           int
		expectedResult []int
	}

	testCases := []testCase{
		{
			name:           "Step 1, right notation",
			step:           1,
			expectedResult: []int{6, 1, 2, 3, 4, 5},
		},
		{
			name:           "Step 6, right notation",
			step:           6,
			expectedResult: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:           "Step 0, right notation",
			step:           0,
			expectedResult: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:           "Step 4, right notation",
			step:           4,
			expectedResult: []int{3, 4, 5, 6, 1, 2},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := rotateRight(slice, tc.step)
			require.Equal(t, tc.expectedResult, actual)
		})
	}
}

func TestRotateLeft(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5, 6}

	type testCase struct {
		name           string
		step           int
		expectedResult []int
	}

	testCases := []testCase{
		{
			name:           "Step 1, left notation",
			step:           1,
			expectedResult: []int{2, 3, 4, 5, 6, 1},
		},
		{
			name:           "Step 6, left notation",
			step:           6,
			expectedResult: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:           "Step 0, left notation",
			step:           0,
			expectedResult: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:           "Step 4, left notation",
			step:           4,
			expectedResult: []int{5, 6, 1, 2, 3, 4},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := rotateLeft(slice, tc.step)
			require.Equal(t, tc.expectedResult, actual)
		})
	}
}
