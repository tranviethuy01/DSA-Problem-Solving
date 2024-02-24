package main

import (
	"fmt"
	"sort"
	"time"
)

//approach Backtrack
/*
Time Complexity:

In the worst-case scenario, the algorithm explores all possible combinations of candidates to find those that sum up to the target. Since each candidate can be included or excluded in each step of the backtracking process, there are 2^n possible combinations, where n is the number of candidates.
Within each recursive call, the algorithm performs constant-time operations (such as appending to slices and copying slices). Therefore, the time complexity of each recursive call is O(1).
Hence, the overall time complexity of the algorithm is O(2^n), where n is the number of candidates.
Space Complexity:

The space complexity mainly comes from the recursion stack and the additional space used to store the valid combinations.
In the worst case, the depth of the recursion stack is equal to the number of candidates, which contributes O(n) space complexity.
Additionally, the space required to store each valid combination is also considered. Since there can be up to 2^n valid combinations (in the worst case), the space complexity for storing these combinations is also O(2^n).
Therefore, the overall space complexity of the algorithm is O(2^n).
*/
func combinationSum_Backtrack(candidates []int, target int) [][]int {
	var result [][]int
	backtrack(candidates, target, []int{}, &result, 0)
	return result
}

func backtrack(candidates []int, target int, current []int, result *[][]int, start int) {
	if target < 0 {
		return
	}
	if target == 0 {
		temp := make([]int, len(current))
		copy(temp, current)
		*result = append(*result, temp)
		return
	}
	for i := start; i < len(candidates); i++ {
		current = append(current, candidates[i])
		backtrack(candidates, target-candidates[i], current, result, i)
		current = current[:len(current)-1]
	}
}

// approach DP
/*
Time Complexity:

Sorting the candidates takes O(n log n), where n is the number of candidates.
The dynamic programming approach fills a table of size target, and for each entry in the table, it iterates through the candidates. The maximum possible value for target is target, and the number of candidates can be up to n.
Therefore, the time complexity of the dynamic programming part is O(n * target).
Overall, considering both sorting and dynamic programming, the time complexity is O(n log n + n * target).
Space Complexity:

The space complexity is primarily determined by the dynamic programming table dp. It has dimensions target+1 by number of combinations.
Since we store combinations in the DP table, the space required can be significant. In the worst case, if all numbers from 1 to target are present in the input candidates, there can be target combinations.
Therefore, the space complexity of the DP table is O(target * number of combinations).
Additionally, sorting the candidates requires O(n) extra space.
Overall, the space complexity is O(n + target * number of combinations).
In summary, the time complexity of the provided dynamic programming solution is O(n log n + n * target), and the space complexity is O(n + target * number of combinations).

*/

func combinationSum_DP(candidates []int, target int) [][]int {
	sort.Ints(candidates) // Sort the candidates to handle duplicates
	dp := make([][][]int, target+1)
	for i := 1; i <= target; i++ {
		var combinations [][]int
		for _, num := range candidates {
			if num > i {
				break
			}
			if num == i {
				combinations = append(combinations, []int{num})
			} else {
				for _, prevComb := range dp[i-num] {
					if num <= prevComb[0] { // Ensure no duplicates
						combinations = append(combinations, append([]int{num}, prevComb...))
					}
				}
			}
		}
		dp[i] = combinations
	}
	return dp[target]
}

//approach : Brute Force . This brute-force solution recursively generates all possible combinations of candidates and checks if their sum equals the target. If so, it adds the combination to the result.

/*
Algorithm:

The algorithm explores all possible combinations of candidates recursively.
At each step of the recursion, it either includes or excludes a candidate, exploring both possibilities.
When the current sum equals the target, the combination is added to the result.
Time Complexity:

In the worst case, the algorithm explores all possible combinations of candidates.
Each candidate can be included or excluded, so there are 2 choices for each candidate.
Therefore, the time complexity is exponential, O(2^n), where n is the number of candidates.
Space Complexity:

The space complexity is primarily determined by the recursion stack.
In the worst case, the depth of the recursion stack is equal to the number of candidates.
Additionally, the space required to store each valid combination is considered.
Therefore, the space complexity is also exponential, O(2^n).


*/

func combinationSum_BruteForce(candidates []int, target int) [][]int {
	var result [][]int
	sort.Ints(candidates) // Sort the candidates to handle duplicates
	findCombinations_BruteForce(candidates, target, []int{}, &result, 0)
	return result
}

func findCombinations_BruteForce(candidates []int, target int, current []int, result *[][]int, index int) {
	if target == 0 {
		temp := make([]int, len(current))
		copy(temp, current)
		*result = append(*result, temp)
		return
	}

	for i := index; i < len(candidates) && candidates[i] <= target; i++ {
		current = append(current, candidates[i])
		findCombinations_BruteForce(candidates, target-candidates[i], current, result, i)
		current = current[:len(current)-1]
	}
}

