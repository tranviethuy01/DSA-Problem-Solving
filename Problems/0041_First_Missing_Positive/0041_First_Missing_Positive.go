package main

import (
	"fmt"
	"time"
)

//approach
/*
The algorithm used here is a variation of cycling sort. It iterates through the array, placing each number in its correct position. Then, it checks for the first position where the number doesn't match its index, indicating the smallest missing positive integer.

Let's analyze the time and space complexity:

Time Complexity:

The algorithm involves two passes through the array:
The first pass iterates through the array to place each number in its correct position, which takes O(n) time.
The second pass iterates through the array to find the first position where the number doesn't match its index, which also takes O(n) time.
Thus, the overall time complexity is O(n).
Space Complexity:

The algorithm uses only a constant amount of additional space for variables, regardless of the size of the input array. Hence, the space complexity is O(1).
Therefore, the algorithm meets the requirements specified in the problem statement, running in O(n) time and using O(1) auxiliary space.
*/

func firstMissingPositive_Approach1(nums []int) int {
	n := len(nums)
	for i := 0; i < n; {
		if nums[i] > 0 && nums[i] <= n && nums[nums[i]-1] != nums[i] {
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
		} else {
			i++
		}
	}
	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return n + 1
}

//approach 2: simulate the code here
/*
 def firstMissingPositive(self, nums):
    """
    :type nums: List[int]
    :rtype: int
     Basic idea:
    1. for any array whose length is l, the first missing positive must be in range [1,...,l+1],
        so we only have to care about those elements in this range and remove the rest.
    2. we can use the array index as the hash to restore the frequency of each number within
         the range [1,...,l+1]
    """
    nums.append(0)
    n = len(nums)
    for i in range(len(nums)): #delete those useless elements
        if nums[i]<0 or nums[i]>=n:
            nums[i]=0
    for i in range(len(nums)): #use the index as the hash to record the frequency of each number
        nums[nums[i]%n]+=n
    for i in range(1,len(nums)):
        if nums[i]/n==0:
            return i
    return n
*/

func firstMissingPositive_Approach2(nums []int) int {
	nums = append(nums, 0)
	n := len(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 || nums[i] >= n {
			nums[i] = 0
		}
	}
	for i := 0; i < len(nums); i++ {
		nums[nums[i]%n] += n
	}
	for i := 1; i < len(nums); i++ {
		if nums[i]/n == 0 {
			return i
		}
	}
	return n
}

func main() {
	timeStartWholeProgram := time.Now()

	testInput := []TestCase{
		{
			Nums: []int{1, 2, 0},
			Result: `
    3
            `,
		},
		{
			Nums: []int{3, 4, -1, 1},
			Result: `
    2
            `,
		},
		{
			Nums: []int{7, 8, 9, 11, 12},
			Result: `
    1
            `,
		},
	}
	for count, value := range testInput {

		fmt.Println("===============")
		fmt.Println("Test count ", count, "for node", value)
		fmt.Println("Solution 1: ")
		timeStart := time.Now()
		result := firstMissingPositive_Approach1(value.Nums)
		timeLapse := time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

		fmt.Println("Solution 2: ")
		timeStart = time.Now()
		result = firstMissingPositive_Approach2(value.Nums)
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
Test count  0 for node {[1 2 0]
    3
            }
Solution 1:
>Solution result 3
Correct result is
    3

TimeLapse 463ns
Solution 2:
>Solution result 3
Correct result is
    3

TimeLapse 1.037µs
===============
Test count  1 for node {[3 4 -1 1]
    2
            }
Solution 1:
>Solution result 2
Correct result is
    2

TimeLapse 297ns
Solution 2:
>Solution result 2
Correct result is
    2

TimeLapse 648ns
===============
Test count  2 for node {[7 8 9 11 12]
    1
            }
Solution 1:
>Solution result 1
Correct result is
    1

TimeLapse 186ns
Solution 2:
>Solution result 1
Correct result is
    1

TimeLapse 574ns
===============
TimeLapse Whole Program 467.534µs

*/
