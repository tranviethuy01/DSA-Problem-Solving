package main

import (
	"fmt"
	"time"
)

// approach : use binary search with optimize
// Time complexity of O(N + Q), where N is the length of the string s and Q is the total number of queries
// space complexity of the algorithm is O(N + Q), dominated by the space used to store the candleIndices array and the result array.
func platesBetweenCandles_BinarySearch(s string, queries [][]int) []int {
	//n := len(s)
	candleIndices := make([]int, 0)

	// Store indices of candles
	for i, c := range s {
		if c == '|' {
			candleIndices = append(candleIndices, i)
		}
	}

	result := make([]int, len(queries))

	// Process each query
	for i, query := range queries {
		left, right := query[0], query[1]
		count := 0

		// Find candles inside the query substring
		leftCandle := -1
		rightCandle := -1
		for _, idx := range candleIndices {
			if idx >= left && idx <= right {
				if leftCandle == -1 {
					leftCandle = idx
				} else {
					rightCandle = idx
					// Calculate plates between candles
					for j := leftCandle + 1; j < rightCandle; j++ {
						if s[j] == '*' {
							count++
						}
					}
					leftCandle = rightCandle
				}
			}
		}

		result[i] = count
	}

	return result
}

//approach : BruteForce : note: this solution is failure, need to check code again
//Time Complexity: The time complexity of this solution is O(N * Q), where N is the length of the string s and Q is the total number of queries. For each query, it potentially scans the entire substring, leading to a quadratic time complexity.

//Space Complexity: The space complexity of this solution is O(Q), where Q is the total number of queries. This is because the result array stores the output for each query. The space complexity is not directly dependent on the length of the string s.

func platesBetweenCandles_BruteForce(s string, queries [][]int) []int {
	result := make([]int, len(queries))

	for i, query := range queries {
		left, right := query[0], query[1]
		count := 0
		candleLeft := false

		// Iterate through the substring to count plates between candles
		for j := left; j <= right; j++ {
			if s[j] == '|' {
				if candleLeft {
					candleLeft = false
				} else {
					candleLeft = true
				}
			} else if s[j] == '*' && candleLeft {
				count++
			}
		}

		result[i] = count
	}

	return result
}

func main() {
	timeStartWholeProgram := time.Now()

	testInput := []TestCase{
		{
			S:       "**|**|***|",
			Queries: [][]int{{2, 5}, {5, 9}},
			Result: `
      [2,3]
            `,
		},

		{
			S:       "***|**|*****|**||**|*",
			Queries: [][]int{{1, 17}, {4, 5}, {14, 17}, {5, 11}, {15, 16}},
			Result: `
      [9,0,0,0,0]
            `,
		},
	}
	for count, value := range testInput {
		fmt.Println("===============")
		fmt.Println("Test count ", count, "for node", value)
		fmt.Println("Solution 1: Binary Search ")
		timeStart := time.Now()
		result := platesBetweenCandles_BinarySearch(value.S, value.Queries)
		timeLapse := time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

		fmt.Println("Solution 2: BruteForce: note: solution failed, need check")
		timeStart = time.Now()
		result = platesBetweenCandles_BruteForce(value.S, value.Queries)
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
	S       string
	Queries [][]int
	Result  string
}

/*


===============
Test count  0 for node {**|**|***| [[2 5] [5 9]]
      [2,3]
            }
Solution 1: Binary Search
>Solution result [2 3]
Correct result is
      [2,3]

TimeLapse 2.093µs
Solution 2: BruteForce: note: solution failed, need check
>Solution result [2 3]
Correct result is
      [2,3]

TimeLapse 871ns
===============
Test count  1 for node {***|**|*****|**||**|* [[1 17] [4 5] [14 17] [5 11] [15 16]]
      [9,0,0,0,0]
            }
Solution 1: Binary Search
>Solution result [9 0 0 0 0]
Correct result is
      [9,0,0,0,0]

TimeLapse 14.204µs
Solution 2: BruteForce: note: solution failed, need check
>Solution result [5 0 0 5 0]
Correct result is
      [9,0,0,0,0]

TimeLapse 889ns
===============
TimeLapse Whole Program 427.105µs


*/
