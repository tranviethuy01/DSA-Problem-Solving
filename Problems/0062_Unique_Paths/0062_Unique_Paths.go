package main

import (
	"fmt"
	"runtime"
	"time"
)

//approach : dynamic programing

func uniquePaths_DP(m int, n int) int {
	// Create a 2D array to store the number of unique paths
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	// Fill the first column with 1 since there is only one way to reach any cell in the first column
	for i := 0; i < m; i++ {
		dp[i][0] = 1
	}

	// Fill the first row with 1 since there is only one way to reach any cell in the first row
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}

	// Calculate the number of unique paths for each cell
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	// Return the number of unique paths to reach the bottom-right corner
	return dp[m-1][n-1]
}

func uniquePaths_DFS(m int, n int) int {
	count := 0
	dfs(0, 0, m, n, &count)
	return count
}

func dfs(row, col, m, n int, count *int) {
	// Base case: if we reached the bottom-right corner, increment the count
	if row == m-1 && col == n-1 {
		*count++
		return
	}

	// Move right if within bounds
	if col+1 < n {
		dfs(row, col+1, m, n, count)
	}
	// Move down if within bounds
	if row+1 < m {
		dfs(row+1, col, m, n, count)
	}
}


//approach BFS => this solution go wrong, need check again

func uniquePaths_BFS(m int, n int) int {
	// Create a queue for BFS
	queue := make([][2]int, 0)
	queue = append(queue, [2]int{0, 0})

	// Create a 2D array to keep track of visited cells
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	// Mark the starting cell as visited
	visited[0][0] = true

	// Direction arrays for moving right and down
	directions := [][2]int{{0, 1}, {1, 0}}

	// BFS traversal
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		// Check if we reached the destination
		if current[0] == m-1 && current[1] == n-1 {
			return 1
		}

		// Explore neighbors (move right or down)
		for _, dir := range directions {
			newRow, newCol := current[0]+dir[0], current[1]+dir[1]

			// Check if the new position is within bounds and not visited
			if newRow < m && newCol < n && !visited[newRow][newCol] {
				// Mark the new cell as visited
				visited[newRow][newCol] = true
				// Add the new cell to the queue
				queue = append(queue, [2]int{newRow, newCol})
			}
		}
	}

	return 0
}

func main() {
	timeStartWholeProgram := time.Now()

	testInput := []TestCase{
		{
			M: 3,
			N: 7,
			Result: `
28
            `,
		},
		{
			M: 3,
			N: 2,
			Result: `
3
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
		fmt.Println("Solution 1: dynamic programing")
		timeStart := time.Now()
		result := uniquePaths_DP(value.M, value.N)
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
		result = uniquePaths_DFS(value.M, value.N)
		timeLapse = time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

		// Memory after allocation
		runtime.ReadMemStats(&m)
		memAfter = m.Alloc
		fmt.Println("Memory before", memBefore, "bytes", "Memory after", memAfter, "bytes", "Memory used:", memAfter-memBefore, "bytes")

		fmt.Printf("Memory usage (HeapAlloc) after Test Case i %d, : %v bytes\n", count, m.HeapAlloc)

		fmt.Println("Solution 3: BFS => failed solution, neee check")
		timeStart = time.Now()
		result = uniquePaths_BFS(value.M, value.N)
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
	Result string
}

/*

===============
Test count  0 for node {3 7
28
            }
Solution 1: dynamic programing
>Solution result 28
Correct result is
28

TimeLapse 4.833µs
Memory before 69224 bytes Memory after 70400 bytes Memory used: 1176 bytes
Memory usage (HeapAlloc) after Test Case i 0, : 70400 bytes
Solution 2: DFS
>Solution result 28
Correct result is
28

TimeLapse 2.185µs
Memory before 69224 bytes Memory after 70592 bytes Memory used: 1368 bytes
Memory usage (HeapAlloc) after Test Case i 0, : 70592 bytes
Solution 3: BFS => failed solution, neee check
>Solution result 1
Correct result is
28

TimeLapse 8.74µs
Memory before 69224 bytes Memory after 71408 bytes Memory used: 2184 bytes
Memory usage (HeapAlloc) after Test Case i 0, : 71408 bytes
===============
Test count  1 for node {3 2
3
            }
Solution 1: dynamic programing
>Solution result 3
Correct result is
3

TimeLapse 2.296µs
Memory before 69224 bytes Memory after 71632 bytes Memory used: 2408 bytes
Memory usage (HeapAlloc) after Test Case i 1, : 71632 bytes
Solution 2: DFS
>Solution result 3
Correct result is
3

TimeLapse 537ns
Memory before 69224 bytes Memory after 71696 bytes Memory used: 2472 bytes
Memory usage (HeapAlloc) after Test Case i 1, : 71696 bytes
Solution 3: BFS => failed solution, neee check
>Solution result 1
Correct result is
3

TimeLapse 3.259µs
Memory before 69224 bytes Memory after 72016 bytes Memory used: 2792 bytes
Memory usage (HeapAlloc) after Test Case i 1, : 72016 bytes
===============
TimeLapse Whole Program 1.16333ms

*/
//REF
