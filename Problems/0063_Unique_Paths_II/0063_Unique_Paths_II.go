package main

import (
	"fmt"
	"runtime"
	"time"
)

//approach  dynamic programing
/*
Algorithm Explanation:

We initialize a 2D array dp to store the number of unique paths to each cell.
We iterate through each cell of the grid. For each cell, if it's not an obstacle, we calculate the number of unique paths to that cell by summing up the number of paths from the cell above and the cell to the left.
We return the value at the bottom-right cell of the dp array, which represents the number of unique paths from the top-left corner to the bottom-right corner.
Time Complexity:

We iterate through each cell of the grid exactly once to fill the dp array. This requires traversing the entire grid, so the time complexity is O(m*n), where m is the number of rows and n is the number of columns in the grid.
Space Complexity:

We use an additional 2D array dp to store the number of unique paths to each cell. The size of this array is the same as the size of the input grid, so the space complexity is O(m*n), where m is the number of rows and n is the number of columns in the grid.

*/

func uniquePathsWithObstacles_DP(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])

	// If the starting point or ending point has obstacle, then no path is possible.
	if obstacleGrid[0][0] == 1 || obstacleGrid[m-1][n-1] == 1 {
		return 0
	}

	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	// Initialize the first row and first column.
	dp[0][0] = 1
	for i := 1; i < m; i++ {
		if obstacleGrid[i][0] == 0 {
			dp[i][0] = dp[i-1][0]
		}
	}
	for j := 1; j < n; j++ {
		if obstacleGrid[0][j] == 0 {
			dp[0][j] = dp[0][j-1]
		}
	}

	// Populate the dp table.
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 0 {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}

	return dp[m-1][n-1]
}

//approach BFS
/*
The time complexity of the BFS solution is O(m * n), where m is the number of rows and n is the number of columns in the grid. This is because in the worst-case scenario, we might have to visit every cell of the grid once.

The space complexity of the BFS solution is also O(m * n). This is because we are using a queue to perform BFS, and in the worst-case scenario, the queue might contain all the cells of the grid. Additionally, we are not using any extra data structures proportional to the size of the grid, so the space complexity remains O(m * n).
*/

type Pair struct {
	x, y, count int
}

func uniquePathsWithObstacles_BFS(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])

	// If the starting point or ending point has obstacle, then no path is possible.
	if obstacleGrid[0][0] == 1 || obstacleGrid[m-1][n-1] == 1 {
		return 0
	}

	// Initialize a queue for BFS.
	queue := []Pair{{0, 0, 1}}
	count := 0

	// Define the direction vectors for moving down and right.
	dx := []int{1, 0}
	dy := []int{0, 1}

	// Perform BFS.
	for len(queue) > 0 {
		front := queue[0]
		queue = queue[1:]

		if front.x == m-1 && front.y == n-1 {
			count += front.count
		}

		for i := 0; i < 2; i++ {
			nx := front.x + dx[i]
			ny := front.y + dy[i]

			if nx >= 0 && nx < m && ny >= 0 && ny < n && obstacleGrid[nx][ny] == 0 {
				queue = append(queue, Pair{nx, ny, front.count})
			}
		}
	}

	return count
}

//approach DFS
/*

The time complexity of the DFS solution is also O(m * n), where m is the number of rows and n is the number of columns in the grid. This is because in the worst-case scenario, we might have to visit every cell of the grid once.

The space complexity of the DFS solution depends on the depth of the recursion stack. In the worst-case scenario, where there are no obstacles and the path goes from the top-left corner to the bottom-right corner, the recursion depth would be at most m + n - 1. Therefore, the space complexity is O(m + n) due to the recursion stack space. Additionally, since we're not using any extra data structures proportional to the size of the grid, the overall space complexity remains O(m + n).

*/

func uniquePathsWithObstacles_DFS(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])

	// If the starting point or ending point has obstacle, then no path is possible.
	if obstacleGrid[0][0] == 1 || obstacleGrid[m-1][n-1] == 1 {
		return 0
	}

	count := 0

	// Define the direction vectors for moving down and right.
	dx := []int{1, 0}
	dy := []int{0, 1}

	var dfs func(x, y int)
	dfs = func(x, y int) {
		if x == m-1 && y == n-1 {
			count++
			return
		}

		for i := 0; i < 2; i++ {
			nx := x + dx[i]
			ny := y + dy[i]

			if nx >= 0 && nx < m && ny >= 0 && ny < n && obstacleGrid[nx][ny] == 0 {
				dfs(nx, ny)
			}
		}
	}

	// Start DFS from the top-left corner.
	dfs(0, 0)

	return count
}

func main() {
	timeStartWholeProgram := time.Now()

	testInput := []TestCase{
		{
			ObstacleGrid: [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}},
			Result: `
2
            `,
		},
		{
			ObstacleGrid: [][]int{{0, 1}, {0, 0}},
			Result: `
1
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
		result := uniquePathsWithObstacles_DP(value.ObstacleGrid)
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
		result = uniquePathsWithObstacles_DFS(value.ObstacleGrid)
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
		result = uniquePathsWithObstacles_BFS(value.ObstacleGrid)
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
	ObstacleGrid [][]int
	Result       string
}

/*

===============
Test count  0 for node {[[0 0 0] [0 1 0] [0 0 0]]
2
            }
Solution 1: dynamic programing
>Solution result 2
Correct result is
2

TimeLapse 4.981µs
Memory before 69248 bytes Memory after 70600 bytes Memory used: 1352 bytes
Memory usage (HeapAlloc) after Test Case i 0, : 70600 bytes
Solution 2: DFS
>Solution result 2
Correct result is
2

TimeLapse 1.944µs
Memory before 69248 bytes Memory after 70664 bytes Memory used: 1416 bytes
Memory usage (HeapAlloc) after Test Case i 0, : 70664 bytes
Solution 3: BFS
>Solution result 2
Correct result is
2

TimeLapse 5.926µs
Memory before 69248 bytes Memory after 71088 bytes Memory used: 1840 bytes
Memory usage (HeapAlloc) after Test Case i 0, : 71088 bytes
===============
Test count  1 for node {[[0 1] [0 0]]
1
            }
Solution 1: dynamic programing
>Solution result 1
Correct result is
1

TimeLapse 2.389µs
Memory before 69248 bytes Memory after 71360 bytes Memory used: 2112 bytes
Memory usage (HeapAlloc) after Test Case i 1, : 71360 bytes
Solution 2: DFS
>Solution result 1
Correct result is
1

TimeLapse 1.111µs
Memory before 69248 bytes Memory after 71424 bytes Memory used: 2176 bytes
Memory usage (HeapAlloc) after Test Case i 1, : 71424 bytes
Solution 3: BFS
>Solution result 1
Correct result is
1

TimeLapse 2.37µs
Memory before 69248 bytes Memory after 71536 bytes Memory used: 2288 bytes
Memory usage (HeapAlloc) after Test Case i 1, : 71536 bytes
===============
TimeLapse Whole Program 1.223549ms

*/
//REF
