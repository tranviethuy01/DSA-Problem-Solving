package main

import (
	"fmt"
	"time"
)

//approach StraightForward
/*
Time Complexity:
The algorithm iterates through each element of the matrix once, which gives O(n^2) operations.
Within the nested loops, the operations performed are constant time operations (e.g., map lookups, assignments), so each iteration contributes O(1) to the overall time complexity.
Therefore, the total time complexity is O(n^2).
Space Complexity:

Two maps are created, one for rows and one for columns, each with a size of n.
Thus, the space required for these maps is O(n) each.
Since there are two such maps and they both scale with the size of the input matrix, the overall space complexity is O(n^2).
*/
func checkValid_StraightForward(matrix [][]int) bool {
	n := len(matrix)

	// Create maps to track occurrences of numbers in rows and columns
	rowMap := make([]map[int]bool, n)
	colMap := make([]map[int]bool, n)

	// Initialize maps
	for i := 0; i < n; i++ {
		rowMap[i] = make(map[int]bool)
		colMap[i] = make(map[int]bool)
	}

	// Iterate through the matrix
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			num := matrix[i][j]
			// Check if the number is within valid range
			if num < 1 || num > n {
				return false
			}
			// Check if the number already exists in the current row or column
			if rowMap[i][num] || colMap[j][num] {
				return false
			}
			// Mark the number as seen in the current row and column
			rowMap[i][num] = true
			colMap[j][num] = true
		}
	}

	// Check if every row and column contains all numbers from 1 to n
	for i := 0; i < n; i++ {
		for num := 1; num <= n; num++ {
			if !rowMap[i][num] || !colMap[i][num] {
				return false
			}
		}
	}

	return true
}

// approach DFS : failure, need check code
func checkValid_DFS(matrix [][]int) bool {
	n := len(matrix)
	visited := make([][]bool, n)

	// Initialize visited matrix
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	// Start DFS from the top-left corner
	return dfs(matrix, visited, 0, 0, 1, n)
}

func dfs(matrix [][]int, visited [][]bool, row, col, num, n int) bool {
	// Base case: if the row or column index is out of bounds or the current number is incorrect
	if row < 0 || row >= n || col < 0 || col >= n || visited[row][col] || matrix[row][col] != num {
		return false
	}

	// Mark the current cell as visited
	visited[row][col] = true

	// If it's the last number in the matrix, return true
	if num == n*n {
		return true
	}

	// Check in all four directions
	directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for _, dir := range directions {
		newRow, newCol := row+dir[0], col+dir[1]
		if dfs(matrix, visited, newRow, newCol, num+1, n) {
			return true
		}
	}

	// Backtrack: mark the current cell as unvisited
	visited[row][col] = false

	return false
}

//approach : ...
/*

 */

func checkValid_Approach2(matrix [][]int) bool {
	n := len(matrix)
	target := n * (n + 1) / 2 // Sum of integers from 1 to n

	// Check rows
	for i := 0; i < n; i++ {
		sum := 0
		for j := 0; j < n; j++ {
			sum += matrix[i][j]
		}
		if sum != target {
			return false
		}
	}

	// Check columns
	for j := 0; j < n; j++ {
		sum := 0
		for i := 0; i < n; i++ {
			sum += matrix[i][j]
		}
		if sum != target {
			return false
		}
	}

	return true
}

//approach : Backtrack
/*
Time Complexity:
The algorithm iterates through each element of the matrix once, which gives O(n^2) operations, where n is the size of the matrix.
Within the nested loops, the operations performed are constant time operations (e.g., array accesses, assignments), so each iteration contributes O(1) to the overall time complexity.
Therefore, the total time complexity is O(n^2).
Space Complexity:
The space complexity is also O(n^2) because two 2D arrays (rowVisited and colVisited) of size n x n are created to keep track of visited numbers in rows and columns.
Additional space is constant, as there are no other significant data structures used.
Thus, the overall space complexity is O(n^2).

*/
func checkValid_Backtrack(matrix [][]int) bool {
	n := len(matrix)
	rowVisited := make([][]bool, n)
	colVisited := make([][]bool, n)

	// Initialize visited matrix for rows and columns
	for i := range rowVisited {
		rowVisited[i] = make([]bool, n)
		colVisited[i] = make([]bool, n)
	}

	// Mark numbers in the matrix as visited based on rows and columns
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			num := matrix[i][j] - 1
			if rowVisited[i][num] || colVisited[j][num] {
				return false
			}
			rowVisited[i][num] = true
			colVisited[j][num] = true
		}
	}

	return true
}

