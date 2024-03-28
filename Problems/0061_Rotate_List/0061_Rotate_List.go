package main

import (
	"fmt"
	"runtime"
	"time"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return head
	}

	// Find the length of the linked list
	length := 1
	tail := head
	for tail.Next != nil {
		tail = tail.Next
		length++
	}

	// Calculate the effective rotation amount
	k = k % length
	if k == 0 {
		return head // No rotation needed
	}

	// Traverse to the node that will become the new tail
	newTailPosition := length - k - 1
	newTail := head
	for i := 0; i < newTailPosition; i++ {
		newTail = newTail.Next
	}

	// Adjust pointers to perform rotation
	newHead := newTail.Next
	newTail.Next = nil
	tail.Next = head

	return newHead
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val, " ")
		head = head.Next
	}
	fmt.Println()
}

func main() {
	timeStartWholeProgram := time.Now()

	testInput := []TestCase{
		{
			ListNode: ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}},
			K:        2,
			Result: `
  [4,5,1,2,3]

            `,
		},
		{
			ListNode: ListNode{Val: 0, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2}}},
			K:        4,
			Result: `
[2,0,1]

            `,
		},
	}

	// Memory before allocation
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	memBefore := m.Alloc

	for count, value := range testInput {
		fmt.Println("===============")
		fmt.Println("Test count ", count, "for node", value)
		fmt.Println("Solution 1: rotateRight")
		timeStart := time.Now()
		result := rotateRight(&value.ListNode, value.K)
		timeLapse := time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		printList(result)
		fmt.Println("TimeLapse", timeLapse)

		// Memory after allocation
		runtime.ReadMemStats(&m)
		memAfter := m.Alloc
		fmt.Println("Memory before", memBefore, "bytes", "Memory after", memAfter, "bytes", "Memory used:", memAfter-memBefore, "bytes")

		fmt.Printf("Memory usage (HeapAlloc) after Test Case i %d, : %v bytes\n", count, m.HeapAlloc)

	}

	timeLapsedWholeProgram := time.Since(timeStartWholeProgram)
	fmt.Println("===============")
	fmt.Println("TimeLapse Whole Program", timeLapsedWholeProgram)
}

type TestCase struct {
	ListNode ListNode
	K        int
	Result   string
}

/*


===============
Test count  0 for node {{1 0x4000010250} 2 
  [4,5,1,2,3]

            }
Solution 1: rotateRight
>Solution result &{4 0x4000010280}
Correct result is  
  [4,5,1,2,3]

            
4 5 1 2 3 
TimeLapse 834ns
Memory before 69112 bytes Memory after 70240 bytes Memory used: 1128 bytes
Memory usage (HeapAlloc) after Test Case i 0, : 70240 bytes
===============
Test count  1 for node {{0 0x4000010290} 4 
[2,0,1]

            }
Solution 1: rotateRight
>Solution result &{2 0x4000072150}
Correct result is  
[2,0,1]

            
2 0 1 
TimeLapse 204ns
Memory before 69112 bytes Memory after 70368 bytes Memory used: 1256 bytes
Memory usage (HeapAlloc) after Test Case i 1, : 70368 bytes
===============
TimeLapse Whole Program 792.625µs


 */
//REF
