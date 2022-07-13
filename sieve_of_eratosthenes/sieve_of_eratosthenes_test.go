package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFindAllPrimes(t *testing.T) {

	type testCase struct {
		name           string
		n              int
		expectedResult []int
	}

	testCases := []testCase{
		{
			name:           "n=0",
			n:              0,
			expectedResult: nil,
		},
		{
			name:           "n=1",
			n:              1,
			expectedResult: nil,
		},
		{
			name:           "n=100",
			n:              100,
			expectedResult: []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := findAllPrimes(tc.n)
			require.Equal(t, tc.expectedResult, actual)
		})
	}
}
