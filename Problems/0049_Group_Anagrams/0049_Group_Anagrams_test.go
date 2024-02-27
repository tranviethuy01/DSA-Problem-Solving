package main

import (
	"reflect"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	// Define test cases
	testCases := []struct {
		input  []string
		output [][]string
	}{
		{
			input:  []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			output: [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}},
		},
		{
			input:  []string{""},
			output: [][]string{{""}},
		},
		{
			input:  []string{"a"},
			output: [][]string{{"a"}},
		},
	}

	// Iterate through test cases
	for _, tc := range testCases {
		// Call the function
		result := groupAnagrams(tc.input)
		// Check if the result matches the expected output
		if !reflect.DeepEqual(result, tc.output) {
			t.Errorf("Input: %v, Expected: %v, Got: %v", tc.input, tc.output, result)
		}
	}
}
