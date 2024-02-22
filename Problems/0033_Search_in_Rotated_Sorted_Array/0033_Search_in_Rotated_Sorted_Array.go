package main

import (
	"fmt"
	"time"
)

//approach : binary search algo
//Time Complexity: O(log n) - This is because the algorithm halves the search space in each iteration, leading to a logarithmic time complexity.
//Space Complexity: O(1) - The algorithm uses only a constant amount of extra space for variables such as left, right, and mid. The space complexity is independent of the input size n, making it O(1).

func search_BinarySearch(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			return mid
		}

		if nums[left] <= nums[mid] { // left half is sorted
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else { // right half is sorted
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return -1
}

// approach : Brute Force

func search_BruteForce(nums []int, target int) int {
	for i, num := range nums {
		if num == target {
			return i
		}
	}
	return -1
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
		result := search_BinarySearch(value.Nums, value.Target)
		timeLapse := time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

		fmt.Println("Solution 2: BruteForce")
		timeStart = time.Now()
		result = search_BruteForce(value.Nums, value.Target)
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
>Solution result 4
Correct result is
      4

TimeLapse 611ns
Solution 2: BruteForce
>Solution result 4
Correct result is
      4

TimeLapse 370ns
===============
Test count  1 for node {[4 5 6 7 0 1 2] 3
      -1
            }
Solution 1: Binary Search
>Solution result -1
Correct result is
      -1

TimeLapse 204ns
Solution 2: BruteForce
>Solution result -1
Correct result is
      -1

TimeLapse 148ns
===============
Test count  2 for node {[1] 0
			-1
            }
Solution 1: Binary Search
>Solution result -1
Correct result is
			-1

TimeLapse 148ns
Solution 2: BruteForce
>Solution result -1
Correct result is
			-1

TimeLapse 130ns
===============
TimeLapse Whole Program 539.697µs

*/
