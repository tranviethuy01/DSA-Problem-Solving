package main

import (
	"fmt"
	"time"
)

//approach : Backtrack
/*
The time complexity: O(N×N!), where N is the number of elements in the input array. This is because there are N! permutations, and for each permutation, we iterate over the array of length N to build it.

The space complexity is O(N×N!) due to the space used to store all the permutations. Additionally, there is auxiliary space used for the recursion stack, which grows up to O(N) due to the depth of recursive calls.
*/
func permute_Backtrack(nums []int) [][]int {
	var result [][]int
	used := make([]bool, len(nums))
	backtrack(nums, []int{}, &result, used)
	return result
}

func backtrack(nums []int, current []int, result *[][]int, used []bool) {
	if len(current) == len(nums) {
		temp := make([]int, len(current))
		copy(temp, current)
		*result = append(*result, temp)
		return
	}

	for i := 0; i < len(nums); i++ {
		if !used[i] {
			used[i] = true
			current = append(current, nums[i])
			backtrack(nums, current, result, used)
			current = current[:len(current)-1]
			used[i] = false
		}
	}
}

// approach : DFS
// The time and space complexity for this DFS-based solution remain the same as the previous solution: O(N×N!).
func permute_DFS(nums []int) [][]int {
	var result [][]int
	visited := make([]bool, len(nums))
	current := []int{}

	dfs(nums, visited, current, &result)
	return result
}

func dfs(nums []int, visited []bool, current []int, result *[][]int) {
	if len(current) == len(nums) {
		temp := make([]int, len(current))
		copy(temp, current)
		*result = append(*result, temp)
		return
	}

	for i := 0; i < len(nums); i++ {
		if !visited[i] {
			visited[i] = true
			current = append(current, nums[i])
			dfs(nums, visited, current, result)
			current = current[:len(current)-1]
			visited[i] = false
		}
	}
}

// approach : BFS
// The time and space complexity of this solution is:O(N×N!), where N is the number of elements in the input array.
func permute_BFS(nums []int) [][]int {
	var result [][]int
	queue := [][]int{{}}

	for i := 0; i < len(nums); i++ {
		next := nums[i]
		for _, permutation := range queue {
			for j := 0; j <= len(permutation); j++ {
				newPermutation := make([]int, len(permutation)+1)
				copy(newPermutation[:j], permutation[:j])
				newPermutation[j] = next
				copy(newPermutation[j+1:], permutation[j:])
				result = append(result, newPermutation)
			}
		}
		queue = result
		result = [][]int{}
	}

	return queue
}

// approach : dynamic programing
// The time and space complexity of this solution is:O(N×N!), where N is the number of elements in the input array.
func permute_DP(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{{}}
	}

	var result [][]int
	for i, num := range nums {
		// Exclude num from nums and get permutations of the rest
		rest := append(append([]int{}, nums[:i]...), nums[i+1:]...)
		permutations := permute_DP(rest)

		// Append num to each permutation of the rest
		for _, p := range permutations {
			result = append(result, append([]int{num}, p...))
		}
	}
	return result
}

//approach : BruteForce
//time complexity of this solution is O(N!), where N is the number of elements in the input array. This is because there are N! possible permutations of N elements. The space complexity is also  O(N!) due to the space required to store all these permutations.

func permute_BruteForce(nums []int) [][]int {
	var result [][]int
	generatePermutations(nums, 0, &result)
	return result
}

func generatePermutations(nums []int, start int, result *[][]int) {
	if start == len(nums)-1 {
		temp := make([]int, len(nums))
		copy(temp, nums)
		*result = append(*result, temp)
		return
	}

	for i := start; i < len(nums); i++ {
		// Swap elements at indices start and i
		nums[start], nums[i] = nums[i], nums[start]
		// Recursively generate permutations for the rest of the array
		generatePermutations(nums, start+1, result)
		// Restore the original order by swapping back
		nums[start], nums[i] = nums[i], nums[start]
	}
}

