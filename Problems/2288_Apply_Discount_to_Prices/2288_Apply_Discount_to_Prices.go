package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

/*

Algorithm:

Split the input sentence into words by spaces.
Iterate through each word:
a. Check if the word starts with '$'.
b. If it does, extract the price value, apply the discount, and update the word.
Join the updated words back into a sentence.
Time Complexity:
Let's denote:

n as the length of the input sentence.
m as the average length of a word in the sentence.
Splitting the sentence into words takes O(n) time.
Iterating through each word takes O(m) time, and applying the discount involves parsing the price, which takes O(1) time per word.
Joining the words back into a sentence takes O(n) time.

Therefore, the overall time complexity is O(n + m), where n is the length of the input sentence and m is the average length of a word.

Space Complexity:
The space complexity primarily depends on the space required to store the list of words after splitting the sentence and the space required for the updated sentence.

O(n) space is required to store the list of words after splitting the sentence.
O(n) space is required to store the updated sentence.
Therefore, the overall space complexity is O(n).

*/

func discountPrices(sentence string, discount int) string {
	words := strings.Split(sentence, " ")
	for i, word := range words {
		if strings.HasPrefix(word, "$") {
			priceStr := word[1:]
			price, err := strconv.ParseFloat(priceStr, 64)
			if err == nil {
				discountedPrice := price * (1 - float64(discount)/100)
				// Format the discounted price to have exactly two decimal places
				discountedPriceStr := fmt.Sprintf("$%.2f", discountedPrice)
				words[i] = discountedPriceStr
			}
		}
	}
	return strings.Join(words, " ")
}

func main() {
	timeStartWholeProgram := time.Now()

	testInput := []TestCase{

		{
			Sentence: "there are $1 $2 and 5$ candies in the shop",
			Discount: 50,
			Result: `
      "there are $0.50 $1.00 and 5$ candies in the shop"
            `,
		},
		{
			Sentence: "1 2 $3 4 $5 $6 7 8$ $9 $10$",
			Discount: 100,
			Result: `
"1 2 $0.00 4 $0.00 $0.00 7 8$ $0.00 $10$"

            `,
		},
	}
	for count, value := range testInput {
		fmt.Println("===============")
		fmt.Println("Test count ", count, "for node", value)
		fmt.Println("Solution 1: straightforward")
		timeStart := time.Now()
		result := discountPrices(value.Sentence, value.Discount)
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
	Sentence string
	Discount int
	Result   string
}

/*


===============
Test count  0 for node {there are $1 $2 and 5$ candies in the shop 50
      "there are $0.50 $1.00 and 5$ candies in the shop"
            }
Solution 1: straightforward
>Solution result there are $0.50 $1.00 and 5$ candies in the shop
Correct result is
      "there are $0.50 $1.00 and 5$ candies in the shop"

TimeLapse 36.166µs
===============
Test count  1 for node {1 2 $3 4 $5 $6 7 8$ $9 $10$ 100
"1 2 $0.00 4 $0.00 $0.00 7 8$ $0.00 $10$"

            }
Solution 1: straightforward
>Solution result 1 2 $0.00 4 $0.00 $0.00 7 8$ $0.00 $10$
Correct result is
"1 2 $0.00 4 $0.00 $0.00 7 8$ $0.00 $10$"


TimeLapse 7.722µs
===============
TimeLapse Whole Program 336.7µs

*/
//REF
//
