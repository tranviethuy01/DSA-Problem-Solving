package main

import (
	"fmt"
	"math"
	"time"
)

//NOTE: should use DP or 2 pointers approach

//approach DP
/*
Algorithm:
Iterate over the array once to calculate the maximum height of bars to the left of each element.
Iterate over the array again to calculate the maximum height of bars to the right of each element.
Iterate over the array once more to calculate the amount of water that can be trapped at each position by taking the minimum of the left and right maximum heights and subtracting the height of the current element.
Sum up the trapped water for all positions.
Time Complexity:
Let's denote
n as the number of elements in the input array.

Calculating the left maximum heights:
O(n)
Calculating the right maximum heights:
O(n)
Calculating the trapped water:
O(n)
Overall, the time complexity is
O(n).

Space Complexity:
The solution uses additional space to store the left and right maximum heights, each requiring an array of size
n. Thus, the space complexity is
O(n) for the two arrays.
Therefore, the overall space complexity of the solution is
O(n).
*/

func trap_DP(height []int) int {
	n := len(height)
	if n == 0 {
		return 0
	}

	leftMax := make([]int, n)
	rightMax := make([]int, n)

	leftMax[0] = height[0]
	for i := 1; i < n; i++ {
		leftMax[i] = max(leftMax[i-1], height[i])
	}

	rightMax[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], height[i])
	}

	totalWater := 0
	for i := 0; i < n; i++ {
		totalWater += min(leftMax[i], rightMax[i]) - height[i]
	}

	return totalWater
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//approach  2 pointers
/*
Time Complexity:
The solution iterates through the height array using two pointers (left and right) until they meet, which requires a single pass through the array.
Hence, the time complexity is linear,
O(n), where
n is the number of elements in the height array.
Space Complexity:

The solution uses only a constant amount of additional space, regardless of the input size. It maintains a few variables (left, right, leftMax, rightMax, total), but these consume a constant amount of memory.
Thus, the space complexity is constant,
O(1).
In summary, the time complexity of the two-pointer solution is
O(n), and the space complexity is
O(1), where
n is the number of elements in the height array. This makes the solution efficient in terms of both time and space.

*/
func trap_2Pointer(height []int) int {
	if len(height) < 3 {
		return 0
	}

	left, right := 0, len(height)-1
	leftMax, rightMax := 0, 0
	total := 0

	for left < right {
		if height[left] < height[right] {
			if height[left] >= leftMax {
				leftMax = height[left]
			} else {
				total += leftMax - height[left]
			}
			left++
		} else {
			if height[right] >= rightMax {
				rightMax = height[right]
			} else {
				total += rightMax - height[right]
			}
			right--
		}
	}

	return total
}

//approach backtrack
/*
Algorithm:
Iterate through each bar of the elevation map.
For each bar, find the maximum height to its left and right.
Calculate the trapped water at the current bar as the minimum of the left and right maximum heights minus the height of the current bar.
Sum up the trapped water for all bars.
Time Complexity:
Let
n be the number of elements in the input array.

For each element in the array, the algorithm iterates through all elements to its left and right to find the maximum heights, resulting in
O(n) operations for each element.
Therefore, the overall time complexity of the solution is
O(n ^ 2).
Space Complexity:
The space complexity of the provided solution is
O(1) since it only uses a constant amount of extra space for variables such as maxLeft, maxRight, and total. It doesn't use any additional data structures that scale with the input size.

Therefore, the time complexity of the provided solution is O(n ^ 2), and the space complexity is O(1).
*/
func trap_Backtrack(height []int) int {
	if len(height) < 3 {
		return 0
	}
	total := 0
	for i := 1; i < len(height)-1; i++ {
		maxLeft := 0
		maxRight := 0
		// Find the maximum height to the left of the current bar
		for j := i; j >= 0; j-- {
			maxLeft = max(maxLeft, height[j])
		}
		// Find the maximum height to the right of the current bar
		for j := i; j < len(height); j++ {
			maxRight = max(maxRight, height[j])
		}
		// Calculate the trapped water at the current bar
		total += min(maxLeft, maxRight) - height[i]
	}
	return total
}

//func max(a, b int) int {
//    if a > b {
//        return a
//    }
//    return b
//}
//
//func min(a, b int) int {
//    if a < b {
//        return a
//    }
//    return b
//}
//

//approach adapt another solution here
/*
class Solution {
public:
    //total water is trapped into the bars
    int trap(vector<int>& h) {
        int l=0,r=h.size()-1,lmax=INT_MIN,rmax=INT_MIN,ans=0;
        while(l<r){
            lmax=max(lmax,h[l]);
            rmax=max(rmax,h[r]);
            ans+=(lmax<rmax)?lmax-h[l++]:rmax-h[r--];
        }
        return ans;
    }
};
*/

func trap_Adapt1(height []int) int {
	l, r := 0, len(height)-1
	lmax, rmax := math.MinInt32, math.MinInt32
	ans := 0
	for l < r {
		lmax = max(lmax, height[l])
		rmax = max(rmax, height[r])
		if lmax < rmax {
			ans += lmax - height[l]
			l++
		} else {
			ans += rmax - height[r]
			r--
		}
	}
	return ans
}

//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}
//

//approach dfs .  it might not be the most efficient approach compared to dynamic programming or two-pointer methods.
/*
Time Complexity:

The trap function iterates through each bar of the elevation map, resulting in
O(n) time complexity.
For each bar, the DFS functions (dfsLeft and dfsRight) traverse through the height array. In the worst case, each DFS function may traverse up to
O(n) elements.
Thus, for each bar, the total time complexity of DFS operations is
O(n).
Overall, the time complexity of the solution is
O(n
2
 ), considering that the DFS operations are performed for each bar.
Space Complexity:

The space complexity primarily comes from the recursive calls in the DFS functions. However, the depth of recursion is limited by the height of the elevation map, which can be at most
O(n) in the worst case.
Each DFS function uses a constant amount of additional space for variables and function call stack.
Therefore, the space complexity of the solution is
O(n).
In summary, the time complexity of the DFS solution is
O(n
2
 ), and the space complexity is
O(n).

*/

