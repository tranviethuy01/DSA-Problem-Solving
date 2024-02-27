package main

import (
	"fmt"
	"sort"
	"time"
)

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
      S: []string{"eat","tea","tan","ate","nat","bat"},
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

	}

	timeLapsedWholeProgram := time.Since(timeStartWholeProgram)
	fmt.Println("===============")
	fmt.Println("TimeLapse Whole Program", timeLapsedWholeProgram)
}

type TestCase struct {
	S []string
	Result string
}

/*


===============
Test count  0 for node {[eat tea tan ate nat bat] 
[["bat"],["nat","tan"],["ate","eat","tea"]]

            }
Solution 1: StraightForward
>Solution result [[tan nat] [bat] [eat tea ate]]
Correct result is  
[["bat"],["nat","tan"],["ate","eat","tea"]]

            
TimeLapse 22.815µs
===============
Test count  1 for node {[] 
      [[""]]

            }
Solution 1: StraightForward
>Solution result [[]]
Correct result is  
      [[""]]

            
TimeLapse 2.352µs
===============
Test count  2 for node {[a] 
      [["a"]]

            }
Solution 1: StraightForward
>Solution result [[a]]
Correct result is  
      [["a"]]

            
TimeLapse 2µs
===============
TimeLapse Whole Program 404.528µs
*/
//REF
//
