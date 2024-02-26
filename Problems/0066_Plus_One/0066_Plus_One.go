package main

import (
	"fmt"
	"time"
)

//time complexity of O(n),
//The space complexity of the algorithm is also O(n). In the worst case scenario where all digits in the array are 9, an additional digit needs to be added at the beginning of the array, increasing the space required linearly with the size of the input array.

func plusOne(digits []int) []int {
	n := len(digits)

	//loop from end to begin
	for i := n - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			//should return
			return digits
		}
		//case digits[i] == 9, should add + 1, mean 10, then do the next loop, plus +1 for the previous number
		digits[i] = 0

	}
	// If all digits were 9, we need to add an additional digit at the beginning.
	return append([]int{1}, digits...)
}

func main() {
	timeStartWholeProgram := time.Now()

	testInput := []TestCase{

		{
			Digits: []int{1, 2, 3},
			Result: `
[1,2,4]
            `,
		},
		{
			Digits: []int{4, 3, 2, 1},
			Result: `
[4,3,2,2]
            `,
		},
		{
			Digits: []int{9},
			Result: `
[1,0]
            `,
		},

		{

			Digits: []int{4, 3, 2, 9},
			Result: `
[4,3,3,0]
            `,
		},
	}
	for count, value := range testInput {
		fmt.Println("===============")
		fmt.Println("Test count ", count, "for node", value)
		fmt.Println("Solution 1: manual multiply")
		timeStart := time.Now()
		result := plusOne(value.Digits)
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
	Digits []int
	Result string
}

/*



===============
Test count  0 for node {[1 2 3]
[1,2,4]
            }
Solution 1: manual multiply
>Solution result [1 2 4]
Correct result is
[1,2,4]

TimeLapse 389ns
===============
Test count  1 for node {[4 3 2 1]
[4,3,2,2]
            }
Solution 1: manual multiply
>Solution result [4 3 2 2]
Correct result is
[4,3,2,2]

TimeLapse 130ns
===============
Test count  2 for node {[9]
[1,0]
            }
Solution 1: manual multiply
>Solution result [1 0]
Correct result is
[1,0]

TimeLapse 814ns
===============
Test count  3 for node {[4 3 2 9]
[4,3,3,0]
            }
Solution 1: manual multiply
>Solution result [4 3 3 0]
Correct result is
[4,3,3,0]

TimeLapse 129ns
===============
TimeLapse Whole Program 472.679µs

*/
//REF
//