//approach BruteForce
/*
Time Complexity:

The algorithm iterates through each element of the matrix twice: once for rows and once for columns. This gives O(2 * n^2) operations, where n is the size of the matrix.
Within the nested loops, the operations performed are constant time operations (e.g., array accesses, assignments), so each iteration contributes O(1) to the overall time complexity.
Therefore, the total time complexity is O(n^2).
Space Complexity:

The space complexity is O(n) for each row and O(n) for each column, as a boolean array of size n+1 (seen) is created to keep track of seen numbers.
Since the algorithm iterates through each row and column separately and there are no other significant data structures used, the overall space complexity is O(n) + O(n) = O(n).

*/

func checkValid_BruteForce(matrix [][]int) bool {
	n := len(matrix)

	// Check rows
	for i := 0; i < n; i++ {
		seen := make([]bool, n+1)
		for j := 0; j < n; j++ {
			num := matrix[i][j]
			if num < 1 || num > n || seen[num] {
				return false
			}
			seen[num] = true
		}
	}

	// Check columns
	for j := 0; j < n; j++ {
		seen := make([]bool, n+1)
		for i := 0; i < n; i++ {
			num := matrix[i][j]
			if num < 1 || num > n || seen[num] {
				return false
			}
			seen[num] = true
		}
	}

	return true
}

func main() {
	timeStartWholeProgram := time.Now()

	testInput := []TestCase{
		{
			Board: [][]int{{1, 2, 3}, {3, 1, 2}, {2, 3, 1}},
			Result: `
	true
            `,
		},

		{
			Board: [][]int{{1, 1, 1}, {1, 2, 3}, {1, 2, 3}},
			Result: `
	false
            `,
		},
	}
	for count, value := range testInput {
		fmt.Println("===============")
		fmt.Println("Test count ", count, "for node", value)
		fmt.Println("Solution 1: Backtrack")
		timeStart := time.Now()
		result := checkValid_StraightForward(value.Board)
		timeLapse := time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

		fmt.Println("Solution 2: DFS : this is a failure solution, need check")
		timeStart = time.Now()
		result = checkValid_DFS(value.Board)
		timeLapse = time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

		fmt.Println("Solution 3: BruteForce")
		timeStart = time.Now()
		result = checkValid_Backtrack(value.Board)
		timeLapse = time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

		fmt.Println("Solution 4: Backtrack")
		timeStart = time.Now()
		result = checkValid_Backtrack(value.Board)
		timeLapse = time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

		fmt.Println("Solution 5: Another approach")
		timeStart = time.Now()
		result = checkValid_Approach2(value.Board)
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
	Board  [][]int
	Result string
}

/*


===============
Test count  0 for node {[[1 2 3] [3 1 2] [2 3 1]]
	true
            }
Solution 1: Backtrack
>Solution result true
Correct result is
	true

TimeLapse 7.814µs
Solution 2: DFS : this is a failure solution, need check
>Solution result false
Correct result is
	true

TimeLapse 2.481µs
Solution 3: BruteForce
>Solution result true
Correct result is
	true

TimeLapse 1.815µs
Solution 4: Backtrack
>Solution result true
Correct result is
	true

TimeLapse 1.204µs
Solution 5: Another approach
>Solution result true
Correct result is
	true

TimeLapse 778ns
===============
Test count  1 for node {[[1 1 1] [1 2 3] [1 2 3]]
	false
            }
Solution 1: Backtrack
>Solution result false
Correct result is
	false

TimeLapse 2.352µs
Solution 2: DFS : this is a failure solution, need check
>Solution result false
Correct result is
	false

TimeLapse 1.13µs
Solution 3: BruteForce
>Solution result false
Correct result is
	false

TimeLapse 1.167µs
Solution 4: Backtrack
>Solution result false
Correct result is
	false

TimeLapse 1.166µs
Solution 5: Another approach
>Solution result false
Correct result is
	false

TimeLapse 167ns
===============
TimeLapse Whole Program 636.001µs

*/
