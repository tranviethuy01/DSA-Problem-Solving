package main

import (
	"fmt"
	"strconv"
	"sort"
	"time"
)

//approach : Backtrack
/*
Time Complexity:
Calculating the factorial of each number up to n requires iterating from 1 to n, resulting in O(n) operations.
The main loop iterates n times, and inside each iteration, there's a calculation of k modulo and division by the factorial, which can be done in constant time.
Additionally, there's a loop to construct the resulting string by appending digits, which also takes O(n) time.
Therefore, the overall time complexity is O(n + n) = O(n).

Space Complexity:
The space complexity is dominated by the arrays factorial and nums, each of size n, thus occupying O(n) space.
The additional space used inside the function is independent of the input size and can be considered constant.
*/



func getPermutation(n int, k int) string {
	factorial := make([]int, n)
	nums := make([]int, n)

	factorial[0] = 1
	for i := 1; i < n; i++ {
		factorial[i] = factorial[i-1] * i
	}

	for i := 1; i <= n; i++ {
		nums[i-1] = i
	}

	k--

	var result string
	for i := n - 1; i >= 0; i-- {
		idx := k / factorial[i]
		k %= factorial[i]
		result += strconv.Itoa(nums[idx])
		nums = append(nums[:idx], nums[idx+1:]...)
	}

	return result
}

//approach : DFS
/*
Time Complexity:

Calculating the factorial array takes O(n) time.
The recursive function dfs is called recursively for each digit in the permutation. The function performs constant time operations inside its body.
In each recursive call, we're reducing the size of the remaining array by one.
Since the dfs function is called recursively n times and each call performs operations that are O(1), the overall time complexity is O(n).
Space Complexity:

The space complexity is dominated by the recursive calls on the call stack. In the worst case, the recursion depth could be n, corresponding to the length of the remaining array.
Additionally, we use an array remaining to keep track of the digits, which requires O(n) space.
Therefore, the overall space complexity is O(n).
*/
func getPermutation_DFS(n, k int) string {
	// Calculate factorial
	fact := make([]int, n)
	fact[0] = 1
	for i := 1; i < n; i++ {
		fact[i] = fact[i-1] * i
	}

	// Recursive function to get kth permutation
	var dfs func(remaining []int, k int) string
	dfs = func(remaining []int, k int) string {
		if len(remaining) == 0 {
			return ""
		}

		idx := (k - 1) / fact[len(remaining)-1]
		nextK := k - idx*fact[len(remaining)-1]

		num := remaining[idx]
		remaining = append(remaining[:idx], remaining[idx+1:]...)

		return strconv.Itoa(num) + dfs(remaining, nextK)
	}

	// Create initial array
	remaining := make([]int, n)
	for i := 0; i < n; i++ {
		remaining[i] = i + 1
	}

	// Start recursion
	return dfs(remaining, k)
}


// approach: BFS
// NOTE: this approach is failure, need check code again
/*
Time Complexity:

Calculating the factorial array takes O(n) time.
In the worst case, the BFS algorithm will generate all permutations of length n before finding the kth permutation. Generating each permutation involves iterating over n elements and performing constant-time operations.
Therefore, the time complexity of generating all permutations is O(n! * n).
However, since we're only interested in finding the kth permutation, we can terminate the BFS search once we find it. Thus, in the best-case scenario where we find the kth permutation early, the time complexity would be less than O(n! * n), but in the worst-case scenario, it's O(n! * n).
Space Complexity:

We use a factorial array of size n, which requires O(n) space.
We maintain a visited map to keep track of visited permutations. In the worst case, the map could contain all permutations, which would require O(n!) space.
The BFS queue can potentially store all permutations before finding the kth permutation, which again is of the order of O(n!).
Therefore, the overall space complexity is dominated by the visited map and the BFS queue, resulting in O(n!) space.


*/

func getPermutation_BFS(n, k int) string {
	factorial := make([]int, n)
	factorial[0] = 1
	for i := 1; i < n; i++ {
		factorial[i] = factorial[i-1] * i
	}

	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = i + 1
	}

	visited := make(map[string]bool)

	queue := make([]string, 0)
	queue = append(queue, "")

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if len(current) == n {
			k--
			if k == 0 {
				return current
			}
			continue
		}

		for i := 0; i < n; i++ {
			num := nums[i]
			if visited[current+strconv.Itoa(num)] {
				continue
			}

			newK := k - factorial[n-len(current)-1]
			if newK > 0 {
				k = newK
				continue
			}

			visited[current+strconv.Itoa(num)] = true
			queue = append(queue, current+strconv.Itoa(num))
		}
	}

	return ""
}



// approach: dynamic programing
/*
Time Complexity:

Precomputing the factorial array takes O(n) time.
The main loop iterates over each digit position from right to left, which requires O(n) iterations.
Inside the loop, we iterate over the array of size n to find the correct digit, which takes O(n) time.
Therefore, the overall time complexity is O(n^2).
Space Complexity:

We use additional space for the factorial array, which requires O(n) space.
We use a boolean array used to keep track of used digits, which also requires O(n) space.
The space used for other variables is negligible and can be considered constant.
Therefore, the overall space complexity is O(n).

*/


