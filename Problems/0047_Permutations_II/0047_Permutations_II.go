package main

import (
	"fmt"
	"sort"
	"time"
)

//approach Backtrack
/*

Time Complexity:
Sorting the input array initially takes O(n * log(n)) time, where n is the number of elements in the array.
Generating all permutations using backtracking takes O(n!), where n is the number of elements in the array. However, due to the duplicate pruning, the actual number of recursive calls is less than n! in the worst case.
So, the overall time complexity is O(n! * n * log(n)).
Space Complexity:

The space complexity is dominated by the recursive call stack during backtracking, which can go up to O(n) in depth.
Additionally, we have additional space for the output result, which can contain up to O(n!) permutations.
Therefore, the overall space complexity is O(n! * n).
*/

func permuteUnique_Backtrack(nums []int) [][]int {
	var res [][]int
	sort.Ints(nums)
	backtrack(&res, nums, []int{}, make([]bool, len(nums)))
	return res
}

func backtrack(res *[][]int, nums, perm []int, used []bool) {
	if len(perm) == len(nums) {
		*res = append(*res, append([]int{}, perm...))
		return
	}
	for i := 0; i < len(nums); i++ {
		if used[i] || (i > 0 && nums[i] == nums[i-1] && !used[i-1]) {
			continue
		}
		used[i] = true
		perm = append(perm, nums[i])
		backtrack(res, nums, perm, used)
		used[i] = false
		perm = perm[:len(perm)-1]
	}
}

//approach DFS
/*
Time Complexity:
Sorting the input array initially takes O(n * log(n)) time, where n is the number of elements in the array.
Generating all permutations using DFS takes O(n!), where n is the number of elements in the array. However, due to the duplicate pruning, the actual number of recursive calls is less than n! in the worst case.
So, the overall time complexity is O(n! * n * log(n)).
Space Complexity:
The space complexity is dominated by the recursive call stack during DFS, which can go up to O(n) in depth.
Additionally, we have additional space for the output result, which can contain up to O(n!) permutations.
Therefore, the overall space complexity is O(n! * n).
*/
func permuteUnique_DFS(nums []int) [][]int {
    var result [][]int
    var path []int
    var visited = make([]bool, len(nums))
    sort.Ints(nums)
    dfs(nums, &result, path, visited)
    return result
}

func dfs(nums []int, result *[][]int, path []int, visited []bool) {
    if len(path) == len(nums) {
        *result = append(*result, append([]int{}, path...))
        return
    }
    for i := 0; i < len(nums); i++ {
        if visited[i] || (i > 0 && nums[i] == nums[i-1] && !visited[i-1]) {
            continue
        }
        visited[i] = true
        dfs(nums, result, append(path, nums[i]), visited)
        visited[i] = false
    }
}

// approach BFS
/*

Time Complexity:
Sorting the input array initially takes O(n * log(n)) time, where n is the number of elements in the array.
Generating permutations using BFS involves exploring all possible permutations level by level. At each level, the number of permutations that need to be generated is proportional to the number of permutations at the previous level.
Considering there are n! permutations in total, and each permutation generation operation takes O(n) time, the overall time complexity is O(n! * n).
Space Complexity:

The space complexity is dominated by the result array that stores all the permutations. At each level of the BFS traversal, the number of permutations stored in the result array grows exponentially.
Additionally, there's space required for the temporary permutations generated at each level.
Therefore, the overall space complexity is also O(n! * n), where n is the number of elements in the array.
In summary, the provided BFS solution has a time complexity of O(n! * n) and a space complexity of O(n! * n).

*/

func permuteUnique_BFS(nums []int) [][]int {
    sort.Ints(nums)
    result := [][]int{{}}
    for _, num := range nums {
        nextLevel := [][]int{}
        for _, perm := range result {
            for i := 0; i <= len(perm); i++ {
                if i > 0 && perm[i-1] == num {
                    break
                }
                nextPerm := make([]int, len(perm)+1)
                copy(nextPerm[:i], perm[:i])
                nextPerm[i] = num
                copy(nextPerm[i+1:], perm[i:])
                nextLevel = append(nextLevel, nextPerm)
            }
        }
        result = nextLevel
    }
    return result
}

