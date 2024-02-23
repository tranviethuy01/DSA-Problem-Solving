package main

import (
	"fmt"
	"time"
)

// NOTE  : passed solution : DFS
/* NOTE
A right C++ solution here

class Solution {
public:
int ways=0;
int cnt;
vector<int>delr={0,1,0,-1};
vector<int>delc={1,0,-1,0};
vector<vector<int>> visited;
void dfs(int r,int c,vector<vector<int>>&grid,vector<vector<int>> &visited,int count){
    count++;
    visited[r][c]=1;

   if(grid[r][c]==2){
       if(count==cnt) ++ways;
       visited[r][c]=0;
       return;
   }

   for(int i=0;i<4;i++){
       int nr=r+delr[i];
       int nc=c+delc[i];
       if(nr>=0 && nc>=0 && nr<grid.size() && nc<grid[0].size()&& grid[nr][nc]!=-1 && visited[nr][nc]==0){
       dfs(nr,nc,grid,visited,count);
       visited[nr][nc]=0;
       }
   }
}
    int uniquePathsIII(vector<vector<int>>&grid) {
         int str=0;
         int stc=0;
         int en=0;
         cnt=0;
          int  n = grid.size(), m = grid[0].size();
         visited.resize(n, vector<int> (m, 0));
        for(int i=0;i<grid.size();i++){
             for(int j=0;j<grid[0].size();j++){
                 if(grid[i][j]==1) {
                     str=i;
                     stc=j;
                 }
                 else if (grid[i][j]==-1) {visited[i][j] = 1; cnt++;}
            }
        }
            cnt=n*m -cnt;

         dfs(str,stc,grid,visited,0);
         return ways;

    }
};

*/

// approach :  (adapt the code above)
// NOTE: need check code again, failure resul

var (
	ways    int
	cnt     int
	delr    = []int{0, 1, 0, -1}
	delc    = []int{1, 0, -1, 0}
	visited [][]int
)

func dfs_AdaptSolution(r, c int, grid [][]int, visited [][]int, count int) {
	count++
	visited[r][c] = 1

	if grid[r][c] == 2 {
		if count == cnt {
			ways++
		}
		visited[r][c] = 0
		return
	}

	for i := 0; i < 4; i++ {
		nr := r + delr[i]
		nc := c + delc[i]
		if nr >= 0 && nc >= 0 && nr < len(grid) && nc < len(grid[0]) && grid[nr][nc] != -1 && visited[nr][nc] == 0 {
			dfs_AdaptSolution(nr, nc, grid, visited, count)
			visited[nr][nc] = 0
		}
	}
}

func uniquePathsIII_AdaptSolution(grid [][]int) int {
	str, stc := 0, 0
	cnt = 0
	n, m := len(grid), len(grid[0])
	visited = make([][]int, n)
	for i := range visited {
		visited[i] = make([]int, m)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 1 {
				str, stc = i, j
			} else if grid[i][j] == -1 {
				visited[i][j] = 1
				cnt++
			}
		}
	}
	cnt = n*m - cnt
	dfs_AdaptSolution(str, stc, grid, visited, 0)
	return ways
}

//approach : Backtrack => wrong solution, need check code

func uniquePathsIII_Backtrack(grid [][]int) int {
	// Call the main function uniquePaths to find the unique paths
	return uniquePaths(grid)
}

func uniquePaths(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	count := 0

	// Find the starting position
	var startX, startY int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				startX, startY = i, j
				break
			}
		}
	}

	// Define directions: up, down, left, right
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	var backtrack func(x, y, remainEmpty int)
	backtrack = func(x, y, remainEmpty int) {
		// If we reached the ending position and all empty squares are visited
		if grid[x][y] == 2 && remainEmpty == 0 {
			count++
			return
		}

		// Mark the current position as visited
		grid[x][y] = -1

		// Explore in four directions
		for _, dir := range directions {
			newX, newY := x+dir[0], y+dir[1]
			if newX >= 0 && newX < m && newY >= 0 && newY < n && grid[newX][newY] != -1 {
				backtrack(newX, newY, remainEmpty-1)
			}
		}

		// Backtrack
		grid[x][y] = 0
	}

	// Start backtracking from the starting position
	backtrack(startX, startY, m*n-1)

	return count
}

