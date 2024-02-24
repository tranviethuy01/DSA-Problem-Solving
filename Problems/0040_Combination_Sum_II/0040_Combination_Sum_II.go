package main

import (
	"fmt"
	"sort"
	"time"
)

// approach Backtrack
/*
The algorithm used here is a backtracking algorithm. Backtracking is a systematic way to explore all possible combinations of a solution space. In this algorithm, we recursively try different candidates, backtrack if we reach an invalid state, and continue exploring other candidates until we find a valid combination that sums up to the target.

Let's analyze the time and space complexity:

- **Time Complexity**:
  - In the worst-case scenario, we might have to explore all possible combinations.
  - Since at each step, we have multiple choices (all the remaining candidates), the time complexity is exponential.
  - However, with pruning (skipping duplicates and stopping early if the sum exceeds the target), we can significantly reduce the search space.
  - The time complexity can be roughly expressed as O(2^N), where N is the number of candidates.

- **Space Complexity**:
  - The space complexity is determined by the recursion depth and the auxiliary space used for storing the combinations.
  - At any point during the recursion, the maximum depth of the call stack would be the length of the candidates array.
  - Additionally, the space used to store each combination would be proportional to the length of the combinations.
  - Therefore, the space complexity can be expressed as O(N * M), where N is the number of candidates and M is the average length of the combinations.

*/
func combinationSum2_Backtrack(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var result [][]int
	backtrack(&result, []int{}, candidates, target, 0)
	return result
}

func backtrack(result *[][]int, temp []int, candidates []int, remain int, start int) {
	if remain < 0 {
		return
	} else if remain == 0 {
		tempCopy := make([]int, len(temp))
		copy(tempCopy, temp)
		*result = append(*result, tempCopy)
		return
	} else {
		for i := start; i < len(candidates); i++ {
			if i > start && candidates[i] == candidates[i-1] {
				continue
			}
			temp = append(temp, candidates[i])
			backtrack(result, temp, candidates, remain-candidates[i], i+1)
			temp = temp[:len(temp)-1]
		}
	}
}

//approach DP
/*
Algorithm:
We create a 3-dimensional dynamic programming table dp, where dp[i][j] stores all combinations of numbers from the candidates array that sum up to j using the first i elements of the candidates array.
We initialize dp[0][0] with an empty combination, indicating that there's one way to obtain a sum of 0 without using any numbers.
We iterate through each candidate number and update dp[i][j] based on the previous combinations stored in dp[i-1][j] and dp[i-1][j-num].
Finally, we return the unique combinations stored in dp[len(candidates)][target].
Time Complexity:

The time complexity is determined by the nested loops used to fill in the dynamic programming table.
We have three nested loops: one for iterating through each candidate number, one for iterating through the target values, and one for iterating through the previous combinations.
Therefore, the time complexity is O(N * T * M), where N is the number of candidates, T is the target value, and M is the average length of combinations.
Space Complexity:

The space complexity is determined by the dynamic programming table dp.
We create a 3-dimensional array of size (N+1) * (T+1) * K, where K is the average length of combinations.
Therefore, the space complexity is O(N * T * K).
In summary:

Algorithm: Dynamic Programming
Time Complexity: O(N * T * M)
Space Complexity: O(N * T * K)
*/

func combinationSum2_DP(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	dp := make([][][]int, target+1)
	for i := range dp {
		dp[i] = make([][]int, 0)
	}
	dp[0] = append(dp[0], []int{})

	for _, num := range candidates {
		for j := target; j >= num; j-- {
			for _, prev := range dp[j-num] {
				temp := make([]int, len(prev))
				copy(temp, prev)
				temp = append(temp, num)
				dp[j] = append(dp[j], temp)
			}
		}
	}
	return uniqueCombinations(dp[target])
}

func uniqueCombinations(combos [][]int) [][]int {
	seen := make(map[string]bool)
	result := make([][]int, 0)
	for _, combo := range combos {
		key := fmt.Sprintf("%v", combo)
		if !seen[key] {
			result = append(result, combo)
			seen[key] = true
		}
	}
	return result
}

