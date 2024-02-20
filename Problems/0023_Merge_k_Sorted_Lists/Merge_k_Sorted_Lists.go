package main

import (
	"container/heap"
	"fmt"
	"sort"
	"time"
)

//NOTE: doing

// ListNode represents a node in a linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// ListNodeHeap is a min-heap for ListNode pointers.
type ListNodeHeap []*ListNode

// Implement the heap.Interface methods for ListNodeHeap.
func (h ListNodeHeap) Len() int           { return len(h) }
func (h ListNodeHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h ListNodeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *ListNodeHeap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}

func (h *ListNodeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// Time Complexity:O(n log k)
// Space Complexity: O(k + n)
// mergeKLists merges k sorted linked lists into one sorted linked list.
//
//approach:merge sort
func mergeKLists(lists []*ListNode) *ListNode {
	h := &ListNodeHeap{}
	heap.Init(h)

	// Push the head of each list onto the heap.
	for _, list := range lists {
		if list != nil {
			heap.Push(h, list)
		}
	}

	dummy := &ListNode{}
	current := dummy

	// Pop the smallest element from the heap and append it to the merged list.
	for h.Len() > 0 {
		smallest := heap.Pop(h).(*ListNode)
		current.Next = smallest
		current = current.Next
		if smallest.Next != nil {
			heap.Push(h, smallest.Next)
		}
	}

	return dummy.Next
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val, ",")
		head = head.Next
	}
}

//==== approach brute force
//Time Complexity: O(n log n)
//Space Complexity: O(n)

// mergeKListsBruteForce merges k sorted linked lists into one sorted linked list using brute force.
func mergeKListsBruteForce(lists []*ListNode) *ListNode {
	// Collect all values from the linked lists.
	var values []int
	for _, list := range lists {
		for list != nil {
			values = append(values, list.Val)
			list = list.Next
		}
	}

	// Sort the collected values.
	sort.Ints(values)

	// Construct the sorted linked list.
	dummy := &ListNode{}
	current := dummy
	for _, val := range values {
		current.Next = &ListNode{Val: val}
		current = current.Next
	}

	return dummy.Next
}

//====

//==== approach : priority queue

//Time Complexity: O(n log k)
//Space Complexity: O(k + n)

// ListNodeHeap is a min-heap for ListNode pointers.
type ListNodeHeap []*ListNode

// Implement the heap.Interface methods for ListNodeHeap.
func (h ListNodeHeap) Len() int           { return len(h) }
func (h ListNodeHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h ListNodeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *ListNodeHeap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}

func (h *ListNodeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// mergeKLists merges k sorted linked lists into one sorted linked list.
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	dummyHead := &ListNode{}
	dummyTail := dummyHead

	h := &ListNodeHeap{}
	heap.Init(h)

	for _, head := range lists {
		if head != nil {
			heap.Push(h, head)
		}
	}

	for h.Len() > 0 {
		minNode := heap.Pop(h).(*ListNode)
		if minNode.Next != nil {
			heap.Push(h, minNode.Next)
		}
		dummyTail.Next = minNode
		dummyTail = dummyTail.Next
	}

	return dummyHead.Next
}

//=====

//==== approach :  divide and conquer
//Time Complexity: O(k log k)
//Space Complexity: O(log k)

// ListNode represents a node in a singly-linked list.
//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

// mergeKLists merges k sorted linked lists into one sorted linked list.
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	return mergeKListsHelper(lists, 0, len(lists)-1)
}

// mergeKListsHelper is a helper function to recursively merge k sorted linked lists.
func mergeKListsHelper(lists []*ListNode, start, end int) *ListNode {
	if start > end {
		return nil
	}
	if start == end {
		return lists[start]
	}

	mid := start + (end-start)/2
	left := mergeKListsHelper(lists, start, mid)
	right := mergeKListsHelper(lists, mid+1, end)
	return merge(left, right)
}

