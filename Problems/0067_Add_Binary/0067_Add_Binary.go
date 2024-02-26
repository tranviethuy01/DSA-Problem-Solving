package main

import (
	"fmt"
	"time"
)

//

/*
Time complexity of O(n), where n is the maximum length between the two input strings. This is because it iterates through the strings once, performing constant-time operations for each character.

The space complexity is also O(n), where n is the length of the longer input string. This is because the algorithm creates a new string to store the result, which can be of length n in the worst case.

*/

func addBinary(a string, b string) string {
	lenA, lenB := len(a), len(b)
	maxLen := max(lenA, lenB)

	// Make the lengths of a and b equal by adding leading zeros if necessary
	if lenA < maxLen {
		a = addLeadingZeros(a, maxLen-lenA)
	} else if lenB < maxLen {
		b = addLeadingZeros(b, maxLen-lenB)
	}
	fmt.Printf("a %s b %s \n", a, b)
	carry := 0
	sum := ""
	for i := maxLen - 1; i >= 0; i-- {
		digitA := int(a[i] - '0')
		digitB := int(b[i] - '0')
		currentSum := digitA + digitB + carry

		fmt.Printf("i %d, digitA %d digitB %d , currentSum %d , currentCarry %d \n", i, digitA, digitB, currentSum, carry)
		sum = fmt.Sprintf("%d%s", currentSum%2, sum)
		carry = currentSum / 2

		fmt.Printf("sum %s , carry after %d \n", sum, carry)

	}

	if carry > 0 {
		fmt.Println("carry >0, carry =", carry)
		sum = fmt.Sprintf("%d%s", carry, sum)
	}

	return sum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func addLeadingZeros(s string, count int) string {
	for i := 0; i < count; i++ {
		s = "0" + s
	}
	return s
}

func main() {
	timeStartWholeProgram := time.Now()

	testInput := []TestCase{

		{
			A: "11",
			B: "1",

			Result: `
"100"
            `,
		},
		{
			A: "1010",
			B: "1011",
			Result: `
"10101"
            `,
		},
	}
	for count, value := range testInput {
		fmt.Println("===============")
		fmt.Println("Test count ", count, "for node", value)
		fmt.Println("Solution 1: manual add binary")
		timeStart := time.Now()
		result := addBinary(value.A, value.B)
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
	A      string
	B      string
	Result string
}

/*

===============
Test count  0 for node {11 1
"100"
            }
Solution 1: manual add binary
a 11 b 01
i 1, digitA 1 digitB 1 , currentSum 2 , currentCarry 0
sum 0 , carry after 1
i 0, digitA 1 digitB 0 , currentSum 2 , currentCarry 1
sum 00 , carry after 1
carry >0, carry = 1
>Solution result 100
Correct result is
"100"

TimeLapse 65.056µs
===============
Test count  1 for node {1010 1011
"10101"
            }
Solution 1: manual add binary
a 1010 b 1011
i 3, digitA 0 digitB 1 , currentSum 1 , currentCarry 0
sum 1 , carry after 0
i 2, digitA 1 digitB 1 , currentSum 2 , currentCarry 0
sum 01 , carry after 1
i 1, digitA 0 digitB 0 , currentSum 1 , currentCarry 1
sum 101 , carry after 0
i 0, digitA 1 digitB 1 , currentSum 2 , currentCarry 0
sum 0101 , carry after 1
carry >0, carry = 1
>Solution result 10101
Correct result is
"10101"

TimeLapse 96.833µs
===============
TimeLapse Whole Program 472.593µs

*/
//REF
//
