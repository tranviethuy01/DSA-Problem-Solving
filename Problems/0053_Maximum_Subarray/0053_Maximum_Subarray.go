package main

import (
	"fmt"
	"runtime"
	"time"
)

// NOTE: choose the Kadane's algorithm approach (dynamic programming)

// approach dynamic programing  :  Kadane's algorithm
/*

Time Complexity: The algorithm iterates through the input array once, visiting each element exactly once. Therefore, the time complexity is O(n), where n is the number of elements in the input array.

Space Complexity: The algorithm uses only a constant amount of additional space for variables such as maxSum and currentSum, regardless of the size of the input array. Hence, the space complexity is O(1), constant space complexity.


*/

func maxSubArray_DP(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxSum := nums[0]
	currentSum := nums[0]

	for i := 1; i < len(nums); i++ {
		currentSum = max(nums[i], currentSum+nums[i])
		maxSum = max(maxSum, currentSum)
	}

	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//approach : DivideAndConquer
/*
Time Complexity: The divide and conquer approach recursively divides the array into halves until the subarray size becomes 1. At each level of recursion, it performs O(n) operations to find the maximum subarray crossing the midpoint. Since the recursion depth is log n (where n is the number of elements in the input array), and at each level, it performs O(n) operations, the overall time complexity is O(n log n).

Space Complexity: The space complexity is determined by the recursive calls made during the divide and conquer process. Since the algorithm divides the array into halves recursively, the maximum recursion depth is log n. At each level of recursion, some additional space is required for function call stack frames. Therefore, the space complexity is O(log n).
*/

func maxSubArray_DivideAndConquer(nums []int) int {
	return maxSubArrayHelper(nums, 0, len(nums)-1)
}

func maxSubArrayHelper(nums []int, left, right int) int {
	if left == right {
		return nums[left]
	}

	mid := (left + right) / 2

	leftMax := maxSubArrayHelper(nums, left, mid)
	rightMax := maxSubArrayHelper(nums, mid+1, right)
	crossMax := maxCrossingSum(nums, left, mid, right)

	return max(leftMax, max(rightMax, crossMax))
}

func maxCrossingSum(nums []int, left, mid, right int) int {
	leftSum := -2147483648 // -infinity
	sum := 0
	for i := mid; i >= left; i-- {
		sum += nums[i]
		if sum > leftSum {
			leftSum = sum
		}
	}

	rightSum := -2147483648 // -infinity
	sum = 0
	for i := mid + 1; i <= right; i++ {
		sum += nums[i]
		if sum > rightSum {
			rightSum = sum
		}
	}

	return leftSum + rightSum
}

//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}
//

//approach BruteForce
/*
Algorithm: The algorithm iterates through every possible subarray of the input array and calculates the sum of each subarray. It then compares each sum to find the maximum one, thus determining the maximum sum subarray.

Time Complexity: The outer loop iterates through each starting index of the subarray, and the inner loop iterates through each ending index. For an array of size
n, there are
O(n ^ 2 ) subarrays to consider. Within each subarray, the sum calculation is constant time. Therefore, the overall time complexity is
O(n ^ 2), where
n is the number of elements in the input array.
Space Complexity: The algorithm uses only a constant amount of additional space for variables such as maxSum, currentSum, i, and j, regardless of the size of the input array. Therefore, the space complexity is
O(1), constant space complexity.

*/

func maxSubArray_BruteForce(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxSum := nums[0]

	for i := 0; i < len(nums); i++ {
		currentSum := 0
		for j := i; j < len(nums); j++ {
			currentSum += nums[j]
			if currentSum > maxSum {
				maxSum = currentSum
			}
		}
	}

	return maxSum
}

// approach Recursively

/*
Algorithm: The solve function is a recursive function that explores all possible subarrays of the input array and calculates the sum of each subarray. It returns the maximum sum encountered.

Time Complexity: Let n be the number of elements in the input array. The solve function is called recursively for each element in the array. At each recursive call, there are two branches, representing whether the current element is included in the sum or not. This results in a binary tree of recursive calls with a maximum depth of
n, corresponding to exploring all possible combinations of including or excluding each element. Therefore, the time complexity is
O(2 ^ n), which is exponential.
Space Complexity: The space complexity is determined by the maximum depth of the recursion stack. Since each recursive call consumes constant space and the maximum depth of the recursion tree is
n, the space complexity is
O(n).

*/

func maxSubArray_Recursive(nums []int) int {
	return solve(nums, 0, false)
}

func solve(A []int, i int, mustPick bool) int {
	if i >= len(A) {
		if mustPick {
			return 0
		}
		return -1e5
	}
	if mustPick {
		return max(0, A[i]+solve(A, i+1, true))
	}
	return max(solve(A, i+1, false), A[i]+solve(A, i+1, true))
}

//
//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}

//NOTE: Depth-First Search (DFS) is primarily used for traversing graphs, it is not the most suitable approach for solving the maximum subarray problem directly. The maximum subarray problem typically involves finding the contiguous subarray within a one-dimensional array of numbers that has the largest sum.

//DFS is better suited for exploring all possible paths in a graph, which may not directly translate to finding the maximum sum subarray efficiently.
//the code below just use for education purpose, think about it as a failure note
/*
package main

import "fmt"

type Node struct {
	Value int
	Children []*Node
}

func maxSumPathDFS(node *Node, currentSum int, maxSum *int) {
	if node == nil {
		return
	}
	// Update current sum by adding current node's value
	currentSum += node.Value
	// Update maxSum if currentSum is greater
	if currentSum > *maxSum {
		*maxSum = currentSum
	}
	// Recursively explore children nodes
	for _, child := range node.Children {
		maxSumPathDFS(child, currentSum, maxSum)
	}
}

func maxSumPath(root *Node) int {
	if root == nil {
		return 0
	}
	maxSum := root.Value // Initialize maxSum with root value
	maxSumPathDFS(root, 0, &maxSum)
	return maxSum
}

func main() {
	root := &Node{
		Value: 1,
		Children: []*Node{
			{Value: -2},
			{Value: 3},
			{Value: 4},
		},
	}
	fmt.Println("Maximum sum path:", maxSumPath(root)) // Output: 6
}


*/

func main() {
	timeStartWholeProgram := time.Now()

	testInput := []TestCase{
		{
			Nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			Result: `
6

            `,
		},
		{
			Nums: []int{1},
			Result: `
1
            `,
		},
		{
			Nums: []int{5, 4, -1, 7, 8},
			Result: `
23
            `,
		},
	}

	// Memory before allocation
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	memBefore := m.Alloc

	for count, value := range testInput {
		fmt.Println("===============")
		fmt.Println("Test count ", count, "for node", value)
		fmt.Println("Solution 1: Kadane - Dynamic Programming")
		timeStart := time.Now()
		result := maxSubArray_DP(value.Nums)
		timeLapse := time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

		// Memory after allocation
		runtime.ReadMemStats(&m)
		memAfter := m.Alloc
		fmt.Println("Memory before", memBefore, "bytes", "Memory after", memAfter, "bytes", "Memory used:", memAfter-memBefore, "bytes")

		fmt.Printf("Memory usage (HeapAlloc) after Test Case i %d, : %v bytes\n", count, m.HeapAlloc)

		fmt.Println("Solution 2: DivideAndConquer ")
		timeStart = time.Now()
		result = maxSubArray_DivideAndConquer(value.Nums)
		timeLapse = time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

		// Memory after allocation
		runtime.ReadMemStats(&m)
		memAfter = m.Alloc
		fmt.Println("Memory before", memBefore, "bytes", "Memory after", memAfter, "bytes", "Memory used:", memAfter-memBefore, "bytes")

		fmt.Printf("Memory usage (HeapAlloc) after Test Case i %d, : %v bytes\n", count, m.HeapAlloc)

		fmt.Println("Solution 3: Recursively")
		timeStart = time.Now()
		result = maxSubArray_Recursive(value.Nums)
		timeLapse = time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

		// Memory after allocation
		runtime.ReadMemStats(&m)
		memAfter = m.Alloc
		fmt.Println("Memory before", memBefore, "bytes", "Memory after", memAfter, "bytes", "Memory used:", memAfter-memBefore, "bytes")

		fmt.Printf("Memory usage (HeapAlloc) after Test Case i %d, : %v bytes\n", count, m.HeapAlloc)

		fmt.Println("Solution 4: BruteForce")
		timeStart = time.Now()
		result = maxSubArray_BruteForce(value.Nums)
		timeLapse = time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

		// Memory after allocation
		runtime.ReadMemStats(&m)
		memAfter = m.Alloc
		fmt.Println("Memory before", memBefore, "bytes", "Memory after", memAfter, "bytes", "Memory used:", memAfter-memBefore, "bytes")

		fmt.Printf("Memory usage (HeapAlloc) after Test Case i %d, : %v bytes\n", count, m.HeapAlloc)

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
Test count  0 for node {[-2 1 -3 4 -1 2 1 -5 4]
6

            }
Solution 1: Kadane - Dynamic Programming
>Solution result 6
Correct result is
6


TimeLapse 777ns
Memory before 69392 bytes Memory after 70520 bytes Memory used: 1128 bytes
Memory usage (HeapAlloc) after Test Case i 0, : 70520 bytes
Solution 2: DivideAndConquer
>Solution result 6
Correct result is
6


TimeLapse 1.13µs
Memory before 69392 bytes Memory after 70584 bytes Memory used: 1192 bytes
Memory usage (HeapAlloc) after Test Case i 0, : 70584 bytes
Solution 3: Recursively
>Solution result 6
Correct result is
6


TimeLapse 1.352µs
Memory before 69392 bytes Memory after 70648 bytes Memory used: 1256 bytes
Memory usage (HeapAlloc) after Test Case i 0, : 70648 bytes
Solution 4: BruteForce
>Solution result 6
Correct result is
6


TimeLapse 685ns
Memory before 69392 bytes Memory after 70712 bytes Memory used: 1320 bytes
Memory usage (HeapAlloc) after Test Case i 0, : 70712 bytes
===============
Test count  1 for node {[1]
1
            }
Solution 1: Kadane - Dynamic Programming
>Solution result 1
Correct result is
1

TimeLapse 204ns
Memory before 69392 bytes Memory after 70824 bytes Memory used: 1432 bytes
Memory usage (HeapAlloc) after Test Case i 1, : 70824 bytes
Solution 2: DivideAndConquer
>Solution result 1
Correct result is
1

TimeLapse 223ns
Memory before 69392 bytes Memory after 70888 bytes Memory used: 1496 bytes
Memory usage (HeapAlloc) after Test Case i 1, : 70888 bytes
Solution 3: Recursively
>Solution result 1
Correct result is
1

TimeLapse 408ns
Memory before 69392 bytes Memory after 70952 bytes Memory used: 1560 bytes
Memory usage (HeapAlloc) after Test Case i 1, : 70952 bytes
Solution 4: BruteForce
>Solution result 1
Correct result is
1

TimeLapse 241ns
Memory before 69392 bytes Memory after 71016 bytes Memory used: 1624 bytes
Memory usage (HeapAlloc) after Test Case i 1, : 71016 bytes
===============
Test count  2 for node {[5 4 -1 7 8]
23
            }
Solution 1: Kadane - Dynamic Programming
>Solution result 23
Correct result is
23

TimeLapse 222ns
Memory before 69392 bytes Memory after 71160 bytes Memory used: 1768 bytes
Memory usage (HeapAlloc) after Test Case i 2, : 71160 bytes
Solution 2: DivideAndConquer
>Solution result 23
Correct result is
23

TimeLapse 426ns
Memory before 69392 bytes Memory after 71224 bytes Memory used: 1832 bytes
Memory usage (HeapAlloc) after Test Case i 2, : 71224 bytes
Solution 3: Recursively
>Solution result 23
Correct result is
23

TimeLapse 556ns
Memory before 69392 bytes Memory after 71288 bytes Memory used: 1896 bytes
Memory usage (HeapAlloc) after Test Case i 2, : 71288 bytes
Solution 4: BruteForce
>Solution result 23
Correct result is
23

TimeLapse 333ns
Memory before 69392 bytes Memory after 71352 bytes Memory used: 1960 bytes
Memory usage (HeapAlloc) after Test Case i 2, : 71352 bytes
===============
TimeLapse Whole Program 1.730841ms


*/
//REF
//
