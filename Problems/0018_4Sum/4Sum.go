package main

import (
	"fmt"
	"sort"
	"time"
)

// approach: hash table
//Time Complexity: O(n^2)
//Space Complexity: O(n^2)

func fourSum_Hash(nums []int, target int) [][]int {
	// Sort the input array
	sort.Ints(nums)

	var result [][]int
	n := len(nums)

	// Create a hash table to store sums of pairs
	pairSum := make(map[int][][]int)

	// Populate the pairSum hash table
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			sum := nums[i] + nums[j]
			pairSum[sum] = append(pairSum[sum], []int{i, j})
		}
	}

	// Iterate through pairs and find complementary pairs
	for i := 0; i < n-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue // Skip duplicates
		}
		for j := i + 1; j < n-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue // Skip duplicates
			}
			complement := target - nums[i] - nums[j]
			if pairs, ok := pairSum[complement]; ok {
				for _, pair := range pairs {
					if pair[0] > j {
						quadruplet := []int{nums[i], nums[j], nums[pair[0]], nums[pair[1]]}
						result = append(result, quadruplet)
						break
					}
				}
			}
		}
	}

	return result
}

// approach: Sort + 2 pointers
//Time Complexity: O(n^2 * log n)
//Space Complexity: O(n^2)

func fourSum_2Pointers(nums []int, target int) [][]int {
	// Sort the input array
	sort.Ints(nums)

	var result [][]int

	// Loop through the array with two pointers
	for i := 0; i < len(nums)-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue // Skip duplicates
		}
		for j := i + 1; j < len(nums)-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue // Skip duplicates
			}
			left := j + 1
			right := len(nums) - 1
			for left < right {
				sum := nums[i] + nums[j] + nums[left] + nums[right]
				if sum == target {
					result = append(result, []int{nums[i], nums[j], nums[left], nums[right]})
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
	}

	return result
}

// approach: brute force
// Time Complexity: O(n^4)
// Space Complexity: O(n^4)
func fourSum_BruteForce(nums []int, target int) [][]int {
	var result [][]int

	// Brute-force approach
	for i := 0; i < len(nums)-3; i++ {
		for j := i + 1; j < len(nums)-2; j++ {
			for k := j + 1; k < len(nums)-1; k++ {
				for l := k + 1; l < len(nums); l++ {
					if nums[i]+nums[j]+nums[k]+nums[l] == target {
						quadruplet := []int{nums[i], nums[j], nums[k], nums[l]}
						sort.Ints(quadruplet) // Sort the quadruplet for uniqueness
						if !contains(result, quadruplet) {
							result = append(result, quadruplet)
						}
					}
				}
			}
		}
	}

	return result
}

// Helper function to check if a quadruplet already exists in the result
func contains(result [][]int, quadruplet []int) bool {
	for _, quad := range result {
		if quad[0] == quadruplet[0] && quad[1] == quadruplet[1] &&
			quad[2] == quadruplet[2] && quad[3] == quadruplet[3] {
			return true
		}
	}
	return false
}

//

func main() {
	timeStartWholeProgram := time.Now()
	testInput := []TestCase{
		{
			Nums:   []int{1, 0, -1, 0, -2, 2},
			Target: 0,
			Result: `
			[[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
            `,
		},
		{
			Nums:   []int{2, 2, 2, 2, 2},
			Target: 8,
			Result: `
         []
            `,
		},
	}
	for count, value := range testInput {
		fmt.Println("===============")
		fmt.Println("Test count ", count, "for node", value)
		fmt.Println("Solution 1: use 2 Pointers")
		timeStart := time.Now()
		result := fourSum_2Pointers(value.Nums, value.Target)
		timeLapse := time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

		fmt.Println("Solution 2: use Hash Table")
		timeStart = time.Now()
		result = fourSum_Hash(value.Nums, value.Target)
		timeLapse = time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

		fmt.Println("Solution 2: use Brute Force")
		timeStart = time.Now()
		result = fourSum_BruteForce(value.Nums, value.Target)
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
Test count  0 for node {[1 0 -1 0 -2 2] 0
			[[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
            }
Solution 1: use 2 Pointers
>Solution result [[-2 -1 1 2] [-2 0 0 2] [-1 0 0 1]]
Correct result is
			[[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]

TimeLapse 5.536µs
Solution 2: use Hash Table
>Solution result [[-2 -1 1 2] [-2 0 0 2] [-1 0 0 1]]
Correct result is
			[[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]

TimeLapse 10.756µs
Solution 2: use Brute Force
>Solution result [[-2 -1 1 2] [-2 0 0 2] [-1 0 0 1]]
Correct result is
			[[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]

TimeLapse 3.776µs
===============
Test count  1 for node {[2 2 2 2 2] 8
         []
            }
Solution 1: use 2 Pointers
>Solution result [[2 2 2 2]]
Correct result is
         []

TimeLapse 1.481µs
Solution 2: use Hash Table
>Solution result [[2 2 2 2]]
Correct result is
         []

TimeLapse 22.254µs
Solution 2: use Brute Force
>Solution result [[2 2 2 2]]
Correct result is
         []

TimeLapse 2.388µs
===============
TimeLapse Whole Program 574.737µs

*/
//REF
//
