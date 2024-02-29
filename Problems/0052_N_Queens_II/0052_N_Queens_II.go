package main

import (
	"fmt"
	"runtime"
	"time"
)

//approach : backtrack
/*
Time Complexity:
The time complexity of the backtracking algorithm depends on the number of possible configurations of placing the queens on the chessboard. In the worst case, the algorithm explores all possible configurations.
Let's denote N as the size of the chessboard (number of rows or columns), which also represents the number of queens to be placed. The time complexity can be approximated as O(N!), where N! represents the factorial of N. This is because at each row, we have N choices to place a queen, and for each choice, we recursively explore N-1 choices for the next row, then N-2 choices for the row after, and so on. Thus, the total number of possibilities is N * (N-1) * (N-2) * ... * 1, which is N!.

Space Complexity:
The space complexity primarily depends on the space used by the recursive call stack and the additional arrays to keep track of the columns and diagonals. We use three arrays:

cols to keep track of whether a column is occupied or not, which requires O(N) space.
diag1 to keep track of whether a diagonal () is occupied or not, which requires O(2N-1) space.
diag2 to keep track of whether a diagonal (/) is occupied or not, which also requires O(2N-1) space.
Therefore, the total space complexity is O(N) + O(2N-1) + O(2N-1), which simplifies to O(N).

=>
Time Complexity: O(N!)
Space Complexity: O(N)


*/
func totalNQueens_Backtrack(n int) int {
	var count int
	cols := make([]bool, n)
	diag1 := make([]bool, 2*n-1) // diagonal \
	diag2 := make([]bool, 2*n-1) // diagonal /
	backtrack(0, n, cols, diag1, diag2, &count)
	return count
}

func backtrack(row, n int, cols, diag1, diag2 []bool, count *int) {
	if row == n {
		*count++
		return
	}
	for col := 0; col < n; col++ {
		if !cols[col] && !diag1[row+col] && !diag2[row-col+n-1] {
			cols[col] = true
			diag1[row+col] = true
			diag2[row-col+n-1] = true
			backtrack(row+1, n, cols, diag1, diag2, count)
			cols[col] = false
			diag1[row+col] = false
			diag2[row-col+n-1] = false
		}
	}
}

// approach  DFS
/*
Time Complexity:
The time complexity remains the same as the backtracking approach. It depends on the number of possible configurations of placing the queens on the chessboard. In the worst case, the algorithm explores all possible configurations.
Let N be the size of the chessboard (number of rows or columns), which also represents the number of queens to be placed. The time complexity can be approximated as O(N!), where N! represents the factorial of N. This is because at each row, we have N choices to place a queen, and for each choice, we recursively explore N-1 choices for the next row, then N-2 choices for the row after, and so on. Thus, the total number of possibilities is N * (N-1) * (N-2) * ... * 1, which is N!.

Space Complexity:
The space complexity analysis is also the same as the backtracking approach. It depends on the space used by the recursive call stack and the additional arrays to keep track of the columns and diagonals. We use three arrays:

cols to keep track of whether a column is occupied or not, which requires O(N) space.
diag1 to keep track of whether a diagonal () is occupied or not, which requires O(2N-1) space.
diag2 to keep track of whether a diagonal (/) is occupied or not, which also requires O(2N-1) space.
Therefore, the total space complexity is O(N) + O(2N-1) + O(2N-1), which simplifies to O(N).
=>
Time Complexity: O(N!)
Space Complexity: O(N)

*/

func totalNQueens_DFS(n int) int {
	var count int
	cols := make([]bool, n)
	diag1 := make([]bool, 2*n-1) // diagonal \
	diag2 := make([]bool, 2*n-1) // diagonal /
	dfs(0, n, cols, diag1, diag2, &count)
	return count
}

func dfs(row, n int, cols, diag1, diag2 []bool, count *int) {
	if row == n {
		*count++
		return
	}
	for col := 0; col < n; col++ {
		if !cols[col] && !diag1[row+col] && !diag2[row-col+n-1] {
			cols[col] = true
			diag1[row+col] = true
			diag2[row-col+n-1] = true
			dfs(row+1, n, cols, diag1, diag2, count)
			cols[col] = false
			diag1[row+col] = false
			diag2[row-col+n-1] = false
		}
	}
}

