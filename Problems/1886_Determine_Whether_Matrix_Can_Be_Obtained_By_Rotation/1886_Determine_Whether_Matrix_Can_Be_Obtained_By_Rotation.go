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

//approach : adapt leetcode solution
/*
bool findRotation(vector<vector<int>>& mat, vector<vector<int>>& target)
{
	bool c[4];
	memset(c,true,sizeof(c));
	int n=mat.size();
	for(int i=0;i<n;i++)
	{
		for(int j=0;j<n;j++)
		{
			if(mat[i][j]!=target[i][j]) c[0]=false;
			if(mat[i][j]!=target[n-j-1][i]) c[1]=false;
			if(mat[i][j]!=target[n-i-1][n-j-1]) c[2]=false;
			if(mat[i][j]!=target[j][n-i-1]) c[3]=false;
		}
	}
	return c[0]||c[1]||c[2]||c[3];
}
*/

func findRotation_Approach2(mat [][]int, target [][]int) bool {
	c := [4]bool{true, true, true, true}
	n := len(mat)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] != target[i][j] {
				c[0] = false
			}
			if mat[i][j] != target[n-j-1][i] {
				c[1] = false
			}
			if mat[i][j] != target[n-i-1][n-j-1] {
				c[2] = false
			}
			if mat[i][j] != target[j][n-i-1] {
				c[3] = false
			}
		}
	}
	return c[0] || c[1] || c[2] || c[3]
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

		fmt.Printf("Memory usage (HeapAlloc) after Test Case i %d, : %v bytes\n", count, m.HeapAlloc)

		fmt.Println("Solution 2: adapt leetcode solution")
		timeStart = time.Now()
		result = findRotation_Approach2(value.Matrix, value.Target)
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

TimeLapse 3.092µs
Memory before 69664 bytes Memory after 70968 bytes Memory used: 1304 bytes
Memory usage (HeapAlloc) after Test Case i 0, : 70968 bytes
Solution 2: adapt leetcode solution
>Solution result true
Correct result is
true

TimeLapse 888ns
===============
Test count  1 for node {[[0 1] [1 1]] [[1 0] [0 1]]
false
            }
Solution 1: StraightForward
>Solution result false
Correct result is
false

TimeLapse 3.259µs
Memory before 69664 bytes Memory after 71608 bytes Memory used: 1944 bytes
Memory usage (HeapAlloc) after Test Case i 1, : 71608 bytes
Solution 2: adapt leetcode solution
>Solution result false
Correct result is
false

TimeLapse 444ns
===============
Test count  2 for node {[[0 0 0] [0 1 0] [1 1 1]] [[1 1 1] [0 1 0] [0 0 0]]
      true
            }
Solution 1: StraightForward
>Solution result true
Correct result is
      true

TimeLapse 3.445µs
Memory before 69664 bytes Memory after 72360 bytes Memory used: 2696 bytes
Memory usage (HeapAlloc) after Test Case i 2, : 72360 bytes
Solution 2: adapt leetcode solution
>Solution result true
Correct result is
      true

TimeLapse 722ns
===============
TimeLapse Whole Program 982.532µs
*/
//REF
//
