package main

import (
	"fmt"
	"time"
)

//approach : binary search
/*
Time Complexity:
The time complexity of each binary search operation is O(log n), where n is the number of elements in the array.
We perform two binary searches (searchLeft and searchRight), both of which have a time complexity of O(log n).
Thus, the overall time complexity of the algorithm is O(log n).
Space Complexity:
The space complexity of the algorithm is O(1) because it uses only a constant amount of extra space regardless of the input size.
We are not using any additional data structures that grow with the size of the input array.
Therefore, the space complexity is O(1).
*/

func searchRange_BinarySearch(nums []int, target int) []int {
	leftIndex := searchLeft(nums, target)
	rightIndex := searchRight(nums, target)

	if leftIndex > rightIndex {
		return []int{-1, -1}
	}

	return []int{leftIndex, rightIndex}
}

func searchLeft(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}

func searchRight(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] <= target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return right
}

//approach Brute Force
/*
Time Complexity:
The algorithm iterates through the entire array once, checking each element to see if it matches the target value.
In the worst case, it requires scanning through all elements of the array.
Therefore, the time complexity is O(n), where n is the number of elements in the array.
Space Complexity:
The space complexity is O(1) because the algorithm only uses a constant amount of extra space regardless of the input size.
We're not using any additional data structures that grow with the size of the input array.
Therefore, the space complexity is O(1).
*/
func searchRange_BruteForce(nums []int, target int) []int {
	start, end := -1, -1

	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			if start == -1 {
				start = i
			}
			end = i
		}
	}

	return []int{start, end}
}

func main() {
	timeStartWholeProgram := time.Now()

	testInput := []TestCase{

		{
			Nums:   []int{4, 5, 6, 7, 0, 1, 2},
			Target: 0,
			Result: `
      4
            `,
		},
		{
			Nums:   []int{4, 5, 6, 7, 0, 1, 2},
			Target: 3,
			Result: `
      -1
            `,
		},
		{
			Nums:   []int{1},
			Target: 0,
			Result: `
			-1
            `,
		},
	}
	for count, value := range testInput {
		fmt.Println("===============")
		fmt.Println("Test count ", count, "for node", value)
		fmt.Println("Solution 1: Binary Search ")
		timeStart := time.Now()
		result := searchRange_BinarySearch(value.Nums, value.Target)
		timeLapse := time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

		fmt.Println("Solution 2: BruteForce")
		timeStart = time.Now()
		result = searchRange_BruteForce(value.Nums, value.Target)
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
	Target int
	Result string
}

/*

===============
Test count  0 for node {[4 5 6 7 0 1 2] 0
      4
            }
Solution 1: Binary Search
>Solution result [-1 -1]
Correct result is
      4

TimeLapse 1µs
Solution 2: BruteForce
>Solution result [4 4]
Correct result is
      4

TimeLapse 629ns
===============
Test count  1 for node {[4 5 6 7 0 1 2] 3
      -1
            }
Solution 1: Binary Search
>Solution result [-1 -1]
Correct result is
      -1

TimeLapse 463ns
Solution 2: BruteForce
>Solution result [-1 -1]
Correct result is
      -1

TimeLapse 389ns
===============
Test count  2 for node {[1] 0
			-1
            }
Solution 1: Binary Search
>Solution result [-1 -1]
Correct result is
			-1

TimeLapse 334ns
Solution 2: BruteForce
>Solution result [-1 -1]
Correct result is
			-1

TimeLapse 333ns
===============
TimeLapse Whole Program 517.327µs

*/
