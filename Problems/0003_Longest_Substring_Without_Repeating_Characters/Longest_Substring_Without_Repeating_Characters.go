package main

import (
	"fmt"
	"time"
	//"sort"
)

// sliding window approach
// Time complexity: O(n)
// Space complexity: O(n)
func lengthOfLongestSubstring(s string) int {
	// Create a map to store the last index of each character encountered
	lastSeen := make(map[byte]int)
	maxLength, start := 0, 0

	for end := 0; end < len(s); end++ {
		fmt.Println("lastSeen", lastSeen)
		// If the character is already seen, update the start index
		if idx, found := lastSeen[s[end]]; found && idx >= start {
			start = idx + 1
		}
		// Update the last seen index of the character
		lastSeen[s[end]] = end
		// Update the maximum length
		if end-start+1 > maxLength {
			maxLength = end - start + 1
		}
	}

	return maxLength
}

//====== brute force approach
//Time complexity: O(n^2)
//Space complexity: O(min(n, m))

func lengthOfLongestSubstringBruteForce(s string) int {
	maxLength := 0

	for i := 0; i < len(s); i++ {
		seen := make(map[byte]bool)
		currentLength := 0

		for j := i; j < len(s); j++ {
			if seen[s[j]] {
				break
			}
			seen[s[j]] = true
			currentLength++
		}

		if currentLength > maxLength {
			maxLength = currentLength
		}
	}

	return maxLength
}

//======

func main() {
	timeStartWholeProgram := time.Now()
	testInput := []TestCase{
		{
			S: "abcabcbb",
			Result: `
                3
            `,
		},
		{
			S: "bbbbb",
			Result: `
                1
            `,
		},
		{
			S: "pwwkew",
			Result: `
                3
            `,
		},
	}
	for count, value := range testInput {
		fmt.Println("===============")
		fmt.Println("Test count ", count, "for node", value)
		fmt.Println("Solution 1")
		timeStart := time.Now()
		result := lengthOfLongestSubstring(value.S)
		timeLapse := time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

		fmt.Println("Solution 2: brute force")
		timeStart = time.Now()
		result = lengthOfLongestSubstringBruteForce(value.S)
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
Test count  0 for node {abcabcbb
                3
            }
Solution 1
>Solution result 3
Correct result is
                3

TimeLapse 4.852µs
Solution 2: brute force
>Solution result 3
Correct result is
                3

TimeLapse 4.278µs
===============
Test count  1 for node {bbbbb
                1
            }
Solution 1
>Solution result 1
Correct result is
                1

TimeLapse 907ns
Solution 2: brute force
>Solution result 1
Correct result is
                1

TimeLapse 1.167µs
===============
Test count  2 for node {pwwkew
                3
            }
Solution 1
>Solution result 3
Correct result is
                3

TimeLapse 1.278µs
Solution 2: brute force
>Solution result 3
Correct result is
                3

TimeLapse 1.982µs
===============
TimeLapse Whole Program 471.254µs

*/
