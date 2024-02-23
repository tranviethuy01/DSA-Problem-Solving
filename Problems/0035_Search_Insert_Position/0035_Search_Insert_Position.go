package main

import (
	"fmt"
	"time"
)

//approach : Binary Search
//the time O(log n)
//space complexity : O(1)
//

func searchInsert_BinarySearch(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return left
}

//approach BruteForce
//the time O( n)
//space complexity : O(1)/
//

func searchInsert_BruteForce(nums []int, target int) int {
	for i, num := range nums {
		if num >= target {
			return i
		}
	}
	return len(nums)
}

func main() {
	timeStartWholeProgram := time.Now()

	testInput := []TestCase{

		{
			Nums:   []int{1, 3, 5, 6},
			Target: 5,
			Result: `
      2
            `,
		},

		{
			Nums:   []int{1, 3, 5, 6},
			Target: 2,
			Result: `
      1
            `,
		},

		{
			Nums:   []int{1, 3, 5, 6},
			Target: 7,
			Result: `
      4
            `,
		},
	}
	for count, value := range testInput {
		fmt.Println("===============")
		fmt.Println("Test count ", count, "for node", value)
		fmt.Println("Solution 1: Binary Search")
		timeStart := time.Now()
		result := searchInsert_BinarySearch(value.Nums, value.Target)
		timeLapse := time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

		fmt.Println("Solution 2: BruteForce")
		timeStart = time.Now()
		result = searchInsert_BruteForce(value.Nums, value.Target)
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
Test count  0 for node {[1 3 5 6] 5
      2
            }
Solution 1: Binary Search
>Solution result 2
Correct result is
      2

TimeLapse 556ns
Solution 2: BruteForce
>Solution result 2
Correct result is
      2

TimeLapse 148ns
===============
Test count  1 for node {[1 3 5 6] 2
      1
            }
Solution 1: Binary Search
>Solution result 1
Correct result is
      1

TimeLapse 148ns
Solution 2: BruteForce
>Solution result 1
Correct result is
      1

TimeLapse 111ns
===============
Test count  2 for node {[1 3 5 6] 7
      4
            }
Solution 1: Binary Search
>Solution result 4
Correct result is
      4

TimeLapse 130ns
Solution 2: BruteForce
>Solution result 4
Correct result is
      4

TimeLapse 111ns
===============
TimeLapse Whole Program 485.434µs

*/
