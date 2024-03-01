package main

import (
	"fmt"
	"runtime"
	"strings"
	"time"
)

//approach : StraightForward , use split and get the length of the last word
// remember to remove the TrimRight
/*
Algorithm:
Trim the trailing spaces from the input string using strings.TrimRight.
Split the trimmed string into a slice of words using strings.Split, with space (" ") as the delimiter.
Retrieve the last word from the slice of words.
Return the length of the last word.
Time Complexity:
strings.TrimRight: This function iterates over the string to remove trailing spaces. Its time complexity is O(n), where n is the length of the input string.
strings.Split: This function splits the string into substrings based on the space delimiter. It also has a time complexity of O(n), where n is the length of the input string.
Accessing the last element of the words slice (words[len(words)-1]) takes O(1) time.
Overall, the time complexity is O(n), where n is the length of the input string.
Space Complexity:
strings.TrimRight and strings.Split both create new strings or slices, but they don't scale with the input size. They only use additional space proportional to the size of the input string.
The space complexity of the words slice is O(m), where m is the number of words in the input string.
The space complexity of the entire function is O(n + m), where n is the length of the input string and m is the number of words.
*/

func lengthOfLastWord_StraightForward(s string) int {
	// Remove trailing spaces
	s = strings.TrimRight(s, " ")
	words := strings.Split(s, " ")
	return len(words[len(words)-1])
}

//approach : another way without use Split method, just check for the character
/*

Algorithm:
Trim the trailing spaces from the input string.
Iterate through the string in reverse order until a space is encountered or the beginning of the string is reached.
Count the characters until a space is found, indicating the end of the last word.
Return the count as the length of the last word.
Time Complexity:
Trimming the trailing spaces takes O(n) time, where n is the length of the input string.
Iterating through the string in reverse order also takes O(n) time since we are traversing the string once.
Overall, the time complexity is O(n).
Space Complexity:
The space complexity is O(1) because we are not using any additional data structures that scale with the input size. We are only using a constant amount of extra space for variables like length.

*/

func lengthOfLastWord_StraightForward2(s string) int {
	// Remove trailing spaces
	s = strings.TrimRight(s, " ")

	length := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			fmt.Println("meet the space here ' ' , so should bread and return the length")
			break
		}
		length++
	}

	return length
}

func main() {
	timeStartWholeProgram := time.Now()

	testInput := []TestCase{
		{
			S: "Hello World",
			Result: `
5
            `,
		},
		{
			S: "   fly me   to   the moon  ",
			Result: `
4
            `,
		},
		{
			S: "luffy is still joyboy",
			Result: `
6

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
		result := lengthOfLastWord_StraightForward(value.S)
		timeLapse := time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

		// Memory after allocation
		runtime.ReadMemStats(&m)
		memAfter := m.Alloc
		fmt.Println("Memory before", memBefore, "bytes", "Memory after", memAfter, "bytes", "Memory used:", memAfter-memBefore, "bytes")

		fmt.Printf("Memory usage (HeapAlloc) after Test Case i %d, : %v bytes\n", count, m.HeapAlloc)

		fmt.Println("Solution 2: try to use another StraightForward approach ")
		timeStart = time.Now()
		result = lengthOfLastWord_StraightForward2(value.S)
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
	S      string
	Result string
}

/*

===============
Test count  0 for node {Hello World
5
            }
Solution 1: StraightForward
>Solution result 5
Correct result is
5

TimeLapse 3.278µs
Memory before 69248 bytes Memory after 70184 bytes Memory used: 936 bytes
Memory usage (HeapAlloc) after Test Case i 0, : 70184 bytes
Solution 2: try to use another StraightForward approach
meet the space here ' ' , so should bread and return the length
>Solution result 5
Correct result is
5

TimeLapse 9.148µs
Memory before 69248 bytes Memory after 70376 bytes Memory used: 1128 bytes
Memory usage (HeapAlloc) after Test Case i 0, : 70376 bytes
===============
Test count  1 for node {   fly me   to   the moon
4
            }
Solution 1: StraightForward
>Solution result 4
Correct result is
4

TimeLapse 18.759µs
Memory before 69248 bytes Memory after 70680 bytes Memory used: 1432 bytes
Memory usage (HeapAlloc) after Test Case i 1, : 70680 bytes
Solution 2: try to use another StraightForward approach
meet the space here ' ' , so should bread and return the length
>Solution result 4
Correct result is
4

TimeLapse 8.556µs
Memory before 69248 bytes Memory after 70744 bytes Memory used: 1496 bytes
Memory usage (HeapAlloc) after Test Case i 1, : 70744 bytes
===============
Test count  2 for node {luffy is still joyboy
6

            }
Solution 1: StraightForward
>Solution result 6
Correct result is
6


TimeLapse 1.833µs
Memory before 69248 bytes Memory after 70904 bytes Memory used: 1656 bytes
Memory usage (HeapAlloc) after Test Case i 2, : 70904 bytes
Solution 2: try to use another StraightForward approach
meet the space here ' ' , so should bread and return the length
>Solution result 6
Correct result is
6


TimeLapse 8.666µs
Memory before 69248 bytes Memory after 70968 bytes Memory used: 1720 bytes
Memory usage (HeapAlloc) after Test Case i 2, : 70968 bytes
===============
TimeLapse Whole Program 1.054634ms

*/
//REF
//
