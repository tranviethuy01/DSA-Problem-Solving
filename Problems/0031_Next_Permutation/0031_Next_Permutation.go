package main

import (
	"fmt"
	"sort"
	"time"
)

//
/*
Time Complexity:

Finding the first decreasing element takes O(n) time, where n is the number of elements in the array.
Finding the next greater element takes O(n) time, where n is the number of elements in the array.
Reversing the suffix takes O(n/2) time, which simplifies to O(n).
Overall, the time complexity is O(n).
Space Complexity:

The algorithm uses only constant extra memory, regardless of the size of the input array.
Thus, the space complexity is O(1).
This algorithm efficiently finds the next permutation in place with a time complexity of O(n) and constant space complexity.
*/

func nextPermutation(nums []int) {
	// Find the first decreasing element from the end
	i := len(nums) - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}

	// If such element is found, find the next greater element to swap with
	if i >= 0 {
		j := len(nums) - 1
		for j >= 0 && nums[j] <= nums[i] {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}

	// Reverse the suffix from i+1 to the end
	reverse(nums[i+1:])
}

// Helper function to reverse a slice
func reverse(nums []int) {
	left, right := 0, len(nums)-1
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}

//approach 2: optimize code more

func nextPermutation_Approach2(nums []int) {
	if len(nums) <= 1 {
		return // No permutation possible for single element or empty array
	}

	// Find the largest index k such that nums[k] < nums[k+1]
	k := len(nums) - 2
	for k >= 0 && nums[k] >= nums[k+1] {
		k--
	}

	if k == -1 {
		// If no such index exists, reverse the whole array
		reverse_Approach2(nums)
		return
	}

	// Find the largest index l greater than k such that nums[k] < nums[l]
	l := len(nums) - 1
	for l >= 0 && nums[l] <= nums[k] {
		l--
	}

	// Swap nums[k] and nums[l]
	nums[k], nums[l] = nums[l], nums[k]

	// Reverse the elements from k+1 to the end
	reverse_Approach2(nums[k+1:])
}

func reverse_Approach2(nums []int) {
	left, right := 0, len(nums)-1
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}

//approach: brute force : failure solution, need check again

func nextPermutation_BruteForce(nums []int) []int {
	// Generate all permutations
	permutations := [][]int{}
	generatePermutations(nums, 0, &permutations)

	// Sort permutations lexicographically
	sort.Slice(permutations, func(i, j int) bool {
		for k := range permutations[i] {
			if permutations[i][k] != permutations[j][k] {
				return permutations[i][k] < permutations[j][k]
			}
		}
		return false
	})

	// Find nums in permutations and return the next permutation
	for i := range permutations {
		if isEqual(nums, permutations[i]) {
			return permutations[(i+1)%len(permutations)]
		}
	}

	return nil
}

func generatePermutations(nums []int, start int, result *[][]int) {
	if start == len(nums)-1 {
		*result = append(*result, append([]int(nil), nums...))
		return
	}

	for i := start; i < len(nums); i++ {
		nums[start], nums[i] = nums[i], nums[start]
		generatePermutations(nums, start+1, result)
		nums[start], nums[i] = nums[i], nums[start]
	}
}

func isEqual(nums1, nums2 []int) bool {
	if len(nums1) != len(nums2) {
		return false
	}
	for i := range nums1 {
		if nums1[i] != nums2[i] {
			return false
		}
	}
	return true
}

func main() {
	timeStartWholeProgram := time.Now()

	testInput := []TestCase{

		{
			Nums: []int{1, 2, 3},
			Result: `
[1,3,2]

            `,
		},
		{
			Nums: []int{3, 2, 1},
			Result: `
      [1,2,3]
            `,
		},
		{
			Nums: []int{1, 1, 5},
			Result: `
      [1,5,1]
            `,
		},
	}
	for count, value := range testInput {
		fmt.Println("===============")
		fmt.Println("Test count ", count, "for node", value)
		fmt.Println("Solution 1: ")
		timeStart := time.Now()
		nextPermutation(value.Nums)
		timeLapse := time.Since(timeStart)
		fmt.Println(">Solution result")
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

		fmt.Println("Solution 2: Optimize a little bit")
		timeStart = time.Now()
		nextPermutation_Approach2(value.Nums)
		timeLapse = time.Since(timeStart)
		fmt.Println(">Solution result")
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

		fmt.Println("Solution 2: BruteForce, failure solution, need check")
		timeStart = time.Now()
		nextPermutation_BruteForce(value.Nums)
		timeLapse = time.Since(timeStart)
		fmt.Println(">Solution result")
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
Test count  0 for node {[1 2 3] 
[1,3,2]

            }
Solution 1: 
>Solution result
Correct result is  
[1,3,2]

            
TimeLapse 1.018µs
Solution 2: Optimize a little bit
>Solution result
Correct result is  
[1,3,2]

            
TimeLapse 593ns
Solution 2: BruteForce, failure solution, need check
>Solution result
Correct result is  
[1,3,2]

            
TimeLapse 34.832µs
===============
Test count  1 for node {[3 2 1] 
      [1,2,3]
            }
Solution 1: 
>Solution result
Correct result is  
      [1,2,3]
            
TimeLapse 277ns
Solution 2: Optimize a little bit
>Solution result
Correct result is  
      [1,2,3]
            
TimeLapse 407ns
Solution 2: BruteForce, failure solution, need check
>Solution result
Correct result is  
      [1,2,3]
            
TimeLapse 4.148µs
===============
Test count  2 for node {[1 1 5] 
      [1,5,1]
            }
Solution 1: 
>Solution result
Correct result is  
      [1,5,1]
            
TimeLapse 297ns
Solution 2: Optimize a little bit
>Solution result
Correct result is  
      [1,5,1]
            
TimeLapse 297ns
Solution 2: BruteForce, failure solution, need check
>Solution result
Correct result is  
      [1,5,1]
            
TimeLapse 4.611µs
===============
TimeLapse Whole Program 665.373µs

 */
//REF
//
