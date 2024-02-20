package main

import (
	"fmt"
	"sort"
	"time"
)

//===== Fix the first element and use two pointers for the rest of the array
//Time Complexity: O(n^2) .
//Space Complexity: O(1).

func threeSum_2Pointer(nums []int) [][]int {
	var res [][]int
	sort.Ints(nums)
	n := len(nums)

	for i := 0; i < n-2; i++ {
		// Skip duplicates
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left, right := i+1, n-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum < 0 {
				left++
			} else if sum > 0 {
				right--
			} else {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			}
		}
	}
	return res
}

//=== use two sum approach

func threeSum_UseWith2Sum(nums []int) [][]int {
	// Sort the array
	sort.Ints(nums)

	var result [][]int

	// Fix the first element and use two pointers for the rest of the array
	for i := 0; i < len(nums)-2; i++ {
		// Skip duplicates
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		//now it is 2sum approach
		target := -nums[i]
		left, right := i+1, len(nums)-1

		for left < right {
			sum := nums[left] + nums[right]
			if sum == target {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				// Skip duplicates
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if sum < target {
				left++
			} else {
				right--
			}
		}
	}

	return result
}

//=== use hash map
//REF
//https://www.nileshblog.tech/leet-code-three-3-sum-java-cpp-python-solution/#Python_HashMap_Approach
//Time Complexity: O(n log n) due to the sorting operation.
//Space Complexity: O(n) due to the hash map used to store indices.

func threeSum_HashMap(nums []int) [][]int {
	sort.Ints(nums)    // Sorted Array
	if len(nums) < 3 { // Base Case 1
		return [][]int{}
	}
	if nums[0] > 0 { // Base Case 2
		return [][]int{}
	}

	hashMap := make(map[int]int)
	answer := [][]int{}

	// Hashing of Indices
	for i := 0; i < len(nums); i++ {
		hashMap[nums[i]] = i
	}

	// Traversing the array to fix the number
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 { // If number fixed is positive, stop there because we can't make it zero by searching after it.
			break
		}
		// Fixing another number after first number
		for j := i + 1; j < len(nums)-1; j++ {
			required := -1 * (nums[i] + nums[j])
			// To make sum 0, we would require the negative sum of both fixed numbers.
			if index, found := hashMap[required]; found && index > j {
				answer = append(answer, []int{nums[i], nums[j], required})
			}
			// Update j to last occurrence of 2nd fixed number to avoid duplicate triplets.
			j = hashMap[nums[j]]
		}
		// Update i to last occurrence of 1st fixed number to avoid duplicate triplets.
		i = hashMap[nums[i]]
	}
	return answer // Return answer vector.
}

//====

//===

// ==== approach brute force
// need to check this solution
func threeSum_BruteForce(nums []int) [][]int {
	var result [][]int
	n := len(nums)

	// Brute force approach to find all triplets
	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			for k := j + 1; k < n; k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					// Check if the triplet is already in the result
					found := false
					for _, res := range result {
						if equalSlices(res, []int{nums[i], nums[j], nums[k]}) {
							found = true
							break
						}
					}
					if !found {
						result = append(result, []int{nums[i], nums[j], nums[k]})
					}
				}
			}
		}
	}

	return result
}

// equalSlices checks if two slices are equal
func equalSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

//====

//=====

//==== two sum problem

func twoSum_Hash(nums []int, target int) []int {
	// Create a map to store the indices of elements
	indexMap := make(map[int]int)

	// Iterate through the array
	for i, num := range nums {
		// Check if the complement exists in the map
		complement := target - num
		if index, found := indexMap[complement]; found {
			// If found, return the indices
			return []int{index, i}
		}
		// Otherwise, store the current element's index
		indexMap[num] = i
	}

	// If no solution is found, return an empty slice
	return nil
}

//===

// ==== twoSum brute force
func twoSum_BruteForce(nums []int, target int) []int {
	n := len(nums)

	// Iterate through every pair of elements
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			// Check if the sum of elements equals the target
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	// If no solution is found, return an empty slice
	return nil
}

//===

func main() {
	timeStartWholeProgram := time.Now()
	testInput := []TestCase{
		{
			Nums: []int{-1, 0, 1, 2, -1, -4},
			Result: `
          [[-1,-1,2],[-1,0,1]]
            `,
		},
		{
			Nums: []int{0, 1, 1},
			Result: `
          []
            `,
		},

		{
			Nums: []int{0, 0, 0},
			Result: `
          [[0,0,0]]
            `,
		},
	}
	for count, value := range testInput {
		fmt.Println("===============")
		fmt.Println("Test count ", count, "for node", value)
		fmt.Println("Solution 1: fix first element and use 2 pointers for the rest")
		timeStart := time.Now()
		result := threeSum_2Pointer(value.Nums)
		timeLapse := time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

		fmt.Println("Solution 2: use 2 sum approach (organize code a little compare to solution 1)")
		timeStart = time.Now()
		result = threeSum_UseWith2Sum(value.Nums)
		timeLapse = time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

		fmt.Println("Solution 3: Hash map approach")
		timeStart = time.Now()
		result = threeSum_HashMap(value.Nums)
		timeLapse = time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

		fmt.Println("Solution 4: Brute Force approach")
		timeStart = time.Now()
		result = threeSum_BruteForce(value.Nums)
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
Test count  0 for node {[-1 0 1 2 -1 -4]
          [[-1,-1,2],[-1,0,1]]
            }
Solution 1: fix first element and use 2 pointers for the rest
>Solution result [[-1 -1 2] [-1 0 1]]
Correct result is
          [[-1,-1,2],[-1,0,1]]

TimeLapse 5.074µs
Solution 2: use 2 sum approach (organize code a little compare to solution 1)
>Solution result [[-1 -1 2] [-1 0 1]]
Correct result is
          [[-1,-1,2],[-1,0,1]]

TimeLapse 2.611µs
Solution 3: Hash map approach
>Solution result [[-1 -1 2] [-1 0 1]]
Correct result is
          [[-1,-1,2],[-1,0,1]]

TimeLapse 6.667µs
Solution 4: Brute Force approach
>Solution result [[-1 -1 2] [-1 0 1]]
Correct result is
          [[-1,-1,2],[-1,0,1]]

TimeLapse 2.203µs
===============
Test count  1 for node {[0 1 1]
          []
            }
Solution 1: fix first element and use 2 pointers for the rest
>Solution result []
Correct result is
          []

TimeLapse 908ns
Solution 2: use 2 sum approach (organize code a little compare to solution 1)
>Solution result []
Correct result is
          []

TimeLapse 740ns
Solution 3: Hash map approach
>Solution result []
Correct result is
          []

TimeLapse 1.389µs
Solution 4: Brute Force approach
>Solution result []
Correct result is
          []

TimeLapse 167ns
===============
Test count  2 for node {[0 0 0]
          [[0,0,0]]
            }
Solution 1: fix first element and use 2 pointers for the rest
>Solution result [[0 0 0]]
Correct result is
          [[0,0,0]]

TimeLapse 1.167µs
Solution 2: use 2 sum approach (organize code a little compare to solution 1)
>Solution result [[0 0 0]]
Correct result is
          [[0,0,0]]

TimeLapse 1.148µs
Solution 3: Hash map approach
>Solution result [[0 0 0]]
Correct result is
          [[0,0,0]]

TimeLapse 1.648µs
Solution 4: Brute Force approach
>Solution result [[0 0 0]]
Correct result is
          [[0,0,0]]

TimeLapse 833ns
===============
TimeLapse Whole Program 741.256µs

*/
//REF
//
