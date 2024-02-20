package main

import (
	"fmt"
	"math"
	"time"
)

//NOTE: considering overflow conditions
//The maximum value for a 32-bit signed integer (assuming two's complement representation) is  2^31 − 1 = 2147483647 , and the minimum value is −2^31 = −2147483648
// => x = 2147483643 => reverse : 3463847412 => overflow case => should return 0

//Time Complexity: O(log|x|)
//Space Complexity: O(1)

func reverse_StraightForward(x int) int {
	var result int

	for x != 0 {
		digit := x % 10
		x /= 10

		// Checking for overflow before adding the new digit
		if result > math.MaxInt32/10 || (result == math.MaxInt32/10 && digit > 7) {
			return 0
		}
		if result < math.MinInt32/10 || (result == math.MinInt32/10 && digit < -8) {
			return 0
		}

		result = result*10 + digit
	}

	return result
}

//===== brute force approach

func reverse_BruteForce(x int) int {
	// Handle edge case where x is 0
	if x == 0 {
		return 0
	}

	// Initialize a variable to store the result
	var result int

	// Get the absolute value of x
	absX := int(math.Abs(float64(x)))

	// Reverse the digits
	for absX > 0 {
		digit := absX % 10
		result = result*10 + digit
		absX /= 10
	}

	// Check if the original number was negative
	if x < 0 {
		result *= -1
	}

	// Check for overflow and return 0 if necessary
	if result > math.MaxInt32 || result < math.MinInt32 {
		return 0
	}

	return result
}

//=====

func main() {
	timeStartWholeProgram := time.Now()
	testInput := []TestCase{
		{
			N: 123,
			Result: `
          321
            `,
		},
		{
			N: -123,
			Result: `
          -321
            `,
		},
		{
			N: 120,
			Result: `
              21
            `,
		},
		{
			N: 2147483641,
			Result: `
      1463847412

            `,
		},
		{
			N: 2147483643,
			Result: `
      => 3463847412 =>overflow case => should result 0

            `,
		},
	}
	for count, value := range testInput {
		fmt.Println("===============")
		fmt.Println("Test count ", count, "for node", value)
		fmt.Println("Solution 1: use straightforward iterative approach")
		timeStart := time.Now()
		result := reverse_StraightForward(value.N)
		timeLapse := time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

		fmt.Println("Solution 2: use Bture Force")
		timeStart = time.Now()
		result = reverse_BruteForce(value.N)
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
Test count  0 for node {123
          321
            }
Solution 1: use straightforward iterative approach
>Solution result 321
Correct result is
          321

TimeLapse 611ns
Solution 2: use Bture Force
>Solution result 321
Correct result is
          321

TimeLapse 333ns
===============
Test count  1 for node {-123
          -321
            }
Solution 1: use straightforward iterative approach
>Solution result -321
Correct result is
          -321

TimeLapse 167ns
Solution 2: use Bture Force
>Solution result -321
Correct result is
          -321

TimeLapse 166ns
===============
Test count  2 for node {120
              21
            }
Solution 1: use straightforward iterative approach
>Solution result 21
Correct result is
              21

TimeLapse 186ns
Solution 2: use Bture Force
>Solution result 21
Correct result is
              21

TimeLapse 148ns
===============
Test count  3 for node {2147483641
      1463847412

            }
Solution 1: use straightforward iterative approach
>Solution result 1463847412
Correct result is
      1463847412


TimeLapse 241ns
Solution 2: use Bture Force
>Solution result 1463847412
Correct result is
      1463847412


TimeLapse 185ns
===============
Test count  4 for node {2147483643
      => 3463847412 =>overflow case => should result 0

            }
Solution 1: use straightforward iterative approach
>Solution result 0
Correct result is
      => 3463847412 =>overflow case => should result 0


TimeLapse 222ns
Solution 2: use Bture Force
>Solution result 0
Correct result is
      => 3463847412 =>overflow case => should result 0


TimeLapse 204ns
===============
TimeLapse Whole Program 699.963µs

*/