func getPermutation_DP(n, k int) string {
	factorial := make([]int, n)
	factorial[0] = 1
	for i := 1; i < n; i++ {
		factorial[i] = factorial[i-1] * i
	}

	used := make([]bool, n)
	var result string

	k-- // Convert k to 0-based index

	for i := n - 1; i >= 0; i-- {
		digitIndex := k / factorial[i]
		k %= factorial[i]

		for j := 0; j < n; j++ {
			if !used[j] {
				if digitIndex == 0 {
					result += strconv.Itoa(j + 1)
					used[j] = true
					break
				}
				digitIndex--
			}
		}
	}

	return result
}

// approach : BruteForce
/*
Time Complexity:
Generating all permutations of the given set involves backtracking, where at each step, we have 
n choices to make for the current position and repeat the process for the remaining positions. This results in 
n! permutations.
For each permutation, the conversion to a string takes linear time, but it's negligible compared to the permutation generation process.
Therefore, the time complexity is O(n!), where 
n is the given integer representing the size of the set.

Space Complexity:
We use additional space to store the permutations. Since there are
n! permutations, the space complexity is 
O(n!⋅n) to store all permutations. Each permutation has a length of 
n characters.
Additionally, during the recursive generation of permutations, there's a stack space required for the recursion, which is proportional to the depth of recursion. In the worst case, this depth is 
n, so the space complexity for recursion is 
O(n).
Hence, the overall space complexity is O(n!⋅n)+O(n), which simplifies to O(n!⋅n).
In summary, the brute-force solution has a time complexity of O(n!) and a space complexity of 
O(n!⋅n), making it inefficient for large values of n.

*/

func getPermutation_BruteForce(n, k int) string {
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = i + 1
	}

	permutations := make([]string, 0)
	generatePermutations(nums, &permutations, 0, n)

	sort.Strings(permutations)

	return permutations[k-1]
}

func generatePermutations(nums []int, permutations *[]string, index, n int) {
	if index == n {
		*permutations = append(*permutations, convertToString(nums))
		return
	}

	for i := index; i < n; i++ {
		swap(nums, index, i)
		generatePermutations(nums, permutations, index+1, n)
		swap(nums, index, i)
	}
}

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}

func convertToString(nums []int) string {
	result := ""
	for _, num := range nums {
		result += strconv.Itoa(num)
	}
	return result
}



func main() {
	timeStartWholeProgram := time.Now()

	testInput := []TestCase{

		{
			N: 3,
			K: 3,

			Result: `
"213"
            `,
		},
		{
			N: 4,
			K: 9,
			Result: `
"2314"
            `,
		},
		{
			N: 3,
			K: 1,
			Result: `
"123"

            `,
		},
	}
	for count, value := range testInput {
		fmt.Println("===============")
		fmt.Println("Test count ", count, "for node", value)
		fmt.Println("Solution 1: Backtrack")
		timeStart := time.Now()
		result := getPermutation(value.N, value.K)
		timeLapse := time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

		
		fmt.Println("Solution 2: DFS")
		timeStart = time.Now()
		result = getPermutation_DFS(value.N, value.K)
		timeLapse = time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)
	
	
		fmt.Println("Solution 3: BFS : note, failure solution, need check code")
		timeStart = time.Now()
		result = getPermutation_BFS(value.N, value.K)
		timeLapse = time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

	
		fmt.Println("Solution 4: DP")
		timeStart = time.Now()
		result = getPermutation_DP(value.N, value.K)
		timeLapse = time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

		
		fmt.Println("Solution 5: BruteForce")
		timeStart = time.Now()
		result = getPermutation_BruteForce(value.N, value.K)
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
	N      int
	K      int
	Result string
}

/*

===============
Test count  0 for node {3 3 
"213"
            }
Solution 1: Backtrack
>Solution result 213
Correct result is  
"213"
            
TimeLapse 11.684µs
Solution 2: DFS
>Solution result 213
Correct result is  
"213"
            
TimeLapse 2.056µs
Solution 3: BFS : note, failure solution, need check code
>Solution result 211
Correct result is  
"213"
            
TimeLapse 69.572µs
Solution 4: DP
>Solution result 213
Correct result is  
"213"
            
TimeLapse 1.611µs
Solution 5: BruteForce
>Solution result 213
Correct result is  
"213"
            
TimeLapse 27.351µs
===============
Test count  1 for node {4 9 
"2314"
            }
Solution 1: Backtrack
>Solution result 2314
Correct result is  
"2314"
            
TimeLapse 1.481µs
Solution 2: DFS
>Solution result 2314
Correct result is  
"2314"
            
TimeLapse 1.574µs
Solution 3: BFS : note, failure solution, need check code
>Solution result 2211
Correct result is  
"2314"
            
TimeLapse 249.844µs
Solution 4: DP
>Solution result 2314
Correct result is  
"2314"
            
TimeLapse 1.444µs
Solution 5: BruteForce
>Solution result 2314
Correct result is  
"2314"
            
TimeLapse 26.574µs
===============
Test count  2 for node {3 1 
"123"

            }
Solution 1: Backtrack
>Solution result 123
Correct result is  
"123"

            
TimeLapse 1.241µs
Solution 2: DFS
>Solution result 123
Correct result is  
"123"

            
TimeLapse 1.333µs
Solution 3: BFS : note, failure solution, need check code
>Solution result 111
Correct result is  
"123"

            
TimeLapse 51.609µs
Solution 4: DP
>Solution result 123
Correct result is  
"123"

            
TimeLapse 1.277µs
Solution 5: BruteForce
>Solution result 123
Correct result is  
"123"

            
TimeLapse 4.574µs
===============
TimeLapse Whole Program 1.409824ms

*/
//REF
//
