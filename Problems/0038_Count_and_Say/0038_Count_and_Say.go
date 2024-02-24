package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

//approach Recursive
/*

Time Complexity:

The time complexity of this algorithm is O(2^n), where n is the input parameter. This is because each recursive call potentially doubles the length of the string.
Space Complexity:

The space complexity is also O(2^n), as each recursive call adds a new string of potentially double the length compared to the previous one. Additionally, there is space used for the function call stack due to recursion.

*/

func countAndSay_Recursive(n int) string {
	if n == 1 {
		return "1"
	}

	prev := countAndSay_Recursive(n - 1)
	var result strings.Builder
	count := 1

	for i := 0; i < len(prev); i++ {
		if i+1 < len(prev) && prev[i] == prev[i+1] {
			count++
		} else {
			result.WriteString(strconv.Itoa(count))
			result.WriteByte(prev[i])
			count = 1
		}
	}

	return result.String()
}

//approach DP
/*
Time Complexity:

Generating each term of the count-and-say sequence requires iterating through the previous term, which can have a maximum length proportional to the current term. So, for each term from 2 to n, the time complexity is proportional to the length of the previous term, and since we iterate up to n terms, the overall time complexity is O(n * m), where n is the input parameter and m is the maximum length of the string generated at any point.
Space Complexity:

The space complexity of this dynamic programming solution is O(n * m), where n is the input parameter and m is the maximum length of the string generated at any point. We store the count-and-say sequences up to the nth term in the dp array, and each term can have a maximum length of m. Additionally, there is space used for other variables, but they are constant relative to the input size.
*/

func countAndSay_DP(n int) string {
	if n == 1 {
		return "1"
	}

	dp := make([]string, n)
	dp[0] = "1"

	for i := 1; i < n; i++ {
		prev := dp[i-1]
		var result strings.Builder
		count := 1

		for j := 0; j < len(prev); j++ {
			if j+1 < len(prev) && prev[j] == prev[j+1] {
				count++
			} else {
				result.WriteString(strconv.Itoa(count))
				result.WriteByte(prev[j])
				count = 1
			}
		}

		dp[i] = result.String()
	}

	return dp[n-1]
}

func main() {
	timeStartWholeProgram := time.Now()

	testInput := []TestCase{

		{
			N: 1,
			Result: `
1
            `,
		},

		{
			N: 4,
			Result: `
1211
            `,
		},
	}
	for count, value := range testInput {

		fmt.Println("===============")
		fmt.Println("Test count ", count, "for node", value)
		fmt.Println("Solution 1: Backtrack")
		timeStart := time.Now()
		result := countAndSay_Recursive(value.N)
		timeLapse := time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

		fmt.Println("Solution 2: DFS")
		timeStart = time.Now()
		result = countAndSay_DP(value.N)
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
	N      int
	Result string
}

/*

===============
Test count  0 for node {1
1
            }
Solution 1: Backtrack
>Solution result 1
Correct result is
1

TimeLapse 297ns
Solution 2: DFS
>Solution result 1
Correct result is
1

TimeLapse 389ns
===============
Test count  1 for node {4
1211
            }
Solution 1: Backtrack
>Solution result 1211
Correct result is
1211

TimeLapse 1.999µs
Solution 2: DFS
>Solution result 1211
Correct result is
1211

TimeLapse 1.981µs
===============
TimeLapse Whole Program 361.353µs

*/