// approach DFS
/*
Algorithm:
The algorithm explores all possible combinations of candidates recursively.
At each step of the recursion, it either includes or excludes a candidate, exploring both possibilities.
When the current sum equals the target, the combination is added to the result.
Time Complexity:

In the worst case, the algorithm explores all possible combinations of candidates.
Each candidate can be included or excluded, so there are 2 choices for each candidate.
Therefore, the time complexity is exponential, O(2^n), where n is the number of candidates.
Space Complexity:

The space complexity is primarily determined by the recursion stack.
In the worst case, the depth of the recursion stack is equal to the number of candidates.
Additionally, the space required to store each valid combination is considered.
Therefore, the space complexity is also exponential, O(2^n).

*/

func combinationSum_DFS(candidates []int, target int) [][]int {
	var result [][]int
	dfs(candidates, target, []int{}, &result, 0)
	return result
}

func dfs(candidates []int, target int, current []int, result *[][]int, start int) {
	if target < 0 {
		return
	}
	if target == 0 {
		temp := make([]int, len(current))
		copy(temp, current)
		*result = append(*result, temp)
		return
	}

	for i := start; i < len(candidates); i++ {
		current = append(current, candidates[i])
		dfs(candidates, target-candidates[i], current, result, i)
		current = current[:len(current)-1]
	}
}

// approach BFS  : need to check code again
/*
func combinationSum(candidates []int, target int) [][]int {
    queue := make([][][]int, 0)
    result := make([][]int, 0)
    queue = append(queue, [][]int{{}, {target}})

    for len(queue) > 0 {
        node := queue[0]
        queue = queue[1:]

        prev := node[0]
        remaining := node[1][0]

        for _, candidate := range candidates {
            if candidate > remaining {
                continue
            }
            if candidate < prev[len(prev)-1] {
                continue
            }

            if candidate == remaining {
                temp := make([]int, len(prev))
                copy(temp, prev)
                temp = append(temp, candidate)
                result = append(result, temp)
            } else {
                next := make([]int, len(prev))
                copy(next, prev)
                next = append(next, candidate)
                queue = append(queue, [][]int{next, []int{candidate, remaining - candidate}})
            }
        }
    }
    return result
}

*/

func main() {
	timeStartWholeProgram := time.Now()

	testInput := []TestCase{
		{
			Candidates: []int{2, 3, 6, 7},
			Target:     7,
			Result: `
[[2,2,3],[7]]
            `,
		},

		{
			Candidates: []int{2, 3, 5},
			Target:     8,
			Result: `
[[2,2,2,2],[2,3,3],[3,5]]
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
		result := combinationSum_Backtrack(value.Candidates, value.Target)
		timeLapse := time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

		fmt.Println("Solution 2: DP")
		timeStart = time.Now()
		result = combinationSum_DP(value.Candidates, value.Target)
		timeLapse = time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

		fmt.Println("Solution 3: BruteForce")
		timeStart = time.Now()
		result = combinationSum_BruteForce(value.Candidates, value.Target)
		timeLapse = time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

		fmt.Println("Solution 4: DFS")
		timeStart = time.Now()
		result = combinationSum_DFS(value.Candidates, value.Target)
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
Test count  0 for node {[2 3 6 7] 7
[[2,2,3],[7]]
            }
Solution 1: Backtrack
>Solution result [[2 2 3] [7]]
Correct result is
[[2,2,3],[7]]

TimeLapse 5µs
Solution 2: DP
>Solution result [[2 2 3] [7]]
Correct result is
[[2,2,3],[7]]

TimeLapse 18.481µs
Solution 3: BruteForce
>Solution result [[2 2 3] [7]]
Correct result is
[[2,2,3],[7]]

TimeLapse 3.296µs
Solution 4: DFS
>Solution result [[2 2 3] [7]]
Correct result is
[[2,2,3],[7]]

TimeLapse 3.574µs
===============
Test count  1 for node {[2 3 5] 8
[[2,2,2,2],[2,3,3],[3,5]]
            }
Solution 1: Backtrack
>Solution result [[2 2 2 2] [2 3 3] [3 5]]
Correct result is
[[2,2,2,2],[2,3,3],[3,5]]

TimeLapse 3.759µs
Solution 2: DP
>Solution result [[2 2 2 2] [2 3 3] [3 5]]
Correct result is
[[2,2,2,2],[2,3,3],[3,5]]

TimeLapse 14.648µs
Solution 3: BruteForce
>Solution result [[2 2 2 2] [2 3 3] [3 5]]
Correct result is
[[2,2,2,2],[2,3,3],[3,5]]

TimeLapse 3.296µs
Solution 4: DFS
>Solution result [[2 2 2 2] [2 3 3] [3 5]]
Correct result is
[[2,2,2,2],[2,3,3],[3,5]]

TimeLapse 3.333µs
===============
Test count  2 for node {[2] 1
[]
            }
Solution 1: Backtrack
>Solution result []
Correct result is
[]

TimeLapse 704ns
Solution 2: DP
>Solution result []
Correct result is
[]

TimeLapse 1.037µs
Solution 3: BruteForce
>Solution result []
Correct result is
[]

TimeLapse 648ns
Solution 4: DFS
>Solution result []
Correct result is
[]

TimeLapse 556ns
===============
TimeLapse Whole Program 974.454µs


*/
