package main

import (
	"container/list"
	"fmt"
	"runtime"
	"time"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

//approach : StraightForward
/*
Time Complexity: The time complexity of this algorithm is O(m * n), where m is the number of rows and n is the number of columns in the generated matrix. This complexity arises because we need to visit each cell in the matrix once to fill it with the values from the linked list.

Space Complexity: The space complexity is also O(m * n) because we are creating a matrix of size m x n to store the values. Additionally, we have a constant amount of extra space used for variables, such as directions and indices, which does not depend on the size of the input.
*/
func spiralMatrix_StraightForward(m int, n int, head *ListNode) [][]int {
	matrix := make([][]int, m)
	for i := range matrix {
		matrix[i] = make([]int, n)
		for j := range matrix[i] {
			matrix[i][j] = -1
		}
	}

	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	dirIdx := 0
	row, col := 0, 0

	current := head
	for current != nil {
		matrix[row][col] = current.Val
		current = current.Next

		nextRow, nextCol := row+directions[dirIdx][0], col+directions[dirIdx][1]
		if nextRow < 0 || nextRow >= m || nextCol < 0 || nextCol >= n || matrix[nextRow][nextCol] != -1 {
			dirIdx = (dirIdx + 1) % 4
			nextRow, nextCol = row+directions[dirIdx][0], col+directions[dirIdx][1]
		}
		row, col = nextRow, nextCol
	}

	return matrix
}

//approach : DFS
/*
Time Complexity: The time complexity of the DFS traversal is O(m * n), where m is the number of rows and n is the number of columns in the generated matrix. This is because each cell in the matrix is visited once during the traversal.

Space Complexity: The space complexity is O(m * n) because we are using a matrix of size m x n to store the values. Additionally, the DFS function uses recursion, which can consume stack space proportional to the depth of the recursion, but since the matrix traversal depth will not exceed m + n - 1, the overall space complexity remains O(m * n).


*/
func spiralMatrix_DFS(m int, n int, head *ListNode) [][]int {
	matrix := make([][]int, m)
	for i := range matrix {
		matrix[i] = make([]int, n)
		for j := range matrix[i] {
			matrix[i][j] = -1
		}
	}

	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	dfs(matrix, directions, 0, 0, 0, head)

	return matrix
}

func dfs(matrix [][]int, directions [][]int, row, col, dirIdx int, current *ListNode) {
	if row < 0 || row >= len(matrix) || col < 0 || col >= len(matrix[0]) || matrix[row][col] != -1 || current == nil {
		return
	}

	matrix[row][col] = current.Val

	nextRow, nextCol := row+directions[dirIdx][0], col+directions[dirIdx][1]
	if nextRow < 0 || nextRow >= len(matrix) || nextCol < 0 || nextCol >= len(matrix[0]) || matrix[nextRow][nextCol] != -1 {
		dirIdx = (dirIdx + 1) % 4
		nextRow, nextCol = row+directions[dirIdx][0], col+directions[dirIdx][1]
	}

	dfs(matrix, directions, nextRow, nextCol, dirIdx, current.Next)
}

//approach : BFS
/*
Algorithm: Breadth-First Search (BFS) is used to traverse the matrix. The traversal explores neighboring cells in a layer-by-layer manner, ensuring that each cell is visited exactly once.

Time Complexity: The time complexity of this BFS solution is O(m * n), where m is the number of rows and n is the number of columns in the generated matrix. This complexity arises because each cell in the matrix is visited once during the traversal.

Space Complexity: The space complexity of this BFS solution is also O(m * n). This space is primarily consumed by the matrix to store the values. Additionally, the queue used for BFS traversal may also consume space proportional to the maximum number of cells in the queue at any point during the traversal. Since the maximum size of the queue will not exceed the number of cells in the matrix (m * n), the overall space complexity remains O(m * n).
*/

func spiralMatrix_BFS(m int, n int, head *ListNode) [][]int {
	matrix := make([][]int, m)
	for i := range matrix {
		matrix[i] = make([]int, n)
		for j := range matrix[i] {
			matrix[i][j] = -1
		}
	}

	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	queue := list.New()
	queue.PushBack([3]int{0, 0, 0}) // [row, col, direction index]

	for current := head; current != nil; current = current.Next {
		element := queue.Front()
		queue.Remove(element)
		cell := element.Value.([3]int)
		row, col, dirIdx := cell[0], cell[1], cell[2]
		matrix[row][col] = current.Val

		nextRow, nextCol := row+directions[dirIdx][0], col+directions[dirIdx][1]
		if nextRow >= 0 && nextRow < m && nextCol >= 0 && nextCol < n && matrix[nextRow][nextCol] == -1 {
			queue.PushBack([3]int{nextRow, nextCol, dirIdx})
		} else {
			dirIdx = (dirIdx + 1) % 4
			nextRow, nextCol = row+directions[dirIdx][0], col+directions[dirIdx][1]
			queue.PushBack([3]int{nextRow, nextCol, dirIdx})
		}
	}

	return matrix
}

//approach : dynamic programing note
/*
The problem of generating a matrix in spiral order from a linked list doesn't lend itself well to dynamic programming because dynamic programming typically involves breaking down a problem into smaller subproblems and then solving those subproblems in order to build up to the final solution.

In this problem, the task is more about iterating through the linked list and filling the matrix in a specific order (spiral order) rather than finding optimal substructures or overlapping subproblems that can be memoized or solved independently.

Therefore, while dynamic programming is a powerful technique for many problems, it's not a natural fit for this particular problem. The straightforward iterative approach provided earlier is more suitable.
*/

func main() {
	timeStartWholeProgram := time.Now()

	testInput := []TestCase{
		{
			M:    3,
			N:    5,
			Head: &ListNode{3, &ListNode{0, &ListNode{2, &ListNode{6, &ListNode{8, &ListNode{1, &ListNode{7, &ListNode{9, &ListNode{4, &ListNode{2, &ListNode{5, &ListNode{5, &ListNode{0, nil}}}}}}}}}}}}},
			Result: `
[[3,0,2,6,8],[5,0,-1,-1,1],[5,2,4,9,7]]

            `,
		},
		{
			M:    1,
			N:    4,
			Head: &ListNode{0, &ListNode{1, &ListNode{2, nil}}},
			Result: `
			[[0,1,2,-1]]

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
		result := spiralMatrix_StraightForward(value.M, value.N, value.Head)
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
		result = spiralMatrix_DFS(value.M, value.N, value.Head)
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
		result = spiralMatrix_BFS(value.M, value.N, value.Head)
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
	M      int
	N      int
	Head   *ListNode
	Result string
}

/*



===============
Test count  0 for node {3 5 0x400008a230
[[3,0,2,6,8],[5,0,-1,-1,1],[5,2,4,9,7]]

            }
Solution 1: StraightForward
>Solution result [[3 0 2 6 8] [5 0 -1 -1 1] [5 2 4 9 7]]
Correct result is
[[3,0,2,6,8],[5,0,-1,-1,1],[5,2,4,9,7]]


TimeLapse 5.333µs
Memory before 67440 bytes Memory after 68936 bytes Memory used: 1496 bytes
Memory usage (HeapAlloc) after Test Case i 0, : 68936 bytes
Solution 2: DFS
>Solution result [[3 0 2 6 8] [5 0 -1 -1 1] [5 2 4 9 7]]
Correct result is
[[3,0,2,6,8],[5,0,-1,-1,1],[5,2,4,9,7]]


TimeLapse 4.574µs
Memory before 67440 bytes Memory after 69448 bytes Memory used: 2008 bytes
Memory usage (HeapAlloc) after Test Case i 0, : 69448 bytes
Solution 3: BFS
>Solution result [[3 0 2 6 8] [5 0 -1 -1 1] [5 2 4 9 7]]
Correct result is
[[3,0,2,6,8],[5,0,-1,-1,1],[5,2,4,9,7]]


TimeLapse 13.555µs
Memory before 67440 bytes Memory after 71016 bytes Memory used: 3576 bytes
Memory usage (HeapAlloc) after Test Case i 0, : 71016 bytes
===============
Test count  1 for node {1 4 0x400008a300
			[[0,1,2,-1]]

            }
Solution 1: StraightForward
>Solution result [[0 1 2 -1]]
Correct result is
			[[0,1,2,-1]]


TimeLapse 3.092µs
Memory before 67440 bytes Memory after 71264 bytes Memory used: 3824 bytes
Memory usage (HeapAlloc) after Test Case i 1, : 71264 bytes
Solution 2: DFS
>Solution result [[0 1 2 -1]]
Correct result is
			[[0,1,2,-1]]


TimeLapse 2.37µs
Memory before 67440 bytes Memory after 71464 bytes Memory used: 4024 bytes
Memory usage (HeapAlloc) after Test Case i 1, : 71464 bytes
Solution 3: BFS
>Solution result [[0 1 2 -1]]
Correct result is
			[[0,1,2,-1]]


TimeLapse 3.926µs
Memory before 67440 bytes Memory after 72000 bytes Memory used: 4560 bytes
Memory usage (HeapAlloc) after Test Case i 1, : 72000 bytes
===============
TimeLapse Whole Program 9.351618ms

*/
//REF
//