func main() {
	timeStartWholeProgram := time.Now()

	testInput := []TestCase{

		{
			Nums: []int{1, 2, 3},
			Result: `
[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
            `,
		},
		{
			Nums: []int{0, 1},
			Result: `
			[[0,1],[1,0]]
            `,
		},
		{
			Nums: []int{1},
			Result: `
			[[1]]
            `,
		},
	}
	for count, value := range testInput {
		fmt.Println("===============")
		fmt.Println("Test count ", count, "for node", value)
		fmt.Println("Solution 1: Backtrack")
		timeStart := time.Now()
		result := permute_Backtrack(value.Nums)
		timeLapse := time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

		fmt.Println("Solution 2: DFS")
		timeStart = time.Now()
		result = permute_DFS(value.Nums)
		timeLapse = time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

		fmt.Println("Solution 3: BFS")
		timeStart = time.Now()
		result = permute_BFS(value.Nums)
		timeLapse = time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

		fmt.Println("Solution 4: DP")
		timeStart = time.Now()
		result = permute_DP(value.Nums)
		timeLapse = time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

		fmt.Println("Solution 5: BruteForce")
		timeStart = time.Now()
		result = permute_BruteForce(value.Nums)
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
Test count  0 for node {[1 2 3] 
[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
            }
Solution 1: Backtrack
>Solution result [[1 2 3] [1 3 2] [2 1 3] [2 3 1] [3 1 2] [3 2 1]]
Correct result is  
[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
            
TimeLapse 12.685µs
Solution 2: DFS
>Solution result [[1 2 3] [1 3 2] [2 1 3] [2 3 1] [3 1 2] [3 2 1]]
Correct result is  
[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
            
TimeLapse 4.685µs
Solution 3: BFS
>Solution result [[3 2 1] [2 3 1] [2 1 3] [3 1 2] [1 3 2] [1 2 3]]
Correct result is  
[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
            
TimeLapse 4.667µs
Solution 4: DP
>Solution result [[1 2 3] [1 3 2] [2 1 3] [2 3 1] [3 1 2] [3 2 1]]
Correct result is  
[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
            
TimeLapse 10.093µs
Solution 5: BruteForce
>Solution result [[1 2 3] [1 3 2] [2 1 3] [2 3 1] [3 2 1] [3 1 2]]
Correct result is  
[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
            
TimeLapse 2.537µs
===============
Test count  1 for node {[0 1] 
			[[0,1],[1,0]]
            }
Solution 1: Backtrack
>Solution result [[0 1] [1 0]]
Correct result is  
			[[0,1],[1,0]]
            
TimeLapse 1.741µs
Solution 2: DFS
>Solution result [[0 1] [1 0]]
Correct result is  
			[[0,1],[1,0]]
            
TimeLapse 1.667µs
Solution 3: BFS
>Solution result [[1 0] [0 1]]
Correct result is  
			[[0,1],[1,0]]
            
TimeLapse 1.814µs
Solution 4: DP
>Solution result [[0 1] [1 0]]
Correct result is  
			[[0,1],[1,0]]
            
TimeLapse 2.649µs
Solution 5: BruteForce
>Solution result [[0 1] [1 0]]
Correct result is  
			[[0,1],[1,0]]
            
TimeLapse 1.333µs
===============
Test count  2 for node {[1] 
			[[1]]
            }
Solution 1: Backtrack
>Solution result [[1]]
Correct result is  
			[[1]]
            
TimeLapse 1.019µs
Solution 2: DFS
>Solution result [[1]]
Correct result is  
			[[1]]
            
TimeLapse 1.074µs
Solution 3: BFS
>Solution result [[1]]
Correct result is  
			[[1]]
            
TimeLapse 907ns
Solution 4: DP
>Solution result [[1]]
Correct result is  
			[[1]]
            
TimeLapse 1.018µs
Solution 5: BruteForce
>Solution result [[1]]
Correct result is  
			[[1]]
            
TimeLapse 704ns
===============
TimeLapse Whole Program 924.015µs


 */
//REF
//
