package main

import (
	"fmt"
	"sort"
	"time"
  "math"
)

// approach: 2 pointer
//Time complexity: O(nlogn)
//Space complexity: O(1)
func threeSumClosest_2Pointer(nums []int, target int) int {
    // Sort the array
    sort.Ints(nums)
    
    // Initialize variables for result and minimum difference
    closestSum := math.MaxInt32
    minDiff := math.MaxInt32
    
    // Iterate over the array
    for i := 0; i < len(nums)-2; i++ {
        left, right := i+1, len(nums)-1
        
        // Use two pointers approach
        for left < right {
            sum := nums[i] + nums[left] + nums[right]
            diff := abs(sum - target)
            
            // Update closestSum and minDiff if current sum is closer to target
            if diff < minDiff {
                closestSum = sum
                minDiff = diff
            }
            
            // Move pointers based on the comparison with target
            if sum < target {
                left++
            } else if sum > target {
                right--
            } else {
                return sum // Exact match found
            }
        }
    }
    
    return closestSum
}


func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}



//==== hash map approach
//NOTE: need to check later
/*
func threeSumClosest_HashMap(nums []int, target int) int {
    // Sort the array
    sort.Ints(nums)
    
    // Initialize variables for result and minimum difference
    closestSum := math.MaxInt32
    minDiff := math.MaxInt32
    
    // Iterate over the array
    for i := 0; i < len(nums)-2; i++ {
        // Use hash map to store differences between target and sum of two elements
        diffMap := make(map[int]int)
        for j := i+1; j < len(nums); j++ {
            diff := target - (nums[i] + nums[j])
            if _, ok := diffMap[diff]; ok {
                // Found a sum closer to target
                sum := target - diff
                if sum < closestSum {
                    closestSum = sum
                    minDiff = diff
                }
            }
            diffMap[nums[j]] = j
        }
    }
    
    return closestSum
}

*/

//====


func main() {
	timeStartWholeProgram := time.Now()
	testInput := []TestCase{
		{
			Nums:   []int{-1, 2, 1, -4},
			Target: 1,
			Result: `
          2
            `,
		},
		{
			Nums:   []int{0, 0, 0},
			Target: 1,
			Result: `
         0
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
		fmt.Println("Solution 1: use 2 pointer")
		timeStart := time.Now()
		result := threeSumClosest_2Pointer(value.Nums, value.Target)
		timeLapse := time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

    /*
		fmt.Println("Solution 2: use hash map")
		timeStart = time.Now()
		result = threeSumClosest_HashMap(value.Nums, value.Target)
		timeLapse = time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)
    */
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



 */
//REF
//
