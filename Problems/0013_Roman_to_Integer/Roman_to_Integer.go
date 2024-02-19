package main

import (
	"fmt"
	"time"
)


//approach Straightforward, loop from back to front and using a map
 func romanToInt_MapString(s string) int {
  romanMap := map[string]int {
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

//===== 
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

	}

	timeLapsedWholeProgram := time.Since(timeStartWholeProgram)
	fmt.Println("===============")
	fmt.Println("TimeLapse Whole Program", timeLapsedWholeProgram)
}

type TestCase struct {
  S string
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
            
TimeLapse 3.777µs
Solution 2: use Straightforward map[byte]int
>Solution result 3
Correct result is  
          3
            
TimeLapse 4.666µs
===============
Test count  1 for node {LVIII 
           58 
             }
Solution 1: use Straightforward approach with map[string]int
>Solution result 58
Correct result is  
           58 
             
TimeLapse 1.703µs
Solution 2: use Straightforward map[byte]int
>Solution result 58
Correct result is  
           58 
             
TimeLapse 1.463µs
===============
Test count  2 for node {MCMXCIV 
           1994 
             }
Solution 1: use Straightforward approach with map[string]int
>Solution result 1994
Correct result is  
           1994 
             
TimeLapse 1.648µs
Solution 2: use Straightforward map[byte]int
>Solution result 1994
Correct result is  
           1994 
             
TimeLapse 1.648µs
===============
TimeLapse Whole Program 601.99µs

*/
//REF
//

