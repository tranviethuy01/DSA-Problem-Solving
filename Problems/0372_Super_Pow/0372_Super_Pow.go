package main

import (
	"fmt"
	"runtime"
	"time"
)

//approach : modular exponentiation. It iteratively calculates the result ofa ^ b mod1337 by taking advantage of the properties of modular arithmetic.
/*
Time Complexity:

The superPow function iterates through each digit in the array b, performing a modular exponentiation operation for each digit.
The time complexity of modular exponentiation is
O(n), where
n is the exponent.
Since the length of array b represents the exponent, the time complexity of the superPow function is
O(m), where
m is the length of array b.
The helper function pow is called
O(n) times for each digit in array b.
Therefore, the overall time complexity is
O(m⋅n).

Space Complexity:
The space complexity is
O(1) because the algorithm uses a constant amount of extra space regardless of the size of the input.
It does not use any additional data structures that scale with the input size.

*/
func superPow_Modular_Exponentiation(a int, b []int) int {
	const mod = 1337
	a %= mod
	result := 1
	for _, digit := range b {
		result = (pow(result, 10) * pow(a, digit)) % mod
	}
	return result
}

func pow(x, n int) int {
	const mod = 1337
	result := 1
	for i := 0; i < n; i++ {
		result = (result * x) % mod
	}
	return result
}

//approach Euler's totient
/*
Algorithm:
The algorithm computes the exponent by combining the digits in array b and then calculates the exponent modulo
ϕ(1337), which is 1140.
It utilizes binary exponentiation to efficiently compute
a ^ b mod1337.
The binary exponentiation algorithm reduces the exponentiation process to logarithmic time by repeatedly squaring the base.
Time Complexity:
Calculating the exponent by combining the digits in array b takes linear time,
O(m), where
m is the length of array b.
Binary exponentiation has a time complexity of
O(logb), where
b is the exponent.
Therefore, the overall time complexity is
O(m+logb).

Space Complexity:
The space complexity is
O(1) because th
e algorithm uses a constant amount of extra space regardless of the size of the input.
It does not use any additional data structures that scale with the input size.
=>
Time Complexity:
O(m+logb)
Space Complexity:
O(1)





*/

func superPow_Euler(a int, b []int) int {
	const mod = 1337
	a %= mod
	exponent := 0
	for _, digit := range b {
		exponent = (exponent*10 + digit) % 1140 // Euler's totient function of 1337 is 1140
	}
	if exponent == 0 {
		exponent = 1140 // handle the case when exponent is multiple of 1140
	}
	return power(a, exponent, mod)
}

func power(a, b, mod int) int {
	if b == 0 {
		return 1
	}
	if b == 1 {
		return a % mod
	}
	partial := power(a, b/2, mod)
	result := (partial * partial) % mod
	if b%2 == 1 {
		result = (result * a) % mod
	}
	return result
}

func main() {
	timeStartWholeProgram := time.Now()

	testInput := []TestCase{
		{
			A: 2,
			B: []int{3},
			Result: `
      8
            `,
		},
		{
			A: 2,
			B: []int{1, 0},
			Result: `
      1024
            `,
		},

		{
			A: 1,
			B: []int{4, 3, 3, 8, 5, 2},
			Result: `
      1
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
		fmt.Println("Solution 1: modular exponentiation")
		timeStart := time.Now()
		result := superPow_Modular_Exponentiation(value.A, value.B)
		timeLapse := time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

		// Memory after allocation
		runtime.ReadMemStats(&m)
		memAfter := m.Alloc
		fmt.Println("Memory before", memBefore, "bytes", "Memory after", memAfter, "bytes", "Memory used:", memAfter-memBefore, "bytes")
		fmt.Printf("Memory usage (HeapAlloc) after Test Case i %d, : %v bytes\n", count, m.HeapAlloc)

		fmt.Println("Solution 2: Euler")
		timeStart = time.Now()
		result = superPow_Euler(value.A, value.B)
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
	A      int
	B      []int
	Result string
}

/*


===============
Test count  0 for node {2 [3]
      8
            }
Solution 1: modular exponentiation
>Solution result 8
Correct result is
      8

TimeLapse 852ns
Memory before 69104 bytes Memory after 70040 bytes Memory used: 936 bytes
Memory usage (HeapAlloc) after Test Case i 0, : 70040 bytes
Solution 2: Euler
>Solution result 8
Correct result is
      8

TimeLapse 796ns
Memory before 69104 bytes Memory after 70232 bytes Memory used: 1128 bytes
Memory usage (HeapAlloc) after Test Case i 0, : 70232 bytes
===============
Test count  1 for node {2 [1 0]
      1024
            }
Solution 1: modular exponentiation
>Solution result 1024
Correct result is
      1024

TimeLapse 444ns
Memory before 69104 bytes Memory after 70376 bytes Memory used: 1272 bytes
Memory usage (HeapAlloc) after Test Case i 1, : 70376 bytes
Solution 2: Euler
>Solution result 1024
Correct result is
      1024

TimeLapse 260ns
Memory before 69104 bytes Memory after 70456 bytes Memory used: 1352 bytes
Memory usage (HeapAlloc) after Test Case i 1, : 70456 bytes
===============
Test count  2 for node {1 [4 3 3 8 5 2]
      1
            }
Solution 1: modular exponentiation
>Solution result 1
Correct result is
      1

TimeLapse 1.13µs
Memory before 69104 bytes Memory after 70616 bytes Memory used: 1512 bytes
Memory usage (HeapAlloc) after Test Case i 2, : 70616 bytes
Solution 2: Euler
>Solution result 1
Correct result is
      1

TimeLapse 389ns
Memory before 69104 bytes Memory after 70680 bytes Memory used: 1576 bytes
Memory usage (HeapAlloc) after Test Case i 2, : 70680 bytes
===============
TimeLapse Whole Program 1.004164ms


*/
//REF
//
