package main

import (
	"fmt"
	"math"
	"sort"
	"time"
)

//approach: binary search
//Time Complexity: O(log(min(m, n)))

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}
	x := len(nums1)
	y := len(nums2)

	low := 0
	high := x

	for low <= high {
		partitionX := (low + high) / 2
		partitionY := (x+y+1)/2 - partitionX

		maxLeftX := math.MinInt64
		if partitionX > 0 {
			maxLeftX = nums1[partitionX-1]
		}

		minRightX := math.MaxInt64
		if partitionX < x {
			minRightX = nums1[partitionX]
		}

		maxLeftY := math.MinInt64
		if partitionY > 0 {
			maxLeftY = nums2[partitionY-1]
		}

		minRightY := math.MaxInt64
		if partitionY < y {
			minRightY = nums2[partitionY]
		}

		if maxLeftX <= minRightY && maxLeftY <= minRightX {
			if (x+y)%2 == 0 {
				return float64(max(maxLeftX, maxLeftY)+min(minRightX, minRightY)) / 2.0
			} else {
				return float64(max(maxLeftX, maxLeftY))
			}
		} else if maxLeftX > minRightY {
			high = partitionX - 1
		} else {
			low = partitionX + 1
		}
	}
	return 0.0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// ===== approach: brute force
// Time complexity of O((m + n) log(m + n))
func findMedianSortedArrays_BruteForce(nums1 []int, nums2 []int) float64 {
	merged := append(nums1, nums2...)
	sort.Ints(merged)
	n := len(merged)
	if n%2 == 0 {
		return float64(merged[n/2-1]+merged[n/2]) / 2
	}
	return float64(merged[n/2])
}

//=====

func main() {
	timeStartWholeProgram := time.Now()
	testInput := []TestCase{
		{
			Num1: []int{1, 3},
			Num2: []int{2},
			Result: `
                2.00000
            `,
		},
		{
			Num1: []int{1, 2},
			Num2: []int{3, 4},
			Result: `
                2.50000
            `,
		},
	}
	for count, value := range testInput {
		fmt.Println("===============")
		fmt.Println("Test count ", count, "for node", value)
		fmt.Println("Solution 1: use binary search")
		timeStart := time.Now()
		result := findMedianSortedArrays(value.Num1, value.Num2)
		timeLapse := time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

		fmt.Println("Solution 2: use brute force solution")
		timeStart = time.Now()
		result = findMedianSortedArrays_BruteForce(value.Num1, value.Num2)
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
	Num1   []int
	Num2   []int
	Result string
}

/*


===============
Test count  0 for node {[1 3] [2]
                2.00000
            }
Solution 1: use binary search
>Solution result 2
Correct result is
                2.00000

TimeLapse 1.13µs
Solution 2: use brute force solution
>Solution result 2
Correct result is
                2.00000

TimeLapse 3.092µs
===============
Test count  1 for node {[1 2] [3 4]
                2.50000
            }
Solution 1: use binary search
>Solution result 2.5
Correct result is
                2.50000

TimeLapse 241ns
Solution 2: use brute force solution
>Solution result 2.5
Correct result is
                2.50000

TimeLapse 1.019µs
===============
TimeLapse Whole Program 400.137µs

*/