/*
//approach Dynamic Program : this code need to fix
time complexity is O(n! * n^2).
space complexity is O(n! * n)

func permuteUnique_DP(nums []int) [][]int {
	sort.Ints(nums)
	dp := make([][][]int, len(nums)+1)
	dp[0] = [][]int{{}}

	for i := 1; i <= len(nums); i++ {
		dp[i] = [][]int{}
		for j := 0; j < i; j++ {
			for _, prev := range dp[j] {
				newPerm := make([]int, len(prev)+1)
				copy(newPerm, prev[:j])
				newPerm[j] = nums[i-1]
				copy(newPerm[j+1:], prev[j:])
				if !contains(dp[i], newPerm) {
					dp[i] = append(dp[i], newPerm)
				}
			}
		}
	}

	return dp[len(nums)]
}

func contains(arr [][]int, perm []int) bool {
	for _, p := range arr {
		if equal(p, perm) {
			return true
		}
	}
	return false
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

*/
//approach BruteForce
/*
Time Complexity:
Sorting the input array initially takes O(n * log(n)) time, where n is the number of elements in the array.
Generating permutations using backtracking involves exploring all possible combinations, which results in a time complexity of O(n!), where n is the number of elements in the array.
However, due to duplicate pruning, the actual number of recursive calls is less than n! in the worst case.
Therefore, the overall time complexity is O(n! * n * log(n)).
Space Complexity:
The space complexity is dominated by the recursive call stack during backtracking, which can go up to O(n) in depth.
Additionally, we have additional space for the output result, which can contain up to O(n!) permutations.
Therefore, the overall space complexity is O(n! * n).
In summary, the time complexity of the provided solution is O(n! * n * log(n)), and the space complexity is O(n! * n).
*/


func permuteUnique_BruteForce(nums []int) [][]int {
    var result [][]int
    var visited = make([]bool, len(nums))
    var current []int
    sort.Ints(nums)
    generatePermutations(nums, visited, current, &result)
    return result
}

func generatePermutations(nums []int, visited []bool, current []int, result *[][]int) {
    if len(current) == len(nums) {
        *result = append(*result, append([]int{}, current...))
        return
    }
    for i := 0; i < len(nums); i++ {
        if visited[i] || (i > 0 && nums[i] == nums[i-1] && !visited[i-1]) {
            continue
        }
        visited[i] = true
        generatePermutations(nums, visited, append(current, nums[i]), result)
        visited[i] = false
    }
}





func main() {
	timeStartWholeProgram := time.Now()

	testInput := []TestCase{

		{
			Nums: []int{1, 1, 2},
			Result: `
[[1,1,2],[1,2,1],[2,1,1]]
            `,
		},
		{
			Nums: []int{1, 2, 3},
			Result: `
			[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
            `,
		},
	}
	for count, value := range testInput {
		fmt.Println("===============")
		fmt.Println("Test count ", count, "for node", value)
		fmt.Println("Solution 1: Backtrack")
		timeStart := time.Now()
		result := permuteUnique_Backtrack(value.Nums)
		timeLapse := time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

		fmt.Println("Solution 2: DFS")
		timeStart = time.Now()
		result = permuteUnique_DFS(value.Nums)
		timeLapse = time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

		fmt.Println("Solution 3: BFS")
		timeStart = time.Now()
		result = permuteUnique_BFS(value.Nums)
		timeLapse = time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)


		fmt.Println("Solution 4: BruteForce")
		timeStart = time.Now()
		result = permuteUnique_BruteForce(value.Nums)
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
	Nums   []int
	Result string
}

/*


===============
Test count  0 for node {[1 1 2] 
[[1,1,2],[1,2,1],[2,1,1]]
            }
Solution 1: Backtrack
>Solution result [[1 1 2] [1 2 1] [2 1 1]]
Correct result is  
[[1,1,2],[1,2,1],[2,1,1]]
            
TimeLapse 6.444µs
Solution 2: DFS
>Solution result [[1 1 2] [1 2 1] [2 1 1]]
Correct result is  
[[1,1,2],[1,2,1],[2,1,1]]
            
TimeLapse 4.277µs
Solution 3: BFS
>Solution result [[2 1 1] [1 2 1] [1 1 2]]
Correct result is  
[[1,1,2],[1,2,1],[2,1,1]]
            
TimeLapse 3.426µs
Solution 4: BruteForce
>Solution result [[1 1 2] [1 2 1] [2 1 1]]
Correct result is  
[[1,1,2],[1,2,1],[2,1,1]]
            
TimeLapse 4.333µs
===============
Test count  1 for node {[1 2 3] 
			[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
            }
Solution 1: Backtrack
>Solution result [[1 2 3] [1 3 2] [2 1 3] [2 3 1] [3 1 2] [3 2 1]]
Correct result is  
			[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
            
TimeLapse 11.63µs
Solution 2: DFS
>Solution result [[1 2 3] [1 3 2] [2 1 3] [2 3 1] [3 1 2] [3 2 1]]
Correct result is  
			[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
            
TimeLapse 4.352µs
Solution 3: BFS
>Solution result [[3 2 1] [2 3 1] [2 1 3] [3 1 2] [1 3 2] [1 2 3]]
Correct result is  
			[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
            
TimeLapse 3.277µs
Solution 4: BruteForce
>Solution result [[1 2 3] [1 3 2] [2 1 3] [2 3 1] [3 1 2] [3 2 1]]
Correct result is  
			[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
            
TimeLapse 4.222µs
===============
TimeLapse Whole Program 712.881µs

 */
//REF
//
