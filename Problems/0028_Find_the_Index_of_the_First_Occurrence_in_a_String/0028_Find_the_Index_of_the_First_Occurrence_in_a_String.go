package main

import (
	"fmt"
	"time"
)

// approach : sliding window
// Time Complexity: O(n * m), where n is the length of the haystack and m is the length of the needle
// Space Complexity: O(1)
func strStr_SlidingWindow(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	for i := 0; i <= len(haystack)-len(needle); i++ {
		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}
	return -1
}

//approach : brute force
//Time Complexity: O((n - m + 1) * m).
//Space Complexity: O(1)

func strStr_BruteForce(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	for i := 0; i <= len(haystack)-len(needle); i++ {
		found := true
		for j := 0; j < len(needle); j++ {
			if haystack[i+j] != needle[j] {
				found = false
				break
			}
		}
		if found {
			return i
		}
	}
	return -1
}

func main() {
	timeStartWholeProgram := time.Now()

	testInput := []TestCase{

		{
			Haystack: "sadbutsad",
			Needle:   "sad",
			Result: `
0

            `,
		},
		{
			Haystack: "leetcode",
			Needle:   "leeto",
			Result: `
      -1
            `,
		},
	}
	for count, value := range testInput {
		fmt.Println("===============")
		fmt.Println("Test count ", count, "for node", value)
		fmt.Println("Solution 1: 2 pointer")
		timeStart := time.Now()
		result := strStr_SlidingWindow(value.Haystack, value.Needle)
		timeLapse := time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

		fmt.Println("Solution 2: brute force")
		timeStart = time.Now()
		result = strStr_BruteForce(value.Haystack, value.Needle)
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
	Haystack string
	Needle   string
	Result   string
}

/*



===============
Test count  0 for node {sadbutsad sad
0

            }
Solution 1: 2 pointer
>Solution result 0
Correct result is
0


TimeLapse 944ns
Solution 2: brute force
>Solution result 0
Correct result is
0


TimeLapse 315ns
===============
Test count  1 for node {leetcode leeto
      -1
            }
Solution 1: 2 pointer
>Solution result -1
Correct result is
      -1

TimeLapse 241ns
Solution 2: brute force
>Solution result -1
Correct result is
      -1

TimeLapse 186ns
===============
TimeLapse Whole Program 389.497µs

*/
//REF
//
