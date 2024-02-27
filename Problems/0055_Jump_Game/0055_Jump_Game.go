package main

import (
	"fmt"
	"time"
)

//approach Linear Scan
/*
Time Complexity: The algorithm has a time complexity of O(n), where n is the length of the input array nums. This is because the algorithm iterates through the array only once, performing a constant number of operations per element.

Space Complexity: The algorithm has a space complexity of O(1), meaning it uses constant extra space. This is because it only uses a fixed amount of extra space for storing variables such as maxReach, n, and loop counters, regardless of the size of the input array. Therefore, the space complexity does not grow with the size of the input.
*/

func canJump_LinearScan(nums []int) bool {
	maxReach := 0
	n := len(nums)

	for i := 0; i < n; i++ {
		fmt.Println("i", i, "currentmaxReach", maxReach)
		if i > maxReach {
			fmt.Println("i > maxReach", i, maxReach, "should be false")
			return false
		}
		// the key here is this: maxReach = max (maxReach, i + num[i])
		maxReach = max(maxReach, i+nums[i])
		fmt.Println("i", i, "nums[i]", nums[i], "i+ nums[i]", i+nums[i], "maxReach", maxReach)
		if maxReach >= n-1 {
			fmt.Println("maxReach >= n-1", maxReach, n-1, "should reach")
			return true
		}
	}

	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	timeStartWholeProgram := time.Now()

	testInput := []TestCase{
		{
			Nums: []int{2, 3, 1, 1, 4},
			Result: `
true
            `,
		},

		{
			Nums: []int{3, 2, 1, 0, 4},
			Result: `
false
            `,
		},
	}
	for count, value := range testInput {
		fmt.Println("===============")
		fmt.Println("Test count ", count, "for node", value)
		fmt.Println("Solution 1: Linear Scan")
		timeStart := time.Now()
		result := canJump_LinearScan(value.Nums)
		timeLapse := time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

	}

	timeLapsedWholeProgram := time.Since(timeStartWholeProgram)
	fmt.Println("===============")
	fmt.Println("TimeLapse Whole Program", timeLapsedWholeProgram)
}

type TestCase struct {
	Nums   []int
	Result string
}

/*

===============
Test count  0 for node {[2 3 1 1 4]
true
            }
Solution 1: Linear Scan
i 0 currentmaxReach 0
i 0 nums[i] 2 i+ nums[i] 2 maxReach 2
i 1 currentmaxReach 2
i 1 nums[i] 3 i+ nums[i] 4 maxReach 4
maxReach >= n-1 4 4 should reach
>Solution result true
Correct result is
true

TimeLapse 53.147µs
===============
Test count  1 for node {[3 2 1 0 4]
false
            }
Solution 1: Linear Scan
i 0 currentmaxReach 0
i 0 nums[i] 3 i+ nums[i] 3 maxReach 3
i 1 currentmaxReach 3
i 1 nums[i] 2 i+ nums[i] 3 maxReach 3
i 2 currentmaxReach 3
i 2 nums[i] 1 i+ nums[i] 3 maxReach 3
i 3 currentmaxReach 3
i 3 nums[i] 0 i+ nums[i] 3 maxReach 3
i 4 currentmaxReach 3
i > maxReach 4 3 should be false
>Solution result false
Correct result is
false

TimeLapse 100.202µs
===============
TimeLapse Whole Program 546.878µs

*/
//REF
//