//approach DFS => ok => pass
/*
The time complexity of the DFS solution for this problem is O(4^(m*n)), where m is the number of rows and n is the number of columns in the grid. This is because, in the worst case, each cell can have up to four directions to explore, and the DFS explores all possible paths.

The space complexity is O(mn) to store the visited array, where m is the number of rows and n is the number of columns in the grid. Additionally, the recursive call stack can go up to O(mn) in the worst case, as it depends on the number of cells in the grid. Therefore, the overall space complexity is O(mn) for the visited array and O(mn) for the recursive call stack, which simplifies to O(m*n).
*/
var (
	directions_DFS = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
)

func dfs(grid [][]int, x, y, emptyCount int, visited [][]bool) int {
	if grid[x][y] == 2 {
		if emptyCount == 0 {
			return 1
		}
		return 0
	}

	visited[x][y] = true
	count := 0

	for _, dir := range directions_DFS {
		newX, newY := x+dir[0], y+dir[1]
		if newX >= 0 && newX < len(grid) && newY >= 0 && newY < len(grid[0]) &&
			!visited[newX][newY] && grid[newX][newY] != -1 {
			count += dfs(grid, newX, newY, emptyCount-1, visited)
		}
	}

	visited[x][y] = false
	return count
}

func uniquePathsIII_DFS(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	startX, startY := 0, 0
	emptyCount := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				startX, startY = i, j
			} else if grid[i][j] == 0 {
				emptyCount++
			}
		}
	}

	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	return dfs(grid, startX, startY, emptyCount+1, visited)
}

//approach: BFS => failure solution, need check
/*
The time complexity of the BFS solution for this problem is O(m * n), where m is the number of rows and n is the number of columns in the grid. This is because each cell is visited at most once, and each visit involves constant-time operations.

The space complexity of the BFS solution is also O(m * n). This is because we use a queue to perform breadth-first search, which can store at most O(m * n) cells in the worst case. Additionally, we use a 2D boolean array to track visited cells, which also takes O(m * n) space. Therefore, the overall space complexity is O(m * n).
*/
var (
	directions_BFS = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
)

type pair struct {
	x, y int
}

func uniquePathsIII_BFS(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	startX, startY := 0, 0
	emptyCount := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				startX, startY = i, j
			} else if grid[i][j] == 0 {
				emptyCount++
			}
		}
	}

	queue := []pair{{startX, startY}}
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	visited[startX][startY] = true
	steps := 0

	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			curr := queue[0]
			queue = queue[1:]

			if grid[curr.x][curr.y] == 2 && emptyCount == 0 {
				return steps
			}

			for _, dir := range directions_BFS {
				newX, newY := curr.x+dir[0], curr.y+dir[1]
				if newX >= 0 && newX < m && newY >= 0 && newY < n &&
					!visited[newX][newY] && grid[newX][newY] != -1 {
					visited[newX][newY] = true
					queue = append(queue, pair{newX, newY})
					if grid[newX][newY] == 0 {
						emptyCount--
					}
				}
			}
		}
		steps++
	}

	return 0
}

//approach : BruteForce => failure solution, need check code again
/*
The time complexity of the brute force solution is exponential, as it generates all possible paths.

Let's denote m as the number of rows and n as the number of columns in the grid. In the worst case scenario, the number of valid paths can be as high as O((m * n)!), as each cell in the grid can be visited or not visited, leading to factorial complexity.

The space complexity of this solution is also exponential. In the worst case, the space complexity is dominated by the storage of all possible paths, which can be O((m * n)!).

Additionally, we are using auxiliary space to store the visited array for each path generation, which requires O(m * n) space. Therefore, the overall space complexity is O((m * n)!) + O(m * n), which simplifies to O((m * n)!).

Due to the exponential time and space complexity, this brute force solution is not efficient for large grid sizes.

*/

var (
	directions_BruteForce = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
)

func isValidPath(grid [][]int, path []int) bool {
	m, n := len(grid), len(grid[0])
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	emptyCount := 0
	for _, pos := range path {
		x, y := pos/n, pos%n
		if grid[x][y] == -1 || visited[x][y] {
			return false
		}
		if grid[x][y] == 0 {
			emptyCount++
		}
		visited[x][y] = true
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 && !visited[i][j] {
				return false
			}
		}
	}

	return emptyCount == 0
}

func generatePaths(grid [][]int, path []int, i, j int, paths *[][]int) {
	m, n := len(grid), len(grid[0])
	if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == -1 {
		return
	}

	if grid[i][j] == 2 {
		*paths = append(*paths, append([]int(nil), path...))
		return
	}

	grid[i][j] = -1
	path = append(path, i*n+j)

	for _, dir := range directions_BruteForce {
		generatePaths(grid, path, i+dir[0], j+dir[1], paths)
	}

	grid[i][j] = 0
	path = path[:len(path)-1]
}

