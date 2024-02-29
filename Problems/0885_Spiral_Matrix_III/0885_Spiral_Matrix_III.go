package main

import (
	"fmt"
	"runtime"
	"time"
)

//
//approach : StraightForward  => this solution is a bad one, time limit Exceeded => look like it loop forever, need check code again
/*

Time Complexity:
The algorithm iterates through each cell in the grid exactly once.
Thus, the time complexity is O(rows * cols), where rows and cols are the dimensions of the grid.
Space Complexity:
The space complexity mainly arises from the output array storing the coordinates visited in the spiral order.
The output array has a size of rows * cols, where each cell stores two integers (coordinates).
Therefore, the space complexity is O(rows * cols).

*/

func spiralMatrixIII_StraightForward(rows int, cols int, rStart int, cStart int) [][]int {
	// Directions: east, south, west, north
	dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	res := make([][]int, rows*cols)
	res[0] = []int{rStart, cStart}
	steps := 1
	idx := 1
	dir := 0

	for idx < rows*cols {
		// Determine number of steps in current direction
		if dir == 0 || dir == 2 {
			steps++
		}

		// Move in the current direction
		for i := 0; i < steps; i++ {
			rStart += dirs[dir][0]
			cStart += dirs[dir][1]

			// Check if position is within grid
			if rStart >= 0 && rStart < rows && cStart >= 0 && cStart < cols {
				res[idx] = []int{rStart, cStart}
				idx++
			}
		}

		// Change direction
		dir = (dir + 1) % 4
	}

	return res
}

//approach : adapt solution from leetcode
/*
Intuition:
Take steps one by one.
If the location is inside of grid, add it to res.
But how to simulate the path?

It seems to be annoying, but if we observer the path:

move right 1 step, turn right
move down 1 step, turn right
move left 2 steps, turn right
move top 2 steps, turn right,
move right 3 steps, turn right
move down 3 steps, turn right
move left 4 steps, turn right
move top 4 steps, turn right,

we can find the sequence of steps: 1,1,2,2,3,3,4,4,5,5....

So there are two thing to figure out:

how to generate sequence 1,1,2,2,3,3,4,4,5,5
how to turn right?
Generate sequence 1,1,2,2,3,3,4,4,5,5
Let n be index of this sequence.
Then A0 = 1, A1 = 1, A2 = 2 ......
We can find that An = n / 2 + 1


How to turn right?
By cross product:
Assume current direction is (x, y) in plane, which is (x, y, 0) in space.
Then the direction after turn right (x, y, 0) × (0, 0, 1) = (y, -x, 0)
Translate to code: tmp = x; x = y; y = -tmp;

By arrays of arrays:
The directions order is (0,1),(1,0),(0,-1),(-1,0), then repeat.
Just define a variable.


Time Complexity:
Time O(max(R,C)^2)
Space O(R*C) for output
*/

func spiralMatrixIII_Adapt2(rows int, cols int, rStart int, cStart int) [][]int {
	res := [][]int{{rStart, cStart}}
	dx, dy := 0, 1
	for n := 0; len(res) < rows*cols; n++ {
		for i := 0; i < n/2+1; i++ {
			rStart += dx
			cStart += dy
			if 0 <= rStart && rStart < rows && 0 <= cStart && cStart < cols {
				res = append(res, []int{rStart, cStart})
			}
		}
		tmp := dx
		dx = dy
		dy = -tmp
	}
	return res
}

