package main

import (
	"fmt"
	"time"
	//"sort"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

//approach: use stack base approach
//Time Complexity:
//
//Building the stacks takes O(n) time, where n is the number of nodes in the longer of the two input linked lists.
//Constructing the result linked list also takes O(n) time, as it requires iterating through the digits once.
//Therefore, the overall time complexity is O(n), where n is the maximum number of nodes in either of the input linked lists.
//Space Complexity:
//
//The algorithm uses extra space for the two stacks to store the digits in reverse order, each taking O(n) space.
//Additionally, the space required for the result linked list also grows linearly with the number of digits, taking O(n) space.
//Therefore, the overall space complexity is O(n), where n is the maximum number of nodes in either of the input linked lists.

// Time complexity: O(n)
// Space complexity: O(n)
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	stack1 := buildStack(l1)
	stack2 := buildStack(l2)
	fmt.Println("stack1", stack1)
	fmt.Println("stack2", stack2)
	dummyHead := &ListNode{}
	carry := 0

	for len(stack1) > 0 || len(stack2) > 0 || carry > 0 {
		sum := carry
		if len(stack1) > 0 {
			sum += stack1[len(stack1)-1]
			stack1 = stack1[:len(stack1)-1]
		}
		if len(stack2) > 0 {
			sum += stack2[len(stack2)-1]
			stack2 = stack2[:len(stack2)-1]
		}

		digit := sum % 10
		carry = sum / 10

		newNode := &ListNode{Val: digit, Next: dummyHead.Next}
		dummyHead.Next = newNode
	}

	return dummyHead.Next
}

func buildStack(head *ListNode) []int {
	stack := make([]int, 0)
	for head != nil {
		stack = append(stack, head.Val)
		head = head.Next
	}
	return stack
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
			L1: ListNode{Val: 7, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}},
			L2: ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}},
			Result: `
                [7,8,0,7]
            `,
		},
		{
			L1: ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}},
			L2: ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}},
			Result: `
                [8,0,7]
            `,
		},
		{
			L1: ListNode{Val: 0},
			L2: ListNode{Val: 0},
			Result: `
                [0]
            `,
		},

		{
			L1: ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9}}}}}}},
			L2: ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9}}}},
			Result: `
            [1, 0, 0 , 0 , 9 , 9 , 9 , 8]               
            `,
		},
	}
	for count, value := range testInput {
		fmt.Println("===============")
		fmt.Println("Test count ", count, "for node", value)
		printLinkedList(&value.L1)
		printLinkedList(&value.L2)
		fmt.Println("Solution 1: use stack base approach")
		timeStart := time.Now()
		result := addTwoNumbers(&value.L1, &value.L2)
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
	L1     ListNode
	L2     ListNode
	Result string
}

/*

===============
Test count  0 for node {{7 0x4000010270} {5 0x40000102a0}
                [7,8,0,7]
            }
7 -> 2 -> 4 -> 3 -> nil
5 -> 6 -> 4 -> nil
Solution 1: use stack base approach
>Solution result &{7 0x40000103b0}
7 -> 8 -> 0 -> 7 -> nil
Correct result is
                [7,8,0,7]

TimeLapse 2.685µs
===============
Test count  1 for node {{2 0x40000102c0} {5 0x40000102e0}
                [8,0,7]
            }
2 -> 4 -> 3 -> nil
5 -> 6 -> 4 -> nil
Solution 1: use stack base approach
>Solution result &{8 0x4000010400}
8 -> 0 -> 7 -> nil
Correct result is
                [8,0,7]

TimeLapse 1.926µs
===============
Test count  2 for node {{0 <nil>} {0 <nil>}
                [0]
            }
0 -> nil
0 -> nil
Solution 1: use stack base approach
>Solution result &{0 <nil>}
0 -> nil
Correct result is
                [0]

TimeLapse 907ns
===============
Test count  3 for node {{9 0x4000010300} {9 0x4000010360}
            [1, 0, 0 , 0 , 9 , 9 , 9 , 8]
            }
9 -> 9 -> 9 -> 9 -> 9 -> 9 -> 9 -> nil
9 -> 9 -> 9 -> 9 -> nil
Solution 1: use stack base approach
>Solution result &{1 0x40000104d0}
1 -> 0 -> 0 -> 0 -> 9 -> 9 -> 9 -> 8 -> nil
Correct result is
            [1, 0, 0 , 0 , 9 , 9 , 9 , 8]

TimeLapse 2.426µs
===============
TimeLapse Whole Program 963.903µs

*/

//REF
//https://www.geeksforgeeks.org/add-two-numbers-represented-by-linked-list/
//https://leetcode.com/problems/add-two-numbers-ii/
//https://leetcode.com/problems/add-two-numbers/description/
