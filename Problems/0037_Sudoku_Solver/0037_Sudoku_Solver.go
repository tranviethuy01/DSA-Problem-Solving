package main

import (
	"fmt"
	"time"
)

//approach Backtrack
/*
Time Complexity:
The time complexity of the backtracking algorithm for Sudoku depends on the number of empty cells that need to be filled and the branching factor at each empty cell.
In the worst case, where the Sudoku is entirely empty, there are 81 empty cells.
At each empty cell, we try up to 9 possibilities (numbers from 1 to 9).
Therefore, the time complexity can be expressed as O(9^(n^2)), where n is the size of the Sudoku board (9 in this case).
However, in practice, the actual time complexity tends to be much lower because Sudoku puzzles are designed to have unique solutions and can be solved relatively quickly.
Space Complexity:

The space complexity is determined by the recursive stack space used during backtracking.
In this algorithm, we use a constant amount of additional space for variables and the input Sudoku board.
The recursive depth of the backtracking can be at most 81 (the total number of cells in the board), corresponding to the worst-case scenario where every cell is empty and we recursively explore each possibility.
Therefore, the space complexity is O(1) in terms of auxiliary space, but O(n^2) if we consider the space required for the input Sudoku board.
In summary:

Time Complexity: O(9^(n^2)) (though practically much lower)
Space Complexity: O(1) auxiliary space, O(n^2) including input space.

*/
func solveSudoku_Backtrack(board [][]byte) {
	solve_Backtrack(board)
}

func solve_Backtrack(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				for num := byte('1'); num <= byte('9'); num++ {
					if isValid_Backtrack(board, i, j, num) {
						board[i][j] = num
						if solve_Backtrack(board) {
							return true
						}
						board[i][j] = '.'
					}
				}
				return false
			}
		}
	}
	return true
}

func isValid_Backtrack(board [][]byte, row, col int, num byte) bool {
	for i := 0; i < 9; i++ {
		if board[row][i] == num {
			return false
		}
		if board[i][col] == num {
			return false
		}
		if board[3*(row/3)+i/3][3*(col/3)+i%3] == num {
			return false
		}
	}
	return true
}

//approach : DFS
/*
Time Complexity:

In the worst case scenario, where every cell is empty, the algorithm tries all possible combinations of numbers for each empty cell.
Each empty cell can have at most 9 possibilities (numbers from 1 to 9).
Therefore, the time complexity is O(9^(n^2)), where n is the size of the Sudoku board (9 in this case).
However, in practice, Sudoku puzzles are designed to have unique solutions and can be solved much faster.
Space Complexity:

The space complexity is determined by the recursive stack space used during DFS.
In this algorithm, we use a constant amount of additional space for variables and the input Sudoku board.
The maximum recursive depth of the DFS can be at most 81 (the total number of cells in the board), corresponding to the worst-case scenario where every cell is empty and we recursively explore each possibility.
Therefore, the space complexity is O(1) in terms of auxiliary space, but O(n^2) including the input space required for the Sudoku board.
In summary:

Time Complexity: O(9^(n^2)) (though practically much lower)
Space Complexity: O(1) auxiliary space, O(n^2) including input space.
*/
func solveSudoku_DFS(board [][]byte) {
	solve_DFS(board)
}

func solve_DFS(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				for num := byte('1'); num <= byte('9'); num++ {
					if isValid_DFS(board, i, j, num) {
						board[i][j] = num
						if solve_DFS(board) {
							return true
						}
						board[i][j] = '.'
					}
				}
				return false
			}
		}
	}
	return true
}

func isValid_DFS(board [][]byte, row, col int, num byte) bool {
	for i := 0; i < 9; i++ {
		if board[row][i] == num {
			return false
		}
		if board[i][col] == num {
			return false
		}
		if board[3*(row/3)+i/3][3*(col/3)+i%3] == num {
			return false
		}
	}
	return true
}

//approach : BFS => not suitable

