package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// approach manual multiplication algorithm
/*
Time Complexity:
The algorithm iterates through each digit of both input numbers, so it has a time complexity of O(n * m), where n is the length of num1 and m is the length of num2.
Within the nested loops, there are basic arithmetic operations and updates to the result array, which are constant time operations.
Therefore, the overall time complexity is O(n * m).
Space Complexity:
The space complexity primarily depends on the size of the result array, which is of length n + m.
Hence, the space complexity is O(n + m).
Additionally, there's a StringBuilder used to build the resulting string, but its space usage is proportional to the length of the output string, which is at most n + m.
Therefore, the dominant factor in terms of space complexity is still the result array.
*/
func multiply(num1 string, num2 string) string {
	n1, n2 := len(num1), len(num2)
	result := make([]int, n1+n2)

	for i := n1 - 1; i >= 0; i-- {
		for j := n2 - 1; j >= 0; j-- {
			mul := int(num1[i]-'0') * int(num2[j]-'0')
			p1, p2 := i+j, i+j+1
			sum := mul + result[p2]
			result[p1] += sum / 10
			result[p2] = sum % 10

			fmt.Println("i,j, num1[i], num2[j], num1[i]-'0', num2[j]-'0',  mul, p1, p2, sum, result[p1], result[p2]", i, j, num1[i], num2[j], int(num1[i]-'0'), int(num2[j]-'0'), mul, p1, p2, sum, result[p1], result[p2])
		}
	}
	fmt.Println("result slice", result)
	sb := strings.Builder{}
	for _, digit := range result {
		if !(len(sb.String()) == 0 && digit == 0) {
			sb.WriteString(strconv.Itoa(digit))
		}
	}
	if sb.Len() == 0 {
		return "0"
	}
	return sb.String()
}

func main() {
	timeStartWholeProgram := time.Now()

	testInput := []TestCase{

		{
			N: "2",
			K: "3",

			Result: `
"6"
            `,
		},
		{
			N: "123",
			K: "456",
			Result: `
"56088"
            `,
		},
	}
	for count, value := range testInput {
		fmt.Println("===============")
		fmt.Println("Test count ", count, "for node", value)
		fmt.Println("Solution 1: manual multiply")
		timeStart := time.Now()
		result := multiply(value.N, value.K)
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
	N      string
	K      string
	Result string
}

/*


===============
Test count  0 for node {2 3
"6"
            }
Solution 1: manual multiply
i,j, num1[i], num2[j], num1[i]-'0', num2[j]-'0',  mul, p1, p2, sum, result[p1], result[p2] 0 0 50 51 2 3 6 0 1 6 0 6
result slice [0 6]
>Solution result 6
Correct result is
"6"

TimeLapse 44.517µs
===============
Test count  1 for node {123 456
"56088"
            }
Solution 1: manual multiply
i,j, num1[i], num2[j], num1[i]-'0', num2[j]-'0',  mul, p1, p2, sum, result[p1], result[p2] 2 2 51 54 3 6 18 4 5 18 1 8
i,j, num1[i], num2[j], num1[i]-'0', num2[j]-'0',  mul, p1, p2, sum, result[p1], result[p2] 2 1 51 53 3 5 15 3 4 16 1 6
i,j, num1[i], num2[j], num1[i]-'0', num2[j]-'0',  mul, p1, p2, sum, result[p1], result[p2] 2 0 51 52 3 4 12 2 3 13 1 3
i,j, num1[i], num2[j], num1[i]-'0', num2[j]-'0',  mul, p1, p2, sum, result[p1], result[p2] 1 2 50 54 2 6 12 3 4 18 4 8
i,j, num1[i], num2[j], num1[i]-'0', num2[j]-'0',  mul, p1, p2, sum, result[p1], result[p2] 1 1 50 53 2 5 10 2 3 14 2 4
i,j, num1[i], num2[j], num1[i]-'0', num2[j]-'0',  mul, p1, p2, sum, result[p1], result[p2] 1 0 50 52 2 4 8 1 2 10 1 0
i,j, num1[i], num2[j], num1[i]-'0', num2[j]-'0',  mul, p1, p2, sum, result[p1], result[p2] 0 2 49 54 1 6 6 2 3 10 1 0
i,j, num1[i], num2[j], num1[i]-'0', num2[j]-'0',  mul, p1, p2, sum, result[p1], result[p2] 0 1 49 53 1 5 5 1 2 6 1 6
i,j, num1[i], num2[j], num1[i]-'0', num2[j]-'0',  mul, p1, p2, sum, result[p1], result[p2] 0 0 49 52 1 4 4 0 1 5 0 5
result slice [0 5 6 0 8 8]
>Solution result 56088
Correct result is
"56088"

TimeLapse 99.461µs
===============
TimeLapse Whole Program 430.27µs

*/
//REF
//
