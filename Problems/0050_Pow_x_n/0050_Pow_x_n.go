package main

import (
	"fmt"
	"runtime"
	"time"
)

//approach StraightForward
/*
Time Complexity: The time complexity of this algorithm is O(n) in the worst case, where n is the absolute value of the exponent. This is because the loop runs n times to perform the multiplication.

Space Complexity: The space complexity is O(1) since the algorithm uses a constant amount of extra space regardless of the input size. It only creates a single variable (result) to store the intermediate result.

*/
func myPow_StraightForward(x float64, n int) float64 {
	result := float64(1)
	if n < 0 {
		x = 1 / x
		n = -n
	}

	for i := 1; i <= n; i++ {
		result *= x
	}

	return result
}

//approach binary exponentiation algorithm
/*

Algorithm: The algorithm recursively computes the power by dividing the exponent by 2 at each step and using the fact that x^(2n) = (x^n)^2 to optimize the computation.

Time Complexity: The time complexity of this algorithm is O(log n), where n is the exponent. This is because at each step, the exponent is halved.

Space Complexity: The space complexity is also O(log n) due to the recursive calls, as the maximum depth of the recursion is logarithmic in terms of the exponent.

*/

func myPow_BinaryExponentiation(x float64, n int) float64 {
	if n < 0 {
		x = 1 / x
		n = -n
	}
	return power(x, n)
}

func power(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	half := power(x, n/2)
	if n%2 == 0 {
		return half * half
	} else {
		return half * half * x
	}
}

//approach : iteratively computes the power by reducing the exponent by half at each step and squaring the base x accordingly.
// Linear approach
/*

Algorithm: This algorithm iteratively computes the power by reducing the exponent by half at each step and squaring the base x accordingly.

Time Complexity: The time complexity of this algorithm is O(log n), where n is the absolute value of the exponent. This is because the number of iterations required to reach 0 is logarithmic in terms of the exponent.

Space Complexity: The space complexity is O(1) since the algorithm uses a constant amount of extra space regardless of the input size. It only creates a few variables (result, x, n) to store the intermediate result and loop counters.



*/

func myPow_Iterative(x float64, n int) float64 {
	if n < 0 {
		x = 1 / x
		n = -n
	}
	result := float64(1)
	for n > 0 {
		if n%2 == 1 {
			result *= x
		}
		x *= x
		n /= 2
	}
	return result
}

func main() {
	timeStartWholeProgram := time.Now()

	testInput := []TestCase{
		{
			X: 2,
			N: 10,
			Result: `
      1024.00000
            `,
		},
		{
			X: 2.1,
			N: 3,
			Result: `
9.26100

            `,
		},
		{
			X: 2.00000,
			N: -2,
			Result: `
      0.25000
            `,
		},
	}

	// Memory before allocation
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	memBefore := m.Alloc

	for count, value := range testInput {
		fmt.Println("===============")
		fmt.Println("Test count ", count, "for node", value)
		fmt.Println("Solution 1: StraightForward")
		timeStart := time.Now()
		result := myPow_StraightForward(value.X, value.N)
		timeLapse := time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

		// Memory after allocation
		runtime.ReadMemStats(&m)
		memAfter := m.Alloc
		fmt.Println("Memory before", memBefore, "bytes", "Memory after", memAfter, "bytes", "Memory used:", memAfter-memBefore, "bytes")
		fmt.Printf("Memory usage (HeapAlloc) after Test Case i %d, : %v bytes\n", count, m.HeapAlloc)

		fmt.Println("Solution 2: Binary exponentiation")
		timeStart = time.Now()
		result = myPow_BinaryExponentiation(value.X, value.N)
		timeLapse = time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

		// Memory after allocation
		runtime.ReadMemStats(&m)
		memAfter = m.Alloc
		fmt.Println("Memory before", memBefore, "bytes", "Memory after", memAfter, "bytes", "Memory used:", memAfter-memBefore, "bytes")
		fmt.Printf("Memory usage (HeapAlloc) after Test Case i %d, : %v bytes\n", count, m.HeapAlloc)

		fmt.Println("Solution 3: Linear, Iterative ")
		timeStart = time.Now()
		result = myPow_Iterative(value.X, value.N)
		timeLapse = time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

		// Memory after allocation
		runtime.ReadMemStats(&m)
		memAfter = m.Alloc
		fmt.Println("Memory before", memBefore, "bytes", "Memory after", memAfter, "bytes", "Memory used:", memAfter-memBefore, "bytes")
		fmt.Printf("Memory usage (HeapAlloc) after Test Case i %d, : %v bytes\n", count, m.HeapAlloc)

	}

	timeLapsedWholeProgram := time.Since(timeStartWholeProgram)
	fmt.Println("===============")
	fmt.Println("TimeLapse Whole Program", timeLapsedWholeProgram)
}