func uniquePathsIII_BruteForce(grid [][]int) int {
	paths := make([][]int, 0)
	m, n := len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				generatePaths(grid, nil, i, j, &paths)
			}
		}
	}

	count := 0
	for _, path := range paths {
		if isValidPath(grid, path) {
			count++
		}
	}
	return count
}

func main() {
	timeStartWholeProgram := time.Now()

	testInput := []TestCase{
		{
			Board: [][]int{{1, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 2, -1}},
			Result: `
	2
            `,
		},

		{
			Board: [][]int{{1, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 2}},
			Result: `
	4
            `,
		},

		{
			Board: [][]int{{0, 1}, {2, 0}},
			Result: `
	0
            `,
		},
	}
	for count, value := range testInput {
		fmt.Println("===============")
		fmt.Println("Test count ", count, "for node", value)
		fmt.Println("Solution 1: DFS => success ")
		timeStart := time.Now()
		result := uniquePathsIII_DFS(value.Board)
		timeLapse := time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

		fmt.Println("Solution 2: DFS : adapt a leetcode solution => failure, need check code")
		timeStart = time.Now()
		result = uniquePathsIII_AdaptSolution(value.Board)
		timeLapse = time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

		fmt.Println("Solution 3: Backtrack => failured, need check")
		timeStart = time.Now()
		result = uniquePathsIII_Backtrack(value.Board)
		timeLapse = time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

		fmt.Println("Solution 4: BFS => failure")
		timeStart = time.Now()
		result = uniquePathsIII_BFS(value.Board)
		timeLapse = time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

		fmt.Println("Solution 5: BruteForce => failed, need check solution")
		timeStart = time.Now()
		result = uniquePathsIII_BruteForce(value.Board)
		timeLapse = time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

	}

	timeLapsedWholeProgram := time.Since(timeStartWholeProgram)
	fmt.Println("===============")
	fmt.Println("TimeLapse Whole Program", timeLapsedWholeProgram)
}

type TestCase struct {
	Board  [][]int
	Result string
}

/*


===============
Test count  0 for node {[[1 0 0 0] [0 0 0 0] [0 0 2 -1]]
	2
            }
Solution 1: DFS => success
>Solution result 2
Correct result is
	2

TimeLapse 10.444µs
Solution 2: DFS : adapt a leetcode solution => failure, need check code
>Solution result 2
Correct result is
	2

TimeLapse 10.166µs
Solution 3: Backtrack => failured, need check
>Solution result 0
Correct result is
	2

TimeLapse 14.518µs
Solution 4: BFS => failure
>Solution result 0
Correct result is
	2

TimeLapse 4.963µs
Solution 5: BruteForce => failed, need check solution
>Solution result 0
Correct result is
	2

TimeLapse 796ns
===============
Test count  1 for node {[[1 0 0 0] [0 0 0 0] [0 0 0 2]]
	4
            }
Solution 1: DFS => success
>Solution result 4
Correct result is
	4

TimeLapse 14.388µs
Solution 2: DFS : adapt a leetcode solution => failure, need check code
>Solution result 6
Correct result is
	4

TimeLapse 14.518µs
Solution 3: Backtrack => failured, need check
>Solution result 0
Correct result is
	4

TimeLapse 22.999µs
Solution 4: BFS => failure
>Solution result 0
Correct result is
	4

TimeLapse 3.111µs
Solution 5: BruteForce => failed, need check solution
>Solution result 0
Correct result is
	4

TimeLapse 297ns
===============
Test count  2 for node {[[0 1] [2 0]]
	0
            }
Solution 1: DFS => success
>Solution result 0
Correct result is
	0

TimeLapse 1.611µs
Solution 2: DFS : adapt a leetcode solution => failure, need check code
>Solution result 6
Correct result is
	0

TimeLapse 1.389µs
Solution 3: Backtrack => failured, need check
>Solution result 0
Correct result is
	0

TimeLapse 1.13µs
Solution 4: BFS => failure
>Solution result 0
Correct result is
	0

TimeLapse 1.926µs
Solution 5: BruteForce => failed, need check solution
>Solution result 0
Correct result is
	0

TimeLapse 278ns
===============
TimeLapse Whole Program 960.677µs

*/
