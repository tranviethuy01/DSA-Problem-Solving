package main

import (
	"fmt"
	"sort"
	"time"
)

// ListNode definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// Straightforward approach
//Time Complexity: O(n)
//Space complexity:  O(1)

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// Dummy node to start the merged list
	dummy := &ListNode{}
	// Pointer to traverse the merged list
	current := dummy

	// While both lists are not empty
	for l1 != nil && l2 != nil {
		// Choose the smaller value between l1 and l2
		if l1.Val <= l2.Val {
			current.Next = l1
			l1 = l1.Next
		} else {
			current.Next = l2
			l2 = l2.Next
		}
		// Move the current pointer forward
		current = current.Next
	}

	// If any list still has remaining elements, append them
	if l1 != nil {
		current.Next = l1
	} else {
		current.Next = l2
	}

	// Return the merged list starting from the next of dummy node
	return dummy.Next
}

// Function to print the linked list
func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d -> ", head.Val)
		head = head.Next
	}
	fmt.Println("nil")
}

//approach : Brute Force
//Time complexity: O(n log n), where n is the total number of nodes in both input lists. This is because sorting an array of n elements typically requires O(n log n) time complexity using comparison-based sorting algorithms like quicksort or mergesort.
//Space complexity: O(n)

// Function to merge two lists and sort the merged list
func mergeTwoLists_BruteForce(l1 *ListNode, l2 *ListNode) *ListNode {
	// Create an array to store all the values from both lists
	values := []int{}

	// Traverse the first list and append values to the array
	for l1 != nil {
		values = append(values, l1.Val)
		l1 = l1.Next
	}

	// Traverse the second list and append values to the array
	for l2 != nil {
		values = append(values, l2.Val)
		l2 = l2.Next
	}

	// Sort the array of values
	sort.Ints(values)

	// Create a dummy node to build the merged list
	dummy := &ListNode{}
	current := dummy

	// Create nodes for the merged list using the sorted values
	for _, val := range values {
		current.Next = &ListNode{Val: val}
		current = current.Next
	}

	// Return the merged list starting from the next of dummy node
	return dummy.Next
}

//=====

func main() {
	timeStartWholeProgram := time.Now()
	testInput := []TestCase{
		{
			L1: 	ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: nil}}},
	L2: ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}},
			Result: `
[1,1,2,3,4,4]

            `,
		},

	}
	for count, value := range testInput {
		fmt.Println("===============")
		fmt.Println("Test count ", count, "for node", value)
		fmt.Println("Solution 1: ")
		timeStart := time.Now()
		result := mergeTwoLists(&value.L1, &value.L2)
		timeLapse := time.Since(timeStart)
		fmt.Println(">Solution result", result)
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)

		fmt.Println("Solution 2: ")
		timeStart = time.Now()
		result = mergeTwoLists_BruteForce(&value.L1, &value.L2)
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
	L1     ListNode
	L2     ListNode
	Result string
}

/*


===============
Test count  0 for node {{1 0x4000118230} {1 0x4000118250} 
[1,1,2,3,4,4]

            }
Solution 1: 
>Solution result &{1 0x400011a160}
Correct result is  
[1,1,2,3,4,4]

            
TimeLapse 889ns
Solution 2: 
>Solution result &{1 0x40001182a0}
Correct result is  
[1,1,2,3,4,4]

            
TimeLapse 4.926µs
===============
TimeLapse Whole Program 320.793µs

*/
