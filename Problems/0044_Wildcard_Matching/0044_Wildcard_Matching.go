package main

import (
	"fmt"
	"time"
)

//approach : dynamic programming
/*
Time Complexity: The time complexity of the algorithm is O(m * n), where m is the length of string s and n is the length of pattern p. This is because we iterate through each character of s and p once, and each cell of the dynamic programming table requires constant time to update.

Space Complexity: The space complexity is O(m * n) as well, where m is the length of string s and n is the length of pattern p. This is because we create a dynamic programming table of size (len(s) + 1) * (len(p) + 1) to store the results of subproblems.
*/

func isMatch_DP(s string, p string) bool {
	// Initialize dp table
	dp := make([][]bool, len(s)+1)
	for i := range dp {
		dp[i] = make([]bool, len(p)+1)
	}

	// Empty string and empty pattern match
	dp[0][0] = true

	// Handling patterns like "*", "*", "a*" etc.
	for j := 1; j <= len(p); j++ {
		if p[j-1] == '*' {
			dp[0][j] = dp[0][j-1]
		}
	}

	// Fill the dp table
	for i := 1; i <= len(s); i++ {
		for j := 1; j <= len(p); j++ {
			if p[j-1] == '?' || s[i-1] == p[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else if p[j-1] == '*' {
				dp[i][j] = dp[i-1][j] || dp[i][j-1]
			}
		}
	}

	return dp[len(s)][len(p)]
}

//approach backtrack
/*
Algorithm: The algorithm explores all possible matches between the string s and the pattern p using backtracking. It advances the pointers in s and p based on the matching conditions of characters and *. The function recursively explores all possible combinations until a match is found or all possibilities are exhausted.

Time Complexity: In the worst case, the time complexity of this algorithm is exponential, O(2^(m+n)), where m is the length of string s and n is the length of pattern p. This is because for each character in p, if it's not a wildcard '*', we explore both matching and non-matching cases, resulting in exponential time complexity.

Space Complexity: The space complexity is O(m + n), where m is the length of string s and n is the length of pattern p. This space is used for the recursive call stack during the backtracking process.

*/
func isMatch_Backtrack(s string, p string) bool {
	return backtrack(s, p, 0, 0)
}

func backtrack(s, p string, sIndex, pIndex int) bool {
	if pIndex == len(p) {
		return sIndex == len(s)
	}

	if p[pIndex] == '*' {
		// Move the pattern pointer, or the string pointer
		for ; pIndex < len(p) && p[pIndex] == '*'; pIndex++ {
		}

		// If we reached the end of the pattern, it matches regardless of the remaining string
		if pIndex == len(p) {
			return true
		}

		// Start from the current position in the string, try to match rest of string and pattern
		for i := sIndex; i < len(s); i++ {
			if backtrack(s, p, i, pIndex) {
				return true
			}
		}
		return false
	} else if sIndex < len(s) && (p[pIndex] == '?' || p[pIndex] == s[sIndex]) {
		// If the current characters match or pattern is '?', move both pointers
		return backtrack(s, p, sIndex+1, pIndex+1)
	}

	return false
}

//approach DFS
/*

Algorithm: The algorithm explores all possible matches between the string s and the pattern p using depth-first search. It recursively explores all possible combinations until a match is found or all possibilities are exhausted.

Time Complexity: In the worst case, the time complexity of this algorithm is exponential, O(2^(m+n)), where m is the length of string s and n is the length of pattern p. This is because for each character in p, if it's not a wildcard '*', the algorithm explores all possible matches in the string s.

Space Complexity: The space complexity is O(m + n), where m is the length of string s and n is the length of pattern p. This space is used for the recursive call stack during the depth-first search process. Additionally, there is a constant amount of space used for other variables in the function.
*/

func isMatch_DFS(s string, p string) bool {
	return dfs(s, p, 0, 0)
}

func dfs(s, p string, sIndex, pIndex int) bool {
	if pIndex == len(p) {
		return sIndex == len(s)
	}

	if p[pIndex] == '*' {
		// Skip consecutive '*' in pattern
		for pIndex < len(p) && p[pIndex] == '*' {
			pIndex++
		}

		// If '*' is the last character, it matches the remaining string
		if pIndex == len(p) {
			return true
		}

		// Try all possible matches for the remaining string
		for i := sIndex; i <= len(s); i++ {
			if dfs(s, p, i, pIndex) {
				return true
			}
		}
		return false
	} else if sIndex < len(s) && (p[pIndex] == '?' || p[pIndex] == s[sIndex]) {
		// If current characters match or pattern is '?', move both pointers
		return dfs(s, p, sIndex+1, pIndex+1)
	}

	return false
}

//approach BFS
/*
Algorithm: The algorithm iteratively explores all possible matches between the string s and the pattern p using breadth-first search. It maintains a queue of states (sIndex, pIndex) representing the current indices in s and p being compared. At each step, it examines the current state, updates the queue with possible next states, and continues until a match is found or all possibilities are exhausted.

Time Complexity: In the worst case, the time complexity of this algorithm is exponential, O(2^(m+n)), where m is the length of string s and n is the length of pattern p. This is because for each character in p, if it's not a wildcard '*', the algorithm explores all possible matches in the string s.

Space Complexity: The space complexity is also exponential, O(2^(m+n)), due to the queue that can potentially hold up to 2^(m+n) states. Each state requires constant space, but the number of states grows exponentially with the input sizes. Additionally, there is a constant amount of space used for other variables in the function.

*/

func isMatch_BFS(s string, p string) bool {
	queue := [][]int{{0, 0}}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		sIndex, pIndex := node[0], node[1]

		if pIndex == len(p) {
			if sIndex == len(s) {
				return true
			}
			continue
		}

		if p[pIndex] == '*' {
			// Move to the next character in the pattern
			queue = append(queue, []int{sIndex, pIndex + 1})

			// Move to the next character in the string
			if sIndex < len(s) {
				queue = append(queue, []int{sIndex + 1, pIndex})
			}
		} else if sIndex < len(s) && (p[pIndex] == '?' || p[pIndex] == s[sIndex]) {
			// Move to the next characters in both string and pattern
			queue = append(queue, []int{sIndex + 1, pIndex + 1})
		}
	}

	return false
}

