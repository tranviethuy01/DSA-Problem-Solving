package main

import (
	"fmt"
	"sort"
	"time"
)

// best solution: use additional space, hash map to speed up the search
// Time complexity: O(n)
// Space complexity: O(n)
func twoSum(nums []int, target int) []int {
	numIndices := make(map[int]int)
	for i, num := range nums {
		complement := target - num
		if j, ok := numIndices[complement]; ok {
			return []int{j, i}
		}
		numIndices[num] = i
	}
	return nil
}

// the simplest approach : Brute Force
// Time Complexity:  O(n2).
// Space complexity: O(1)
func twoSum_BruteForce(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			//skip the same element
			if i == j {
				continue
			}

			//pair existed
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}

			//optmized more on sorted array
			// if nums[i] + nums[j] > target {
			//     break
			// }

		}
	}

	//not found
	return nil
}

// this 2 pointer should used on a sorted array
// Sorting the array: O(n log n)
// Two-pointer traversal: O(n)
// => Time complexity: O(n log n)
func twoSum_2pointer(nums []int, target int) []int {
	// Sort the array
	sort.Ints(nums)

	// Initialize pointers
	left, right := 0, len(nums)-1

	// Two pointer approach
	for left < right {
		sum := nums[left] + nums[right]
		if sum == target {
			return []int{left, right}
		} else if sum < target {
			left++
		} else {
			right--
		}
	}

	// If no solution found
	return nil
}

func main() {
	timeStartWholeProgram := time.Now()
	testInput := []TestCase{
		{
			Target: 9,
			Nums:   []int{2, 7, 11, 15},
			Result: `
                [0,1]
            `,
		},

		{
			Target: 9,
			Nums:   []int{15, 2, 7, 11},
			Result: `
                [1,2]
            `,
		},

		{
			Target: 6,
			Nums:   []int{3, 2, 4},
			Result: `
                [1,2]
            `,
		},
		{
			Target: 6,
			Nums:   []int{3, 3},
			Result: `
                [0,1]
            `,
		},
	}
	for count, value := range testInput {
		fmt.Println("===============")
		fmt.Println("Test count ", count, "for node", value)
		fmt.Println("Solution 1: best method: use hash table")
		timeStart := time.Now()
		result := twoSum(value.Nums, value.Target)
		timeLapse := time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

		fmt.Println("Solution 2: the simplest approach: Brute Force")
		timeStart = time.Now()
		result = twoSum_BruteForce(value.Nums, value.Target)
		timeLapse = time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

		fmt.Println("Solution 3: 2 pointer")
		timeStart = time.Now()
		result = twoSum_2pointer(value.Nums, value.Target)
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
	Target int
	Nums   []int
	Result string
}

/*

===============
Test count  0 for node {9 [2 7 11 15]
                [0,1]
            }
Solution 1: best method: use hash table
>Solution result [0 1]
Correct result is
                [0,1]

TimeLapse 3.148µs
Solution 2: the simplest approach: Brute Force
>Solution result [0 1]
Correct result is
                [0,1]

TimeLapse 352ns
Solution 3: 2 pointer
>Solution result [0 1]
Correct result is
                [0,1]

TimeLapse 2.074µs
===============
Test count  1 for node {9 [15 2 7 11]
                [1,2]
            }
Solution 1: best method: use hash table
>Solution result [1 2]
Correct result is
                [1,2]

TimeLapse 1.056µs
Solution 2: the simplest approach: Brute Force
>Solution result [1 2]
Correct result is
                [1,2]

TimeLapse 407ns
Solution 3: 2 pointer
>Solution result [0 1]
Correct result is
                [1,2]

TimeLapse 1.092µs
===============
Test count  2 for node {6 [3 2 4]
                [1,2]
            }
Solution 1: best method: use hash table
>Solution result [1 2]
Correct result is
                [1,2]

TimeLapse 760ns
Solution 2: the simplest approach: Brute Force
>Solution result [1 2]
Correct result is
                [1,2]

TimeLapse 352ns
Solution 3: 2 pointer
>Solution result [0 2]
Correct result is
                [1,2]

TimeLapse 814ns
===============
Test count  3 for node {6 [3 3]
                [0,1]
            }
Solution 1: best method: use hash table
>Solution result [0 1]
Correct result is
                [0,1]

TimeLapse 685ns
Solution 2: the simplest approach: Brute Force
>Solution result [0 1]
Correct result is
                [0,1]

TimeLapse 334ns
Solution 3: 2 pointer
>Solution result [0 1]
Correct result is
                [0,1]

TimeLapse 741ns
===============
TimeLapse Whole Program 766.245µs

*/
