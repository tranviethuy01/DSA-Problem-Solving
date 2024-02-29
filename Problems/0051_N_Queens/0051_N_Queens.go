package main

import (
	"fmt"
	"runtime"
	"time"
)

//approach : backtrack
/*
Time Complexity:
In the worst case, the backtracking algorithm explores all possible configurations of queen placements. For each row, it tries to place a queen in every column. So, for the first row, there are n choices, for the second row, there are n-1 choices, and so on. Therefore, the total number of configurations explored is n * (n-1) * (n-2) * ... * 1, which is n!.
Since there are n! possible configurations, the time complexity of the backtracking algorithm is O(n!).
Space Complexity:
The space complexity primarily comes from the storage of the chessboard itself, which is represented as a 2D array of size n x n. Therefore, the space complexity is O(n^2).
Additionally, the result array stores all distinct solutions, which can have a maximum of n! solutions in the worst case. Therefore, the space complexity of the result array is also O(n!).

*/

func solveNQueens_Backtrack(n int) [][]string {
	result := make([][]string, 0)
	board := make([][]rune, n)
	for i := range board {
		board[i] = make([]rune, n)
		for j := range board[i] {
			board[i][j] = '.'
		}
	}

	var backtrack func(int)
	backtrack = func(row int) {
		if row == n {
			temp := make([]string, n)
			for i := range board {
				temp[i] = string(board[i])
			}
			result = append(result, temp)
			return
		}
		for col := 0; col < n; col++ {
			if isSafe(board, row, col, n) {
				board[row][col] = 'Q'
				backtrack(row + 1)
				board[row][col] = '.'
			}
		}
	}

	backtrack(0)
	return result
}

func isSafe(board [][]rune, row, col, n int) bool {
	// Check upper columns in the same row
	for i := 0; i < row; i++ {
		if board[i][col] == 'Q' {
			return false
		}
	}
	// Check upper left diagonal
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == 'Q' {
			return false
		}
	}
	// Check upper right diagonal
	for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
		if board[i][j] == 'Q' {
			return false
		}
	}
	return true
}

//approach : dfs
/*

 */

func solveNQueens_DFS(n int) [][]string {
	result := make([][]string, 0)
	board := make([][]rune, n)
	for i := range board {
		board[i] = make([]rune, n)
		for j := range board[i] {
			board[i][j] = '.'
		}
	}

	var dfs func(int)
	dfs = func(row int) {
		if row == n {
			temp := make([]string, n)
			for i := range board {
				temp[i] = string(board[i])
			}
			result = append(result, temp)
			return
		}
		for col := 0; col < n; col++ {
			if isValid(board, row, col, n) {
				board[row][col] = 'Q'
				dfs(row + 1)
				board[row][col] = '.'
			}
		}
	}

	dfs(0)
	return result
}

func isValid(board [][]rune, row, col, n int) bool {
	// Check upper columns in the same row
	for i := 0; i < row; i++ {
		if board[i][col] == 'Q' {
			return false
		}
	}
	// Check upper left diagonal
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == 'Q' {
			return false
		}
	}
	// Check upper right diagonal
	for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
		if board[i][j] == 'Q' {
			return false
		}
	}
	return true
}

/* approach : dynamic programming , not a good solution, just to educate
Time Complexity: O(n^(n+1))
Space Complexity: O(n^4)


func solveNQueens(n int) [][]string {
	// Define a memoization table to store intermediate results
	memo := make(map[int][][]string)
	return solveNQueensDP(n, memo)
}

func solveNQueensDP(n int, memo map[int][][]string) [][]string {
	// Check if the result for this `n` has already been calculated
	if memo[n] != nil {
		return memo[n]
	}

	// Base case: For n=1, only one solution exists
	if n == 1 {
		return [][]string{{"Q"}}
	}

	// Recursively solve smaller subproblems
	subSolutions := solveNQueensDP(n-1, memo)
	result := make([][]string, 0)

	// For each solution in subSolutions, append a new row
	for _, subSolution := range subSolutions {
		for j := 0; j < n; j++ {
			if isValid(subSolution, n-1, j) {
				newSolution := make([]string, n)
				for k, row := range subSolution {
					newSolution[k] = row
				}
				newRow := ""
				for k := 0; k < n; k++ {
					if k == n-1 {
						newRow += "Q"
					} else {
						newRow += "."
					}
				}
				newSolution[n-1] = newRow
				result = append(result, newSolution)
			}
		}
	}

	// Store the result for future use
	memo[n] = result
	return result
}

func isValid(solution []string, row, col int) bool {
	// Check column
	for i := 0; i < row; i++ {
		if solution[i][col] == 'Q' {
			return false
		}
	}
	// Check diagonal
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if solution[i][j] == 'Q' {
			return false
		}
	}
	for i, j := row-1, col+1; i >= 0 && j < len(solution); i, j = i-1, j+1 {
		if solution[i][j] == 'Q' {
			return false
		}
	}
	return true
}

*/

func main() {
	timeStartWholeProgram := time.Now()

	testInput := []TestCase{
		{
			N: 4,
			Result: `
[[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]


            `,
		},
		{
			N: 1,
			Result: `
[["Q"]]
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
		fmt.Println("Solution 1: Backtrack")
		timeStart := time.Now()
		result := solveNQueens_Backtrack(value.N)
		timeLapse := time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

		// Memory after allocation
		runtime.ReadMemStats(&m)
		memAfter := m.Alloc
		fmt.Println("Memory before", memBefore, "bytes", "Memory after", memAfter, "bytes", "Memory used:", memAfter-memBefore, "bytes")

		fmt.Printf("Memory usage (HeapAlloc) after Test Case i %d, : %v bytes\n", count, m.HeapAlloc)

		fmt.Println("Solution 2: DFS")
		timeStart = time.Now()
		result = solveNQueens_DFS(value.N)
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
	N      int
	Result string
}

/*

===============
Test count  0 for node {4
[[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]


            }
Solution 1: Backtrack
>Solution result [[.Q.. ...Q Q... ..Q.] [..Q. Q... ...Q .Q..]]
Correct result is
[[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]



TimeLapse 13.852µs
Memory before 67184 bytes Memory after 68768 bytes Memory used: 1584 bytes
Memory usage (HeapAlloc) after Test Case i 0, : 68768 bytes
Solution 2: DFS
>Solution result [[.Q.. ...Q Q... ..Q.] [..Q. Q... ...Q .Q..]]
Correct result is
[[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]



TimeLapse 11.278µs
Memory before 67184 bytes Memory after 69472 bytes Memory used: 2288 bytes
Memory usage (HeapAlloc) after Test Case i 0, : 69472 bytes
===============
Test count  1 for node {1
[["Q"]]
            }
Solution 1: Backtrack
>Solution result [[Q]]
Correct result is
[["Q"]]

TimeLapse 2.63µs
Memory before 67184 bytes Memory after 69704 bytes Memory used: 2520 bytes
Memory usage (HeapAlloc) after Test Case i 1, : 69704 bytes
Solution 2: DFS
>Solution result [[Q]]
Correct result is
[["Q"]]

TimeLapse 2.593µs
Memory before 67184 bytes Memory after 69912 bytes Memory used: 2728 bytes
Memory usage (HeapAlloc) after Test Case i 1, : 69912 bytes
===============
TimeLapse Whole Program 905.433µs

*/
//REF
//
