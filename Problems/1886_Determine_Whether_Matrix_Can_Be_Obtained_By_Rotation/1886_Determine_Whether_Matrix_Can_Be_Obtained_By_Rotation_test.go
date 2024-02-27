package main

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestFindRotationEqual(t *testing.T) {
	testCases := []struct {
		mat      [][]int
		target   [][]int
		expected bool
	}{
		{[][]int{{0, 1}, {1, 0}}, [][]int{{1, 0}, {0, 1}}, true},
		{[][]int{{0, 1}, {1, 1}}, [][]int{{1, 0}, {0, 1}}, false},
		{[][]int{{0, 0, 0}, {0, 1, 0}, {1, 1, 1}}, [][]int{{1, 1, 1}, {0, 1, 0}, {0, 0, 0}}, true},
	}

	for _, tc := range testCases {
		start := time.Now()
		result := findRotationEqual(tc.mat, tc.target)
		elapsed := time.Since(start)

		if result != tc.expected {
			t.Errorf("For mat=%v, target=%v, expected %t but got %t", tc.mat, tc.target, tc.expected, result)
		}

		fmt.Printf("Time taken to process test case: %v\n", elapsed)

		var m runtime.MemStats
		runtime.ReadMemStats(&m)
		fmt.Printf("Memory usage (HeapAlloc): %v bytes\n", m.HeapAlloc)
	}
}
