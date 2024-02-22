package main

import (
	"fmt"
	"time"
)

//append : use stack
/*
Time Complexity (O(n)): The algorithm iterates through the string s once, processing each character in constant time. Each character is either pushed onto or popped off the stack at most once. Hence, the time complexity is linear in the size of the input string.

Space Complexity (O(n)): The stack used in the algorithm can have a maximum of n/2 elements, where n is the length of the input string s. This occurs when all opening parentheses are followed by closing parentheses. Therefore, the space complexity is linear in the size of the input string.
*/
func longestValidParentheses_Stack(s string) int {
	stack := make([]int, 0)
	maxLength := 0
	stack = append(stack, -1) // Push -1 to the stack as a base for the valid substring length calculation

	for i, char := range s {
		if char == '(' {
			stack = append(stack, i) // Push the index of '(' to the stack
		} else { // Encountered ')'
			stack = stack[:len(stack)-1] // Pop the last element from the stack
			if len(stack) == 0 {         // If the stack is empty
				stack = append(stack, i) // Push the current index to the stack as a base for the next valid substring
			} else { // If the stack is not empty
				length := i - stack[len(stack)-1] // Calculate the length of the current valid substring
				if length > maxLength {           // Update maxLength if necessary
					maxLength = length
				}
			}
		}
	}

	return maxLength
}

//approach dfs : this solution is a failure, need check code

func longestValidParentheses_DFS(s string) int {
	maxLength := 0
	dfs(s, 0, 0, &maxLength)
	return maxLength
}

func dfs(s string, start, openCount int, maxLength *int) {
	if start == len(s) {
		if openCount == 0 && *maxLength < len(s)-start {
			*maxLength = len(s) - start
		}
		return
	}

	if s[start] == '(' {
		// Include current '(' in the substring
		dfs(s, start+1, openCount+1, maxLength)
		// Exclude current '(' in the substring
		dfs(s, start+1, openCount, maxLength)
	} else if s[start] == ')' {
		if openCount > 0 {
			// Include current ')' in the substring
			dfs(s, start+1, openCount-1, maxLength)
		}
		// Exclude current ')' in the substring
		dfs(s, start+1, openCount, maxLength)
	} else {
		// If the character is neither '(' nor ')', move to the next character
		dfs(s, start+1, openCount, maxLength)
	}
}

//approach recursive
/*
Time Complexity:
The time complexity of the recursive solution is harder to analyze due to its nature of exploring all possible combinations of parentheses. However, we can provide an upper bound.
Let n be the length of the input string s. In the worst case, the recursive function is called
O(2 ^ n ) times because at each position, we have two choices: either to include or exclude the character. However, not all of these calls are executed to completion. Many of them are pruned due to constraints.

Space Complexity:
The space complexity of the recursive solution is O(n) where
n is the length of the input string s. This space is used for the recursive call stack.
Additionally, the recursive function maintains a few variables (index, openCount, closeCount, maxLength) which occupy constant space, hence not affecting the overall space complexity.

=> Time Complexity: The upper bound is
O(2 ^ n), but it's often much less due to pruning.
Space Complexity:
O(n) due to the space used by the recursive call stack.
*/

func longestValidParentheses_Recursive(s string) int {
	maxLength := 0
	findLongestValid(s, 0, 0, 0, &maxLength)
	return maxLength
}

func findLongestValid(s string, index, openCount, closeCount int, maxLength *int) {
	if index == len(s) {
		if openCount == closeCount && openCount*2 > *maxLength {
			*maxLength = openCount * 2
		}
		return
	}

	if s[index] == '(' {
		findLongestValid(s, index+1, openCount+1, closeCount, maxLength) // Include current '('
		findLongestValid(s, index+1, openCount, closeCount, maxLength)   // Exclude current '('
	} else if s[index] == ')' {
		if openCount > closeCount {
			findLongestValid(s, index+1, openCount, closeCount+1, maxLength) // Include current ')'
		}
		findLongestValid(s, index+1, openCount, closeCount, maxLength) // Exclude current ')'
	} else {
		findLongestValid(s, index+1, openCount, closeCount, maxLength) // Other characters, move to the next character
	}
}

