package main

import (
	"fmt"
	"time"
	"strings"
)

// approach: Expand From Center
// Time complexity: O(n^2).
// Space complexity: O(1)
func longestPalindrome_ExpandFromCenter(s string) string {
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

//===== approach: dynamic programming
// approach: dynamic programming
//Time complexity : O(n^2). 
//Space complexity : O(n^2). It uses O(n^2) space to store the table.
func longestPalindrome_DP(s string) string {
    if len(s) <= 1 {
        return s
    }
    
    maxLen := 1
    maxStr := s[0:1]
    dp := make([][]bool, len(s))
    for i := range dp {
        dp[i] = make([]bool, len(s))
    }
    
    for i := 0; i < len(s); i++ {
        dp[i][i] = true
        for j := 0; j < i; j++ {
            if s[j] == s[i] && (i-j <= 2 || dp[j+1][i-1]) {
                dp[j][i] = true
                if i-j+1 > maxLen {
                    maxLen = i - j + 1
                    maxStr = s[j : i+1]
                }
            }
        }
    }
    
    return maxStr
}



//=====


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

//===== approach: Manacher approach
//Time complexity : O(n). Since expanding a palindrome around its center could take O(n) time, the overall complexity is O(n).
//Space complexity : O(n). It uses O(n) space to store the table.

func longestPalindrome_Manacher(s string) string {
    if len(s) <= 1 {
        return s
    }
    
    maxLen := 1
    maxStr := s[0:1]
    
    s = "#" + strings.Join(strings.Split(s, ""), "#") + "#"
    dp := make([]int, len(s))
    center := 0
    right := 0
    
    for i := 0; i < len(s); i++ {
        if i < right {
            dp[i] = min(right-i, dp[2*center-i])
        }
        
        for i-dp[i]-1 >= 0 && i+dp[i]+1 < len(s) && s[i-dp[i]-1] == s[i+dp[i]+1] {
            dp[i]++
        }
        
        if i+dp[i] > right {
            center = i
            right = i + dp[i]
        }
        
        if dp[i] > maxLen {
            maxLen = dp[i]
            maxStr = strings.ReplaceAll(s[i-dp[i]:i+dp[i]+1], "#", "")
        }
    }
    
    return maxStr
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}


//=====


//===== approach : recursive 
//Time complexity : O(n^3).
//Space complexity : O(n).

func longestPalindrome_Recursive(s string) string {
    if isPalindrome_R(s) {
        return s
    }
    
    left := longestPalindrome_Recursive(s[1:])
    right := longestPalindrome_Recursive(s[:len(s)-1])
    
    if len(left) > len(right) {
        return left
    } else {
        return right
    }
}

func isPalindrome_R(s string) bool {
    for i := 0; i < len(s)/2; i++ {
        if s[i] != s[len(s)-1-i] {
            return false
        }
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

		fmt.Println("Solution 2: use Expand From Center")
		timeStart = time.Now()
		result = longestPalindrome_ExpandFromCenter(value.S)
		timeLapse = time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

		fmt.Println("Solution 3: use brute force solution")
		timeStart = time.Now()
		result = longestPalindrome_BruteForce(value.S)
		timeLapse = time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

		fmt.Println("Solution 4: use Manacher approach")
		timeStart = time.Now()
		result = longestPalindrome_Manacher(value.S)
		timeLapse = time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

		fmt.Println("Solution 5: use Recurisve approach")
		timeStart = time.Now()
		result = longestPalindrome_Recursive(value.S)
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
>Solution result bab
Correct result is  
                bab
            
TimeLapse 13.593µs
Solution 2: use Expand From Center
>Solution result aba
Correct result is  
                bab
            
TimeLapse 963ns
Solution 3: use brute force solution
>Solution result bab
Correct result is  
                bab
            
TimeLapse 963ns
Solution 4: use Manacher approach
>Solution result bab
Correct result is  
                bab
            
TimeLapse 19.389µs
Solution 5: use Recurisve approach
>Solution result bab
Correct result is  
                bab
            
TimeLapse 667ns
===============
Test count  1 for node {cbbd 
                bb
            }
Solution 1: use DP
>Solution result bb
Correct result is  
                bb
            
TimeLapse 1.24µs
Solution 2: use Expand From Center
>Solution result bb
Correct result is  
                bb
            
TimeLapse 463ns
Solution 3: use brute force solution
>Solution result bb
Correct result is  
                bb
            
TimeLapse 426ns
Solution 4: use Manacher approach
>Solution result bb
Correct result is  
                bb
            
TimeLapse 13.371µs
Solution 5: use Recurisve approach
>Solution result bb
Correct result is  
                bb
            
TimeLapse 463ns
===============
TimeLapse Whole Program 795.577µs


*/
