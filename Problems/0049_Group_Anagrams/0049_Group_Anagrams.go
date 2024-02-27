package main

import (
	"fmt"
	"runtime"
	"sort"
	"time"
)

/*
We iterate through each word in the input array once, which takes O(n) time, where n is the total number of characters in all the words.
For each word, we sort the characters, which takes O(k log k) time, where k is the length of the longest word.
Since there are n words, the total time complexity for sorting all the words is O(n * k log k).
We store the sorted words in a hashmap. The space complexity for this hashmap is O(n * k), where n is the number of words and k is the length of the longest word.
We then iterate through the hashmap to collect the anagram groups, which takes O(n) time.
Considering the above analysis, the overall time complexity of the algorithm is O(n * k log k) due to the sorting operation, and the space complexity is O(n * k) due to the hashmap.


*/

func groupAnagrams(strs []string) [][]string {
	// Create a map to store anagrams
	anagrams := make(map[string][]string)

	// Iterate through each word in the input
	for _, word := range strs {
		// Convert word to a sorted string
		sortedWord := sortString(word)
		// Append the word to the anagram group
		anagrams[sortedWord] = append(anagrams[sortedWord], word)
	}

	// Create result slice
	var result [][]string
	// Append each anagram group to the result
	for _, group := range anagrams {
		result = append(result, group)
	}

	return result
}

// Helper function to sort a string
func sortString(str string) string {
	// Convert string to a slice of bytes
	bytes := []byte(str)
	// Sort the slice of bytes
	sort.Slice(bytes, func(i, j int) bool { return bytes[i] < bytes[j] })
	// Convert the sorted slice of bytes back to a string
	return string(bytes)
}

func main() {
	timeStartWholeProgram := time.Now()

	testInput := []TestCase{
		{
			S: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			Result: `
[["bat"],["nat","tan"],["ate","eat","tea"]]

            `,
		},
		{
			S: []string{""},
			Result: `
      [[""]]

            `,
		},
		{
			S: []string{"a"},
			Result: `
      [["a"]]

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
		result := groupAnagrams(value.S)
		timeLapse := time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

		// Memory after allocation
		runtime.ReadMemStats(&m)
		memAfter := m.Alloc
		fmt.Println("Memory before", memBefore, "bytes", "Memory after", memAfter, "bytes", "Memory used:", memAfter-memBefore, "bytes")

		fmt.Printf("Memory usage (HeapAlloc) after Test Case i %d, : %v bytes\n", count, m.HeapAlloc)

	}

	timeLapsedWholeProgram := time.Since(timeStartWholeProgram)
	fmt.Println("===============")
	fmt.Println("TimeLapse Whole Program", timeLapsedWholeProgram)
}

type TestCase struct {
	S      []string
	Result string
}

/*

===============
Test count  0 for node {[eat tea tan ate nat bat] 
[["bat"],["nat","tan"],["ate","eat","tea"]]

            }
Solution 1: StraightForward
>Solution result [[eat tea ate] [tan nat] [bat]]
Correct result is  
[["bat"],["nat","tan"],["ate","eat","tea"]]

            
TimeLapse 20.315µs
Memory before 69376 bytes Memory after 71504 bytes Memory used: 2128 bytes
Memory usage (HeapAlloc) after Test Case i 0, : 71504 bytes
===============
Test count  1 for node {[] 
      [[""]]

            }
Solution 1: StraightForward
>Solution result [[]]
Correct result is  
      [[""]]

            
TimeLapse 3.148µs
Memory before 69376 bytes Memory after 71760 bytes Memory used: 2384 bytes
Memory usage (HeapAlloc) after Test Case i 1, : 71760 bytes
===============
Test count  2 for node {[a] 
      [["a"]]

            }
Solution 1: StraightForward
>Solution result [[a]]
Correct result is  
      [["a"]]

            
TimeLapse 3.092µs
Memory before 69376 bytes Memory after 72032 bytes Memory used: 2656 bytes
Memory usage (HeapAlloc) after Test Case i 2, : 72032 bytes
===============
TimeLapse Whole Program 919.038µs


*/
//REF
//
