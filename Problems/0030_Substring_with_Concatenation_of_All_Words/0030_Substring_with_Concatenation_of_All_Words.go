package main

import (
	"fmt"
	"strings"
	"time"
)

//approach: Sliding Window
//Time Complexity: O(n * m) + O(n * m) * O(totalLen / wordLen), which simplifies to O(n * m * (n / wordLen)).
//Space Complexity: .
//wordMap: We use a hashmap wordMap to store the frequency of each word. The space complexity for this hashmap is O(n * m), where n is the number of words and m is the average length of each word.
//tempMap: Within the isConcatenatedSubstring function, we use a temporary hashmap tempMap. The space complexity for this hashmap is O(totalLen / wordLen).
//Therefore, the total space complexity is O(n * m) + O(totalLen / wordLen).

func findSubstring_SlidingWindow(s string, words []string) []int {
	if len(words) == 0 || len(words[0]) == 0 {
		return []int{}
	}

	wordLen := len(words[0])
	wordsCount := len(words)
	totalLen := wordLen * wordsCount
	result := []int{}

	if len(s) < totalLen {
		return result
	}

	wordMap := make(map[string]int)
	for _, word := range words {
		wordMap[word]++
	}

	for i := 0; i <= len(s)-totalLen; i++ {
		if isConcatenatedSubstring(s[i:i+totalLen], wordMap, wordLen) {
			result = append(result, i)
		}
	}

	return result
}

func isConcatenatedSubstring(sub string, wordMap map[string]int, wordLen int) bool {
	tempMap := make(map[string]int)

	for i := 0; i < len(sub); i += wordLen {
		word := sub[i : i+wordLen]
		if _, ok := wordMap[word]; !ok {
			return false
		}
		tempMap[word]++
		if tempMap[word] > wordMap[word] {
			return false
		}
	}

	return true
}

//approach : dynamic programing
/*

Time Complexity:
Iteration through s: We iterate through the string s once, and for each starting index, we traverse through a substring of length totalLen. This operation takes O(n * m) time, where n is the length of the string s and m is the average length of each word.

Iterating over words: Within each iteration of s, we also iterate over the words to build tempMap. This operation takes O(wordsCount) time.

Overall, the time complexity is O(n * m * wordsCount).

Space Complexity:
wordMap: We use a hashmap wordMap to store the frequency of each word. The space complexity for this hashmap is O(m), where m is the number of unique words in words.

tempMap: We use a temporary hashmap tempMap to keep track of the words encountered so far in the current substring. The space complexity for this hashmap is O(m), where m is the number of unique words in words.

result: The space complexity for the result slice is O(k), where k is the number of concatenated substrings found.

Overall, the space complexity is O(m) + O(m) + O(k), which simplifies to O(m + k).

Note:
In the worst case, if all the words in words are unique, then the space complexity can be O(n), where n is the number of words.
The actual time and space complexity will depend on the distribution and length of words in words and the length of the string s.

*/

func findSubstring_DP(s string, words []string) []int {
	if len(words) == 0 || len(words[0]) == 0 {
		return []int{}
	}

	wordLen := len(words[0])
	wordsCount := len(words)
	totalLen := wordLen * wordsCount
	result := []int{}

	if len(s) < totalLen {
		return result
	}

	wordMap := make(map[string]int)
	for _, word := range words {
		wordMap[word]++
	}

	for i := 0; i < wordLen; i++ {
		left := i
		count := 0
		tempMap := make(map[string]int)

		for j := i; j <= len(s)-wordLen; j += wordLen {
			word := s[j : j+wordLen]
			if _, ok := wordMap[word]; ok {
				tempMap[word]++
				count++
				for tempMap[word] > wordMap[word] {
					tempMap[s[left:left+wordLen]]--
					left += wordLen
					count--
				}
				if count == wordsCount {
					result = append(result, left)
					tempMap[s[left:left+wordLen]]--
					left += wordLen
					count--
				}
			} else {
				tempMap = make(map[string]int)
				count = 0
				left = j + wordLen
			}
		}
	}

	return result
}

// approach : DFS
/*

Time Complexity:
Permutation Generation: Generating all permutations of the words array involves factorial time complexity, which is O(n!), where n is the number of words in the array. This is because there are n! possible permutations of n elements.
Substring Checking: For each permutation, we check if it forms a substring of the given string s. The strings.Contains function has a time complexity of O(n * m), where n is the length of s and m is the length of the concatenated substring.
Overall, the time complexity of the DFS-based approach is O(n! * n * m).

Space Complexity:
visited array: We use a boolean array visited to keep track of visited words during the permutation generation. Its space complexity is O(n), where n is the number of words in the array.
permutation array: We use an array permutation to store the current permutation being constructed. Its space complexity is O(n), where n is the number of words in the array.
result array: The space complexity for the result slice is O(k), where k is the number of concatenated substrings found.
Overall, the space complexity is O(n + k), where n is the number of words in the array and k is the number of concatenated substrings found.

Note:
The DFS-based approach is not efficient due to its factorial time complexity, especially for large inputs. It's provided here mainly for educational purposes to demonstrate an alternative solution using DFS.
*/

