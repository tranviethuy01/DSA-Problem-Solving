package main

import (
	"fmt"
	"time"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

//Time Complexity: O(n)
//Space Complexity: O(1)

func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	prev := dummy

	for prev.Next != nil && prev.Next.Next != nil {
		first := prev.Next
		second := prev.Next.Next

		// Swapping nodes
		first.Next = second.Next
		second.Next = first
		prev.Next = second

		// Move prev to next pair
		prev = first
	}

	return dummy.Next
}

func printLinkedList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d -> ", head.Val)
		head = head.Next
	}
	fmt.Println("nil")
}

func main() {
	timeStartWholeProgram := time.Now()

	testInput := []TestCase{

		{
			Head: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}},
			Result: `
[2,1,4,3]

            `,
		},
		{
			Head: &ListNode{},
			Result: `
         []
            `,
		},
		{
			Head: &ListNode{Val: 1},
			Result: `
         [1]
            `,
		},
	}
	for count, value := range testInput {
		fmt.Println("===============")
		fmt.Println("Test count ", count, "for node", value)
		fmt.Println("Solution 1: ")
		timeStart := time.Now()
		result := swapPairs(value.Head)
		timeLapse := time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("")
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)
	}

	timeLapsedWholeProgram := time.Since(timeStartWholeProgram)
	fmt.Println("===============")
	fmt.Println("TimeLapse Whole Program", timeLapsedWholeProgram)
}

type TestCase struct {
	Head   *ListNode
	Result string
}

/*


===============
Test count  0 for node {0x400009e230
[2,1,4,3]

            }
Solution 1:
>Solution result &{2 0x400009e230}

Correct result is
[2,1,4,3]


TimeLapse 963ns
===============
Test count  1 for node {0x400009e270
         []
            }
Solution 1:
>Solution result &{0 <nil>}

Correct result is
         []

TimeLapse 130ns
===============
Test count  2 for node {0x400009e280
         [1]
            }
Solution 1:
>Solution result &{1 <nil>}

Correct result is
         [1]

TimeLapse 241ns
===============
TimeLapse Whole Program 538.253µs
*/
//REF
//
