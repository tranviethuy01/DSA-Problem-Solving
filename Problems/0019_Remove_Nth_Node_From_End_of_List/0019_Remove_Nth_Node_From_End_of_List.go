package main

import (
	"fmt"

	"time"
)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

//Time Complexity: O(N)
//Space Complexity: O(1)

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Val: 0, Next: head}
	first := dummy
	second := dummy

	// Move the first pointer ahead by n+1 steps
	for i := 0; i <= n; i++ {
		first = first.Next
	}

	// Move both pointers until the first pointer reaches the end
	for first != nil {
		first = first.Next
		second = second.Next
	}

	// Remove the nth node from the end
	second.Next = second.Next.Next

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
			Head: ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}},
			N:    2,
			Result: `
    [1,2,3,5]
            `,
		},
		{
			Head: ListNode{Val: 1},
			N:    1,
			Result: `
         []
            `,
		},
		{
			Head: ListNode{Val: 1, Next: &ListNode{Val: 2}},
			N:    1,
			Result: `
         [1]
            `,
		},
	}
	for count, value := range testInput {
		fmt.Println("===============")
		fmt.Println("Test count ", count, "for node", value)
		fmt.Println("Solution 1: use 2 Pointers")
		timeStart := time.Now()
		result := removeNthFromEnd(&value.Head, value.N)
		timeLapse := time.Since(timeStart)
		fmt.Println(">Solution result", result)
		printLinkedList(result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)
	}

	timeLapsedWholeProgram := time.Since(timeStartWholeProgram)
	fmt.Println("===============")
	fmt.Println("TimeLapse Whole Program", timeLapsedWholeProgram)
}

type TestCase struct {
	Head   ListNode
	N      int
	Result string
}

/*

===============
Test count  0 for node {{1 0x400009e230} 2
    [1,2,3,5]
            }
Solution 1: use 2 Pointers
>Solution result &{1 0x400009e230}
1 -> 2 -> 3 -> 5 -> nil
Correct result is
    [1,2,3,5]

TimeLapse 704ns
===============
Test count  1 for node {{1 <nil>} 1
         []
            }
Solution 1: use 2 Pointers
>Solution result <nil>
nil
Correct result is
         []

TimeLapse 185ns
===============
Test count  2 for node {{1 0x400009e270} 1
         [1]
            }
Solution 1: use 2 Pointers
>Solution result &{1 <nil>}
1 -> nil
Correct result is
         [1]

TimeLapse 167ns
===============
TimeLapse Whole Program 458.754µs


*/
//REF
//