func findSubstring_DFS(s string, words []string) []int {
	if len(words) == 0 || len(words[0]) == 0 {
		return []int{}
	}

	result := []int{}
	visited := make([]bool, len(words))
	permutation := make([]string, len(words))

	dfs(s, words, visited, permutation, 0, &result)

	return result
}

func dfs(s string, words []string, visited []bool, permutation []string, index int, result *[]int) {
	if index == len(words) {
		concatenated := strings.Join(permutation, "")
		if strings.Contains(s, concatenated) {
			*result = append(*result, strings.Index(s, concatenated))
		}
		return
	}

	for i := 0; i < len(words); i++ {
		if !visited[i] {
			visited[i] = true
			permutation[index] = words[i]
			dfs(s, words, visited, permutation, index+1, result)
			visited[i] = false
		}
	}
}

// approach : BFS
//NOTE: wrong, need check code again
/*

Time Complexity:
Iteration through words: Building the wordMap takes O(n * m) time, where n is the number of words and m is the average length of each word.
BFS: We start BFS from each index in the range [0, wordLen) to cover all possible starting points of substrings. Within each BFS iteration, we explore substrings and check if they are valid concatenated substrings. The worst-case time complexity for exploring all possible substrings is O(n * m * k), where n is the length of s, m is the average length of each word, and k is the number of words.
Overall, the time complexity is O(n * m * k).

Space Complexity:
wordMap: We use a hashmap wordMap to store the frequency of each word. The space complexity for this hashmap is O(n * m), where n is the number of words and m is the average length of each word.
Queue: We maintain a queue to store words during BFS. In the worst case, the queue can contain all words, which contributes to the space complexity with O(k), where k is the number of words.
visited map: We use a visited map to keep track of visited words during BFS. In the worst case, the visited map can contain all words, which contributes to the space complexity with O(k), where k is the number of words.
result array: The space complexity for the result slice is O(k'), where k' is the number of concatenated substrings found.
Overall, the space complexity is O(n * m) + O(k) + O(k) + O(k'), which simplifies to O(n * m + k + k').

*/

func findSubstring_BFS(s string, words []string) []int {
	if len(words) == 0 || len(words[0]) == 0 {
		return []int{}
	}

	wordLen := len(words[0])
	totalLen := wordLen * len(words)
	result := []int{}
	wordMap := make(map[string]int)

	for _, word := range words {
		wordMap[word]++
	}

	for i := 0; i < wordLen; i++ {
		queue := make([]string, 0)
		visited := make(map[string]int)
		for _, word := range words {
			queue = append(queue, word)
		}

		for j := i; j <= len(s)-totalLen; j += wordLen {
			temp := append([]string(nil), queue...)
			for len(temp) > 0 {
				substr := temp[0]
				temp = temp[1:]
				if strings.HasPrefix(s[j:], substr) {
					visited[substr]++
					if visited[substr] > wordMap[substr] {
						break
					}
					if len(visited) == len(wordMap) {
						result = append(result, j)
						break
					}
					temp = append(temp, queue...)
				}
			}
		}
	}

	return result
}

// approach : Memoization
/*
Time Complexity:
Iterating through Substrings: The solution iterates through all possible substrings of length totalLen in the string s. This involves iterating n - totalLen + 1 times, where n is the length of the string s. For each iteration, checking if a substring is a concatenated substring takes O(totalLen) time.

Overall, the time complexity is O((n - totalLen + 1) * totalLen), which simplifies to O(n * totalLen).

Space Complexity:
wordMap: We use a hashmap wordMap to store the frequency of each word. The space complexity for this hashmap is O(m), where m is the number of unique words in words.

tempMap: Within the isConcatenatedSubstring function, we use a temporary hashmap tempMap. The space complexity for this hashmap is O(m), where m is the number of unique words in words.

result: The space complexity for the result slice is O(k), where k is the number of concatenated substrings found.

Overall, the space complexity is O(m) + O(m) + O(k), which simplifies to O(m + k).

Note:
The time complexity of this solution can be high, especially for large inputs, due to the brute-force nature of checking all possible substrings.
The actual time and space complexity will depend on the distribution and length of words in words and the length of the string s. If the words are very short compared to the length of s, the complexity can be close to O(n^2).
*/

