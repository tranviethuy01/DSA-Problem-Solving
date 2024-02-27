package main

import (
	"fmt"
	"runtime"
	"time"
)

//approach

/*

Time Complexity:

The rotateMatrix function rotates the matrix by 90 degrees clockwise. It iterates through each element of the matrix once and performs constant-time operations inside the nested loops. So, the time complexity of rotateMatrix is O(n^2), where n is the size of the matrix.
The matricesEqual function compares two matrices element by element. Again, it iterates through each element of the matrices once and performs constant-time operations inside the nested loops. So, the time complexity of matricesEqual is also O(n^2).
The findRotation function calls rotateMatrix function at most 4 times, and for each rotation, it calls matricesEqual function. So, the time complexity of findRotation is O(4 * n^2), which simplifies to O(n^2).
The findRotationEqual function calls findRotation function twice and does not involve any additional loops. So, the time complexity of findRotationEqual is also O(n^2).
Space Complexity:

The rotateMatrix function creates a new matrix to store the rotated matrix. This matrix has the same size as the input matrix. Thus, the space complexity of rotateMatrix is O(n^2).
The space complexity of matricesEqual, findRotation, and findRotationEqual functions is negligible as they only use a constant amount of extra space for variables.

*/

func rotateMatrix(matrix [][]int) [][]int {
	n := len(matrix)
	rotated := make([][]int, n)
	for i := 0; i < n; i++ {
		rotated[i] = make([]int, n)
		for j := 0; j < n; j++ {
			rotated[i][j] = matrix[n-j-1][i]
		}
	}
	return rotated
}

func matricesEqual(mat1, mat2 [][]int) bool {
	for i, row := range mat1 {
		for j, val := range row {
			if val != mat2[i][j] {
				return false
			}
		}
	}
	return true
}

func findRotation(mat [][]int, target [][]int) bool {
	for i := 0; i < 4; i++ {
		if matricesEqual(mat, target) {
			return true
		}
		mat = rotateMatrix(mat)
	}
	return false
}

func findRotationEqual(mat [][]int, target [][]int) bool {
	return findRotation(mat, target) || findRotation(rotateMatrix(mat), target)
}

func main() {
	timeStartWholeProgram := time.Now()

	testInput := []TestCase{
		{
			Matrix: [][]int{{0, 1}, {1, 0}},
			Target: [][]int{{1, 0}, {0, 1}},
			Result: `
true
            `,
		},
		{
			Matrix: [][]int{{0, 1}, {1, 1}},
			Target: [][]int{{1, 0}, {0, 1}},
			Result: `
false
            `,
		},
		{
			Matrix: [][]int{{0, 0, 0}, {0, 1, 0}, {1, 1, 1}},
			Target: [][]int{{1, 1, 1}, {0, 1, 0}, {0, 0, 0}},
			Result: `
      true
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
		fmt.Println("Solution 1: StraightForward")
		timeStart := time.Now()
		result := findRotation(value.Matrix, value.Target)
		timeLapse := time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

		// Memory after allocation
		runtime.ReadMemStats(&m)
		memAfter := m.Alloc
		fmt.Println("Memory before", memBefore, "bytes", "Memory after", memAfter, "bytes", "Memory used:", memAfter-memBefore, "bytes")
	
    fmt.Printf("Memory usage (HeapAlloc) after Test Case i %d, : %v bytes\n", count,  m.HeapAlloc)


	}

	timeLapsedWholeProgram := time.Since(timeStartWholeProgram)
	fmt.Println("===============")
	fmt.Println("TimeLapse Whole Program", timeLapsedWholeProgram)
}

type TestCase struct {
	Matrix [][]int
	Target [][]int
	Result string
}

/*

===============
Test count  0 for node {[[0 1] [1 0]] [[1 0] [0 1]] 
true
            }
Solution 1: StraightForward
>Solution result true
Correct result is  
true
            
TimeLapse 2.703µs
Memory before 67856 bytes Memory after 69160 bytes Memory used: 1304 bytes
Memory usage (HeapAlloc) after Test Case i 0, : 69160 bytes
===============
Test count  1 for node {[[0 1] [1 1]] [[1 0] [0 1]] 
false
            }
Solution 1: StraightForward
>Solution result false
Correct result is  
false
            
TimeLapse 2.87µs
Memory before 67856 bytes Memory after 69768 bytes Memory used: 1912 bytes
Memory usage (HeapAlloc) after Test Case i 1, : 69768 bytes
===============
Test count  2 for node {[[0 0 0] [0 1 0] [1 1 1]] [[1 1 1] [0 1 0] [0 0 0]] 
      true
            }
Solution 1: StraightForward
>Solution result true
Correct result is  
      true
            
TimeLapse 3.037µs
Memory before 67856 bytes Memory after 70488 bytes Memory used: 2632 bytes
Memory usage (HeapAlloc) after Test Case i 2, : 70488 bytes
===============
TimeLapse Whole Program 812.277µs

*/
//REF
//
