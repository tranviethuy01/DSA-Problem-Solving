package main

import (
	"fmt"
	"time"
)

//approach : 2 pointer
//Time Complexity: O(n).
//Space Complexity: O(1)

func removeElement_2Pointers(nums []int, val int) int {
	k := 0 // Counter for elements not equal to val
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[k] = nums[i]
			k++
		}
	}
	return k
}

//
//approach : brute force
//Time Complexity: O(n).
//Space Complexity: O(n): where n is the length of the input array nums. This is because the solution creates a new slice to store elements not equal to val, and the size of this slice could be at most equal to the size of the input array nums.

func removeElement_BruteForce(nums []int, val int) int {
	// Initialize a new slice to store elements not equal to val
	result := make([]int, 0)
	for _, num := range nums {
		if num != val {
			result = append(result, num)
		}
	}

	// Copy the elements back to nums
	for i := range result {
		nums[i] = result[i]
	}

	// Return the length of the new array
	return len(result)
}

func main() {
	timeStartWholeProgram := time.Now()

	testInput := []TestCase{

		{
			Nums: []int{3, 2, 2, 3},
			Val:  3,
			Result: `
[2,2]

            `,
		},
		{
			Nums: []int{0, 1, 2, 2, 3, 0, 4, 2},
			Val:  2,
			Result: `
         [0,1,3,0,4]
            `,
		},
	}
	for count, value := range testInput {
		fmt.Println("===============")
		fmt.Println("Test count ", count, "for node", value)
		fmt.Println("Solution 1: 2 pointer")
		timeStart := time.Now()
		result := removeElement_2Pointers(value.Nums, value.Val)
		timeLapse := time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

		fmt.Println("Solution 2: brute force")
		timeStart = time.Now()
		result = removeElement_BruteForce(value.Nums, value.Val)
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
	Val    int
	Result string
}

/*

===============
Test count  0 for node {[3 2 2 3] 3
[2,2]

            }
Solution 1: 2 pointer
>Solution result 2
Correct result is
[2,2]


TimeLapse 444ns
Solution 2: brute force
>Solution result 3
Correct result is
[2,2]


TimeLapse 1.019µs
===============
Test count  1 for node {[0 1 2 2 3 0 4 2] 2
         [0,1,3,0,4]
            }
Solution 1: 2 pointer
>Solution result 5
Correct result is
         [0,1,3,0,4]

TimeLapse 259ns
Solution 2: brute force
>Solution result 7
Correct result is
         [0,1,3,0,4]

TimeLapse 1.166µs
===============
TimeLapse Whole Program 496.363µs

*/
//REF
//
