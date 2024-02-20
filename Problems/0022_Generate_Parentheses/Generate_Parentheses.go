package main

import (
	"fmt"
	"time"
)

//approach: backtracking
//Time Complexity: O(4^n)
//Space Complexity: O(n)

func generateParenthesis_Backtracking(n int) []string {
	var result []string
	backtrack(&result, "", 0, 0, n)
	return result
}

func backtrack(result *[]string, current string, open, close, max int) {
	if len(current) == max*2 {
		*result = append(*result, current)
		return
	}

	if open < max {
		backtrack(result, current+"(", open+1, close, max)
	}
	if close < open {
		backtrack(result, current+")", open, close+1, max)
	}
}

//==== approach: dynamic programing
//Time Complexity: O(n^3)
//Space Complexity: O(n * k), where k is upper bounded by the total number of valid combinations.

func generateParenthesis_DP(n int) []string {
	if n == 0 {
		return []string{""}
	}

	dp := make([][]string, n+1)
	dp[0] = []string{""}
	dp[1] = []string{"()"}

	for i := 2; i <= n; i++ {
		var combinations []string
		for j := 0; j < i; j++ {
			for _, first := range dp[j] {
				for _, second := range dp[i-j-1] {
					combinations = append(combinations, "("+first+")"+second)
				}
			}
		}
		dp[i] = combinations
	}

	return dp[n]
}

//=====

// ===== approach :dfs
// Time Complexity: O(4^n)
// Space Complexity: O(n)
func generateParenthesis_DFS(n int) []string {
	var result []string
	dfs(&result, "", 0, 0, n)
	return result
}

func dfs(result *[]string, current string, open, close, max int) {
	if len(current) == max*2 {
		*result = append(*result, current)
		return
	}

	if open < max {
		dfs(result, current+"(", open+1, close, max)
	}
	if close < open {
		dfs(result, current+")", open, close+1, max)
	}
}

// ====

//=== approach : BFS

type Node struct {
	str   string
	open  int
	close int
}

func generateParenthesis_BFS(n int) []string {
	var result []string

	queue := []Node{{"", 0, 0}}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if len(node.str) == n*2 {
			result = append(result, node.str)
		} else {
			if node.open < n {
				queue = append(queue, Node{node.str + "(", node.open + 1, node.close})
			}
			if node.close < node.open {
				queue = append(queue, Node{node.str + ")", node.open, node.close + 1})
			}
		}
	}

	return result
}

//===

func main() {
	timeStartWholeProgram := time.Now()

	testInput := []TestCase{
		{
			N: 3,
			Result: `
["((()))","(()())","(())()","()(())","()()()"]

            `,
		},
		{
			N: 1,
			Result: `
["()"]


            `,
		},
	}
	for count, value := range testInput {
		fmt.Println("===============")
		fmt.Println("Test count ", count, "for node", value)
		fmt.Println("Solution 1: ")
		timeStart := time.Now()
		result := generateParenthesis_Backtracking(value.N)
		timeLapse := time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("")
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

		fmt.Println("Solution 2: DP ")
		timeStart = time.Now()
		result = generateParenthesis_DP(value.N)
		timeLapse = time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("")
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

		fmt.Println("Solution 3: DFS ")
		timeStart = time.Now()
		result = generateParenthesis_DFS(value.N)
		timeLapse = time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("")
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

		fmt.Println("Solution 4: BFS")
		timeStart = time.Now()
		result = generateParenthesis_BFS(value.N)
		timeLapse = time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("")
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

	}

	timeLapsedWholeProgram := time.Since(timeStartWholeProgram)
	fmt.Println("===============")
	fmt.Println("TimeLapse Whole Program", timeLapsedWholeProgram)
}

type TestCase struct {
	N      int
	Result string
}

/*

===============
Test count  0 for node {3
["((()))","(()())","(())()","()(())","()()()"]

            }
Solution 1:
>Solution result [((())) (()()) (())() ()(()) ()()()]

Correct result is
["((()))","(()())","(())()","()(())","()()()"]


TimeLapse 19.816µs
Solution 2:
>Solution result [()()() ()(()) (())() (()()) ((()))]

Correct result is
["((()))","(()())","(())()","()(())","()()()"]


TimeLapse 5.463µs
===============
Test count  1 for node {1
["()"]


            }
Solution 1:
>Solution result [()]

Correct result is
["()"]



TimeLapse 1µs
Solution 2:
>Solution result [()]

Correct result is
["()"]



TimeLapse 704ns
===============
TimeLapse Whole Program 451.468µs

*/
//REF
//
