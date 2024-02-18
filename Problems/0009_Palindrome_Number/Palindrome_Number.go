package main

import (
	"fmt"

	"strconv"
	"time"
)

// Simplest: approach: reverse digit of number, then compare the reversed number with the origin number

func isPalindrome(x int) bool {
	// Special cases:
	// If x is negative or ends with 0 (except for 0 itself), it cannot be a palindrome.
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	// Reverse half of the number.
	reversed := 0
	for x > reversed {
		reversed = reversed*10 + x%10
		x /= 10
	}

	// If the number of digits is odd, we can ignore the middle digit.
	// For example, if x = 12321, at the end of the loop we would have
	// reversed = 123, x = 12 (middle digit is 2 which can be ignored).
	// For even digits, reversed and x would be of the same length.
	return x == reversed || x == reversed/10
}

// approach : reverse number

func checkPalindrome_ReverseNumber(n int) bool {
	reverse := 0
	temp := n
	for temp != 0 {
		reverse = (reverse * 10) + (temp % 10)
		temp = temp / 10
	}
	return reverse == n
}

//===== approach: convert to string then iterate through haft of the string

func isPalindrome_ConvertToString(x int) bool {
	// Convert the integer to a string.
	str := strconv.Itoa(x)

	// Iterate through half of the string.
	n := len(str)
	for i := 0; i < n/2; i++ {
		// Compare characters from the beginning and end of the string.
		if str[i] != str[n-1-i] {
			return false
		}
	}
	return true
}

//=====

//===== approach : check for a string
//str = "112233445566778899000000998877665544332211"

func isPalindrome_CheckString(str string) bool {
	// Iterate through half of the string.
	n := len(str)
	for i := 0; i < n/2; i++ {
		// Compare characters from the beginning and end of the string.
		if str[i] != str[n-1-i] {
			return false
		}
	}
	return true
}

//=====

// ===== recursive solution
//
// Function to check if a number is palindrome.
func isPalindrome_Recursive(num int) bool {
	// If num is negative, make it positive.
	if num < 0 {
		num = -num
	}

	// Create a separate copy of num, so that modifications made to address dupNum don't change the input number.
	dupNum := num

	// Call the recursive function to check if the number is palindrome.
	return isPalindrom_RecursiveUtil(num, &dupNum)
}

// A recursive function to check whether num is palindrome or not.
// num contains the current number being checked.
// dupNum contains the original number to compare with.
func isPalindrom_RecursiveUtil(num int, dupNum *int) bool {
	// Base case: If num contains only one digit, compare it with the last digit of dupNum.
	if num < 10 {
		return num == *dupNum%10
	}

	// Recursive call to move up the recursion tree and check each digit.
	if !isPalindrom_RecursiveUtil(num/10, dupNum) {
		return false
	}

	// Update dupNum to the next digit for comparison.
	*dupNum /= 10

	// Check if the current digit of num matches with the corresponding digit of dupNum.
	return num%10 == *dupNum%10
}

//=====