func trap_DFS(height []int) int {
	if len(height) < 3 {
		return 0
	}
	total := 0
	for i := 1; i < len(height)-1; i++ {
		leftMax := dfsLeft(height, i-1, i, height[i])
		rightMax := dfsRight(height, i+1, i, height[i])
		minHeight := min(leftMax, rightMax)
		if minHeight > height[i] {
			total += minHeight - height[i]
		}
	}
	return total
}

func dfsLeft(height []int, index, start, currentMax int) int {
	if index < 0 {
		return currentMax
	}
	currentMax = max(currentMax, height[index])
	return dfsLeft(height, index-1, start, currentMax)
}

func dfsRight(height []int, index, start, currentMax int) int {
	if index >= len(height) {
		return currentMax
	}
	currentMax = max(currentMax, height[index])
	return dfsRight(height, index+1, start, currentMax)
}

//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}
//
//func min(a, b int) int {
//	if a < b {
//		return a
//	}
//	return b
//}
//
//
//

//approach BruteForce

func trap_BruteForce(height []int) int {
	if len(height) < 3 {
		return 0
	}
	total := 0
	for i := 1; i < len(height)-1; i++ {
		leftMax := 0
		rightMax := 0
		// Find the maximum height to the left of the current bar
		for j := i; j >= 0; j-- {
			leftMax = max(leftMax, height[j])
		}
		// Find the maximum height to the right of the current bar
		for j := i; j < len(height); j++ {
			rightMax = max(rightMax, height[j])
		}
		// Calculate the trapped water at the current bar
		total += min(leftMax, rightMax) - height[i]
	}
	return total
}

//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}
//
//func min(a, b int) int {
//	if a < b {
//		return a
//	}
//	return b
//}
//
//

func main() {
	timeStartWholeProgram := time.Now()

	testInput := []TestCase{
		{
			Nums: []int{1, 2, 0},
			Result: `
    3
            `,
		},
		{
			Nums: []int{3, 4, -1, 1},
			Result: `
    2
            `,
		},
		{
			Nums: []int{7, 8, 9, 11, 12},
			Result: `
    1
            `,
		},
	}
	for count, value := range testInput {

		fmt.Println("===============")
		fmt.Println("Test count ", count, "for node", value)
		fmt.Println("Solution 1: DP")
		timeStart := time.Now()
		result := trap_DP(value.Nums)
		timeLapse := time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

		fmt.Println("Solution 2: 2 pointers")
		timeStart = time.Now()
		result = trap_2Pointer(value.Nums)
		timeLapse = time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

		fmt.Println("Solution 3: Backtrack ")
		timeStart = time.Now()
		result = trap_Backtrack(value.Nums)
		timeLapse = time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

		fmt.Println("Solution 4: adpapt code from leetcode ")
		timeStart = time.Now()
		result = trap_Adapt1(value.Nums)
		timeLapse = time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

		fmt.Println("Solution 5: DFS ")
		timeStart = time.Now()
		result = trap_DFS(value.Nums)
		timeLapse = time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

		fmt.Println("Solution 6: BruteForce ")
		timeStart = time.Now()
		result = trap_BruteForce(value.Nums)
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
	Nums   []int
	Result string
}

/*


===============
Test count  0 for node {[1 2 0]
    3
            }
Solution 1: DP
>Solution result 0
Correct result is
    3

TimeLapse 944ns
Solution 2: 2 pointers
>Solution result 0
Correct result is
    3

TimeLapse 481ns
Solution 3: Backtrack
>Solution result 0
Correct result is
    3

TimeLapse 426ns
Solution 4: adpapt code from leetcode
>Solution result 0
Correct result is
    3

TimeLapse 407ns
Solution 5: DFS
>Solution result 0
Correct result is
    3

TimeLapse 944ns
Solution 6: BruteForce
>Solution result 0
Correct result is
    3

TimeLapse 351ns
===============
Test count  1 for node {[3 4 -1 1]
    2
            }
Solution 1: DP
>Solution result 2
Correct result is
    2

TimeLapse 481ns
Solution 2: 2 pointers
>Solution result 2
Correct result is
    2

TimeLapse 203ns
Solution 3: Backtrack
>Solution result 2
Correct result is
    2

TimeLapse 241ns
Solution 4: adpapt code from leetcode
>Solution result 2
Correct result is
    2

TimeLapse 185ns
Solution 5: DFS
>Solution result 2
Correct result is
    2

TimeLapse 259ns
Solution 6: BruteForce
>Solution result 2
Correct result is
    2

TimeLapse 222ns
===============
Test count  2 for node {[7 8 9 11 12]
    1
            }
Solution 1: DP
>Solution result 0
Correct result is
    1

TimeLapse 500ns
Solution 2: 2 pointers
>Solution result 0
Correct result is
    1

TimeLapse 149ns
Solution 3: Backtrack
>Solution result 0
Correct result is
    1

TimeLapse 277ns
Solution 4: adpapt code from leetcode
>Solution result 0
Correct result is
    1

TimeLapse 148ns
Solution 5: DFS
>Solution result 0
Correct result is
    1

TimeLapse 334ns
Solution 6: BruteForce
>Solution result 0
Correct result is
    1

TimeLapse 259ns
===============
TimeLapse Whole Program 846.312µs

*/
