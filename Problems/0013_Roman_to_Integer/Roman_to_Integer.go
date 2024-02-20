package main

import (
	"fmt"
	"time"
)

// approach Straightforward, loop from back to front and using a map
func romanToInt_MapString(s string) int {
	romanMap := map[string]int{
		"M": 1000,
		"D": 500,
		"C": 100,
		"L": 50,
		"X": 10,
		"V": 5,
		"I": 1,
	}

	result := 0

	prev := 0

	for i := len(s) - 1; i >= 0; i-- {
		value := romanMap[string(s[i])]

		if value < prev {
			result -= value
		} else {
			result += value
		}

		prev = value
	}

	return result
}

//=====

// =====
// use map[byte]int instead of map string
func romanToInt_MapByte(s string) int {
	romanMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	result := 0
	prev := 0

	for i := len(s) - 1; i >= 0; i-- {
		value := romanMap[s[i]]

		if value < prev {
			result -= value
		} else {
			result += value
		}

		prev = value
	}

	return result
}

//=====

//===== approach loop front to end

func romanToInt_FrontToEnd(s string) int {
	romanMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	result := 0
	prev := 0

	for i := 0; i < len(s); i++ {
		value := romanMap[s[i]]

		/* Explain
				prev represents the integer value of the previous Roman numeral character encountered.
		If the current value (value) is greater than the previous value (prev), it indicates a special case where subtraction is involved in Roman numerals. In Roman numerals, a smaller value precedes a larger value, and when this happens, it means we need to subtract the smaller value from the larger one. For example, in "IV", the "I" comes before "V", so we subtract 1 from 5 to get 4.
		To perform this subtraction in the code, instead of subtracting prev from value, we subtract it twice. This is because we've already added prev in the previous iteration of the loop, and now we need to subtract it again. Hence, value - 2 * prev.
		*/
		if value > prev {
			result += value - 2*prev
		} else {
			result += value
		}

		prev = value
	}

	return result
}

//=====

//==== another approach : just directly include all the values for roman here

func romanToInt_IncludeAll(s string) int {

	romanMap := map[string]int{
		"I":  1,
		"V":  5,
		"X":  10,
		"L":  50,
		"C":  100,
		"D":  500,
		"M":  1000,
		"IV": 4,
		"IX": 9,
		"XL": 40,
		"XC": 90,
		"CD": 400,
		"CM": 900,
	}

	result := 0

	for i := 0; i < len(s); i++ {
		//check for case the next letter is bigger then the current
		if i+1 < len(s) && romanMap[string(s[i])] < romanMap[string(s[i+1])] {
			result += romanMap[s[i:i+2]]
			i++
		} else {
			result += romanMap[string(s[i])]
		}

		//		if i+1 < len(s) && romanMap[s[i:i+2]] != 0 {
		//			result += romanMap[s[i:i+2]]
		//			i++
		//		} else {
		//			result += romanMap[string(s[i])]
		//		}

	}

	return result
}

//====

func main() {
	timeStartWholeProgram := time.Now()
	testInput := []TestCase{
		{
			S: "III",
			Result: `
          3
            `,
		},
		{

			S: "LVIII",
			Result: `
           58 
             `,
		},
		{

			S: "MCMXCIV",
			Result: `
           1994 
             `,
		},
	}
	for count, value := range testInput {
		fmt.Println("===============")
		fmt.Println("Test count ", count, "for node", value)
		fmt.Println("Solution 1: use Straightforward approach with map[string]int")
		timeStart := time.Now()
		result := romanToInt_MapString(value.S)
		timeLapse := time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

		fmt.Println("Solution 2: use Straightforward map[byte]int")
		timeStart = time.Now()
		result = romanToInt_MapByte(value.S)
		timeLapse = time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

		fmt.Println("Solution 3: use front to end approach")
		timeStart = time.Now()
		result = romanToInt_FrontToEnd(value.S)
		timeLapse = time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

		fmt.Println("Solution 4: just include all roman case")
		timeStart = time.Now()
		result = romanToInt_IncludeAll(value.S)
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
Test count  0 for node {III
          3
            }
Solution 1: use Straightforward approach with map[string]int
>Solution result 3
Correct result is
          3

TimeLapse 3.834µs
Solution 2: use Straightforward map[byte]int
>Solution result 3
Correct result is
          3

TimeLapse 4.148µs
Solution 3: use front to end approach
>Solution result 3
Correct result is
          3

TimeLapse 1.834µs
Solution 4: just include all roman case
>Solution result 3
Correct result is
          3

TimeLapse 4.407µs
===============
Test count  1 for node {LVIII
           58
             }
Solution 1: use Straightforward approach with map[string]int
>Solution result 58
Correct result is
           58

TimeLapse 1.5µs
Solution 2: use Straightforward map[byte]int
>Solution result 58
Correct result is
           58

TimeLapse 1.407µs
Solution 3: use front to end approach
>Solution result 58
Correct result is
           58

TimeLapse 1.204µs
Solution 4: just include all roman case
>Solution result 58
Correct result is
           58

TimeLapse 2.87µs
===============
Test count  2 for node {MCMXCIV
           1994
             }
Solution 1: use Straightforward approach with map[string]int
>Solution result 1994
Correct result is
           1994

TimeLapse 1.629µs
Solution 2: use Straightforward map[byte]int
>Solution result 1994
Correct result is
           1994

TimeLapse 1.537µs
Solution 3: use front to end approach
>Solution result 1994
Correct result is
           1994

TimeLapse 1.444µs
Solution 4: just include all roman case
>Solution result 1994
Correct result is
           1994

TimeLapse 3.778µs
===============
TimeLapse Whole Program 752.439µs
*/
//REF
//