func main() {
	timeStartWholeProgram := time.Now()
	testInput := []TestCase{
		{
			N: 121,
			Result: `
          true
            `,
		},
		{
			N: -121,
			Result: `
             false
             `,
		},
		{
			N: 10,
			Result: `
			false
            `,
		},
		{
			N: 2147483647,
			Result: `
			false
            `,
		},
		{
			N: 9223372036854775806,
			Result: `
			false
            `,
		},
		{
			N: 9223372036854775807,
			Result: `
			false
            `,
		},
		{
			N: 9223372036854775008,
			Result: `
      overflow case: false
            `,
		},
		{
			S: "112233445566778899000000998877665544332211",
			Result: `
      Test case with String => true

            `,
		},
	}
	for count, value := range testInput {
		fmt.Println("===============")
		fmt.Println("Test count ", count, "for node", value)
		fmt.Println("Solution 1: use")
		timeStart := time.Now()
		result := isPalindrome(value.N)
		timeLapse := time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

		fmt.Println("Solution 2: use convert to string")
		timeStart = time.Now()
		result = isPalindrome_ConvertToString(value.N)
		timeLapse = time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

		fmt.Println("Solution 3: recursive")
		timeStart = time.Now()
		result = isPalindrome_Recursive(value.N)
		timeLapse = time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

		fmt.Println("Additional test : test case with string input for large string")
		timeStart = time.Now()
		result = isPalindrome_CheckString(value.S)
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
	N      int
	Result string
}

/*

===============
Test count  0 for node { 121
          true
            }
Solution 1: use
>Solution result true
Correct result is
          true

TimeLapse 537ns
Solution 2: use convert to string
>Solution result true
Correct result is
          true

TimeLapse 2.13µs
Solution 3: recursive
>Solution result true
Correct result is
          true

TimeLapse 704ns
Additional test : test case with string input for large string
>Solution result true
Correct result is
          true

TimeLapse 259ns
===============
Test count  1 for node { -121
             false
             }
Solution 1: use
>Solution result false
Correct result is
             false

TimeLapse 111ns
Solution 2: use convert to string
>Solution result false
Correct result is
             false

TimeLapse 444ns
Solution 3: recursive
>Solution result true
Correct result is
             false

TimeLapse 167ns
Additional test : test case with string input for large string
>Solution result true
Correct result is
             false

TimeLapse 111ns
===============
Test count  2 for node { 10
			false
            }
Solution 1: use
>Solution result false
Correct result is
			false

TimeLapse 111ns
Solution 2: use convert to string
>Solution result false
Correct result is
			false

TimeLapse 185ns
Solution 3: recursive
>Solution result false
Correct result is
			false

TimeLapse 148ns
Additional test : test case with string input for large string
>Solution result true
Correct result is
			false

TimeLapse 74ns
===============
Test count  3 for node { 2147483647
			false
            }
Solution 1: use
>Solution result false
Correct result is
			false

TimeLapse 185ns
Solution 2: use convert to string
>Solution result false
Correct result is
			false

TimeLapse 519ns
Solution 3: recursive
>Solution result false
Correct result is
			false

TimeLapse 241ns
Additional test : test case with string input for large string
>Solution result true
Correct result is
			false

TimeLapse 93ns
===============
Test count  4 for node { 9223372036854775806
			false
            }
Solution 1: use
>Solution result false
Correct result is
			false

TimeLapse 186ns
Solution 2: use convert to string
>Solution result false
Correct result is
			false

TimeLapse 7.537µs
Solution 3: recursive
>Solution result false
Correct result is
			false

TimeLapse 315ns
Additional test : test case with string input for large string
>Solution result true
Correct result is
			false

TimeLapse 112ns
===============
Test count  5 for node { 9223372036854775807
			false
            }
Solution 1: use
>Solution result false
Correct result is
			false

TimeLapse 203ns
Solution 2: use convert to string
>Solution result false
Correct result is
			false

TimeLapse 592ns
Solution 3: recursive
>Solution result false
Correct result is
			false

TimeLapse 277ns
Additional test : test case with string input for large string
>Solution result true
Correct result is
			false

TimeLapse 93ns
===============
Test count  6 for node { 9223372036854775008
      overflow case: false
            }
Solution 1: use
>Solution result false
Correct result is
      overflow case: false

TimeLapse 185ns
Solution 2: use convert to string
>Solution result false
Correct result is
      overflow case: false

TimeLapse 537ns
Solution 3: recursive
>Solution result false
Correct result is
      overflow case: false

TimeLapse 259ns
Additional test : test case with string input for large string
>Solution result true
Correct result is
      overflow case: false

TimeLapse 74ns
===============
Test count  7 for node {112233445566778899000000998877665544332211 0
      Test case with String => true

            }
Solution 1: use
>Solution result true
Correct result is
      Test case with String => true


TimeLapse 130ns
Solution 2: use convert to string
>Solution result true
Correct result is
      Test case with String => true


TimeLapse 315ns
Solution 3: recursive
>Solution result true
Correct result is
      Test case with String => true


TimeLapse 130ns
Additional test : test case with string input for large string
>Solution result true
Correct result is
      Test case with String => true


TimeLapse 204ns
===============
TimeLapse Whole Program 1.528271ms

*/

//REF
//https://www.geeksforgeeks.org/check-if-a-number-is-palindrome/
