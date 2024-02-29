package main

import (
	"fmt"
	"runtime"
	"time"
)

//approach : StraightForward

/*
Time complexity of O(m * n), where 'm' is the number of rows and 'n' is the number of columns in the matrix. This is because the function iterates through each element of the matrix exactly once to build the spiral order output.

The space complexity of the algorithm is O(1) for the result slice, which stores the output. The space complexity does not depend on the size of the input matrix, but rather on the size of the output, which is linearly proportional to the number of elements in the matrix.
*/

func spiralOrder_StraightForward(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}

	var result []int

	top, bottom := 0, len(matrix)-1
	left, right := 0, len(matrix[0])-1

	for top <= bottom && left <= right {
		// Traverse top row
		for i := left; i <= right; i++ {
			result = append(result, matrix[top][i])
		}
		top++

		// Traverse right column
		for i := top; i <= bottom; i++ {
			result = append(result, matrix[i][right])
		}
		right--

		// Ensure we're not on the same row after the first traversal
		if top <= bottom {
			// Traverse bottom row
			for i := right; i >= left; i-- {
				result = append(result, matrix[bottom][i])
			}
			bottom--
		}

		// Ensure we're not on the same column after the second traversal
		if left <= right {
			// Traverse left column
			for i := bottom; i >= top; i-- {
				result = append(result, matrix[i][left])
			}
			left++
		}
	}

	return result
}

// note : more guide
/*

Approach:
We will use a while loop to traverse the matrix in a clockwise spiral order.
We will define four variables: left, right, top, bottom to represent the four boundaries of the current spiral.
We will use four for loops to traverse each edge of the current spiral in clockwise order and add the elements to the result list.
We will update the boundaries of the current spiral and continue the process until all elements have been traversed.
Intuition:
We start with the outermost layer of the matrix and traverse it in a clockwise spiral order, adding the elements to the result list.
Then we move on to the next inner layer of the matrix and repeat the process until we have traversed all layers.
To traverse each layer, we need to keep track of the four boundaries of the current spiral.
We start at the top-left corner of the current spiral and move right until we hit the top-right corner.
Then we move down to the bottom-right corner and move left until we hit the bottom-left corner.
Finally, we move up to the top-left corner of the next spiral and repeat the process until we have traversed all elements in the matrix.

class Solution {
public:
    vector<int> spiralOrder(vector<vector<int>>& matrix) {
        vector<int> result;
        if (matrix.empty() || matrix[0].empty()) {
            return result;
        }

        int rows = matrix.size(), cols = matrix[0].size();
        int left = 0, right = cols-1, top = 0, bottom = rows-1;

        while (left <= right && top <= bottom) {
            for (int i = left; i <= right; i++) {
                result.push_back(matrix[top][i]);
            }
            top++;

            for (int i = top; i <= bottom; i++) {
                result.push_back(matrix[i][right]);
            }
            right--;

            if (top <= bottom) {
                for (int i = right; i >= left; i--) {
                    result.push_back(matrix[bottom][i]);
                }
                bottom--;
            }

            if (left <= right) {
                for (int i = bottom; i >= top; i--) {
                    result.push_back(matrix[i][left]);
                }
                left++;
            }
        }

        return result;
    }
};


*/

func spiralOrder_Adapt2(matrix [][]int) []int {
	var result []int
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return result
	}

	rows, cols := len(matrix), len(matrix[0])
	left, right := 0, cols-1
	top, bottom := 0, rows-1

	for left <= right && top <= bottom {
		// Traverse top row
		for i := left; i <= right; i++ {
			result = append(result, matrix[top][i])
		}
		top++

		// Traverse right column
		for i := top; i <= bottom; i++ {
			result = append(result, matrix[i][right])
		}
		right--

		if top <= bottom {
			// Traverse bottom row
			for i := right; i >= left; i-- {
				result = append(result, matrix[bottom][i])
			}
			bottom--
		}

		if left <= right {
			// Traverse left column
			for i := bottom; i >= top; i-- {
				result = append(result, matrix[i][left])
			}
			left++
		}
	}

	return result
}

func main() {
	timeStartWholeProgram := time.Now()

	testInput := []TestCase{
		{
			Matrix: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			Result: `
[1,2,3,6,9,8,7,4,5]

            `,
		},
		{
			Matrix: [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}},
			Result: `
[1,2,3,4,8,12,11,10,9,5,6,7]

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
		result := spiralOrder_StraightForward(value.Matrix)
		timeLapse := time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

		// Memory after allocation
		runtime.ReadMemStats(&m)
		memAfter := m.Alloc
		fmt.Println("Memory before", memBefore, "bytes", "Memory after", memAfter, "bytes", "Memory used:", memAfter-memBefore, "bytes")

		fmt.Printf("Memory usage (HeapAlloc) after Test Case i %d, : %v bytes\n", count, m.HeapAlloc)

		fmt.Println("Solution 2: another solution from leetcode ")
		timeStart = time.Now()
		result = spiralOrder_Adapt2(value.Matrix)
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
	Matrix [][]int
	Result string
}

/*


===============
Test count  0 for node {[[1 2 3] [4 5 6] [7 8 9]]
[1,2,3,6,9,8,7,4,5]

            }
Solution 1: Kadane - Dynamic Programming
>Solution result [1 2 3 6 9 8 7 4 5]
Correct result is
[1,2,3,6,9,8,7,4,5]


TimeLapse 2.667µs
Memory before 67512 bytes Memory after 69056 bytes Memory used: 1544 bytes
Memory usage (HeapAlloc) after Test Case i 0, : 69056 bytes
Solution 2: another solution from leetcode
>Solution result [1 2 3 6 9 8 7 4 5]
Correct result is
[1,2,3,6,9,8,7,4,5]


TimeLapse 3.481µs
Memory before 67512 bytes Memory after 69464 bytes Memory used: 1952 bytes
Memory usage (HeapAlloc) after Test Case i 0, : 69464 bytes
===============
Test count  1 for node {[[1 2 3 4] [5 6 7 8] [9 10 11 12]]
[1,2,3,4,8,12,11,10,9,5,6,7]

            }
Solution 1: Kadane - Dynamic Programming
>Solution result [1 2 3 4 8 12 11 10 9 5 6 7]
Correct result is
[1,2,3,4,8,12,11,10,9,5,6,7]


TimeLapse 2.704µs
Memory before 67512 bytes Memory after 70120 bytes Memory used: 2608 bytes
Memory usage (HeapAlloc) after Test Case i 1, : 70120 bytes
Solution 2: another solution from leetcode
>Solution result [1 2 3 4 8 12 11 10 9 5 6 7]
Correct result is
[1,2,3,4,8,12,11,10,9,5,6,7]


TimeLapse 2.759µs
Memory before 67512 bytes Memory after 70560 bytes Memory used: 3048 bytes
Memory usage (HeapAlloc) after Test Case i 1, : 70560 bytes
===============
TimeLapse Whole Program 872.392µs

*/
//REF
//