// merge merges two sorted linked lists into one sorted linked list.
func merge(list1Head, list2Head *ListNode) *ListNode {
	dummyHead := &ListNode{}
	dummyTail := dummyHead

	for list1Head != nil && list2Head != nil {
		if list1Head.Val < list2Head.Val {
			dummyTail.Next = list1Head
			list1Head = list1Head.Next
		} else {
			dummyTail.Next = list2Head
			list2Head = list2Head.Next
		}
		dummyTail = dummyTail.Next
	}

	if list1Head != nil {
		dummyTail.Next = list1Head
	} else {
		dummyTail.Next = list2Head
	}

	return dummyHead.Next
}

//====

//=== approach: merge 2 lists at a time

// ListNode represents a node in a singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// mergeKLists merges k sorted linked lists into one sorted linked list by merging two lists at a time.
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	for len(lists) > 1 {
		var mergedLists []*ListNode

		// Merge two lists at a time
		for i := 0; i < len(lists)-1; i += 2 {
			mergedLists = append(mergedLists, merge(lists[i], lists[i+1]))
		}

		// If there's an odd number of lists, append the last one
		if len(lists)%2 != 0 {
			mergedLists = append(mergedLists, lists[len(lists)-1])
		}

		// Update the lists for the next iteration
		lists = mergedLists
	}

	return lists[0]
}

// merge merges two sorted linked lists into one sorted linked list.
func merge(list1Head, list2Head *ListNode) *ListNode {
	dummyHead := &ListNode{}
	current := dummyHead

	for list1Head != nil && list2Head != nil {
		if list1Head.Val < list2Head.Val {
			current.Next = list1Head
			list1Head = list1Head.Next
		} else {
			current.Next = list2Head
			list2Head = list2Head.Next
		}
		current = current.Next
	}

	if list1Head != nil {
		current.Next = list1Head
	} else {
		current.Next = list2Head
	}

	return dummyHead.Next
}

//===

//approach : Compare K elements One By One
//Time Complexity: O(nk)
//Space Complexity: O(k + n)

// ListNode represents a node in a singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// mergeKLists merges k sorted linked lists into one sorted linked list by comparing k elements one by one.
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	// Initialize the result list and pointers for each input list
	var result *ListNode
	pointers := make([]*ListNode, len(lists))
	for i, list := range lists {
		pointers[i] = list
	}

	// Iterate until all lists are merged
	for {
		// Find the minimum node among the current pointers
		minNodeIndex := -1
		for i, node := range pointers {
			if node != nil && (minNodeIndex == -1 || node.Val < pointers[minNodeIndex].Val) {
				minNodeIndex = i
			}
		}

		// Break the loop if no minimum node is found
		if minNodeIndex == -1 {
			break
		}

		// Append the minimum node to the result list
		if result == nil {
			result = pointers[minNodeIndex]
		} else {
			result.Next = pointers[minNodeIndex]
			result = result.Next
		}

		// Move the pointer of the selected list to the next node
		pointers[minNodeIndex] = pointers[minNodeIndex].Next
	}

	return result
}

//====

func main() {
	timeStartWholeProgram := time.Now()

	testInput := []TestCase{
		{
			L: []*ListNode{
				{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}},
				{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}},
				{Val: 2, Next: &ListNode{Val: 6}},
			},
			Result: `
[1,1,2,3,4,4,5,6]
            `,
		},
		{
			L: []*ListNode{},
			Result: `
[]

            `,
		},
		{
			L: []*ListNode{
				&ListNode{}},
			Result: `
[]
            `,
		},
	}
	for count, value := range testInput {
		fmt.Println("===============")
		fmt.Println("Test count ", count, "for node", value)
		fmt.Println("Solution 1: ")
		timeStart := time.Now()
		result := mergeKLists(value.L)
		timeLapse := time.Since(timeStart)
		fmt.Println(">Solution result", result)
		printList(result)
		fmt.Println("")
		fmt.Println("Correct result is ", value.Result)
		fmt.Println("TimeLapse", timeLapse)
	}

	timeLapsedWholeProgram := time.Since(timeStartWholeProgram)
	fmt.Println("===============")
	fmt.Println("TimeLapse Whole Program", timeLapsedWholeProgram)
}

type TestCase struct {
	L      []*ListNode
	Result string
}

/*


 */
//REF
//