type TestCase struct {
	X      float64
	N      int
	Result string
}

/*



===============
Test count  0 for node {2 10
      1024.00000
            }
Solution 1: StraightForward
>Solution result 1024
Correct result is
      1024.00000

TimeLapse 648ns
Memory before 67184 bytes Memory after 68104 bytes Memory used: 920 bytes
Memory usage (HeapAlloc) after Test Case i 0, : 68104 bytes
Solution 2: Binary exponentiation
>Solution result 1024
Correct result is
      1024.00000

TimeLapse 945ns
Memory before 67184 bytes Memory after 68312 bytes Memory used: 1128 bytes
Memory usage (HeapAlloc) after Test Case i 0, : 68312 bytes
Solution 3: Linear, Iterative
>Solution result 1024
Correct result is
      1024.00000

TimeLapse 593ns
Memory before 67184 bytes Memory after 68392 bytes Memory used: 1208 bytes
Memory usage (HeapAlloc) after Test Case i 0, : 68392 bytes
===============
Test count  1 for node {2.1 3
9.26100

            }
Solution 1: StraightForward
>Solution result 9.261000000000001
Correct result is
9.26100


TimeLapse 185ns
Memory before 67184 bytes Memory after 68488 bytes Memory used: 1304 bytes
Memory usage (HeapAlloc) after Test Case i 1, : 68488 bytes
Solution 2: Binary exponentiation
>Solution result 9.261000000000001
Correct result is
9.26100


TimeLapse 445ns
Memory before 67184 bytes Memory after 68568 bytes Memory used: 1384 bytes
Memory usage (HeapAlloc) after Test Case i 1, : 68568 bytes
Solution 3: Linear, Iterative
>Solution result 9.261000000000001
Correct result is
9.26100


TimeLapse 185ns
Memory before 67184 bytes Memory after 68632 bytes Memory used: 1448 bytes
Memory usage (HeapAlloc) after Test Case i 1, : 68632 bytes
===============
Test count  2 for node {2 -2
      0.25000
            }
Solution 1: StraightForward
>Solution result 0.25
Correct result is
      0.25000

TimeLapse 296ns
Memory before 67184 bytes Memory after 68744 bytes Memory used: 1560 bytes
Memory usage (HeapAlloc) after Test Case i 2, : 68744 bytes
Solution 2: Binary exponentiation
>Solution result 0.25
Correct result is
      0.25000

TimeLapse 315ns
Memory before 67184 bytes Memory after 68824 bytes Memory used: 1640 bytes
Memory usage (HeapAlloc) after Test Case i 2, : 68824 bytes
Solution 3: Linear, Iterative
>Solution result 0.25
Correct result is
      0.25000

TimeLapse 203ns
Memory before 67184 bytes Memory after 68888 bytes Memory used: 1704 bytes
Memory usage (HeapAlloc) after Test Case i 2, : 68888 bytes
===============
TimeLapse Whole Program 1.684732ms

*/
//REF
//