//approach : BruteForce
/*
Time Complexity:

In the worst-case scenario, the algorithm tries all possible combinations until it finds a valid solution.
The number of possible combinations depends on the number of empty cells in the Sudoku puzzle.
Each empty cell can potentially have 9 possible values (numbers from 1 to 9).
Therefore, the time complexity can be expressed as O(9^m), where 'm' is the number of empty cells.
However, in practice, the number of empty cells 'm' is usually much less than 81 (the total number of cells), and the algorithm often finds the solution much faster.
Space Complexity:

The space complexity of the brute-force approach is determined by the recursive stack space used during the backtracking process.
In this algorithm, the recursive depth can be at most 'm', corresponding to the number of empty cells that need to be filled.
Therefore, the space complexity is O(m) in terms of auxiliary space.
Additionally, the input Sudoku board itself requires O(1) space.
In summary:

Time Complexity: O(9^m), where 'm' is the number of empty cells.
Space Complexity: O(m), where 'm' is the number of empty cells.
*/

func solveSudoku_BruteForce(board [][]byte) bool {
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if board[row][col] == '.' {
				for num := byte('1'); num <= byte('9'); num++ {
					if isValid_BruteForce(board, row, col, num) {
						board[row][col] = num
						if solveSudoku_BruteForce(board) {
							return true
						}
						board[row][col] = '.' // Backtrack
					}
				}
				return false // If no valid number found for this cell
			}
		}
	}
	return true // All cells are filled
}

func isValid_BruteForce(board [][]byte, row, col int, num byte) bool {
	// Check row
	for i := 0; i < 9; i++ {
		if board[row][i] == num {
			return false
		}
	}
	// Check column
	for i := 0; i < 9; i++ {
		if board[i][col] == num {
			return false
		}
	}
	// Check subgrid
	startRow, startCol := (row/3)*3, (col/3)*3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[startRow+i][startCol+j] == num {
				return false
			}
		}
	}
	return true
}

func main() {
	timeStartWholeProgram := time.Now()

	testInput := []TestCase{

		{
			Board: [][]byte{{'5', '3', '.', '.', '7', '.', '.', '.', '.'}, {'6', '.', '.', '1', '9', '5', '.', '.', '.'}, {'.', '9', '8', '.', '.', '.', '.', '6', '.'}, {'8', '.', '.', '.', '6', '.', '.', '.', '3'}, {'4', '.', '.', '8', '.', '3', '.', '.', '1'}, {'7', '.', '.', '.', '2', '.', '.', '.', '6'}, {'.', '6', '.', '.', '.', '.', '2', '8', '.'}, {'.', '.', '.', '4', '1', '9', '.', '.', '5'}, {'.', '.', '.', '.', '8', '.', '.', '7', '9'}},
			Result: `
      [["5","3","4","6","7","8","9","1","2"],["6","7","2","1","9","5","3","4","8"],["1","9","8","3","4","2","5","6","7"],["8","5","9","7","6","1","4","2","3"],["4","2","6","8","5","3","7","9","1"],["7","1","3","9","2","4","8","5","6"],["9","6","1","5","3","7","2","8","4"],["2","8","7","4","1","9","6","3","5"],["3","4","5","2","8","6","1","7","9"]]
            `,
		},
	}
	for count, value := range testInput {
		//make copied board so that the func will compute on that copy, not change the origin test data
		board := value.Board
		copied1Board := make([][]byte, len(board))
		for i := range board {
			copied1Board[i] = make([]byte, len(board[i]))
			copy(copied1Board[i], board[i])
		}

		copied2Board := make([][]byte, len(board))
		for i := range board {
			copied2Board[i] = make([]byte, len(board[i]))
			copy(copied2Board[i], board[i])
		}

		copied3Board := make([][]byte, len(board))
		for i := range board {
			copied3Board[i] = make([]byte, len(board[i]))
			copy(copied3Board[i], board[i])
		}

		fmt.Println("===============")
		fmt.Println("Test count ", count, "for node", value)
		fmt.Println("Solution 1: Backtrack")
		timeStart := time.Now()
		solveSudoku_Backtrack(copied1Board)
		timeLapse := time.Since(timeStart)
		fmt.Println(">Solution result")
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)
		fmt.Println("Sudoku Solved:")
		for _, row := range copied1Board {
			fmt.Println(string(row))
		}

		//note:need create a new copied board and handle on that
		fmt.Println("Solution 2: DFS")
		timeStart = time.Now()
		solveSudoku_DFS(copied2Board)
		timeLapse = time.Since(timeStart)
		fmt.Println(">Solution result")
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)
		fmt.Println("Sudoku Solved:")
		for _, row := range copied2Board {
			fmt.Println(string(row))
		}

		//note:need create a new copied board and handle on that
		fmt.Println("Solution 3: BruteForce")
		timeStart = time.Now()
		solveSudoku_BruteForce(copied3Board)
		timeLapse = time.Since(timeStart)
		fmt.Println(">Solution result")
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)
		fmt.Println("Sudoku Solved:")
		for _, row := range copied3Board {
			fmt.Println(string(row))
		}

	}

	timeLapsedWholeProgram := time.Since(timeStartWholeProgram)
	fmt.Println("===============")
	fmt.Println("TimeLapse Whole Program", timeLapsedWholeProgram)
}

