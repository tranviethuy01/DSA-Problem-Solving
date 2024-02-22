package main

import (
	"fmt"
	"sort"
	"time"
)


//NOTE: use solution targetIndices_SingleLoop, which have time complexity is O(n) and loop one. 


//approach: this solution fail, need check code
//Time Complexity: dominated by the sorting step, which is O(n log n)
// space complexity is O(n).

func targetIndices(nums []int, target int) []int {
	// Create a map to store indices of target elements
	targetIndicesMap := make(map[int][]int)

	// Traverse the nums array and store indices of target elements
	for i, num := range nums {
		if num == target {
			targetIndicesMap[num] = append(targetIndicesMap[num], i)
		}
	}

	// If there are no target elements, return an empty list
	if len(targetIndicesMap[target]) == 0 {
		return []int{}
	}

	// Sort the nums array
	sort.Ints(nums)

	// Retrieve the sorted indices of target elements
	sortedIndices := targetIndicesMap[target]
	for i := range sortedIndices {
		sortedIndices[i] = indexOf(nums, target, sortedIndices[i])
	}

	// Sort the indices in increasing order
	sort.Ints(sortedIndices)

	return sortedIndices
}

// Function to find the index of the target element in the sorted nums array

func indexOf(nums []int, target int, originalIndex int) int {
	for i, num := range nums {
		if num == target {
			if originalIndex == 0 {
				return i
			}
			originalIndex--
		}
	}
	return -1
}

//approach : Linear approach
//
//Time Complexity: O(n log n).
//Space Complexity:

//The space complexity is determined by the additional space used for the ans slice to store the indices of target elements. In the worst case, if all elements in the array are target elements, this would require O(n) additional space.
//Additionally, the space complexity of sorting is O(log n) for the recursive call stack.
//Therefore, the overall space complexity is O(n).
//


func targetIndices_Linear(nums []int, target int) []int {
    // Sort the nums array
    sort.Ints(nums)
    
    // Initialize an empty list to store the indices of target elements
    var ans []int
    
    // Traverse the sorted array and store indices of target elements
    for i := 0; i < len(nums); i++ {
        if nums[i] == target {
            ans = append(ans, i)
        }
    }
    
    return ans
}

//approach:
//time complexity of the algorithm is O(n)
//Space Complexity:
//The space complexity is determined by the additional space used for the result slice, which has a maximum size of the count of occurrences of the target element. In the worst case, if all elements are equal to the target, the result slice will have a length of n (the length of the nums array).
//Additionally, the space complexity for other variables like count and lessthan is constant and doesn't depend on the input size.
//Therefore, the overall space complexity of the algorithm is O(n).

func targetIndices_SingleLoop(nums []int, target int) []int {
    count := 0
    lessthan := 0
    
    // Count the number of occurrences of target and elements less than target
    for _, n := range nums {
        if n == target {
            count++
        }
        if n < target {
            lessthan++
        }
    }
    
    // Initialize a slice to store the result
    result := make([]int, count)
    
    // Populate the result slice with indices less than target
    for i := 0; i < count; i++ {
        result[i] = lessthan
        lessthan++
    }
    
    return result
}




func main() {
	timeStartWholeProgram := time.Now()

	testInput := []TestCase{

		{
			Nums:   []int{1, 2, 5, 2, 3},
			Target: 2,
			Result: `
      [1,2]
            `,
		},

		{
			Nums:   []int{1, 2, 5, 2, 3},
			Target: 3,
			Result: `
      [3]
            `,
		},

		{
			Nums:   []int{1, 2, 5, 2, 3},
			Target: 5,
			Result: `
      [4]
            `,
		},
	}
	for count, value := range testInput {
		fmt.Println("===============")
		fmt.Println("Test count ", count, "for node", value)
		fmt.Println("Solution 1: note, this solution failure, need check code again")
		timeStart := time.Now()
		result := targetIndices(value.Nums, value.Target)
		timeLapse := time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

		fmt.Println("Solution 2: Linear approach")
		timeStart = time.Now()
		result = targetIndices_Linear(value.Nums, value.Target)
		timeLapse = time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)


		fmt.Println("Solution 3: a Single loop approach")
		timeStart = time.Now()
		result = targetIndices_SingleLoop(value.Nums, value.Target)
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
Test count  0 for node {[1 2 5 2 3] 2 
      [1,2]
            }
Solution 1: note, this solution failure, need check code again
>Solution result [-1 2]
Correct result is  
      [1,2]
            
TimeLapse 6.778µs
Solution 2: Linear approach
>Solution result [1 2]
Correct result is  
      [1,2]
            
TimeLapse 1.722µs
Solution 3: a Single loop approach
>Solution result [1 2]
Correct result is  
      [1,2]
            
TimeLapse 685ns
===============
Test count  1 for node {[1 2 5 2 3] 3 
      [3]
            }
Solution 1: note, this solution failure, need check code again
>Solution result [-1]
Correct result is  
      [3]
            
TimeLapse 1.722µs
Solution 2: Linear approach
>Solution result [3]
Correct result is  
      [3]
            
TimeLapse 1.111µs
Solution 3: a Single loop approach
>Solution result [3]
Correct result is  
      [3]
            
TimeLapse 389ns
===============
Test count  2 for node {[1 2 5 2 3] 5 
      [4]
            }
Solution 1: note, this solution failure, need check code again
>Solution result [-1]
Correct result is  
      [4]
            
TimeLapse 1.63µs
Solution 2: Linear approach
>Solution result [4]
Correct result is  
      [4]
            
TimeLapse 1.018µs
Solution 3: a Single loop approach
>Solution result [4]
Correct result is  
      [4]
            
TimeLapse 259ns
===============
TimeLapse Whole Program 646.386µs


 */
