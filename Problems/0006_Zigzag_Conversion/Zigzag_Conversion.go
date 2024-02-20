package main

import (
	"fmt"
	"strings"
	"time"
)

// Time Complexity: O(n)
// Space Complexity: O(n)
// linear scan
func convert_LinearScan(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	var rows []strings.Builder
	for i := 0; i < numRows; i++ {
		rows = append(rows, strings.Builder{})
	}

	direction := 1 // Direction of movement, 1 for down, -1 for up
	currentRow := 0

	for _, char := range s {
		rows[currentRow].WriteRune(char)
		if currentRow == 0 {
			direction = 1
		} else if currentRow == numRows-1 {
			direction = -1
		}
		currentRow += direction
	}

	var result strings.Builder
	for _, row := range rows {
		result.WriteString(row.String())
	}

	return result.String()
}

// ======= Brute Force
// Time Complexity: O(n)
// Space Complexity: O(numRows * n)
func convert_BruteForce(s string, numRows int) string {
	if numRows == 1 || numRows >= len(s) {
		return s
	}

	// Create numRows number of slices to store characters row by row
	rows := make([][]rune, numRows)
	for i := range rows {
		rows[i] = make([]rune, 0)
	}

	// direction represents the movement of rows, 1 for down, -1 for up
	direction := 1
	// row indicates the current row
	row := 0

	for _, char := range s {
		rows[row] = append(rows[row], char)

		// Change direction if reaching the top or bottom row
		if row == 0 {
			direction = 1
		} else if row == numRows-1 {
			direction = -1
		}

		// Update the current row based on the direction
		row += direction
	}

	// Combine characters row by row to form the zigzag pattern
	result := ""
	for _, r := range rows {
		result += string(r)
	}

	return result
}

//======

func main() {
	timeStartWholeProgram := time.Now()
	testInput := []TestCase{
		{
			S: "PAYPALISHIRING",
			N: 3,
			Result: `
                PAHNAPLSIIGYIR
            `,
		},
		{
			S: "PAYPALISHIRING",
			N: 4,
			Result: `
                PINALSIGYAHRPI
            `,
		},
		{
			S: "A",
			N: 1,
			Result: `
                A
            `,
		},
	}
	for count, value := range testInput {
		fmt.Println("===============")
		fmt.Println("Test count ", count, "for node", value)
		fmt.Println("Solution 1: use Linear Scan")
		timeStart := time.Now()
		result := convert_LinearScan(value.S, value.N)
		timeLapse := time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

		fmt.Println("Solution 2: use Bture Force")
		timeStart = time.Now()
		result = convert_BruteForce(value.S, value.N)
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
	N      int
	Result string
}

/*

===============
Test count  0 for node {PAYPALISHIRING 3
                PAHNAPLSIIGYIR
            }
Solution 1: use Linear Scan
>Solution result PAHNAPLSIIGYIR
Correct result is
                PAHNAPLSIIGYIR

TimeLapse 15.741µs
Solution 2: use Bture Force
>Solution result PAHNAPLSIIGYIR
Correct result is
                PAHNAPLSIIGYIR

TimeLapse 6.074µs
===============
Test count  1 for node {PAYPALISHIRING 4
                PINALSIGYAHRPI
            }
Solution 1: use Linear Scan
>Solution result PINALSIGYAHRPI
Correct result is
                PINALSIGYAHRPI

TimeLapse 2.407µs
Solution 2: use Bture Force
>Solution result PINALSIGYAHRPI
Correct result is
                PINALSIGYAHRPI

TimeLapse 3.259µs
===============
Test count  2 for node {A 1
                A
            }
Solution 1: use Linear Scan
>Solution result A
Correct result is
                A

TimeLapse 112ns
Solution 2: use Bture Force
>Solution result A
Correct result is
                A

TimeLapse 111ns
===============
TimeLapse Whole Program 529.002µs


*/