func findSubstring_Memoization(s string, words []string) []int {
	if len(words) == 0 || len(words[0]) == 0 {
		return []int{}
	}

	wordLen := len(words[0])
	totalLen := wordLen * len(words)
	result := []int{}
	wordMap := make(map[string]int)

	for _, word := range words {
		wordMap[word]++
	}

	for i := 0; i < len(s)-totalLen+1; i++ {
		substr := s[i : i+totalLen]
		if isConcatenatedSubstring_Memoization(substr, wordMap, wordLen) {
			result = append(result, i)
		}
	}

	return result
}

func isConcatenatedSubstring_Memoization(s string, wordMap map[string]int, wordLen int) bool {
	tempMap := make(map[string]int)
	wordsCount := len(wordMap)

	for i := 0; i < len(s); i += wordLen {
		word := s[i : i+wordLen]
		if _, ok := wordMap[word]; !ok {
			return false
		}
		tempMap[word]++
		if tempMap[word] > wordMap[word] {
			return false
		}
	}

	return len(tempMap) == wordsCount
}

func main() {
	timeStartWholeProgram := time.Now()

	testInput := []TestCase{

		{
			S:     "barfoothefoobarman",
			Words: []string{"foo", "bar"},
			Result: `
[0,9]

            `,
		},
		{
			S:     "wordgoodgoodgoodbestword",
			Words: []string{"word", "good", "best", "word"},
			Result: `
      []
            `,
		},
		{
			S:     "barfoofoobarthefoobarman",
			Words: []string{"bar", "foo", "the"},
			Result: `
      [6,9,12]
            `,
		},
	}
	for count, value := range testInput {
		fmt.Println("===============")
		fmt.Println("Test count ", count, "for node", value)
		fmt.Println("Solution 1: Sliding Window")
		timeStart := time.Now()
		result := findSubstring_SlidingWindow(value.S, value.Words)
		timeLapse := time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

		fmt.Println("Solution 2: DP")
		timeStart = time.Now()
		result = findSubstring_DP(value.S, value.Words)
		timeLapse = time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

		fmt.Println("Solution 3: DFS")
		timeStart = time.Now()
		result = findSubstring_DFS(value.S, value.Words)
		timeLapse = time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

		fmt.Println("Solution 4: BFS. NOTE: this solution is not complete, need check code again")
		timeStart = time.Now()
		result = findSubstring_BFS(value.S, value.Words)
		timeLapse = time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

		fmt.Println("Solution 5: Memoization")
		timeStart = time.Now()
		result = findSubstring_Memoization(value.S, value.Words)
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
	Words  []string
	Result string
}

/*

===============
Test count  0 for node {barfoothefoobarman [foo bar]
[0,9]

            }
Solution 1: Sliding Window
>Solution result [0 9]
Correct result is
[0,9]


TimeLapse 6.537µs
Solution 2: DP
>Solution result [0 9]
Correct result is
[0,9]


TimeLapse 6.315µs
Solution 3: DFS
>Solution result [9 0]
Correct result is
[0,9]


TimeLapse 5.444µs
Solution 4: BFS. NOTE: this solution is not complete, need check code again
>Solution result [3]
Correct result is
[0,9]


TimeLapse 7.296µs
Solution 5: Memoization
>Solution result [0 9]
Correct result is
[0,9]


TimeLapse 4.278µs
===============
Test count  1 for node {wordgoodgoodgoodbestword [word good best word]
      []
            }
Solution 1: Sliding Window
>Solution result []
Correct result is
      []

TimeLapse 2.834µs
Solution 2: DP
>Solution result []
Correct result is
      []

TimeLapse 5.314µs
Solution 3: DFS
>Solution result []
Correct result is
      []

TimeLapse 9.574µs
Solution 4: BFS. NOTE: this solution is not complete, need check code again
>Solution result []
Correct result is
      []

TimeLapse 34.129µs
Solution 5: Memoization
>Solution result []
Correct result is
      []

TimeLapse 3.481µs
===============
Test count  2 for node {barfoofoobarthefoobarman [bar foo the]
      [6,9,12]
            }
Solution 1: Sliding Window
>Solution result [6 9 12]
Correct result is
      [6,9,12]

TimeLapse 7.778µs
Solution 2: DP
>Solution result [6 9 12]
Correct result is
      [6,9,12]

TimeLapse 8.296µs
Solution 3: DFS
>Solution result [9 6 12]
Correct result is
      [6,9,12]

TimeLapse 6µs
Solution 4: BFS. NOTE: this solution is not complete, need check code again
>Solution result [12]
Correct result is
      [6,9,12]

TimeLapse 11.777µs
Solution 5: Memoization
>Solution result [6 9 12]
Correct result is
      [6,9,12]

TimeLapse 7.074µs
===============
TimeLapse Whole Program 1.041317ms

*/
//REF
//