//approach DFS

/*

Time Complexity:

The time complexity is dependent on the number of recursive calls made during the DFS traversal.
In the worst case, each candidate can be selected or skipped at each step, leading to a branching factor of 2.
Therefore, the time complexity is exponential, but with pruning (skipping duplicates and stopping early if the sum exceeds the target), we can significantly reduce the search space.
The time complexity can be expressed as O(2^N), where N is the number of candidates.
Space Complexity:

The space complexity is determined by the recursion depth and the auxiliary space used for storing the combinations.
At any point during the recursion, the maximum depth of the call stack would be the length of the candidates array.
Additionally, the space used to store each combination would be proportional to the length of the combinations.
Therefore, the space complexity can be expressed as O(N * M), where N is the number of candidates and M is the average length of the combinations.
In summary:

Time Complexity: O(2^N) where N is the number of candidates.
Space Complexity: O(N * M) whe
*/

func combinationSum2_DFS(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var result [][]int
	dfs(&result, []int{}, candidates, target, 0)
	return result
}

func dfs(result *[][]int, temp []int, candidates []int, remain, start int) {
	if remain < 0 {
		return
	} else if remain == 0 {
		tempCopy := make([]int, len(temp))
		copy(tempCopy, temp)
		*result = append(*result, tempCopy)
		return
	} else {
		for i := start; i < len(candidates); i++ {
			// Skip duplicates
			if i > start && candidates[i] == candidates[i-1] {
				continue
			}
			temp = append(temp, candidates[i])
			// Move to the next index with adjusted remaining value
			dfs(result, temp, candidates, remain-candidates[i], i+1)
			temp = temp[:len(temp)-1]
		}
	}
}

//approach BFS => failure solution => need check code again
/*

func combinationSum2_BFS(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var result [][]int
	queue := [][]int{{}}
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		sumCurr := sum(curr)
		if sumCurr == target {
			result = append(result, curr)
		} else if sumCurr < target {
			for i, num := range candidates {
				if i > 0 && candidates[i] == candidates[i-1] {
					continue
				}
				if len(curr) == 0 || num >= curr[len(curr)-1] {
					newComb := make([]int, len(curr)+1)
					copy(newComb, curr)
					newComb[len(curr)] = num
					queue = append(queue, newComb)
				}
			}
		}
	}
	return result
}

func sum(nums []int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

*/

//approach : BruteForce, a brute-force approach by trying all possible combinations without skipping any candidates or performing any duplicate checking , just to learn
/*
func combinationSum2_BruteForce(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var result [][]int
	var backtrack func([]int, int, int)
	backtrack = func(curr []int, start int, remain int) {
		if remain == 0 {
			temp := make([]int, len(curr))
			copy(temp, curr)
			result = append(result, temp)
			return
		}
		for i := start; i < len(candidates); i++ {
			if i > start && candidates[i] == candidates[i-1] {
				continue
			}
			if candidates[i] > remain {
				break
			}
			backtrack(append(curr, candidates[i]), i+1, remain-candidates[i])
		}
	}
	backtrack([]int{}, 0, target)
	return result
}


//BruteForce, no optimization, another code again
func combinationSum2(candidates []int, target int) [][]int {
	var result [][]int
	var dfs func(int, []int)
	dfs = func(idx int, current []int) {
		if idx == len(candidates) {
			if sum(current) == target {
				temp := make([]int, len(current))
				copy(temp, current)
				result = append(result, temp)
			}
			return
		}
		// Include current candidate
		current = append(current, candidates[idx])
		dfs(idx+1, current)
		// Exclude current candidate
		current = current[:len(current)-1]
		dfs(idx+1, current)
	}
	dfs(0, []int{})
	return result
}

func sum(nums []int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}


*/