//approach BFS
/*
The N-Queens problem is typically solved using backtracking or depth-first search (DFS) due to its recursive nature. While it's not impossible to solve it using breadth-first search (BFS), it's not the most common or efficient approach. However, a BFS-based solution can be used to educate like this code below

Time Complexity:
In the worst-case scenario, where every possible placement of queens needs to be explored, the time complexity of the BFS approach is still exponential. Each queen can be placed in any column of its row, resulting in
N ^ N
  possible configurations to be explored. However, due to the nature of BFS, it might explore fewer configurations than backtracking or DFS because it considers all configurations at the same level before moving to the next level.
Hence, the time complexity is still bounded by
O(N ^ N).

Space Complexity:
The space complexity of the BFS approach depends on the maximum number of states stored in the queue at any point during the execution. In the worst case, the queue might contain all possible states of the chessboard.
At each level, there are
N possible positions for the queen. In the worst case, all rows up to
N−1 are filled, so the number of states at the next level would be
N ^ 2
 . This continues until the last row is filled. Therefore, the maximum number of states stored in the queue is roughly Additionally, each state of the chessboard occupies
O(N ^ 2) space.

Hence, the space complexity is
O(N^N × N^2
 ), which simplifies to
O(N ^ N+2 ) or simply  O(N ^ N)
=>
Algorithm: Breadth-First Search (BFS)
Time Complexity:
O(N ^ N)
Space Complexity:
O(N ^ N)
*/

type State struct {
	row   int
	board [][]byte
}

func totalNQueens_BFS(n int) int {
	count := 0
	queue := []State{{0, make([][]byte, n)}}
	for len(queue) > 0 {
		state := queue[0]
		queue = queue[1:]
		if state.row == n {
			count++
			continue
		}
		for col := 0; col < n; col++ {
			if isValid(state.board, state.row, col, n) {
				newBoard := copyBoard(state.board, n)
				newBoard[state.row][col] = 'Q'
				queue = append(queue, State{state.row + 1, newBoard})
			}
		}
	}
	return count
}

func isValid(board [][]byte, row, col, n int) bool {
	for i := 0; i < row; i++ {
		if board[i][col] == 'Q' {
			return false
		}
	}
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == 'Q' {
			return false
		}
	}
	for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
		if board[i][j] == 'Q' {
			return false
		}
	}
	return true
}

func copyBoard(board [][]byte, n int) [][]byte {
	newBoard := make([][]byte, n)
	for i := range newBoard {
		newBoard[i] = make([]byte, n)
		copy(newBoard[i], board[i])
	}
	return newBoard
}

//approach Dynamic programming
/*
Time Complexity:
Precomputation Step: In the precomputation step, we iterate over i from 1 to n, where n is the size of the chessboard. For each i, we solve the N-Queens problem on an i x i board. The time complexity for solving each subproblem is similar to the backtracking approach and can be approximated as O(i!).
Total Time Complexity: The total time complexity is the sum of solving each subproblem. Therefore, the overall time complexity can be approximated as the sum of the factorial of each number from 1 to n, i.e., 
O(1!)+O(2!)+O(3!)+…+O(n!).

Space Complexity:
We use a dp array to store the number of solutions for each subproblem, which requires O(n) space.
For each subproblem, we also use a chessboard representation (board), columns array (cols), and two diagonal arrays (diag1 and diag2). The space complexity for each of these arrays is O(n).
Total Space Complexity: Considering all the above factors, the total space complexity is O(n) (for dp) + O(n) (for board, cols, diag1, and diag2) = O(n).


 */

func totalNQueens_DP(n int) int {
	if n <= 0 {
		return 0
	}
	// dp[i] stores the number of solutions for i queens on an i x i board
	dp := make([]int, n+1)
	dp[0] = 1 // Base case: 0 queens have 1 solution (empty board)
	for i := 1; i <= n; i++ {
		// Initialize the board with all empty cells
		board := make([][]bool, i)
		for j := range board {
			board[j] = make([]bool, i)
		}
		// Solve for i queens on an i x i board
		dp[i] = solveNQueens(0, i, board, make([]bool, i), make([]bool, 2*i-1), make([]bool, 2*i-1))
	}
	return dp[n]
}

