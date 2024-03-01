package main

import (
	"fmt"
	"runtime"
	"time"
)

//approach Linear Scan
/*
The time complexity of this algorithm is O(n), where n is the number of intervals in the input list.

Space complexity, the algorithm uses additional space to store the result, which could potentially contain all the intervals in the worst-case scenario. Therefore, the space complexity is O(n), where n is the number of intervals in the input list.

*/

func insert_LinearScan(intervals [][]int, newInterval []int) [][]int {
	result := [][]int{}
	merged := newInterval
	for _, interval := range intervals {
		fmt.Println("interval", interval, "merged", merged)
		if interval[1] < merged[0] {
			fmt.Println("intervals[1] < merged[0], need append to result")
			result = append(result, interval)
		} else if interval[0] > merged[1] {
			fmt.Println("intervals[0] > merged[1], need append to result, and change merged = interval")
			result = append(result, merged)
			merged = interval
		} else {
			fmt.Println("need re calculate value of merged")
			fmt.Println("merged before", merged)
			merged[0] = min(merged[0], interval[0])
			merged[1] = max(merged[1], interval[1])
			fmt.Println("merged after", merged)

		}
	}
	result = append(result, merged)
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//approach: try to use Binary search => note, still O(n)
/*


The modified solution still utilizes a linear scan for merging intervals, which results in O(n) time complexity for merging, where n is the number of intervals. The insertion process, however, is optimized using binary search, resulting in O(log n) time complexity for finding the insertion point. Therefore, the overall time complexity of the algorithm is O(n) + O(log n), which simplifies to O(n).

In terms of space complexity, the algorithm uses additional space to store the result, which could potentially contain all the intervals in the worst-case scenario. Therefore, the space complexity is O(n), where n is the number of intervals in the input list.

To summarize:

Time complexity: O(n)
Space complexity: O(n)

*/

func insert_BinarySearch(intervals [][]int, newInterval []int) [][]int {
	start, _ := newInterval[0], newInterval[1]
	left := 0
	right := len(intervals) - 1
	insertIdx := len(intervals)

	// Binary search to find the insertion point
	for left <= right {
		mid := left + (right-left)/2
		if intervals[mid][0] <= start {
			left = mid + 1
		} else {
			insertIdx = mid
			right = mid - 1
		}
	}

	// Insert newInterval at the found index
	intervals = append(intervals[:insertIdx], append([][]int{newInterval}, intervals[insertIdx:]...)...)

	// Merge intervals if necessary
	result := [][]int{}
	for _, interval := range intervals {
		if len(result) == 0 || interval[0] > result[len(result)-1][1] {
			result = append(result, interval)
		} else {
			result[len(result)-1][1] = max(result[len(result)-1][1], interval[1])
		}
	}
	return result
}

//func max(a, b int) int {
//    if a > b {
//        return a
//    }
//    return b
//}
//
//

func main() {
	timeStartWholeProgram := time.Now()

	testInput := []TestCase{
		{
			Intervals:   [][]int{{1, 3}, {6, 9}},
			NewInterval: []int{2, 5},
			Result: `
[[1,5],[6,9]]


            `,
		},
		{
			Intervals:   [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}},
			NewInterval: []int{4, 8},
			Result: `
[[1,2],[3,10],[12,16]]

            `,
		},
		{
			Intervals:   [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}, {17, 20}, {25, 30}},
			NewInterval: []int{4, 8},
			Result: `
[[1,2],[3,10],[12,16] , [17,20], [25,30]]


            `,
		},
	}

	// Memory before allocation
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	memBefore := m.Alloc

	for count, value := range testInput {
		fmt.Println("===============")
		fmt.Println("Test count ", count, "for node", value)
		fmt.Println("Solution 1: Linear Scan")
		timeStart := time.Now()
		result := insert_LinearScan(value.Intervals, value.NewInterval)
		timeLapse := time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

		// Memory after allocation
		runtime.ReadMemStats(&m)
		memAfter := m.Alloc
		fmt.Println("Memory before", memBefore, "bytes", "Memory after", memAfter, "bytes", "Memory used:", memAfter-memBefore, "bytes")

		fmt.Printf("Memory usage (HeapAlloc) after Test Case i %d, : %v bytes\n", count, m.HeapAlloc)

		fmt.Println("Solution 2: try to use Binary Search, but still O(n) ")
		timeStart = time.Now()
		result = insert_BinarySearch(value.Intervals, value.NewInterval)
		timeLapse = time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

		// Memory after allocation
		runtime.ReadMemStats(&m)
		memAfter = m.Alloc
		fmt.Println("Memory before", memBefore, "bytes", "Memory after", memAfter, "bytes", "Memory used:", memAfter-memBefore, "bytes")

		fmt.Printf("Memory usage (HeapAlloc) after Test Case i %d, : %v bytes\n", count, m.HeapAlloc)

	}

	timeLapsedWholeProgram := time.Since(timeStartWholeProgram)
	fmt.Println("===============")
	fmt.Println("TimeLapse Whole Program", timeLapsedWholeProgram)
}