type TestCase struct {
	Board  [][]byte
	Result string
}

/*

===============
Test count  0 for node {[[53 51 46 46 55 46 46 46 46] [54 46 46 49 57 53 46 46 46] [46 57 56 46 46 46 46 54 46] [56 46 46 46 54 46 46 46 51] [52 46 46 56 46 51 46 46 49] [55 46 46 46 50 46 46 46 54] [46 54 46 46 46 46 50 56 46] [46 46 46 52 49 57 46 46 53] [46 46 46 46 56 46 46 55 57]]
      [["5","3","4","6","7","8","9","1","2"],["6","7","2","1","9","5","3","4","8"],["1","9","8","3","4","2","5","6","7"],["8","5","9","7","6","1","4","2","3"],["4","2","6","8","5","3","7","9","1"],["7","1","3","9","2","4","8","5","6"],["9","6","1","5","3","7","2","8","4"],["2","8","7","4","1","9","6","3","5"],["3","4","5","2","8","6","1","7","9"]]
            }
Solution 1: Backtrack
>Solution result
Correct result is
      [["5","3","4","6","7","8","9","1","2"],["6","7","2","1","9","5","3","4","8"],["1","9","8","3","4","2","5","6","7"],["8","5","9","7","6","1","4","2","3"],["4","2","6","8","5","3","7","9","1"],["7","1","3","9","2","4","8","5","6"],["9","6","1","5","3","7","2","8","4"],["2","8","7","4","1","9","6","3","5"],["3","4","5","2","8","6","1","7","9"]]

TimeLapse 2.555653ms
Sudoku Solved:
534678912
672195348
198342567
859761423
426853791
713924856
961537284
287419635
345286179
Solution 2: DFS
>Solution result
Correct result is
      [["5","3","4","6","7","8","9","1","2"],["6","7","2","1","9","5","3","4","8"],["1","9","8","3","4","2","5","6","7"],["8","5","9","7","6","1","4","2","3"],["4","2","6","8","5","3","7","9","1"],["7","1","3","9","2","4","8","5","6"],["9","6","1","5","3","7","2","8","4"],["2","8","7","4","1","9","6","3","5"],["3","4","5","2","8","6","1","7","9"]]

TimeLapse 2.37812ms
Sudoku Solved:
534678912
672195348
198342567
859761423
426853791
713924856
961537284
287419635
345286179
Solution 3: BruteForce
>Solution result
Correct result is
      [["5","3","4","6","7","8","9","1","2"],["6","7","2","1","9","5","3","4","8"],["1","9","8","3","4","2","5","6","7"],["8","5","9","7","6","1","4","2","3"],["4","2","6","8","5","3","7","9","1"],["7","1","3","9","2","4","8","5","6"],["9","6","1","5","3","7","2","8","4"],["2","8","7","4","1","9","6","3","5"],["3","4","5","2","8","6","1","7","9"]]

TimeLapse 2.238845ms
Sudoku Solved:
534678912
672195348
198342567
859761423
426853791
713924856
961537284
287419635
345286179
===============
TimeLapse Whole Program 7.875993ms
*/
