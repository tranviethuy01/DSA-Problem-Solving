package main

import (
	"fmt"
	"time"
)

// approach: dynamic programming
// Time complexity: O(n^2).
// Space complexity: O(1)
func longestPalindrome_DP(s string) string {
	if len(s) <= 1 {
		return s
	}

	var start, end int

	for i := 0; i < len(s); i++ {
		len1 := expandAroundCenter(s, i, i)
		len2 := expandAroundCenter(s, i, i+1)
		maxLength := max(len1, len2)

		if maxLength > end-start {
			start = i - (maxLength-1)/2
			end = i + maxLength/2
		}
	}

	return s[start : end+1]
}

func expandAroundCenter(s string, left, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return right - left - 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//===== approach : brute force
//Time complexity: O(n^3).
//Space complexity: O(1)

func longestPalindrome_BruteForce(s string) string {
	if len(s) <= 1 {
		return s
	}

	var longest string

	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			substr := s[i : j+1]
			if isPalindrome(substr) && len(substr) > len(longest) {
				longest = substr
			}
		}
	}

	return longest
}

func isPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

//=====

func main() {
	timeStartWholeProgram := time.Now()
	testInput := []TestCase{
		{
			S: "babad",
			Result: `
                bab
            `,
		},
		{
			S: "cbbd",
			Result: `
                bb
            `,
		},
	}
	for count, value := range testInput {
		fmt.Println("===============")
		fmt.Println("Test count ", count, "for node", value)
		fmt.Println("Solution 1: use DP")
		timeStart := time.Now()
		result := longestPalindrome_DP(value.S)
		timeLapse := time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

		fmt.Println("Solution 2: use brute force solution")
		timeStart = time.Now()
		result = longestPalindrome_BruteForce(value.S)
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
Test count  0 for node {babad
                bab
            }
Solution 1: use DP
>Solution result aba
Correct result is
                bab

TimeLapse 908ns
Solution 2: use brute force solution
>Solution result bab
Correct result is
                bab

TimeLapse 8.963µs
===============
Test count  1 for node {cbbd
                bb
            }
Solution 1: use DP
>Solution result bb
Correct result is
                bb

TimeLapse 352ns
Solution 2: use brute force solution
>Solution result bb
Correct result is
                bb

TimeLapse 297ns
===============
TimeLapse Whole Program 425.091µs

*/