//approach BruteForce

func isMatch_BruteForce(s string, p string) bool {
	return matchHelper(s, p, 0, 0)
}

func matchHelper(s, p string, sIndex, pIndex int) bool {
	// If both string and pattern have reached the end, it's a match
	if sIndex == len(s) && pIndex == len(p) {
		return true
	}

	// If pattern has reached the end but string hasn't, it's not a match
	if pIndex == len(p) {
		return false
	}

	// If string has reached the end but pattern hasn't, check if remaining pattern contains only '*'
	if sIndex == len(s) {
		for i := pIndex; i < len(p); i++ {
			if p[i] != '*' {
				return false
			}
		}
		return true
	}

	// If current characters match or pattern character is '?', move both pointers
	if s[sIndex] == p[pIndex] || p[pIndex] == '?' {
		return matchHelper(s, p, sIndex+1, pIndex+1)
	}

	// If pattern character is '*', try all possible matches by advancing pattern pointer
	if p[pIndex] == '*' {
		for i := 0; sIndex+i <= len(s); i++ {
			if matchHelper(s, p, sIndex+i, pIndex+1) {
				return true
			}
		}
	}

	return false
}

func main() {
	timeStartWholeProgram := time.Now()

	testInput := []TestCase{

		{
			S: "aa",
			P: "a",
			Result: `
      false
            `,
		},
		{
			S: "aa",
			P: "*",
			Result: `
true
            `,
		},
		{
			S: "cb",
			P: "?a",
			Result: `
false
            `,
		},
	}
	for count, value := range testInput {
		fmt.Println("===============")
		fmt.Println("Test count ", count, "for node", value)
		fmt.Println("Solution 1: DP")
		timeStart := time.Now()
		result := isMatch_DP(value.S, value.P)
		timeLapse := time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

		fmt.Println("Solution 2: Backtrack")
		timeStart = time.Now()
		result = isMatch_Backtrack(value.S, value.P)
		timeLapse = time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

		fmt.Println("Solution 3: DFS")
		timeStart = time.Now()
		result = isMatch_DFS(value.S, value.P)
		timeLapse = time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

		fmt.Println("Solution 4: BFS")
		timeStart = time.Now()
		result = isMatch_BFS(value.S, value.P)
		timeLapse = time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

		fmt.Println("Solution 5: BruteForce")
		timeStart = time.Now()
		result = isMatch_BruteForce(value.S, value.P)
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
	S      string
	P      string
	Result string
}

/*

===============
Test count  0 for node {aa a
      false
            }
Solution 1: DP
>Solution result false
Correct result is
      false

TimeLapse 2.167µs
Solution 2: Backtrack
>Solution result false
Correct result is
      false

TimeLapse 371ns
Solution 3: DFS
>Solution result false
Correct result is
      false

TimeLapse 388ns
Solution 4: BFS
>Solution result false
Correct result is
      false

TimeLapse 1.759µs
Solution 5: BruteForce
>Solution result false
Correct result is
      false

TimeLapse 611ns
===============
Test count  1 for node {aa *
true
            }
Solution 1: DP
>Solution result true
Correct result is
true

TimeLapse 1.056µs
Solution 2: Backtrack
>Solution result true
Correct result is
true

TimeLapse 148ns
Solution 3: DFS
>Solution result true
Correct result is
true

TimeLapse 240ns
Solution 4: BFS
>Solution result true
Correct result is
true

TimeLapse 2.278µs
Solution 5: BruteForce
>Solution result true
Correct result is
true

TimeLapse 333ns
===============
Test count  2 for node {cb ?a
false
            }
Solution 1: DP
>Solution result false
Correct result is
false

TimeLapse 926ns
Solution 2: Backtrack
>Solution result false
Correct result is
false

TimeLapse 186ns
Solution 3: DFS
>Solution result false
Correct result is
false

TimeLapse 148ns
Solution 4: BFS
>Solution result false
Correct result is
false

TimeLapse 778ns
Solution 5: BruteForce
>Solution result false
Correct result is
false

TimeLapse 130ns
===============
TimeLapse Whole Program 775.194µs

*/
//REF
//
