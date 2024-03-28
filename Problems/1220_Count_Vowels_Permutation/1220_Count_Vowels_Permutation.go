package main

import (
	"fmt"
	"runtime"
	"time"
)

const mod = int(1e9) + 7

//approach dynamic Programming

func countVowelPermutation_DP(n int) int {
	// Initialize the DP array to store counts for each vowel ending
	dp := make([][5]int, n+1)

	// Initialize base cases for strings of length 1
	for i := 0; i < 5; i++ {
		dp[1][i] = 1
	}

	// Dynamic Programming: Fill DP table
	for i := 2; i <= n; i++ {
		dp[i][0] = (dp[i-1][1] + dp[i-1][2] + dp[i-1][4]) % mod // 'a' can follow 'e', 'i', or 'u'
		dp[i][1] = (dp[i-1][0] + dp[i-1][2]) % mod              // 'e' can follow 'a' or 'i'
		dp[i][2] = (dp[i-1][1] + dp[i-1][3]) % mod              // 'i' can follow 'e' or 'o'
		dp[i][3] = dp[i-1][2]                                   // 'o' can only follow 'i'
		dp[i][4] = (dp[i-1][2] + dp[i-1][3]) % mod              // 'u' can only follow 'i' or 'o'
	}

	// Sum up counts for all possible strings of length n
	total := 0
	for _, count := range dp[n] {
		total = (total + count) % mod
	}

	return total
}

//approach DFS

func countVowelPermutation_DFS(n int) int {
	// Initialize a counter to store the total count
	count := 0

	// Start DFS from each vowel as the starting point
	vowels := []rune{'a', 'e', 'i', 'o', 'u'}
	for _, vowel := range vowels {
		count += dfs(vowel, 1, n)
		count %= mod
	}

	return count
}

func dfs(pos rune, length, n int) int {
	// If the length of the string reaches n, return 1
	if length == n {
		return 1
	}

	// Initialize a variable to store the total count
	total := 0

	// Try appending each vowel according to the rules
	nextVowels := getNextVowels_DFS(pos)
	for _, next := range nextVowels {
		// Recur with the next position and updated length
		total += dfs(next, length+1, n)
		total %= mod
	}

	return total
}

func getNextVowels_DFS(pos rune) []rune {
	switch pos {
	case 'a':
		return []rune{'e'}
	case 'e':
		return []rune{'a', 'i'}
	case 'i':
		return []rune{'a', 'e', 'o', 'u'}
	case 'o':
		return []rune{'i', 'u'}
	case 'u':
		return []rune{'a'}
	default:
		return nil
	}
}

//approach BFS

func countVowelPermutation_BFS(n int) int {
	// Initialize a queue for BFS
	queue := make([]state, 0)

	// Initialize the count of valid strings
	count := 0

	// Start BFS from each vowel as the starting point
	vowels := []rune{'a', 'e', 'i', 'o', 'u'}
	for _, vowel := range vowels {
		queue = append(queue, state{vowel, 1})
	}

	// Perform BFS
	for len(queue) > 0 {
		// Dequeue a state
		currState := queue[0]
		queue = queue[1:]

		// If the length reaches n, increment the count
		if currState.length == n {
			count = (count + 1) % mod
			continue
		}

		// Get next possible vowels based on current position
		nextVowels := getNextVowels_BFS(currState.pos)

		// Enqueue next states
		for _, next := range nextVowels {
			queue = append(queue, state{next, currState.length + 1})
		}
	}

	return count
}

type state struct {
	pos    rune
	length int
}

func getNextVowels_BFS(pos rune) []rune {
	switch pos {
	case 'a':
		return []rune{'e'}
	case 'e':
		return []rune{'a', 'i'}
	case 'i':
		return []rune{'a', 'e', 'o', 'u'}
	case 'o':
		return []rune{'i', 'u'}
	case 'u':
		return []rune{'a'}
	default:
		return nil
	}
}

// approach brute force
func countVowelPermutation_BruteForce(n int) int {
	// Initialize a counter to store the total count
	count := 0

	// Generate all possible strings of length n
	generate("", n, &count)

	return count
}

func generate(current string, n int, count *int) {
	// If the length of the current string reaches n, check if it's valid
	if len(current) == n {
		if isValid(current) {
			*count = (*count + 1) % mod
		}
		return
	}

	// Append each vowel to the current string and recurse
	vowels := []rune{'a', 'e', 'i', 'o', 'u'}
	for _, vowel := range vowels {
		generate(current+string(vowel), n, count)
	}
}

