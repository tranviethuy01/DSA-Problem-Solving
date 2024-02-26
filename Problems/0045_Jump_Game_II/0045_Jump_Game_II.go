package main

import (
	"fmt"
	"time"
)

//approach : Greedy approach
/*
Time Complexity:
The time complexity of this algorithm is O(n), where n is the length of the input array nums. This is because the algorithm iterates through the array only once.

Space Complexity:
The space complexity of this algorithm is O(1). It uses only a constant amount of extra space for storing variables like maxReach, steps, and lastJump, regardless of the size of the input array nums. Thus, the space complexity is constant or O(1).
*/
func jump_Greedy(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	maxReach, steps, lastJump := nums[0], nums[0], 1

	for i := 1; i < len(nums); i++ {
		if i == len(nums)-1 {
			return lastJump
		}
		maxReach = max(maxReach, i+nums[i])
		steps--
		if steps == 0 {
			steps = maxReach - i
			lastJump++
		}
	}
	return -1 // Should never reach here
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//approach Greedy 2

func jump_Greedy2(nums []int) int {
	for i := 1; i < len(nums); i++ {
		nums[i] = max(nums[i]+i, nums[i-1])
	}

	ind := 0
	ans := 0

	for ind < len(nums)-1 {
		ans++
		ind = nums[ind]
	}

	return ans
}

//
//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}
//
//

func main() {
	timeStartWholeProgram := time.Now()

	testInput := []TestCase{
		{
			Nums: []int{2, 3, 1, 1, 4},
			Result: `
2
            `,
		},

		{
			Nums: []int{2, 3, 0, 1, 4},
			Result: `
2
            `,
		},
	}
	for count, value := range testInput {
		fmt.Println("===============")
		fmt.Println("Test count ", count, "for node", value)
		fmt.Println("Solution 1: Greedy approach")
		timeStart := time.Now()
		result := jump_Greedy(value.Nums)
		timeLapse := time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

		fmt.Println("Solution 2: Greedy approach, change a little bit")
		timeStart = time.Now()
		result = jump_Greedy2(value.Nums)
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
	Nums   []int
	Result string
}

/*

===============
Test count  0 for node {[2 3 1 1 4]
2
            }
Solution 1: Greedy approach
>Solution result 2
Correct result is
2

TimeLapse 500ns
Solution 2: Greedy approach, change a little bit
>Solution result 2
Correct result is
2

TimeLapse 408ns
===============
Test count  1 for node {[2 3 0 1 4]
2
            }
Solution 1: Greedy approach
>Solution result 2
Correct result is
2

TimeLapse 278ns
Solution 2: Greedy approach, change a little bit
>Solution result 2
Correct result is
2

TimeLapse 148ns
===============
TimeLapse Whole Program 415.139µs
*/
//REF
//