type TestCase struct {
	Intervals   [][]int
	NewInterval []int
	Result      string
}

/*


===============
Test count  0 for node {[[1 3] [6 9]] [2 5]
[[1,5],[6,9]]


            }
Solution 1: Linear Scan
interval [1 3] merged [2 5]
need re calculate value of merged
merged before [2 5]
merged after [1 5]
interval [6 9] merged [1 5]
intervals[0] > merged[1], need append to result, and change merged = interval
>Solution result [[1 5] [6 9]]
Correct result is
[[1,5],[6,9]]



TimeLapse 67.203µs
Memory before 69664 bytes Memory after 71256 bytes Memory used: 1592 bytes
Memory usage (HeapAlloc) after Test Case i 0, : 71256 bytes
Solution 2: try to use Binary Search, but still O(n)
>Solution result [[1 5] [6 9]]
Correct result is
[[1,5],[6,9]]



TimeLapse 5.389µs
Memory before 69664 bytes Memory after 71640 bytes Memory used: 1976 bytes
Memory usage (HeapAlloc) after Test Case i 0, : 71640 bytes
===============
Test count  1 for node {[[1 2] [3 5] [6 7] [8 10] [12 16]] [4 8]
[[1,2],[3,10],[12,16]]

            }
Solution 1: Linear Scan
interval [1 2] merged [4 8]
intervals[1] < merged[0], need append to result
interval [3 5] merged [4 8]
need re calculate value of merged
merged before [4 8]
merged after [3 8]
interval [6 7] merged [3 8]
need re calculate value of merged
merged before [3 8]
merged after [3 8]
interval [8 10] merged [3 8]
need re calculate value of merged
merged before [3 8]
merged after [3 10]
interval [12 16] merged [3 10]
intervals[0] > merged[1], need append to result, and change merged = interval
>Solution result [[1 2] [3 10] [12 16]]
Correct result is
[[1,2],[3,10],[12,16]]


TimeLapse 166.054µs
Memory before 69664 bytes Memory after 72952 bytes Memory used: 3288 bytes
Memory usage (HeapAlloc) after Test Case i 1, : 72952 bytes
Solution 2: try to use Binary Search, but still O(n)
>Solution result [[1 2] [3 10] [12 16]]
Correct result is
[[1,2],[3,10],[12,16]]


TimeLapse 16.425µs
Memory before 69664 bytes Memory after 73680 bytes Memory used: 4016 bytes
Memory usage (HeapAlloc) after Test Case i 1, : 73680 bytes
===============
Test count  2 for node {[[1 2] [3 5] [6 7] [8 10] [12 16] [17 20] [25 30]] [4 8]
[[1,2],[3,10],[12,16] , [17,20], [25,30]]


            }
Solution 1: Linear Scan
interval [1 2] merged [4 8]
intervals[1] < merged[0], need append to result
interval [3 5] merged [4 8]
need re calculate value of merged
merged before [4 8]
merged after [3 8]
interval [6 7] merged [3 8]
need re calculate value of merged
merged before [3 8]
merged after [3 8]
interval [8 10] merged [3 8]
need re calculate value of merged
merged before [3 8]
merged after [3 10]
interval [12 16] merged [3 10]
intervals[0] > merged[1], need append to result, and change merged = interval
interval [17 20] merged [12 16]
intervals[0] > merged[1], need append to result, and change merged = interval
interval [25 30] merged [17 20]
intervals[0] > merged[1], need append to result, and change merged = interval
>Solution result [[1 2] [3 10] [12 16] [17 20] [25 30]]
Correct result is
[[1,2],[3,10],[12,16] , [17,20], [25,30]]



TimeLapse 214.589µs
Memory before 69664 bytes Memory after 75760 bytes Memory used: 6096 bytes
Memory usage (HeapAlloc) after Test Case i 2, : 75760 bytes
Solution 2: try to use Binary Search, but still O(n)
>Solution result [[1 2] [3 10] [12 16] [17 20] [25 30]]
Correct result is
[[1,2],[3,10],[12,16] , [17,20], [25,30]]



TimeLapse 16.944µs
Memory before 69664 bytes Memory after 76920 bytes Memory used: 7256 bytes
Memory usage (HeapAlloc) after Test Case i 2, : 76920 bytes
===============
TimeLapse Whole Program 1.653292ms

*/
//REF
//
