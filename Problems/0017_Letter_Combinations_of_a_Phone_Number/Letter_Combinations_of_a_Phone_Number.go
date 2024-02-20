package main

import (
	"fmt"
	"time"
)

//==== backtracking

var phoneMap_Backtracking = map[byte][]string{
	'2': {"a", "b", "c"},
	'3': {"d", "e", "f"},
	'4': {"g", "h", "i"},
	'5': {"j", "k", "l"},
	'6': {"m", "n", "o"},
	'7': {"p", "q", "r", "s"},
	'8': {"t", "u", "v"},
	'9': {"w", "x", "y", "z"},
}

func letterCombinations_Backtracking(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	result := make([]string, 0)
	backtrack(digits, "", &result)
	return result
}

func backtrack(digits string, combination string, result *[]string) {
	if len(digits) == 0 {
		*result = append(*result, combination)
		return
	}

	digit := digits[0]
	letters := phoneMap_Backtracking[digit]
	for _, letter := range letters {
		backtrack(digits[1:], combination+letter, result)
	}
}

//====

//==dfs

var phoneMap_DFS = map[byte][]string{
	'2': {"a", "b", "c"},
	'3': {"d", "e", "f"},
	'4': {"g", "h", "i"},
	'5': {"j", "k", "l"},
	'6': {"m", "n", "o"},
	'7': {"p", "q", "r", "s"},
	'8': {"t", "u", "v"},
	'9': {"w", "x", "y", "z"},
}

func letterCombinations_DFS(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	result := make([]string, 0)
	dfs(digits, "", 0, &result)
	return result
}

func dfs(digits string, combination string, index int, result *[]string) {
	if index == len(digits) {
		*result = append(*result, combination)
		return
	}

	digit := digits[index]
	letters := phoneMap_DFS[digit]
	for _, letter := range letters {
		dfs(digits, combination+letter, index+1, result)
	}
}

//====

//==== bfs

var phoneMap_BFS = map[byte][]string{
	'2': {"a", "b", "c"},
	'3': {"d", "e", "f"},
	'4': {"g", "h", "i"},
	'5': {"j", "k", "l"},
	'6': {"m", "n", "o"},
	'7': {"p", "q", "r", "s"},
	'8': {"t", "u", "v"},
	'9': {"w", "x", "y", "z"},
}

func letterCombinations_BFS(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	queue := []string{""}
	for i := 0; i < len(digits); i++ {
		nextQueue := make([]string, 0)
		digit := digits[i]
		letters := phoneMap_BFS[digit]

		for _, prefix := range queue {
			for _, letter := range letters {
				nextQueue = append(nextQueue, prefix+letter)
			}
		}
		queue = nextQueue
	}

	return queue
}

//===

//==== approach : dynamic programing

var phoneMap_DP = map[byte][]string{
	'2': {"a", "b", "c"},
	'3': {"d", "e", "f"},
	'4': {"g", "h", "i"},
	'5': {"j", "k", "l"},
	'6': {"m", "n", "o"},
	'7': {"p", "q", "r", "s"},
	'8': {"t", "u", "v"},
	'9': {"w", "x", "y", "z"},
}

func letterCombinations_DP(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	memo := make(map[string][]string)
	return dp(digits, memo)
}

func dp(digits string, memo map[string][]string) []string {
	if len(digits) == 0 {
		return []string{""}
	}

	if result, ok := memo[digits]; ok {
		return result
	}

	digit := digits[0]
	letters := phoneMap_DP[digit]
	combinations := make([]string, 0)

	remainingCombinations := dp(digits[1:], memo)
	for _, letter := range letters {
		for _, remaining := range remainingCombinations {
			combinations = append(combinations, string(letter)+remaining)
		}
	}

	memo[digits] = combinations
	return combinations
}

//====

func main() {
	timeStartWholeProgram := time.Now()
	testInput := []TestCase{
		{
			Digits: "23",
			Result: `
          ["ad","ae","af","bd","be","bf","cd","ce","cf"]
            `,
		},
		{
			Digits: "",
			Result: `
         []
            `,
		},

		{
			Digits: "2",
			Result: `
          ["a","b","c"]
            `,
		},
	}
	for count, value := range testInput {
		fmt.Println("===============")
		fmt.Println("Test count ", count, "for node", value)
		fmt.Println("Solution 1: use Backtracking")
		timeStart := time.Now()
		result := letterCombinations_Backtracking(value.Digits)
		timeLapse := time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

		fmt.Println("Solution 2: use DFS")
		timeStart = time.Now()
		result = letterCombinations_DFS(value.Digits)
		timeLapse = time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

		fmt.Println("Solution 3: use BFS")
		timeStart = time.Now()
		result = letterCombinations_BFS(value.Digits)
		timeLapse = time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

		fmt.Println("Solution 4: use DP")
		timeStart = time.Now()
		result = letterCombinations_DP(value.Digits)
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
	Digits string
	Result string
}

/*


===============
Test count  0 for node {23
          ["ad","ae","af","bd","be","bf","cd","ce","cf"]
            }
Solution 1: use Backtracking
>Solution result [ad ae af bd be bf cd ce cf]
Correct result is
          ["ad","ae","af","bd","be","bf","cd","ce","cf"]

TimeLapse 20.148µs
Solution 2: use DFS
>Solution result [ad ae af bd be bf cd ce cf]
Correct result is
          ["ad","ae","af","bd","be","bf","cd","ce","cf"]

TimeLapse 4.241µs
Solution 3: use BFS
>Solution result [ad ae af bd be bf cd ce cf]
Correct result is
          ["ad","ae","af","bd","be","bf","cd","ce","cf"]

TimeLapse 4.445µs
Solution 4: use DP
>Solution result [ad ae af bd be bf cd ce cf]
Correct result is
          ["ad","ae","af","bd","be","bf","cd","ce","cf"]

TimeLapse 5.649µs
===============
Test count  1 for node {
         []
            }
Solution 1: use Backtracking
>Solution result []
Correct result is
         []

TimeLapse 111ns
Solution 2: use DFS
>Solution result []
Correct result is
         []

TimeLapse 111ns
Solution 3: use BFS
>Solution result []
Correct result is
         []

TimeLapse 111ns
Solution 4: use DP
>Solution result []
Correct result is
         []

TimeLapse 92ns
===============
Test count  2 for node {2
          ["a","b","c"]
            }
Solution 1: use Backtracking
>Solution result [a b c]
Correct result is
          ["a","b","c"]

TimeLapse 1.611µs
Solution 2: use DFS
>Solution result [a b c]
Correct result is
          ["a","b","c"]

TimeLapse 1.574µs
Solution 3: use BFS
>Solution result [a b c]
Correct result is
          ["a","b","c"]

TimeLapse 1.666µs
Solution 4: use DP
>Solution result [a b c]
Correct result is
          ["a","b","c"]

TimeLapse 1.926µs
===============
TimeLapse Whole Program 773.146µs


*/
//REF
//