//approach BFS, this code failure, need check
/*

Time Complexity:
Traversal: The BFS solution traverses each character in the input string once, checking if it's part of a valid substring. This traversal occurs in the isValid function. The worst-case scenario is that each character is visited once, leading to a time complexity of O(n), where n is the length of the input string s.
State Generation: At each step, the BFS algorithm generates two new states (representing the next index with or without including the current character). In the worst case, each state is generated once, leading to a time complexity of O(2n)=O(n).
Overall: Therefore, the overall time complexity of the BFS solution is O(n).
Space Complexity:
Queue: The BFS solution uses a queue to store states. In the worst case, the queue can contain at most n states, where n is the length of the input string s. Therefore, the space complexity due to the queue is O(n).
isValid Function: The isValid function uses a stack to check the validity of a substring. The size of the stack can be at most
n in the worst case (if all characters are opening parentheses). Thus, the space complexity due to the stack is also O(n).
Overall: Therefore, the overall space complexity of the BFS solution is O(n).
In summary:
Time Complexity: O(n)
Space Complexity: O(n)
The BFS solution provides a linear time and space complexity, making it efficient for solving the problem within the given constraints.




type State struct {
	index int // Current index in the string
	count int // Length of valid substring found so far
}

func longestValidParentheses(s string) int {
	queue := make([]State, 0)
	queue = append(queue, State{index: -1, count: 0}) // Initial state

	maxLength := 0

	for len(queue) > 0 {
		currState := queue[0]
		queue = queue[1:]

		if isValid(s, currState.index) {
			if currState.count > maxLength {
				maxLength = currState.count
			}

			// Extend the current valid substring by adding one more character
			queue = append(queue, State{index: currState.index + 1, count: currState.count + 1})
		}

		// Move to the next character
		if currState.index+1 < len(s) {
			queue = append(queue, State{index: currState.index + 1, count: currState.count})
		}
	}

	return maxLength
}

func isValid(s string, index int) bool {
	stack := 0
	count := 0

	for i := index; i < len(s); i++ {
		if s[i] == '(' {
			stack++
		} else if s[i] == ')' {
			stack--
			if stack < 0 {
				break
			}
			count += 2 // Valid substring found
			if stack == 0 && count > 0 {
				return true
			}
		}
	}

	return false
}

*/

//approach Dynamic Programing

/*

Time Complexity:
Iteration over String: The solution iterates through the input string once, performing constant-time operations at each step. Therefore, the time complexity of this part is
O(n), where
n is the length of the input string s.
Overall: The overall time complexity is
O(n).
Space Complexity:
Dynamic Programming Array: The solution uses a one-dimensional dynamic programming array of size
�
n, where
�
n is the length of the input string s. Therefore, the space complexity due to the dynamic programming array is
O(n).
Additional Space: Apart from the dynamic programming array, the solution uses only a few constant-space variables (maxLength, loop variables, etc.), which do not contribute significantly to the space complexity.
Overall: The overall space complexity is
O(n).
In summary:

Time Complexity:
O(n)
Space Complexity:
O(n)

*/

func longestValidParentheses_DP(s string) int {
	if len(s) <= 1 {
		return 0
	}

	dp := make([]int, len(s))
	maxLength := 0

	for i := 1; i < len(s); i++ {
		if s[i] == ')' {
			if s[i-1] == '(' {
				// Case: '()'
				dp[i] = 2
				if i >= 2 {
					dp[i] += dp[i-2]
				}
			} else if i-dp[i-1]-1 >= 0 && s[i-dp[i-1]-1] == '(' {
				// Case: '(...)'
				dp[i] = dp[i-1] + 2
				if i-dp[i-1]-2 >= 0 {
					dp[i] += dp[i-dp[i-1]-2]
				}
			}
			if dp[i] > maxLength {
				maxLength = dp[i]
			}
		}
	}

	return maxLength
}

