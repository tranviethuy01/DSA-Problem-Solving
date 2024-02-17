package main

import 
(
    "fmt"
    "time"
)

func twoSum(nums []int, target int) []int {
    numIndices := make(map[int]int)
    for i, num := range nums {
        complement := target - num
        if j, ok := numIndices[complement]; ok {
            return []int{j, i}
        }
        numIndices[num] = i
    }
    return nil
}



func main() {
    timeStartWholeProgram := time.Now()
    testInput := []TestCase{
        {
            Target: 9,
            Nums: []int{2, 7, 11, 15},
            Result: `
                [0,1]
            `,
        },

        {
            Target: 9,
            Nums: []int{15, 2, 7, 11},
            Result: `
                [1,2]
            `,
        },

        {
            Target: 6,
            Nums: []int {3,2,4},
            Result: `
                [1,2]
            `,
        },
        {
            Target: 6,
            Nums: []int{3,3},
            Result: `
                [0,1]
            `,
        },
    } 
    for count, value := range testInput {
        fmt.Println("===============")
        fmt.Println("Solution 1:  Test count ", count, "for node", value)
        timeStart := time.Now()
        result := twoSum(value.Nums , value.Target)
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
    Target       int
    Nums []int
    Result string 
}

/*
Note:
Time complexity: O(n)
Space complexity: O(n)



Result:
===============
Solution 1:  Test count  0 for node {9 [2 7 11 15] 
                [0,1]
            }
>Solution result [0 1]
Correct result is  
                [0,1]
            
TimeLapse 3.111µs
===============
Solution 1:  Test count  1 for node {9 [15 2 7 11] 
                [1,2]
            }
>Solution result [1 2]
Correct result is  
                [1,2]
            
TimeLapse 1.13µs
===============
Solution 1:  Test count  2 for node {6 [3 2 4] 
                [1,2]
            }
>Solution result [1 2]
Correct result is  
                [1,2]
            
TimeLapse 945ns
===============
Solution 1:  Test count  3 for node {6 [3 3] 
                [0,1]
            }
>Solution result [0 1]
Correct result is  
                [0,1]
            
TimeLapse 759ns
===============
TimeLapse Whole Program 474.586µs

*/
