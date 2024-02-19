package main

import (
	"fmt"
	"time"
	"sort"
	"strings"
)

//===== approach verital scanning
//Time Complexity: O(nxm)
//Space Complexity: O(1)
func longestCommonPrefix_VerticalScanning(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		for !strings.HasPrefix(strs[i], prefix) {
			prefix = prefix[:len(prefix)-1]
			if prefix == "" {
				return ""
			}
		}
	}
	return prefix
}


//=====

//===== approach : sort first
//Time Complexity: O(mlogm + n)
//Space Complexity: O(1)
func longestCommonPrefix_WithSort(v []string) string {
	if len(v) == 0 {
		return ""
	}

	sort.Strings(v)
	first := v[0]
	last := v[len(v)-1]

	var ans string
	for i := 0; i < min(len(first), len(last)); i++ {
		if first[i] != last[i] {
			return ans
		}
		ans += string(first[i])
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}



//=====
func main() {
	timeStartWholeProgram := time.Now()
	testInput := []TestCase{
		{
			S: []string{"flower", "flow", "flight"},
			Result: `
          "fl"
            `,
		},
		{

			S: []string{"dog", "racecar", "car"},
			Result: `
           "" 
             `,
		},

	}
	for count, value := range testInput {
		fmt.Println("===============")
		fmt.Println("Test count ", count, "for node", value)
		fmt.Println("Solution 1: use Straightforward approach with map[string]int")
		timeStart := time.Now()
		result := longestCommonPrefix_VerticalScanning(value.S)
		timeLapse := time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

		fmt.Println("Solution 2: use Straightforward map[byte]int")
		timeStart = time.Now()
		result = longestCommonPrefix_WithSort(value.S)
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
	S      []string
	Result string
}

/*



===============
Test count  0 for node {[flower flow flight] 
          "fl"
            }
Solution 1: use Straightforward approach with map[string]int
>Solution result fl
Correct result is  
          "fl"
            
TimeLapse 1.056µs
Solution 2: use Straightforward map[byte]int
>Solution result fl
Correct result is  
          "fl"
            
TimeLapse 6.055µs
===============
Test count  1 for node {[dog racecar car] 
           "" 
             }
Solution 1: use Straightforward approach with map[string]int
>Solution result 
Correct result is  
           "" 
             
TimeLapse 222ns
Solution 2: use Straightforward map[byte]int
>Solution result 
Correct result is  
           "" 
             
TimeLapse 1.074µs
===============
TimeLapse Whole Program 408.809µs
*/
//REF
//