//approach Brute Force
//Iterate over all starting indices of substrings in the input string.
//For each starting index, iterate over all possible lengths of substrings starting from that index.
//Check if each substring is a valid parentheses substring using a helper function.
//Update the maximum length of valid parentheses substrings found so far.
//Time Complexity: O(n^3)
//Space Complexity: O(n ^ 2)

func longestValidParentheses_BruteForce(s string) int {
	maxLength := 0

	for i := 0; i < len(s); i++ {
		for j := i + 2; j <= len(s); j += 2 {
			if isValid(s[i:j]) {
				maxLength = max(maxLength, j-i)
			}
		}
	}

	return maxLength
}

func isValid(s string) bool {
	stack := 0

	for _, char := range s {
		if char == '(' {
			stack++
		} else if char == ')' {
			if stack == 0 {
				return false
			}
			stack--
		}
	}

	return stack == 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	timeStartWholeProgram := time.Now()

	testInput := []TestCase{

		{
			S: "(()",
			Result: `
2
            `,
		},
		{
			S: ")()())",
			Result: `
4
            `,
		},
		{
			S: "",
			Result: `
			0
            `,
		},
	}
	for count, value := range testInput {
		fmt.Println("===============")
		fmt.Println("Test count ", count, "for node", value)
		fmt.Println("Solution 1: use Stack")
		timeStart := time.Now()
		result := longestValidParentheses_Stack(value.S)
		timeLapse := time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

		fmt.Println("Solution 2: DFS")
		timeStart = time.Now()
		result = longestValidParentheses_DFS(value.S)
		timeLapse = time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

		//		fmt.Println("Solution 3: BFS")
		//		timeStart = time.Now()
		//		result = longestValidParentheses_BFS(value.S)
		//		timeLapse = time.Since(timeStart)
		//		fmt.Println(">Solution result", result)
		//		fmt.Println("Correct result is ", value.Result)
		//		fmt.Println("TimeLapse", timeLapse)

		fmt.Println("Solution 4: DP")
		timeStart = time.Now()
		result = longestValidParentheses_DP(value.S)
		timeLapse = time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

		fmt.Println("Solution 5: Recursively")
		timeStart = time.Now()
		result = longestValidParentheses_Recursive(value.S)
		timeLapse = time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

		fmt.Println("Solution 6: BruteForce")
		timeStart = time.Now()
		result = longestValidParentheses_BruteForce(value.S)
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
	Result string
}

/*



===============
Test count  0 for node {(()
2
            }
Solution 1: use Stack
>Solution result 2
Correct result is
2

TimeLapse 1.371µs
Solution 2: DFS
>Solution result 0
Correct result is
2

TimeLapse 926ns
Solution 4: DP
>Solution result 2
Correct result is
2

TimeLapse 1.203µs
Solution 5: Recursively
>Solution result 2
Correct result is
2

TimeLapse 814ns
Solution 6: BruteForce
>Solution result 2
Correct result is
2

TimeLapse 666ns
===============
Test count  1 for node {)()())
4
            }
Solution 1: use Stack
>Solution result 4
Correct result is
4

TimeLapse 815ns
Solution 2: DFS
>Solution result 0
Correct result is
4

TimeLapse 833ns
Solution 4: DP
>Solution result 4
Correct result is
4

TimeLapse 611ns
Solution 5: Recursively
>Solution result 4
Correct result is
4

TimeLapse 834ns
Solution 6: BruteForce
>Solution result 4
Correct result is
4

TimeLapse 500ns
===============
Test count  2 for node {
			0
            }
Solution 1: use Stack
>Solution result 0
Correct result is
			0

TimeLapse 334ns
Solution 2: DFS
>Solution result 0
Correct result is
			0

TimeLapse 111ns
Solution 4: DP
>Solution result 0
Correct result is
			0

TimeLapse 111ns
Solution 5: Recursively
>Solution result 0
Correct result is
			0

TimeLapse 130ns
Solution 6: BruteForce
>Solution result 0
Correct result is
			0

TimeLapse 111ns
===============
TimeLapse Whole Program 793.487µs

*/
//REF
//
