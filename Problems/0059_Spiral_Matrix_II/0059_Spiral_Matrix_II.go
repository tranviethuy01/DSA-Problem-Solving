package main

import (
	"fmt"
	"runtime"
	"time"
)

//approach : StraightForward => Iterate through each cell => Brute Force solution
/*
Time Complexity: The algorithm iterates through each cell in the matrix once to fill in the values. Since the matrix size is n x n, the total number of cells is n^2. Therefore, the time complexity is O(n^2).

Space Complexity: The space complexity is determined by the space required to store the generated matrix. Since the matrix size is n x n, it requires O(n^2) space to store all the elements.
*/

func generateMatrix_StraightForward(n int) [][]int {
	// Create an empty matrix
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	// Define boundaries
	top, bottom, left, right := 0, n-1, 0, n-1
	num := 1

	// Generate matrix in spiral order
	for num <= n*n {
		// Move from left to right
		for i := left; i <= right; i++ {
			matrix[top][i] = num
			num++
		}
		top++

		// Move from top to bottom
		for i := top; i <= bottom; i++ {
			matrix[i][right] = num
			num++
		}
		right--

		// Move from right to left
		for i := right; i >= left; i-- {
			matrix[bottom][i] = num
			num++
		}
		bottom--

		// Move from bottom to top
		for i := bottom; i >= top; i-- {
			matrix[i][left] = num
			num++
		}
		left++
	}

	return matrix
}

func main() {
	timeStartWholeProgram := time.Now()

	testInput := []TestCase{
		{
			N: 3,
			Result: `
[[1,2,3],[8,9,4],[7,6,5]]
            `,
		},
		{
			N: 1,
			Result: `
[[1]]

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
		fmt.Println("Solution 1: StraightForward | Brute Force Solution")
		timeStart := time.Now()
		result := generateMatrix_StraightForward(value.N)
		timeLapse := time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

		// Memory after allocation
		runtime.ReadMemStats(&m)
		memAfter := m.Alloc
		fmt.Println("Memory before", memBefore, "bytes", "Memory after", memAfter, "bytes", "Memory used:", memAfter-memBefore, "bytes")

		fmt.Printf("Memory usage (HeapAlloc) after Test Case i %d, : %v bytes\n", count, m.HeapAlloc)
	}

	timeLapsedWholeProgram := time.Since(timeStartWholeProgram)
	fmt.Println("===============")
	fmt.Println("TimeLapse Whole Program", timeLapsedWholeProgram)
}

type TestCase struct {
	N      int
	Result string
}

/*

===============
Test count  0 for node {3 
[[1,2,3],[8,9,4],[7,6,5]]
            }
Solution 1: StraightForward | Brute Force Solution
>Solution result [[1 2 3] [8 9 4] [7 6 5]]
Correct result is  
[[1,2,3],[8,9,4],[7,6,5]]
            
TimeLapse 3.778µs
Memory before 69248 bytes Memory after 70488 bytes Memory used: 1240 bytes
Memory usage (HeapAlloc) after Test Case i 0, : 70488 bytes
===============
Test count  1 for node {1 
[[1]]

            }
Solution 1: StraightForward | Brute Force Solution
>Solution result [[1]]
Correct result is  
[[1]]

            
TimeLapse 833ns
Memory before 69248 bytes Memory after 70664 bytes Memory used: 1416 bytes
Memory usage (HeapAlloc) after Test Case i 1, : 70664 bytes
===============
TimeLapse Whole Program 820.035µs

 */
//REF
//
