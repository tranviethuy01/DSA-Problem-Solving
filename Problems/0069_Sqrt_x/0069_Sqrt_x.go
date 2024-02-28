package main

import (
	"fmt"
	"runtime"
	"time"
)

// approach : StraightForward . Bad code, just use it for education
// T: O(x)
// S: O(1)
func mySqrt_StraightForward(x int) int {
	if x == 0 {
		return 0
	}

	for i := 0; i <= x; i++ {
		square := i * i
		if square == x {
			return i
		} else if square > x {
			return i - 1
		}
	}
	return x
}

//approach binary search
/*
Time Complexity:
The time complexity of binary search is O(log n), where n is the range of the search space, which in this case is x. Since we're performing binary search over the range from 1 to x, the time complexity is O(log x).

Space Complexity:
The space complexity of the algorithm is O(1) because we are using a constant amount of extra space, regardless of the input size. We only use a few integer variables (left, right, mid, square) to keep track of the search space and perform calculations.

*/

func mySqrt_BinarySearch(x int) int {
	if x == 0 {
		return 0
	}

	left, right := 1, x
	for left < right {
		mid := left + (right-left)/2
		square := mid * mid
		if square == x {
			return mid
		} else if square < x {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left - 1
}

func main() {
	timeStartWholeProgram := time.Now()

	testInput := []TestCase{
		{
			X: 4,
			Result: `
	2
            `,
		},
		{
			X: 8,
			Result: `
2
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
		fmt.Println("Solution 1: StraightForward")
		timeStart := time.Now()
		result := mySqrt_StraightForward(value.X)
		timeLapse := time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

		// Memory after allocation
		runtime.ReadMemStats(&m)
		memAfter := m.Alloc
		fmt.Println("Memory before", memBefore, "bytes", "Memory after", memAfter, "bytes", "Memory used:", memAfter-memBefore, "bytes")
		fmt.Printf("Memory usage (HeapAlloc) after Test Case i %d, : %v bytes\n", count, m.HeapAlloc)

		fmt.Println("Solution 2: Binary Search")
		timeStart = time.Now()
		result = mySqrt_BinarySearch(value.X)
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
	X      int
	Result string
}

/*

===============
Test count  0 for node {4
	2
            }
Solution 1: StraightForward
>Solution result 2
Correct result is
	2

TimeLapse 556ns
Memory before 67184 bytes Memory after 68080 bytes Memory used: 896 bytes
Memory usage (HeapAlloc) after Test Case i 0, : 68080 bytes
Solution 2: Binary Search
>Solution result 2
Correct result is
	2

TimeLapse 185ns
Memory before 67184 bytes Memory after 68272 bytes Memory used: 1088 bytes
Memory usage (HeapAlloc) after Test Case i 0, : 68272 bytes
===============
Test count  1 for node {8
2
            }
Solution 1: StraightForward
>Solution result 2
Correct result is
2

TimeLapse 333ns
Memory before 67184 bytes Memory after 68360 bytes Memory used: 1176 bytes
Memory usage (HeapAlloc) after Test Case i 1, : 68360 bytes
Solution 2: Binary Search
>Solution result 2
Correct result is
2

TimeLapse 241ns
Memory before 67184 bytes Memory after 68424 bytes Memory used: 1240 bytes
Memory usage (HeapAlloc) after Test Case i 1, : 68424 bytes
===============
TimeLapse Whole Program 954.653µs


*/
//REF
//
