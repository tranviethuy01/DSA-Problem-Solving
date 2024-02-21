package main

import (
	"fmt"
"math"
	"time"
)

//approach: binary long division
//Time Complexity: O(log(dividend/divisor)), where dividend and divisor are the input integers.
//Space Complexity: O(1), because it uses only a few integer variables regardless of the input size.
//

func divide(dividend int, divisor int) int {
    if dividend == math.MinInt32 && divisor == -1 {
        return math.MaxInt32
    }

    negative := false
    if (dividend < 0 && divisor > 0) || (dividend > 0 && divisor < 0) {
        negative = true
    }

    dividend = int(math.Abs(float64(dividend)))
    divisor = int(math.Abs(float64(divisor)))

    quotient := 0
    for dividend >= divisor {
        temp := divisor
        multiple := 1
        for dividend >= (temp << 1) {
            temp <<= 1
            multiple <<= 1
        }
        dividend -= temp
        quotient += multiple
    }

    if negative {
        return -quotient
    }
    return quotient
}


// approach: repeatedly subtracts the divisor from the dividend until the dividend becomes less than the divisor
//Time Complexity: O(log(dividend/divisor))
//Space Complexity: O(1)

func divide_RepeatedlySubtract(dividend int, divisor int) int {
    if dividend == math.MinInt32 && divisor == -1 {
        return math.MaxInt32
    }

    negative := false
    if (dividend < 0 && divisor > 0) || (dividend > 0 && divisor < 0) {
        negative = true
    }

    dividend = int(math.Abs(float64(dividend)))
    divisor = int(math.Abs(float64(divisor)))

    quotient := 0
    for dividend >= divisor {
        dividend -= divisor
        quotient++
    }

    if negative {
        return -quotient
    }
    return quotient
}





func main() {
	timeStartWholeProgram := time.Now()

	testInput := []TestCase{

		{
	Dividend: 10,
	Divisor: 3,
			Result: `
3

            `,
		},
		{
Dividend: 7,
Divisor : -3, 
			Result: `
      -2
            `,
		},
	}
	for count, value := range testInput {
		fmt.Println("===============")
		fmt.Println("Test count ", count, "for node", value)
		fmt.Println("Solution 1: 2 pointer")
		timeStart := time.Now()
		result := divide(value.Dividend, value.Divisor)
		timeLapse := time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

		fmt.Println("Solution 2: brute force")
		timeStart = time.Now()
		result = divide_RepeatedlySubtract(value.Dividend, value.Divisor)
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
	Dividend int
	Divisor int
	Result   string
}

/*


===============
Test count  0 for node {10 3 
3

            }
Solution 1: 2 pointer
>Solution result 3
Correct result is  
3

            
TimeLapse 685ns
Solution 2: brute force
>Solution result 3
Correct result is  
3

            
TimeLapse 796ns
===============
Test count  1 for node {7 -3 
      -2
            }
Solution 1: 2 pointer
>Solution result -2
Correct result is  
      -2
            
TimeLapse 167ns
Solution 2: brute force
>Solution result -2
Correct result is  
      -2
            
TimeLapse 148ns
===============
TimeLapse Whole Program 393.099µs

*/
//REF
//
