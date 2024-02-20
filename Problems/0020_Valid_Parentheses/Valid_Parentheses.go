package main

import (
	"fmt"

	"time"
)

//Time Complexity: O(n)
//Space Complexity: O(n)

func isValid(s string) bool {
	stack := make([]rune, 0)

	// Define a mapping of closing brackets to their corresponding opening brackets
	mapping := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, char := range s {
		// If the character is an opening bracket, push it onto the stack
		if char == '(' || char == '[' || char == '{' {
			stack = append(stack, char)
		} else {
			// If the stack is empty or the top of the stack doesn't match the current closing bracket,
			// the string is invalid
			if len(stack) == 0 || stack[len(stack)-1] != mapping[char] {
				return false
			}
			// Pop the top element from the stack
			stack = stack[:len(stack)-1]
		}
	}

	// If the stack is empty, all brackets have been properly matched and string is valid
	return len(stack) == 0
}

func main() {
	timeStartWholeProgram := time.Now()
	testInput := []TestCase{
		{
			S: "()",
			Result: `
true
            `,
		},
		{
			S: "()[]{}",
			Result: `
         true
            `,
		},
		{
			S: "(]",
			Result: `
         false
            `,
		},
	}
	for count, value := range testInput {
		fmt.Println("===============")
		fmt.Println("Test count ", count, "for node", value)
		fmt.Println("Solution 1: ")
		timeStart := time.Now()
		result := isValid(value.S)
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
	S      string
	Result string
}

/*

===============
Test count  0 for node {()
true
            }
Solution 1:
>Solution result true
Correct result is
true

TimeLapse 3.926µs
===============
Test count  1 for node {()[]{}
         true
            }
Solution 1:
>Solution result true
Correct result is
         true

TimeLapse 1.185µs
===============
Test count  2 for node {(]
         false
            }
Solution 1:
>Solution result false
Correct result is
         false

TimeLapse 889ns
===============
TimeLapse Whole Program 347.34µs


*/
//REF
//
