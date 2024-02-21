package main

import (
	"fmt"
	"time"
)

//Time Complexity:  O(n)
//Space complexity: O(1) .

func canBeValid(s string, locked string) bool {
	validate := func(s string, locked string, op byte) bool {
		bal, wild := 0, 0
		for i := 0; i < len(s); i++ {
			if locked[i] == '1' {
				if s[i] == op {
					bal++
				} else {
					bal--
				}
			} else {
				wild++
			}
			if wild+bal < 0 {
				return false
			}
		}
		return bal <= wild
	}

	return len(s)%2 == 0 && validate(s, locked, '(') && validate(reverseString(s), reverseString(locked), ')')
}

func reverseString(s string) string {
	runeS := []rune(s)
	for i, j := 0, len(runeS)-1; i < j; i, j = i+1, j-1 {
		runeS[i], runeS[j] = runeS[j], runeS[i]
	}
	return string(runeS)
}

/*

failure solution and need check

func canBeValid(s string, locked string) bool {
	open := 0
	for i := range s {
		if s[i] == '(' {
			open++
		} else if s[i] == ')' {
			if open > 0 {
				open--
			} else if locked[i] == '0' {
				return false
			}
		} else if locked[i] == '0' {
			open++
		}
	}
	return open == 0
}



*/

func main() {
	timeStartWholeProgram := time.Now()

	testInput := []TestCase{
		{
			S:      "))()))",
			Locked: "010100",
			Result: `
	true
            `,
		},
		{
			S:      "()()",
			Locked: "0000",
			Result: `
      true

            `,
		},
		{

			S:      ")",
			Locked: "0",
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
		result := canBeValid(value.S, value.Locked)
		timeLapse := time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("")
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)
	}

	timeLapsedWholeProgram := time.Since(timeStartWholeProgram)
	fmt.Println("===============")
	fmt.Println("TimeLapse Whole Program", timeLapsedWholeProgram)
}

type TestCase struct {
	S      string
	Locked string
	Result string
}

/*


===============
Test count  0 for node {))())) 010100
	true
            }
Solution 1:
>Solution result true

Correct result is
	true

TimeLapse 3.111µs
===============
Test count  1 for node {()() 0000
      true

            }
Solution 1:
>Solution result true

Correct result is
      true


TimeLapse 796ns
===============
Test count  2 for node {) 0
false
            }
Solution 1:
>Solution result false

Correct result is
false

TimeLapse 167ns
===============
TimeLapse Whole Program 378.403µs

*/
//REF
//