func main() {
	timeStartWholeProgram := time.Now()

	testInput := []TestCase{
		{
			Candidates: []int{10, 1, 2, 7, 6, 1, 5},
			Target:     8,
			Result: `
[[1,1,6],[1,2,5],[1,7],[2,6]]
            `,
		},

		{
			Candidates: []int{2, 5, 2, 1, 2},
			Target:     8,
			Result: `
			[[1,2,2],[5]]
            `,
		},

		{
			Candidates: []int{2},
			Target:     1,
			Result: `
[]
            `,
		},
	}
	for count, value := range testInput {

		fmt.Println("===============")
		fmt.Println("Test count ", count, "for node", value)
		fmt.Println("Solution 1: Backtrack")
		timeStart := time.Now()
		result := combinationSum2_Backtrack(value.Candidates, value.Target)
		timeLapse := time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

		fmt.Println("Solution 2: DP")
		timeStart = time.Now()
		result = combinationSum2_DP(value.Candidates, value.Target)
		timeLapse = time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

		//		fmt.Println("Solution 3: BruteForce")
		//		timeStart = time.Now()
		//		result = combinationSum_BruteForce(value.Candidates, value.Target)
		//		timeLapse = time.Since(timeStart)
		//		fmt.Println(">Solution result", result)
		//		fmt.Println("Correct result is ", value.Result)
		//		fmt.Println("TimeLapse", timeLapse)

		fmt.Println("Solution 4: DFS")
		timeStart = time.Now()
		result = combinationSum2_DFS(value.Candidates, value.Target)
		timeLapse = time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

		//		fmt.Println("Solution 5: BFS => need check code again")
		//		timeStart = time.Now()
		//		result = combinationSum_BFS(value.Candidates, value.Target)
		//		timeLapse = time.Since(timeStart)
		//		fmt.Println(">Solution result", result)
		//		fmt.Println("Correct result is ", value.Result)
		//		fmt.Println("TimeLapse", timeLapse)
		//

	}

	timeLapsedWholeProgram := time.Since(timeStartWholeProgram)
	fmt.Println("===============")
	fmt.Println("TimeLapse Whole Program", timeLapsedWholeProgram)
}

type TestCase struct {
	Candidates []int
	Target     int
	Result     string
}

/*

===============
Test count  0 for node {[10 1 2 7 6 1 5] 8
[[1,1,6],[1,2,5],[1,7],[2,6]]
            }
Solution 1: Backtrack
>Solution result [[1 1 6] [1 2 5] [1 7] [2 6]]
Correct result is
[[1,1,6],[1,2,5],[1,7],[2,6]]

TimeLapse 9.778µs
Solution 2: DP
>Solution result [[1 2 5] [1 1 6] [2 6] [1 7]]
Correct result is
[[1,1,6],[1,2,5],[1,7],[2,6]]

TimeLapse 42.425µs
Solution 4: DFS
>Solution result [[1 1 6] [1 2 5] [1 7] [2 6]]
Correct result is
[[1,1,6],[1,2,5],[1,7],[2,6]]

TimeLapse 5.111µs
===============
Test count  1 for node {[2 5 2 1 2] 8
			[[1,2,2],[5]]
            }
Solution 1: Backtrack
>Solution result [[1 2 5]]
Correct result is
			[[1,2,2],[5]]

TimeLapse 2.796µs
Solution 2: DP
>Solution result [[1 2 5]]
Correct result is
			[[1,2,2],[5]]

TimeLapse 16.611µs
Solution 4: DFS
>Solution result [[1 2 5]]
Correct result is
			[[1,2,2],[5]]

TimeLapse 3.037µs
===============
Test count  2 for node {[2] 1
[]
            }
Solution 1: Backtrack
>Solution result []
Correct result is
[]

TimeLapse 1.389µs
Solution 2: DP
>Solution result []
Correct result is
[]

TimeLapse 1.648µs
Solution 4: DFS
>Solution result []
Correct result is
[]

TimeLapse 852ns
===============
TimeLapse Whole Program 740.194µs

*/
