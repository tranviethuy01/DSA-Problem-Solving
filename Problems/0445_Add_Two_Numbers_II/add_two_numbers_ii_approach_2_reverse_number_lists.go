package main
// NOTED: this solution is failed, review again
import 
(
    "fmt"
    "time"
    //"sort"
)


type ListNode struct {
  Val int
  Next *ListNode

}


// LinkedList represents a linked list
type LinkedList struct{}

// ReverseList reverses the linked list and returns the head of the reversed list
func (ll *LinkedList) ReverseList(list *ListNode) *ListNode {
	var prev *ListNode = nil
	curr := list
	var next *ListNode = nil
	for curr != nil {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

// AddTwoLists calculates and prints the sum of two numbers represented by linked lists
func (ll *LinkedList) AddTwoLists(first *ListNode, second *ListNode) {
	// Reverse both lists
	first = ll.ReverseList(first)
	second = ll.ReverseList(second)

	carry := 0
	//var head *ListNode = nil
	//var prev *ListNode = nil
	var sumList *ListNode = nil

	// Add the two lists and carry over if necessary
	for first != nil || second != nil || carry == 1 {
		newVal := carry
		if first != nil {
			newVal += first.Val
		}
		if second != nil {
			newVal += second.Val
		}
		carry = newVal / 10
		newVal = newVal % 10

		// Create a new node for the sum and append it to the beginning of the final ans list
		newNode := &ListNode{Val: newVal, Next: sumList}
		sumList = newNode

		// Initialize nodes for the next iteration
		if first != nil {
			first = first.Next
		}
		if second != nil {
			second = second.Next
		}
	}

	ll.PrintList(sumList)
}

// PrintList prints the linked list
func (ll *LinkedList) PrintList(head *ListNode) {
	for head.Next != nil {
		fmt.Printf("%d ", head.Val)
		head = head.Next
	}
	fmt.Printf("%d\n", head.Val)
}


func main() {
    timeStartWholeProgram := time.Now()
    testInput := []TestCase{
        {
            L1: ListNode{Val: 7, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}} ,
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
            L1:  ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9}}}}}}} ,
            L2:  ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9}}}},
            Result: `
            [1, 0, 0 , 0 , 9 , 9 , 9 , 8]               
            `,
        },



      } 
    for count, value := range testInput {
        fmt.Println("===============")
        fmt.Println("Test count ", count, "for node", value)
        // Create a LinkedList instance and call AddTwoLists method to add the two lists
	      ll := &LinkedList{}
        ll.PrintList(&value.L1)
        ll.PrintList(&value.L2)
        ll.AddTwoLists(&value.L1, &value.L2)

        fmt.Println("Solution 1: use reverse number lists approach")
        timeStart := time.Now()
        ll.AddTwoLists(&value.L1, &value.L2)
        timeLapse := time.Since(timeStart)
        //fmt.Println(">Solution result", result)
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
7 2 4 3
5 6 4
7 8 0 7
Solution 1: use reverse number lists approach
1 2
Correct result is  
                [7,8,0,7]
            
TimeLapse 18.259µs
===============
Test count  1 for node {{2 0x40000102c0} {5 0x40000102e0} 
                [8,0,7]
            }
2 4 3
5 6 4
8 0 7
Solution 1: use reverse number lists approach
7
Correct result is  
                [8,0,7]
            
TimeLapse 9.241µs
===============
Test count  2 for node {{0 <nil>} {0 <nil>} 
                [0]
            }
0
0
0
Solution 1: use reverse number lists approach
0
Correct result is  
                [0]
            
TimeLapse 9.277µs
===============
Test count  3 for node {{9 0x4000010300} {9 0x4000010360} 
            [1, 0, 0 , 0 , 9 , 9 , 9 , 8]               
            }
9 9 9 9 9 9 9
9 9 9 9
1 0 0 0 9 9 9 8
Solution 1: use reverse number lists approach
1 8
Correct result is  
            [1, 0, 0 , 0 , 9 , 9 , 9 , 8]               
            
TimeLapse 18.073µs
===============
TimeLapse Whole Program 826.472µs


*/

//REF
//https://www.geeksforgeeks.org/add-two-numbers-represented-by-linked-list/
//https://leetcode.com/problems/add-two-numbers-ii/
//https://leetcode.com/problems/add-two-numbers/description/
