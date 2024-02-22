package main

import (
	"fmt"
	"time"
)

//appraach : Binary Search
/*
Time Complexity:

In each step of the binary search, the algorithm reduces the search range by half. Thus, the time complexity of the binary search algorithm is O(log n), where n is the number of versions.
Space Complexity:

The space complexity of the algorithm is O(1) because it uses only a constant amount of extra space regardless of the input size. The algorithm doesn't require any additional data structures that grow with the input size; it only uses a few variables to keep track of indices.
*/

var bad int

func isBadVersion(version int) bool {
	return version >= bad
}

func firstBadVersion_BinarySearch(n int) int {
	left, right := 1, n
	for left < right {
		mid := left + (right-left)/2
		if isBadVersion(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

//approach Linear
//Time Complexity:The time complexity of this solution is O(n), where n is the number of versions. This is because in the worst-case scenario, we may need to iterate through all versions from 1 to n until we find the first bad version.
//Space Complexity:The space complexity of this solution is O(1), which means it uses constant space regardless of the input size. This is because the algorithm only uses a fixed amount of additional space for storing variables like the loop index and function parameters. It doesn't require any extra data structures that grow with the input size.

func firstBadVersion_Linear(n int) int {
	for i := 1; i <= n; i++ {
		if isBadVersion(i) {
			return i
		}
	}
	return -1 // Indicates no bad version found (not expected given the problem constraints)
}

func main() {
	timeStartWholeProgram := time.Now()

	testInput := []TestCase{

		{
			N:   5,
			Bad: 4,
			Result: `
      4
            `,
		},
		{
			N:   1,
			Bad: 1,
			Result: `
      1
            `,
		},
	}
	for count, value := range testInput {
		fmt.Println("===============")
		fmt.Println("Test count ", count, "for node", value)
		fmt.Println("Solution 1: Binary Search ")
		timeStart := time.Now()
		result := firstBadVersion_BinarySearch(value.N)
		timeLapse := time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

		fmt.Println("Solution 2: BruteForce")
		timeStart = time.Now()
		result = firstBadVersion_Linear(value.N)
		timeLapse = time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

	}

	timeLapsedWholeProgram := time.Since(timeStartWholeProgram)
	fmt.Println("===============")
	fmt.Println("TimeLapse Whole Program", timeLapsedWholeProgram)
}

type TestCase struct {
	N      int
	Bad    int
	Result string
}

/*

===============
Test count  0 for node {5 4
      4
            }
Solution 1: Binary Search
>Solution result 1
Correct result is
      4

TimeLapse 555ns
Solution 2: BruteForce
>Solution result 1
Correct result is
      4

TimeLapse 148ns
===============
Test count  1 for node {1 1
      1
            }
Solution 1: Binary Search
>Solution result 1
Correct result is
      1

TimeLapse 92ns
Solution 2: BruteForce
>Solution result 1
Correct result is
      1

TimeLapse 111ns
===============
TimeLapse Whole Program 364.479µs


*/
