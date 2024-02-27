package main

import (
	"fmt"
	"time"
)
//approach straight forward
/*
Time complexity of O(n^2), where n is the number of rows or columns in the input matrix.
The transposition step takes O(n^2) time because each element of the matrix is being accessed once and swapped.
The reversing step also takes O(n^2) time because each row of the matrix is being accessed, and within each row, only half of the elements need to be swapped.

The space complexity of the algorithm is O(1) because it operates in-place, meaning it does not use any additional space that scales with the size of the input. The operations are performed directly on the input matrix without using any extra data structures. Therefore, the space complexity remains constant regardless of the size of the matrix.
*/



func rotate(matrix [][]int) {
    n := len(matrix)

    // Transpose the matrix
    //swap column with rows here
    for i := 0; i < n; i++ {
        for j := i; j < n; j++ {
            matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
        }
    }

    fmt.Println("matrix after transpose", matrix)

    // Reverse each row
    for i := 0; i < n; i++ {
        for j := 0; j < n/2; j++ {
            matrix[i][j], matrix[i][n-j-1] = matrix[i][n-j-1], matrix[i][j]
        }
    }

    fmt.Println("matrix after reverse each row", matrix)

}


/* 
//the code below will Rotate CounterClockwise


func rotateCounterClockwise(matrix [][]int) {
    n := len(matrix)

    // Transpose the matrix
    for i := 0; i < n; i++ {
        for j := i; j < n; j++ {
            matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
        }
    }

    //here, reverse each column instead of rows
    // Reverse each column (equivalent to reversing each row in the transposed matrix)
    for i := 0; i < n/2; i++ {
        for j := 0; j < n; j++ {
            matrix[i][j], matrix[n-i-1][j] = matrix[n-i-1][j], matrix[i][j]
        }
    }
}

*/

func main() {
	timeStartWholeProgram := time.Now()

	testInput := []TestCase{
		{
      Matrix: [][]int{{1,2,3},{4,5,6},{7,8,9}},
			Result: `
[[7,4,1],[8,5,2],[9,6,3]]

            `,
		},

		{
      Matrix: [][]int{{5,1,9,11},{2,4,8,10},{13,3,6,7},{15,14,12,16}},
			Result: `
[[15,13,2,5],[14,3,4,1],[12,6,8,9],[16,7,10,11]]
            `,
		},
	}
	for count, value := range testInput {
		fmt.Println("===============")
		fmt.Println("Test count ", count, "for node", value)
		fmt.Println("Solution 1: StraightForward")
		timeStart := time.Now()
		 rotate(value.Matrix)
		timeLapse := time.Since(timeStart)
		fmt.Println(">Solution result")
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)


  }

	timeLapsedWholeProgram := time.Since(timeStartWholeProgram)
	fmt.Println("===============")
	fmt.Println("TimeLapse Whole Program", timeLapsedWholeProgram)
}

type TestCase struct {
	Matrix   [][]int
	Result string
}

/*

===============
Test count  0 for node {[[1 2 3] [4 5 6] [7 8 9]] 
[[7,4,1],[8,5,2],[9,6,3]]

            }
Solution 1: StraightForward
matrix after transpose [[1 4 7] [2 5 8] [3 6 9]]
matrix after reverse each row [[7 4 1] [8 5 2] [9 6 3]]
>Solution result
Correct result is  
[[7,4,1],[8,5,2],[9,6,3]]

            
TimeLapse 30.777µs
===============
Test count  1 for node {[[5 1 9 11] [2 4 8 10] [13 3 6 7] [15 14 12 16]] 
[[15,13,2,5],[14,3,4,1],[12,6,8,9],[16,7,10,11]]
            }
Solution 1: StraightForward
matrix after transpose [[5 2 13 15] [1 4 3 14] [9 8 6 12] [11 10 7 16]]
matrix after reverse each row [[15 13 2 5] [14 3 4 1] [12 6 8 9] [16 7 10 11]]
>Solution result
Correct result is  
[[15,13,2,5],[14,3,4,1],[12,6,8,9],[16,7,10,11]]
            
TimeLapse 32.721µs
===============
TimeLapse Whole Program 376.084µs

*/
//REF
//
