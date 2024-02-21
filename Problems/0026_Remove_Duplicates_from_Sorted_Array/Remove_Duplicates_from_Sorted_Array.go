package main

import (
	"fmt"
	"time"
)

//approach : 2 pointer
//Time Complexity: O(n).
//Space Complexity: O(1)

func removeDuplicates_2Pointer(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// Initialize pointer to track unique elements
	uniquePointer := 0

	// Iterate through the array
	for i := 1; i < len(nums); i++ {
		// If current element is different from the previous one
		if nums[i] != nums[uniquePointer] {
			// Move the uniquePointer forward
			uniquePointer++
			// Update the unique element at uniquePointer
			nums[uniquePointer] = nums[i]
		}
	}

	// Return the count of unique elements (uniquePointer + 1)
	return uniquePointer + 1
}

//==== approach brute force

func removeDuplicates_BruteForce(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return n
	}

	// Initialize a counter for unique elements
	uniqueCount := 1

	// Iterate through the array
	for i := 1; i < len(nums); i++ {
		// Check if the current element is a duplicate
		duplicate := false
		for j := 0; j < i; j++ {
			if nums[j] == nums[i] {
				duplicate = true
				break
			}
		}

		// If not a duplicate, update the unique count
		if !duplicate {
			nums[uniqueCount] = nums[i]
			uniqueCount++
		}
	}

	return uniqueCount
}

//======

func main() {
	timeStartWholeProgram := time.Now()

	testInput := []TestCase{

		{
			Nums: []int{1, 1, 2},
			Result: `
[1,2]

            `,
		},
		{
			Nums: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			Result: `
         [0,1,2,3,4]
            `,
		},
	}
	for count, value := range testInput {
		fmt.Println("===============")
		fmt.Println("Test count ", count, "for node", value)
		fmt.Println("Solution 1: 2 pointer")
		timeStart := time.Now()
		result := removeDuplicates_2Pointer(value.Nums)
		timeLapse := time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

		fmt.Println("Solution 2: brute force")
		timeStart = time.Now()
		result = removeDuplicates_BruteForce(value.Nums)
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
Test count  0 for node {[1 1 2]
[1,2]

            }
Solution 1: 2 pointer
>Solution result 2
Correct result is
[1,2]


TimeLapse 444ns
Solution 2: brute force
>Solution result 2
Correct result is
[1,2]


TimeLapse 426ns
===============
Test count  1 for node {[0 0 1 1 1 2 2 3 3 4]
         [0,1,2,3,4]
            }
Solution 1: 2 pointer
>Solution result 5
Correct result is
         [0,1,2,3,4]

TimeLapse 240ns
Solution 2: brute force
>Solution result 5
Correct result is
         [0,1,2,3,4]

TimeLapse 371ns
===============
TimeLapse Whole Program 420.944µs
*/
//REF
//
