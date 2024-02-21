package main

import (
	"fmt"
	"time"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// Time complexity: O(n)
// Space complexity: O(1)
func reverseKGroup(head *ListNode, k int) *ListNode {
	headCopy := head

	// line is the group for reverse
	headInLine := head
	tailInLine := head

	fakeHeadToReturn := &ListNode{
		Val:  -1,
		Next: head,
	}

	prevForLine := fakeHeadToReturn
	nextForLine := head

	count := 1
	isFirstLine := true
	for headCopy != nil {
		tailInLine = headCopy
		headCopy = headCopy.Next
		count++
		if count > k {
			// cut line of nodes from first node to last, for last set next to nil
			headInLine = prevForLine.Next
			nextForLine = tailInLine.Next
			tailInLine.Next = nil

			// set next as prev for first node in line
			tPrev := nextForLine
			// move last node in line to first place for prev node
			prevForLine.Next = tailInLine
			// set first node in line as prev for future lines
			prevForLine = headInLine

			// everything set up, just reverse cut line
			for headInLine != nil {
				temp := headInLine.Next
				headInLine.Next = tPrev
				tPrev = headInLine
				headInLine = temp
			}

			if isFirstLine {
				fakeHeadToReturn.Next = tPrev
				isFirstLine = false
			}
			count = 1
		}
	}

	return fakeHeadToReturn.Next
}

func printLinkedList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d -> ", head.Val)
		head = head.Next
	}
	fmt.Println("nil")
}

//==Approach :
//Time Complexity: O(n)
//Space Complexity: O(1)

func reverseKGroup_Approach2(head *ListNode, k int) *ListNode {
	var preTail *ListNode = nil  // stores the tail node of previous LL.
	curHead := head              // stores the head node of current LL
	curTail := head              // stores the tail node of current LL
	var nextHead *ListNode = nil // stores the head node of next LL

	for curHead != nil {
		// initialize count from 1
		count := 1

		// iterate the LL until count becomes k or we reach the last node.
		for curTail.Next != nil && count < k {
			curTail = curTail.Next
			count++
		}

		if count != k {
			break
		}

		// set the nextHead pointer to the head of the next LL.
		nextHead = curTail.Next

		// detach the RHS of the current LL.
		curTail.Next = nil

		// detach the LHS of the current LL.
		if preTail != nil {
			preTail.Next = nil
		}

		// after reversing the current LL, curHead becomes the new tail.
		// and curTail becomes the new head.
		curTail = reverse(curHead)

		// attach the LHS
		if preTail != nil {
			preTail.Next = curTail
		} else {
			// if preTail is nil then we have reversed the first LL
			// so store the reference of curHead in original head pointer.
			head = curTail
		}

		// attach the RHS
		curHead.Next = nextHead

		// update the pointer
		preTail = curHead
		curHead = nextHead
		curTail = nextHead
	}

	return head
}

func reverse(head *ListNode) *ListNode {
	var preNode *ListNode = nil
	curNode := head

	var nextNode *ListNode = head

	for curNode != nil {
		nextNode = nextNode.Next
		curNode.Next = preNode
		preNode = curNode
		curNode = nextNode
	}
	return preNode
}

//====

/* failure solution, need to check again

//Approach : recursive
//Time Complexity: O(n)
//Space Complexity: O(n/k)

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || k == 1 {
		return head
	}

	// Helper function to reverse a sublist of size k.
	reverse := func(start, end *ListNode) *ListNode {
		var prev, next *ListNode = nil, nil
		cur := start
		for cur != end {
			next = cur.Next
			cur.Next = prev
			prev = cur
			cur = next
		}
		return prev
	}

	dummy := &ListNode{Next: head} // Dummy node to handle edge cases
	prev := dummy

	for head != nil {
		tail := prev
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return dummy.Next // If remaining nodes are less than k, return the original list
			}
		}
		next := tail.Next // Save the next node for the next iteration
		head = reverse(head, tail)
		prev.Next = head
		tail.Next = next
		prev = tail
		head = next
	}

	return dummy.Next
}

// Function to print the linked list
func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d ", head.Val)
		head = head.Next
	}
	fmt.Println()
}

*/

func main() {
	timeStartWholeProgram := time.Now()

	testInput := []TestCase{

		{
			Head: ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}},
			K:    2,
			Result: `
      [2,1,4,3,5]
            `,
		},
		{
			Head: ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}},
			K:    3,
			Result: `
[3,2,1,4,5]

            `,
		},
	}
	for count, value := range testInput {
		fmt.Println("===============")
		fmt.Println("Test count ", count, "for node", value)
		fmt.Println("Solution 1: ")
		timeStart := time.Now()
		result := reverseKGroup(&value.Head, value.K)
		timeLapse := time.Since(timeStart)
		fmt.Println(">Solution result", result)
		printLinkedList(result)
		fmt.Println("Correct result is ", value.Result)
		printLinkedList(result)
		fmt.Println("TimeLapse", timeLapse)

		fmt.Println("Solution 2: (note: the test value has change because of the pointer value, need to update the test value again)")
		printLinkedList(&value.Head)
		timeStart = time.Now()
		result = reverseKGroup_Approach2(&value.Head, value.K)
		timeLapse = time.Since(timeStart)
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
	K      int
	Result string
}

/*

===============
Test count  0 for node {{1 0x400009e230} 2
      [2,1,4,3,5]
            }
Solution 1:
>Solution result &{2 0x40000a0150}
2 -> 1 -> 4 -> 3 -> 5 -> nil
Correct result is
      [2,1,4,3,5]

2 -> 1 -> 4 -> 3 -> 5 -> nil
TimeLapse 1.055µs
Solution 2:
1 -> 4 -> 3 -> 5 -> nil
>Solution result &{4 0x40000a0150}
4 -> 1 -> 5 -> 3 -> nil
Correct result is
      [2,1,4,3,5]

TimeLapse 982ns
===============
Test count  1 for node {{1 0x400009e270} 3
[3,2,1,4,5]

            }
Solution 1:
>Solution result &{3 0x400009e270}
3 -> 2 -> 1 -> 4 -> 5 -> nil
Correct result is
[3,2,1,4,5]


3 -> 2 -> 1 -> 4 -> 5 -> nil
TimeLapse 204ns
Solution 2:
1 -> 4 -> 5 -> nil
>Solution result &{5 0x400009e290}
5 -> 4 -> 1 -> nil
Correct result is
[3,2,1,4,5]


TimeLapse 185ns
===============
TimeLapse Whole Program 800.843µs

*/