func main() {
	timeStartWholeProgram := time.Now()

	testInput := []TestCase{
		{
			Rows:   1,
			Cols:   4,
			RStart: 0,
			CStart: 0,
			Result: `
[[0,0],[0,1],[0,2],[0,3]]

            `,
		},
		{
			Rows:   5,
			Cols:   6,
			RStart: 1,
			CStart: 4,
			Result: `
[[1,4],[1,5],[2,5],[2,4],[2,3],[1,3],[0,3],[0,4],[0,5],[3,5],[3,4],[3,3],[3,2],[2,2],[1,2],[0,2],[4,5],[4,4],[4,3],[4,2],[4,1],[3,1],[2,1],[1,1],[0,1],[4,0],[3,0],[2,0],[1,0],[0,0]]
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
		fmt.Println("Solution 1: adapt solution from leetcode")
		timeStart := time.Now()
		result := spiralMatrixIII_Adapt2(value.Rows, value.Cols, value.RStart, value.CStart)
		timeLapse := time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

		// Memory after allocation
		runtime.ReadMemStats(&m)
		memAfter := m.Alloc
		fmt.Println("Memory before", memBefore, "bytes", "Memory after", memAfter, "bytes", "Memory used:", memAfter-memBefore, "bytes")

		fmt.Printf("Memory usage (HeapAlloc) after Test Case i %d, : %v bytes\n", count, m.HeapAlloc)

		fmt.Println("Solution 2: StraightForward => time limit Exceeded => need to check code again")
		//		timeStart = time.Now()
		//		result = spiralMatrixIII_StraightForward(value.Rows, value.Cols, value.RStart, value.CStart)
		//		timeLapse = time.Since(timeStart)
		//		fmt.Println(">Solution result", result)
		//		fmt.Println("Correct result is ", value.Result)
		//		fmt.Println("TimeLapse", timeLapse)

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
	Rows   int
	Cols   int
	RStart int
	CStart int

	Result string
}

/*

===============
Test count  0 for node {1 4 0 0
[[0,0],[0,1],[0,2],[0,3]]

            }
Solution 1: adapt solution from leetcode
>Solution result [[0 0] [0 1] [0 2] [0 3]]
Correct result is
[[0,0],[0,1],[0,2],[0,3]]


TimeLapse 5.499µs
Memory before 69040 bytes Memory after 70392 bytes Memory used: 1352 bytes
Memory usage (HeapAlloc) after Test Case i 0, : 70392 bytes
Solution 2: StraightForward => time limit Exceeded
>Solution result [[0 0] [0 1] [0 2] [0 3]]
Correct result is
[[0,0],[0,1],[0,2],[0,3]]


TimeLapse 3.296µs
Memory before 69040 bytes Memory after 70800 bytes Memory used: 1760 bytes
Memory usage (HeapAlloc) after Test Case i 0, : 70800 bytes
===============
Test count  1 for node {5 6 1 4
[[1,4],[1,5],[2,5],[2,4],[2,3],[1,3],[0,3],[0,4],[0,5],[3,5],[3,4],[3,3],[3,2],[2,2],[1,2],[0,2],[4,5],[4,4],[4,3],[4,2],[4,1],[3,1],[2,1],[1,1],[0,1],[4,0],[3,0],[2,0],[1,0],[0,0]]
            }
Solution 1: adapt solution from leetcode
>Solution result [[1 4] [1 5] [2 5] [2 4] [2 3] [1 3] [0 3] [0 4] [0 5] [3 5] [3 4] [3 3] [3 2] [2 2] [1 2] [0 2] [4 5] [4 4] [4 3] [4 2] [4 1] [3 1] [2 1] [1 1] [0 1] [4 0] [3 0] [2 0] [1 0] [0 0]]
Correct result is
[[1,4],[1,5],[2,5],[2,4],[2,3],[1,3],[0,3],[0,4],[0,5],[3,5],[3,4],[3,3],[3,2],[2,2],[1,2],[0,2],[4,5],[4,4],[4,3],[4,2],[4,1],[3,1],[2,1],[1,1],[0,1],[4,0],[3,0],[2,0],[1,0],[0,0]]

TimeLapse 27.351µs
Memory before 69040 bytes Memory after 74384 bytes Memory used: 5344 bytes
Memory usage (HeapAlloc) after Test Case i 1, : 74384 bytes
Solution 2: StraightForward => time limit Exceeded


*/
//REF
//
