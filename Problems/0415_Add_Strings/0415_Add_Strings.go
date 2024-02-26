package main

import (
	"fmt"
	"strings"
	"time"
)

//approach :
/*

Algorithm:

The algorithm iterates through both input strings once, from right to left.
At each iteration, it performs constant time operations such as converting characters to integers, adding them, and appending the result to a string builder.
The algorithm handles carries properly during addition.
Time Complexity:

Since the algorithm iterates through both input strings once regardless of their lengths, the time complexity is linear with respect to the lengths of the input strings.
Let n be the maximum length between num1 and num2. In the worst case, both strings have length n, so the time complexity is O(n).
Space Complexity:

The space complexity primarily depends on the space used by the string builder to store the result.
Apart from the string builder, the algorithm uses only a few integer variables (i, j, carry, sum), which occupy constant space.
The space used by the string builder is proportional to the length of the result string, which can be at most n + 1 (when there is a carry in the most significant digit).
Thus, the space complexity is O(n) since it scales linearly with the length of the result string.

*/

func addStrings(num1 string, num2 string) string {
	var sb strings.Builder

	i, j := len(num1)-1, len(num2)-1
	carry := 0

	for i >= 0 || j >= 0 || carry > 0 {
		sum := carry

		fmt.Println("i", i, "j", j, "carry", carry, "sum", sum)
		if i >= 0 {
			sum += int(num1[i] - '0')
			i--
		}
		if j >= 0 {
			sum += int(num2[j] - '0')
			j--
		}

		sb.WriteByte(byte(sum%10 + '0'))
		carry = sum / 10
	}

	// Reverse the string in the StringBuilder
	result := sb.String()
	fmt.Println(" sb ", sb, "result before Reverse", result)
	return reverseString(result)
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	timeStartWholeProgram := time.Now()

	testInput := []TestCase{

		{
			Num1: "11",
			Num2: "123",
			Result: `
"134"
            `,
		},

		{
			Num1: "456",
			Num2: "77",
			Result: `
"533"
            `,
		},

		{
			Num1: "0",
			Num2: "0",
			Result: `
"0"
            `,
		},
	}
	for count, value := range testInput {
		fmt.Println("===============")
		fmt.Println("Test count ", count, "for node", value)
		fmt.Println("Solution 1: manual add strings")
		timeStart := time.Now()
		result := addStrings(value.Num1, value.Num2)
		timeLapse := time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

	}

	timeLapsedWholeProgram := time.Since(timeStartWholeProgram)
	fmt.Println("===============")
	fmt.Println("TimeLapse Whole Program", timeLapsedWholeProgram)
}

type TestCase struct {
	Num1   string
	Num2   string
	Result string
}

/*

===============
Test count  0 for node {11 123
"134"
            }
Solution 1: manual add strings
i 1 j 2 carry 0 sum 0
i 0 j 1 carry 0 sum 0
i -1 j 0 carry 0 sum 0
 sb  {0x400008dc38 [52 51 49]} result before Reverse 431
>Solution result 134
Correct result is
"134"

TimeLapse 50.684µs
===============
Test count  1 for node {456 77
"533"
            }
Solution 1: manual add strings
i 2 j 1 carry 0 sum 0
i 1 j 0 carry 1 sum 1
i 0 j -1 carry 1 sum 1
 sb  {0x400008dc38 [51 51 53]} result before Reverse 335
>Solution result 533
Correct result is
"533"

TimeLapse 40.592µs
===============
Test count  2 for node {0 0
"0"
            }
Solution 1: manual add strings
i 0 j 0 carry 0 sum 0
 sb  {0x400008dc38 [48]} result before Reverse 0
>Solution result 0
Correct result is
"0"

TimeLapse 21.482µs
===============
TimeLapse Whole Program 475.644µs


*/
//REF
//
