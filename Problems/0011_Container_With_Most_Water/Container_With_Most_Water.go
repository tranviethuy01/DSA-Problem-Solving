package main

import (
	"fmt"
	"time"
)

// Two Pointer Approach
// Time Complexity: O(n).
// Space Complexity: O(1).
func maxArea_TwoPointer(height []int) int {
	maxArea := 0
	left := 0
	right := len(height) - 1

	for left < right {
		// Calculate the area
		area := min(height[left], height[right]) * (right - left)

		// Update maxArea if current area is greater
		if area > maxArea {
			maxArea = area
		}

		// Move the pointer of the smaller line towards the other line
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return maxArea
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//===== approach: traightfoward
//Time Complexity: O(n^2).
//Space Complexity: O(1),

func maxArea_BruteForce(height []int) int {
	maxArea := 0
	n := len(height)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			// Calculate the area between two lines
			area := min(height[i], height[j]) * (j - i)
			// Update maxArea if current area is greater
			if area > maxArea {
				maxArea = area
			}
		}
	}

	return maxArea
}

//func min(a, b int) int {
//    if a < b {
//        return a
//    }
//    return b
//}
//

//=====

func main() {
	timeStartWholeProgram := time.Now()
	testInput := []TestCase{
		{
			Height: []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			Result: `
          49
            `,
		},
		{

			Height: []int{1, 1},
			Result: `
             1
             `,
		},
	}
	for count, value := range testInput {
		fmt.Println("===============")
		fmt.Println("Test count ", count, "for node", value)
		fmt.Println("Solution 1: use Two Pointer")
		timeStart := time.Now()
		result := maxArea_TwoPointer(value.Height)
		timeLapse := time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

		fmt.Println("Solution 2: use brute force")
		timeStart = time.Now()
		result = maxArea_TwoPointer(value.Height)
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
	Height []int
	Result string
}

/*


===============
Test count  0 for node {[1 8 6 2 5 4 8 3 7]
          49
            }
Solution 1: use Two Pointer
>Solution result 49
Correct result is
          49

TimeLapse 574ns
Solution 2: use brute force
>Solution result 49
Correct result is
          49

TimeLapse 334ns
===============
Test count  1 for node {[1 1]
             1
             }
Solution 1: use Two Pointer
>Solution result 1
Correct result is
             1

TimeLapse 130ns
Solution 2: use brute force
>Solution result 1
Correct result is
             1

TimeLapse 129ns
===============
TimeLapse Whole Program 377.546µs


*/
//REF
//
