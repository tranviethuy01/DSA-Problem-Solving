package main

import (
	"fmt"
	"time"
)

// approach straight forward
/*
The algorithm used in the solution is quite straightforward. It iterates through each cell of the Sudoku board exactly once, performing constant-time operations to check if the current number violates any of the Sudoku rules (row, column, or 3x3 box). Therefore, the time complexity of this algorithm is O(1) per cell, resulting in an overall time complexity of O(1) * 81 cells = O(81), which simplifies to O(1).

In terms of space complexity, the solution uses three sets of maps to keep track of numbers seen in rows, columns, and 3x3 boxes. However, since the Sudoku board is always 9x9, the size of each map is constant, and the total space required by these maps is also constant. Therefore, the space complexity of the algorithm is O(1).

To summarize:
Time complexity: O(1)
Space complexity: O(1)
*/

func isValidSudoku_StraightForward(board [][]byte) bool {
	rows := make([]map[byte]bool, 9)
	columns := make([]map[byte]bool, 9)
	boxes := make([]map[byte]bool, 9)

	for i := 0; i < 9; i++ {
		rows[i] = make(map[byte]bool)
		columns[i] = make(map[byte]bool)
		boxes[i] = make(map[byte]bool)
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				num := board[i][j]
				boxIndex := (i/3)*3 + j/3

				if rows[i][num] || columns[j][num] || boxes[boxIndex][num] {
					return false
				}

				rows[i][num] = true
				columns[j][num] = true
				boxes[boxIndex][num] = true
			}
		}
	}

	return true
}

/*
// NTOE: need check this code, return a failure, but might consider it
func isValidSudoku(board [][]byte) bool {
    seen := make(map[string]bool)

    for i := 0; i < 9; i++ {
        for j := 0; j < 9; j++ {
            number := board[i][j]
            if number != '.' {
                // Check row, column, and 3x3 block
                if !seen[string(number)+" in row "+strconv.Itoa(i)] || !seen[string(number)+" in column "+strconv.Itoa(j)] || !seen[string(number)+" in block "+strconv.Itoa(i/3)+"-"+strconv.Itoa(j/3)] {
                    return false
                }

                // Add current cell to seen set
                seen[string(number)+" in row "+strconv.Itoa(i)] = true
                seen[string(number)+" in column "+strconv.Itoa(j)] = true
                seen[string(number)+" in block "+strconv.Itoa(i/3)+"-"+strconv.Itoa(j/3)] = true
            }
        }
    }

    return true
}


*/

func main() {
	timeStartWholeProgram := time.Now()

	testInput := []TestCase{

		{
			Board: [][]byte{{'5', '3', '.', '.', '7', '.', '.', '.', '.'}, {'6', '.', '.', '1', '9', '5', '.', '.', '.'}, {'.', '9', '8', '.', '.', '.', '.', '6', '.'}, {'8', '.', '.', '.', '6', '.', '.', '.', '3'}, {'4', '.', '.', '8', '.', '3', '.', '.', '1'}, {'7', '.', '.', '.', '2', '.', '.', '.', '6'}, {'.', '6', '.', '.', '.', '.', '2', '8', '.'}, {'.', '.', '.', '4', '1', '9', '.', '.', '5'}, {'.', '.', '.', '.', '8', '.', '.', '7', '9'}},
			Result: `
      true
            `,
		},

		{
			Board: [][]byte{{'8', '3', '.', '.', '7', '.', '.', '.', '.'}, {'6', '.', '.', '1', '9', '5', '.', '.', '.'}, {'.', '9', '8', '.', '.', '.', '.', '6', '.'}, {'8', '.', '.', '.', '6', '.', '.', '.', '3'}, {'4', '.', '.', '8', '.', '3', '.', '.', '1'}, {'7', '.', '.', '.', '2', '.', '.', '.', '6'}, {'.', '6', '.', '.', '.', '.', '2', '8', '.'}, {'.', '.', '.', '4', '1', '9', '.', '.', '5'}, {'.', '.', '.', '.', '8', '.', '.', '7', '9'}},
			Result: `
      1
            `,
		},
	}
	for count, value := range testInput {
		fmt.Println("===============")
		fmt.Println("Test count ", count, "for node", value)
		fmt.Println("Solution 1: StraightForward")
		timeStart := time.Now()
		result := isValidSudoku_StraightForward(value.Board)
		timeLapse := time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)
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
      true
            }
Solution 1: StraightForward
>Solution result true
Correct result is
      true

TimeLapse 21.945µs
===============
Test count  1 for node {[[56 51 46 46 55 46 46 46 46] [54 46 46 49 57 53 46 46 46] [46 57 56 46 46 46 46 54 46] [56 46 46 46 54 46 46 46 51] [52 46 46 56 46 51 46 46 49] [55 46 46 46 50 46 46 46 54] [46 54 46 46 46 46 50 56 46] [46 46 46 52 49 57 46 46 53] [46 46 46 46 56 46 46 55 57]]
      1
            }
Solution 1: StraightForward
>Solution result false
Correct result is
      1

TimeLapse 8.592µs
===============
TimeLapse Whole Program 462.62µs


*/
