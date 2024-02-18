package main

import (
	"fmt"
	"time"
)
//approach: dynamic programing
//Time Complexity: O(n * m)
//Space Complexity: O(n * m)

func isMatch_DP(s string, p string) bool {
	memo := make(map[string]bool)
	var dp func(i, j int) bool
	dp = func(i, j int) bool {
		if j == len(p) {
			return i == len(s)
		}
		if i == len(s) {
			if (len(p)-j)%2 == 1 {
				return false
			}
			for ; j+1 < len(p); j += 2 {
				if p[j+1] != '*' {
					return false
				}
			}
			return true
		}
		key := fmt.Sprintf("%d-%d", i, j)
		if val, ok := memo[key]; ok {
			return val
		}
		var match bool
		if s[i] == p[j] || p[j] == '.' {
			if j < len(p)-1 && p[j+1] == '*' {
				match = dp(i, j+2) || dp(i+1, j)
			} else {
				match = dp(i+1, j+1)
			}
		} else {
			if j < len(p)-1 && p[j+1] == '*' {
				match = dp(i, j+2)
			} else {
				match = false
			}
		}
		memo[key] = match
		return match
	}
	return dp(0, 0)
}


//===== approach: brute force 
//Time Complexity: Exponential, O(2^(n+m))
//Space Complexity: O(n + m)

func isMatch_BruteForce(s string, p string) bool {
    var match func(sIndex, pIndex int) bool
    match = func(sIndex, pIndex int) bool {
        if pIndex == len(p) {
            return sIndex == len(s)
        }
        if sIndex == len(s) {
            if (len(p)-pIndex)%2 == 1 {
                return false
            }
            for i := pIndex + 1; i < len(p); i += 2 {
                if p[i] != '*' {
                    return false
                }
            }
            return true
        }
        if pIndex+1 < len(p) && p[pIndex+1] == '*' {
            return (s[sIndex] == p[pIndex] || p[pIndex] == '.') && (match(sIndex+1, pIndex) || match(sIndex, pIndex+2))
        } else {
            return (s[sIndex] == p[pIndex] || p[pIndex] == '.') && match(sIndex+1, pIndex+1)
        }
    }
    return match(0, 0)
}



//=====

//===== approach recursive
//Time Complexity: Exponential, O(2^(n+m))
//Space Complexity: O(n + m)
func isMatch_Recursive(s string, p string) bool {
    if p == "" {
        return s == ""
    }

    firstMatch := s != "" && (p[0] == '.' || s[0] == p[0])

    if len(p) >= 2 && p[1] == '*' {
        return isMatch_Recursive(s, p[2:]) || (firstMatch && isMatch_Recursive(s[1:], p))
    }

    return firstMatch && isMatch_Recursive(s[1:], p[1:])
}


//=====


func main() {
	timeStartWholeProgram := time.Now()
	testInput := []TestCase{
		{
      S: "aa",
      P: "a",
			Result: `
          false
            `,
		},
		{
      S: "aa",
      P: "a*",
			Result: `
             true
             `,
		},
		{
		  S: "ab",
      P:".*",
			Result: `
			true
            `,
		},
		
	}
	for count, value := range testInput {
		fmt.Println("===============")
		fmt.Println("Test count ", count, "for node", value)
		fmt.Println("Solution 1: use")
		timeStart := time.Now()
		result := isMatch_DP(value.S, value.P)
		timeLapse := time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

		fmt.Println("Solution 2: use brute force")
		timeStart = time.Now()
		result = isMatch_BruteForce(value.S, value.P)
		timeLapse = time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

		fmt.Println("Solution 3: use recursive")
		timeStart = time.Now()
		result = isMatch_Recursive(value.S, value.P)
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
	P      string
	Result string
}

/*

===============
Test count  0 for node {aa a 
          false
            }
Solution 1: use
>Solution result false
Correct result is  
          false
            
TimeLapse 6.852µs
Solution 2: use brute force
>Solution result false
Correct result is  
          false
            
TimeLapse 1.166µs
Solution 3: use recursive
>Solution result false
Correct result is  
          false
            
TimeLapse 537ns
===============
Test count  1 for node {aa a* 
             true
             }
Solution 1: use
>Solution result true
Correct result is  
             true
             
TimeLapse 2.389µs
Solution 2: use brute force
>Solution result true
Correct result is  
             true
             
TimeLapse 389ns
Solution 3: use recursive
>Solution result true
Correct result is  
             true
             
TimeLapse 278ns
===============
Test count  2 for node {ab .* 
			true
            }
Solution 1: use
>Solution result true
Correct result is  
			true
            
TimeLapse 2.259µs
Solution 2: use brute force
>Solution result true
Correct result is  
			true
            
TimeLapse 240ns
Solution 3: use recursive
>Solution result true
Correct result is  
			true
            
TimeLapse 278ns
===============
TimeLapse Whole Program 589.625µs

*/
//REF
//https://www.geeksforgeeks.org/implementing-regular-expression-matching/