func isValid(s string) bool {
	for i := 0; i < len(s)-1; i++ {
		switch s[i] {
		case 'a':
			if s[i+1] != 'e' {
				return false
			}
		case 'e':
			if s[i+1] != 'a' && s[i+1] != 'i' {
				return false
			}
		case 'i':
			if s[i+1] == 'i' {
				return false
			}
		case 'o':
			if s[i+1] != 'i' && s[i+1] != 'u' {
				return false
			}
		case 'u':
			if s[i+1] != 'a' {
				return false
			}
		}
	}
	return true
}

func main() {
	timeStartWholeProgram := time.Now()

	testInput := []TestCase{
		{
			N: 1,
			Result: `
  5
            `,
		},
		{
			N: 2,
			Result: `
10
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
		fmt.Println("Solution 1: DP")
		timeStart := time.Now()
		result := countVowelPermutation_DP(value.N)
		timeLapse := time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

		// Memory after allocation
		runtime.ReadMemStats(&m)
		memAfter := m.Alloc
		fmt.Println("Memory before", memBefore, "bytes", "Memory after", memAfter, "bytes", "Memory used:", memAfter-memBefore, "bytes")

		fmt.Printf("Memory usage (HeapAlloc) after Test Case i %d, : %v bytes\n", count, m.HeapAlloc)

		fmt.Println("Solution 2: DFS")
		timeStart = time.Now()
		result = countVowelPermutation_DFS(value.N)
		timeLapse = time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

		// Memory after allocation
		runtime.ReadMemStats(&m)
		memAfter = m.Alloc
		fmt.Println("Memory before", memBefore, "bytes", "Memory after", memAfter, "bytes", "Memory used:", memAfter-memBefore, "bytes")

		fmt.Printf("Memory usage (HeapAlloc) after Test Case i %d, : %v bytes\n", count, m.HeapAlloc)

		fmt.Println("Solution 3: BFS")
		timeStart = time.Now()
		result = countVowelPermutation_BFS(value.N)
		timeLapse = time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

		// Memory after allocation
		runtime.ReadMemStats(&m)
		memAfter = m.Alloc
		fmt.Println("Memory before", memBefore, "bytes", "Memory after", memAfter, "bytes", "Memory used:", memAfter-memBefore, "bytes")

		fmt.Printf("Memory usage (HeapAlloc) after Test Case i %d, : %v bytes\n", count, m.HeapAlloc)

		fmt.Println("Solution 4: BruteForce")
		timeStart = time.Now()
		result = countVowelPermutation_BruteForce(value.N)
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
	N      int
	Result string
}

/*

===============
Test count  0 for node {1
  5
            }
Solution 1: DP
>Solution result 5
Correct result is
  5

TimeLapse 2.148µs
Memory before 69208 bytes Memory after 70184 bytes Memory used: 976 bytes
Memory usage (HeapAlloc) after Test Case i 0, : 70184 bytes
Solution 2: DFS
>Solution result 5
Correct result is
  5

TimeLapse 722ns
Memory before 69208 bytes Memory after 70376 bytes Memory used: 1168 bytes
Memory usage (HeapAlloc) after Test Case i 0, : 70376 bytes
Solution 3: BFS
>Solution result 5
Correct result is
  5

TimeLapse 3.148µs
Memory before 69208 bytes Memory after 70680 bytes Memory used: 1472 bytes
Memory usage (HeapAlloc) after Test Case i 0, : 70680 bytes
Solution 4: BruteForce
>Solution result 5
Correct result is
  5

TimeLapse 3.574µs
Memory before 69208 bytes Memory after 70760 bytes Memory used: 1552 bytes
Memory usage (HeapAlloc) after Test Case i 0, : 70760 bytes
===============
Test count  1 for node {2
10
            }
Solution 1: DP
>Solution result 10
Correct result is
10

TimeLapse 1.148µs
Memory before 69208 bytes Memory after 70976 bytes Memory used: 1768 bytes
Memory usage (HeapAlloc) after Test Case i 1, : 70976 bytes
Solution 2: DFS
>Solution result 10
Correct result is
10

TimeLapse 815ns
Memory before 69208 bytes Memory after 71040 bytes Memory used: 1832 bytes
Memory usage (HeapAlloc) after Test Case i 1, : 71040 bytes
Solution 3: BFS
>Solution result 10
Correct result is
10

TimeLapse 15.555µs
Memory before 69208 bytes Memory after 71808 bytes Memory used: 2600 bytes
Memory usage (HeapAlloc) after Test Case i 1, : 71808 bytes
Solution 4: BruteForce
>Solution result 10
Correct result is
10

TimeLapse 5.074µs
Memory before 69208 bytes Memory after 71936 bytes Memory used: 2728 bytes
Memory usage (HeapAlloc) after Test Case i 1, : 71936 bytes
===============
TimeLapse Whole Program 1.195322ms

*/
//REF
//https://chat.openai.com/c/58beb139-b62f-490c-9a6f-3e82895fc7b0