func solveNQueens(row, n int, board [][]bool, cols, diag1, diag2 []bool) int {
	if row == n {
		return 1
	}
	count := 0
	for col := 0; col < n; col++ {
		if !cols[col] && !diag1[row+col] && !diag2[row-col+n-1] {
			// Place the queen
			board[row][col] = true
			cols[col] = true
			diag1[row+col] = true
			diag2[row-col+n-1] = true
			// Recur for the next row
			count += solveNQueens(row+1, n, board, cols, diag1, diag2)
			// Remove the queen (backtrack)
			board[row][col] = false
			cols[col] = false
			diag1[row+col] = false
			diag2[row-col+n-1] = false
		}
	}
	return count
}

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
		result := totalNQueens_Backtrack(value.N)
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
		result = totalNQueens_DFS(value.N)
		timeLapse = time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

		// Memory after allocation
		runtime.ReadMemStats(&m)
		memAfter = m.Alloc
		fmt.Println("Memory before", memBefore, "bytes", "Memory after", memAfter, "bytes", "Memory used:", memAfter-memBefore, "bytes")

		fmt.Printf("Memory usage (HeapAlloc) after Test Case i %d, : %v bytes\n", count, m.HeapAlloc)

		fmt.Println("Solution 3: BFS")
		timeStart = time.Now()
		result = totalNQueens_BFS(value.N)
		timeLapse = time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

		// Memory after allocation
		runtime.ReadMemStats(&m)
		memAfter = m.Alloc
		fmt.Println("Memory before", memBefore, "bytes", "Memory after", memAfter, "bytes", "Memory used:", memAfter-memBefore, "bytes")

		fmt.Printf("Memory usage (HeapAlloc) after Test Case i %d, : %v bytes\n", count, m.HeapAlloc)

		fmt.Println("Solution 4: DP")
		timeStart = time.Now()
		result = totalNQueens_DP(value.N)
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
>Solution result 2
Correct result is
[[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]



TimeLapse 3.018µs
Memory before 67392 bytes Memory after 68368 bytes Memory used: 976 bytes
Memory usage (HeapAlloc) after Test Case i 0, : 68368 bytes
Solution 2: DFS
>Solution result 2
Correct result is
[[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]



TimeLapse 2.629µs
Memory before 67392 bytes Memory after 68464 bytes Memory used: 1072 bytes
Memory usage (HeapAlloc) after Test Case i 0, : 68464 bytes
Solution 3: BFS
>Solution result 2
Correct result is
[[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]



TimeLapse 35.37µs
Memory before 67392 bytes Memory after 71472 bytes Memory used: 4080 bytes
Memory usage (HeapAlloc) after Test Case i 0, : 71472 bytes
Solution 4: DP
>Solution result 2
Correct result is
[[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]



TimeLapse 9.259µs
Memory before 67392 bytes Memory after 71912 bytes Memory used: 4520 bytes
Memory usage (HeapAlloc) after Test Case i 0, : 71912 bytes
===============
Test count  1 for node {1
[["Q"]]
            }
Solution 1: Backtrack
>Solution result 1
Correct result is
[["Q"]]

TimeLapse 741ns
Memory before 67392 bytes Memory after 72016 bytes Memory used: 4624 bytes
Memory usage (HeapAlloc) after Test Case i 1, : 72016 bytes
Solution 2: DFS
>Solution result 1
Correct result is
[["Q"]]

TimeLapse 759ns
Memory before 67392 bytes Memory after 72096 bytes Memory used: 4704 bytes
Memory usage (HeapAlloc) after Test Case i 1, : 72096 bytes
Solution 3: BFS
>Solution result 1
Correct result is
[["Q"]]

TimeLapse 3.055µs
Memory before 67392 bytes Memory after 72256 bytes Memory used: 4864 bytes
Memory usage (HeapAlloc) after Test Case i 1, : 72256 bytes
Solution 4: DP
>Solution result 1
Correct result is
[["Q"]]

TimeLapse 1.87µs
Memory before 67392 bytes Memory after 72376 bytes Memory used: 4984 bytes
Memory usage (HeapAlloc) after Test Case i 1, : 72376 bytes
===============
TimeLapse Whole Program 1.196412ms

*/
//REF
//
