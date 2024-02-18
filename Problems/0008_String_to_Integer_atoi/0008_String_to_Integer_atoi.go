package main

import (
	"fmt"
	"math"
	"time"
)

//NOTE: this solution is failed, need check
//
//Time Complexity: O(n)
//Space Complexity: O(1)

func myAtoi(s string) int {
    var sign, result int
    i := 0

    // ignore leading whitespace
    for i < len(s) && s[i] == ' ' {
        i++
    }

    // check sign
    if i < len(s) && (s[i] == '-' || s[i] == '+') {
        if s[i] == '-' {
            sign = -1
        }
        i++
    }

    // convert digits
    for i < len(s) && s[i] >= '0' && s[i] <= '9' {
        result = result*10 + int(s[i]-'0')

        // handle overflow
        if sign == 1 && result > (1<<31-1) {
            return 1<<31 - 1
        } else if sign == -1 && -result < -(1<<31) {
            return -(1 << 31)
        }

        i++
    }

    return result * sign
}

//===== approach Linear Scan
//Time Complexity: O(n)
//Space Complexity: O(1)
func myAtoi_LinearScan(s string) int {
    sign, ans := 1, 0
    n := len(s)
    i := 0

    // Ignore leading spaces
    for i < n && s[i] == ' ' {
        i++
    }

    // Handle the sign
    if i < n && (s[i] == '-' || s[i] == '+') {
        if s[i] == '+' {
            sign = 1
        } else {
            sign = -1
        }
        i++
    }

    // Handle the case where the first non-space character is '0'
    if i < n && s[i] == '0' {
        for i < n && s[i] == '0' {
            i++
        }

        // Check if there's a valid number after leading zeros
        if i < n && s[i] >= '1' && s[i] <= '9' {
            for i < n && s[i] >= '0' && s[i] <= '9' {
                digit := int(s[i] - '0')

                // Check for overflow before updating ans
                if ans > math.MaxInt32/10 || (ans == math.MaxInt32/10 && digit > 7) {
                    if sign == 1 {
                        return math.MaxInt32
                    } else {
                        return math.MinInt32
                    }
                }

                ans = ans*10 + digit
                i++
            }
            return ans * sign
        } else {
            return 0 // Invalid input
        }
    }

    // Handle the general case of non-zero digits
    if i < n && s[i] >= '1' && s[i] <= '9' {
        for i < n && s[i] >= '0' && s[i] <= '9' {
            digit := int(s[i] - '0')

            // Check for overflow before updating ans
            if ans > math.MaxInt32/10 || (ans == math.MaxInt32/10 && digit > 7) {
                if sign == 1 {
                    return math.MaxInt32
                } else {
                    return math.MinInt32
                }
            }

            ans = ans*10 + digit
            i++
        }
        return ans * sign
    } else {
        return 0 // Invalid input
    }
}




//=====


//===== brute force approach
//Time Complexity: O(n)
//Space Complexity: O(1)
func myAtoi_BruteForce(s string) int {
    var result int
    sign := 1
    started := false

    for _, ch := range s {
        if ch == ' ' && !started {
            continue
        } else if ch == '-' && !started {
            sign = -1
            started = true
        } else if ch == '+' && !started {
            started = true
        } else if ch >= '0' && ch <= '9' {
            started = true
            digit := int(ch - '0')
            result = result*10 + digit
            if result*sign > math.MaxInt32 {
                return math.MaxInt32
            } else if result*sign < math.MinInt32 {
                return math.MinInt32
            }
        } else {
            break
        }
    }

    return result * sign
}


//=====

func main() {
	timeStartWholeProgram := time.Now()
	testInput := []TestCase{
		{
			S: "42",
			Result: `
          42
            `,
		},
		{
			S: "   -42",
			Result: `
             -42
             `,
		},
		{
			S:  "4193 with words",
			Result: `
			4193


            `,
		},

	}
	for count, value := range testInput {
		fmt.Println("===============")
		fmt.Println("Test count ", count, "for node", value)
		fmt.Println("Solution 1: use straightforward iterative approach")
		timeStart := time.Now()
		result := myAtoi(value.S)
		timeLapse := time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

		fmt.Println("Solution 2: use Bture Force")
		timeStart = time.Now()
		result = myAtoi_BruteForce(value.S)
		timeLapse = time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

		fmt.Println("Solution 3: Linear Scan")
		timeStart = time.Now()
		result = myAtoi_LinearScan(value.S)
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
	Result string
}

/*

===============
Test count  0 for node {42 
          42
            }
Solution 1: use straightforward iterative approach
>Solution result 0
Correct result is  
          42
            
TimeLapse 648ns
Solution 2: use Bture Force
>Solution result 42
Correct result is  
          42
            
TimeLapse 556ns
Solution 3: Linear Scan
>Solution result 42
Correct result is  
          42
            
TimeLapse 666ns
===============
Test count  1 for node {   -42 
             -42
             }
Solution 1: use straightforward iterative approach
>Solution result -42
Correct result is  
             -42
             
TimeLapse 259ns
Solution 2: use Bture Force
>Solution result -42
Correct result is  
             -42
             
TimeLapse 222ns
Solution 3: Linear Scan
>Solution result -42
Correct result is  
             -42
             
TimeLapse 185ns
===============
Test count  2 for node {4193 with words 
			4193


            }
Solution 1: use straightforward iterative approach
>Solution result 0
Correct result is  
			4193


            
TimeLapse 241ns
Solution 2: use Bture Force
>Solution result 4193
Correct result is  
			4193


            
TimeLapse 203ns
Solution 3: Linear Scan
>Solution result 4193
Correct result is  
			4193


            
TimeLapse 204ns
===============
TimeLapse Whole Program 721.343µs


*/
