package main

import (
	"fmt"
	"time"
)

// Time complexity: O(13 * log(num)), where 13 is the number of symbols/values used in Roman numerals, and log(num) represents the number of times the loop iterates based on the size of the input number num. Since the maximum value of num is 3999, the number of iterations will not exceed log(3999) which is constant.
// Space Complexity: O(1)
func intToRoman_StraightForward(num int) string {
	// Define the symbols and their values
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

	result := ""

	// Iterate through symbols and values
	for i := 0; i < len(values); i++ {
		// Repeat the symbol while num is greater or equal to its corresponding value
		for num >= values[i] {
			num -= values[i]
			result += symbols[i]
		}
	}

	return result
}

//==== approach : copied from here https://leetcode.com/problems/integer-to-roman/solutions/2962674/easiest-o-1-faang-method-ever/
//Time complexity: O(1)
//Space Complexity: O(1)

func intToRoman_Faang(num int) string {
	ones := []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
	tens := []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	hrns := []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	ths := []string{"", "M", "MM", "MMM"}

	return ths[num/1000] + hrns[(num%1000)/100] + tens[(num%100)/10] + ones[num%10]
}

//====

// ===== approach : brute force => failed code
// the code need check again
func intToRoman_BruteForce(num int) string {
	romanMap := map[int]string{
		1:    "I",
		4:    "IV",
		5:    "V",
		9:    "IX",
		10:   "X",
		40:   "XL",
		50:   "L",
		90:   "XC",
		100:  "C",
		400:  "CD",
		500:  "D",
		900:  "CM",
		1000: "M",
	}

	result := ""

	for i := 1000; i >= 1; i-- {
		for num >= i {
			result += romanMap[i]
			num -= i
		}
	}

	return result
}

//=====

// ===== approach straight forward with some optimize
// NOTE: wrong answer, need double check the code again
// this solution is incorrect because the map is unordered, so the loop give the wrong order, then result in the wrong answer
func intToRoman_StraightForward_Optimized(num int) string {
	fmt.Println("NOTE: this is a wrong answer")
	symbols := map[int]string{
		1000: "M",
		900:  "CM",
		500:  "D",
		400:  "CD",
		100:  "C",
		90:   "XC",
		50:   "L",
		40:   "XL",
		10:   "X",
		9:    "IX",
		5:    "V",
		4:    "IV",
		1:    "I",
	}

	result := ""

	for value, symbol := range symbols {
		for num >= value {
			result += symbol
			num -= value
		}
	}

	return result
}

//=====

// ==== approach Straightforward with optimize code
func intToRoman_StraightForward_Optimized_UseInterface(num int) string {
	Roman := ""
	storeIntRoman := [][]interface{}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}
	for i := 0; i < len(storeIntRoman); i++ {
		for num >= storeIntRoman[i][0].(int) {
			Roman += storeIntRoman[i][1].(string)
			num -= storeIntRoman[i][0].(int)
		}
	}
	return Roman
}

//====

//==== approach Straightforward with optimize code for better understanding

func intToRoman_StraightForward_Optimized_UseSliceStruct(num int) string {
	symbols := []RomanSymbol{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	result := ""

	for _, rs := range symbols {
		for num >= rs.value {
			result += rs.symbol
			num -= rs.value
		}
	}

	return result
}

type RomanSymbol struct {
	value  int
	symbol string
}

//=====

func main() {
	timeStartWholeProgram := time.Now()
	testInput := []TestCase{
		{
			N: 3,
			Result: `
          III
            `,
		},
		{

			N: 58,
			Result: `
           LVIII 
             `,
		},
		{

			N: 1994,
			Result: `
           MCMXCIV 
             `,
		},
	}
	for count, value := range testInput {
		fmt.Println("===============")
		fmt.Println("Test count ", count, "for node", value)
		fmt.Println("Solution 1: use Straightforward approach")
		timeStart := time.Now()
		result := intToRoman_StraightForward(value.N)
		timeLapse := time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

		fmt.Println("Solution 2: use Straightforward with optimize code with map => this is a failed solution")
		timeStart = time.Now()
		result = intToRoman_StraightForward_Optimized(value.N)
		timeLapse = time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

		fmt.Println("Solution 3: use Straightforward with optimize code ")
		timeStart = time.Now()
		result = intToRoman_StraightForward_Optimized_UseInterface(value.N)
		timeLapse = time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

		fmt.Println("Solution 4: use Straightforward with optimize code ")
		timeStart = time.Now()
		result = intToRoman_StraightForward_Optimized_UseSliceStruct(value.N)
		timeLapse = time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

		fmt.Println("Solution 5: use faang method")

		timeStart = time.Now()
		result = intToRoman_Faang(value.N)
		timeLapse = time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

		fmt.Println("Solution 6: use brute force")
		timeStart = time.Now()
		result = intToRoman_BruteForce(value.N)
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
	Result string
}

/*


===============
Test count  0 for node {3
          III
            }
Solution 1: use Straightforward approach
>Solution result III
Correct result is
          III

TimeLapse 7.963µs
Solution 2: use Straightforward with optimize code with map => this is a failed solution
NOTE: this is a wrong answer
>Solution result III
Correct result is
          III

TimeLapse 18.963µs
Solution 3: use Straightforward with optimize code
>Solution result III
Correct result is
          III

TimeLapse 2.056µs
Solution 4: use Straightforward with optimize code
>Solution result III
Correct result is
          III

TimeLapse 1.518µs
Solution 5: use faang method
>Solution result III
Correct result is
          III

TimeLapse 2.074µs
Solution 6: use brute force
>Solution result
Correct result is
          III

TimeLapse 6.982µs
===============
Test count  1 for node {58
           LVIII
             }
Solution 1: use Straightforward approach
>Solution result LVIII
Correct result is
           LVIII

TimeLapse 1.241µs
Solution 2: use Straightforward with optimize code with map => this is a failed solution
NOTE: this is a wrong answer
>Solution result IIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIII
Correct result is
           LVIII

TimeLapse 20.759µs
Solution 3: use Straightforward with optimize code
>Solution result LVIII
Correct result is
           LVIII

TimeLapse 1.296µs
Solution 4: use Straightforward with optimize code
>Solution result LVIII
Correct result is
           LVIII

TimeLapse 1.203µs
Solution 5: use faang method
>Solution result LVIII
Correct result is
           LVIII

TimeLapse 704ns
Solution 6: use brute force
>Solution result
Correct result is
           LVIII

TimeLapse 5.315µs
===============
Test count  2 for node {1994
           MCMXCIV
             }
Solution 1: use Straightforward approach
>Solution result MCMXCIV
Correct result is
           MCMXCIV

TimeLapse 1.149µs
Solution 2: use Straightforward with optimize code with map => this is a failed solution
NOTE: this is a wrong answer
>Solution result CCCCCCCCCCCCCCCCCCCXCIIII
Correct result is
           MCMXCIV

TimeLapse 15µs
Solution 3: use Straightforward with optimize code
>Solution result MCMXCIV
Correct result is
           MCMXCIV

TimeLapse 1.463µs
Solution 4: use Straightforward with optimize code
>Solution result MCMXCIV
Correct result is
           MCMXCIV

TimeLapse 982ns
Solution 5: use faang method
>Solution result MCMXCIV
Correct result is
           MCMXCIV

TimeLapse 759ns
Solution 6: use brute force
>Solution result M
Correct result is
           MCMXCIV

TimeLapse 4.574µs
===============
TimeLapse Whole Program 1.087478ms

*/
//REF
//
