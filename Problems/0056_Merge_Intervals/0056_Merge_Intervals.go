package main

import (
	"fmt"
	"runtime"
	"sort"
	"time"
)

// approach : StraightForward
/*
Time Complexity:

Sorting the intervals takes O(n log n) time, where n is the number of intervals.
Iterating through the sorted intervals takes O(n) time.
So, the overall time complexity is O(n log n), dominated by the sorting step.
Space Complexity:

We use additional space to store the sorted intervals and the merged intervals.
The space complexity is O(n) because we might need to store all the intervals in the merged array.
*/
func merge_StraightForward(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	// Sort intervals based on the start value
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	fmt.Println("intervals after sort", intervals)
	merged := [][]int{intervals[0]}

	for i := 1; i < len(intervals); i++ {
		last := &merged[len(merged)-1]
		fmt.Println("i", i, "len", len(intervals), "last", last)
		if intervals[i][0] <= (*last)[1] {
			fmt.Println("overlapping, need merge, i ", i, "intervals[i]", intervals[i], "last", last)
			// Overlapping intervals, merge
			if intervals[i][1] > (*last)[1] {
				(*last)[1] = intervals[i][1]
			}
		} else {
			fmt.Println(" non overlapping, add to result, i ", i, "intervals[i]", intervals[i])
			// Non-overlapping intervals, add to result
			merged = append(merged, intervals[i])
		}
	}

	return merged
}

//approach : adpat a solution from leetcode
/*
class Solution {
public:
    vector<vector<int>> merge(vector<vector<int>>& intervals) {
        int n = intervals.size(); // Size of the array
        vector<vector<int>>ans;

        // sort the given intervals
        sort(intervals.begin(),intervals.end());
        for(int i=0;i<n;i++)
        {
            //if the current interval does not exist in the last interval
            if(ans.empty() || intervals[i][0]> ans.back()[1])
            {
                ans.push_back(intervals[i]);
            }
            else
            {
                //if the current interval lies in the last interval
                ans.back()[1] = max(ans.back()[1] , intervals[i][1]);
            }
        }
        return ans;
    }
};

// If there's any other optimal approach then do let me know :)

*/

func merge_Adapt2(intervals [][]int) [][]int {
	n := len(intervals)
	ans := make([][]int, 0)

	// Sort the given intervals
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	for i := 0; i < n; i++ {
		// If the current interval does not exist in the last interval
		if len(ans) == 0 || intervals[i][0] > ans[len(ans)-1][1] {
			ans = append(ans, intervals[i])
		} else {
			// If the current interval lies in the last interval
			ans[len(ans)-1][1] = max(ans[len(ans)-1][1], intervals[i][1])
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	timeStartWholeProgram := time.Now()

	testInput := []TestCase{
		{
			Intervals: [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
			Result: `
[[1,6],[8,10],[15,18]]

            `,
		},
		{
			Intervals: [][]int{{1, 4}, {4, 5}},
			Result: `
[[1,5]]

            `,
		},
		{
			Intervals: [][]int{{15, 18}, {1, 3}, {2, 6}, {8, 10}},
			Result: `
[[1,6],[8,10],[15,18]]

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
		result := merge_StraightForward(value.Intervals)
		timeLapse := time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

		// Memory after allocation
		runtime.ReadMemStats(&m)
		memAfter := m.Alloc
		fmt.Println("Memory before", memBefore, "bytes", "Memory after", memAfter, "bytes", "Memory used:", memAfter-memBefore, "bytes")

		fmt.Printf("Memory usage (HeapAlloc) after Test Case i %d, : %v bytes\n", count, m.HeapAlloc)

		fmt.Println("Solution 2: another solution from leetcode ")
		timeStart = time.Now()
		result = merge_Adapt2(value.Intervals)
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
	Intervals [][]int
	Result    string
}

/*


===============
Test count  0 for node {[[1 3] [2 6] [8 10] [15 18]]
[[1,6],[8,10],[15,18]]

            }
Solution 1: StraightForward
intervals after sort [[1 3] [2 6] [8 10] [15 18]]
i 1 len 4 last &[1 3]
overlapping, need merge, i  1 intervals[i] [2 6] last &[1 3]
i 2 len 4 last &[1 6]
 non overlapping, add to result, i  2 intervals[i] [8 10]
i 3 len 4 last &[8 10]
 non overlapping, add to result, i  3 intervals[i] [15 18]
>Solution result [[1 6] [8 10] [15 18]]
Correct result is
[[1,6],[8,10],[15,18]]


TimeLapse 97.44µs
Memory before 69648 bytes Memory after 71728 bytes Memory used: 2080 bytes
Memory usage (HeapAlloc) after Test Case i 0, : 71728 bytes
Solution 2: another solution from leetcode
>Solution result [[1 6] [8 10] [15 18]]
Correct result is
[[1,6],[8,10],[15,18]]


TimeLapse 4.481µs
Memory before 69648 bytes Memory after 72200 bytes Memory used: 2552 bytes
Memory usage (HeapAlloc) after Test Case i 0, : 72200 bytes
===============
Test count  1 for node {[[1 4] [4 5]]
[[1,5]]

            }
Solution 1: StraightForward
intervals after sort [[1 4] [4 5]]
i 1 len 2 last &[1 4]
overlapping, need merge, i  1 intervals[i] [4 5] last &[1 4]
>Solution result [[1 5]]
Correct result is
[[1,5]]


TimeLapse 80.126µs
Memory before 69648 bytes Memory after 72816 bytes Memory used: 3168 bytes
Memory usage (HeapAlloc) after Test Case i 1, : 72816 bytes
Solution 2: another solution from leetcode
>Solution result [[1 5]]
Correct result is
[[1,5]]


TimeLapse 2.926µs
Memory before 69648 bytes Memory after 73064 bytes Memory used: 3416 bytes
Memory usage (HeapAlloc) after Test Case i 1, : 73064 bytes
===============
Test count  2 for node {[[15 18] [1 3] [2 6] [8 10]]
[[1,6],[8,10],[15,18]]

            }
Solution 1: StraightForward
intervals after sort [[1 3] [2 6] [8 10] [15 18]]
i 1 len 4 last &[1 3]
overlapping, need merge, i  1 intervals[i] [2 6] last &[1 3]
i 2 len 4 last &[1 6]
 non overlapping, add to result, i  2 intervals[i] [8 10]
i 3 len 4 last &[8 10]
 non overlapping, add to result, i  3 intervals[i] [15 18]
>Solution result [[1 6] [8 10] [15 18]]
Correct result is
[[1,6],[8,10],[15,18]]


TimeLapse 88.274µs
Memory before 69648 bytes Memory after 74224 bytes Memory used: 4576 bytes
Memory usage (HeapAlloc) after Test Case i 2, : 74224 bytes
Solution 2: another solution from leetcode
>Solution result [[1 6] [8 10] [15 18]]
Correct result is
[[1,6],[8,10],[15,18]]


TimeLapse 3.704µs
Memory before 69648 bytes Memory after 74696 bytes Memory used: 5048 bytes
Memory usage (HeapAlloc) after Test Case i 2, : 74696 bytes
===============
TimeLapse Whole Program 1.307941ms


*/
//REF
//
