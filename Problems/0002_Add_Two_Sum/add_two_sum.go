package main

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

//approach: straightforward iteration over the linked lists
//The algorithm iterates through the linked lists once, which takes O(n), where n is the maximum number of nodes in either of the linked lists.
//Time complexity: O(n)
//Space complexity: O(n)
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    dummyHead := &ListNode{}
    current := dummyHead
    carry := 0

    for l1 != nil || l2 != nil {
        sum := carry
        if l1 != nil {
            sum += l1.Val
            l1 = l1.Next
        }
        if l2 != nil {
            sum += l2.Val
            l2 = l2.Next
        }

        current.Next = &ListNode{Val: sum % 10}
        current = current.Next
        carry = sum / 10
    }

    if carry > 0 {
        current.Next = &ListNode{Val: carry}
    }

    return dummyHead.Next
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
            L1: ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}},
            L2: ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}},
            Result: `
                [7,0,8]
            `,
        },
   {
            L1:  ListNode{Val: 0} ,
            L2: ListNode{Val: 0},
            Result: `
                [0]
            `,
        },

   {
            L1:  ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9}}}}}}} ,
            L2:  ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9}}}},
            Result: `
                [8,9,9,9,0,0,0,1]
            `,
        },



      } 
    for count, value := range testInput {
        fmt.Println("===============")
        fmt.Println("Test count ", count, "for node", value)
        printLinkedList(&value.L1)
        printLinkedList(&value.L2)
        fmt.Println("Solution 1: straightforward iteration over linked lists")
        timeStart := time.Now()
        result := addTwoNumbers(&value.L1 , &value.L2)
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
Test count  0 for node {{2 0x4000118230} {5 0x4000118250} 
                [7,0,8]
            }
2 -> 4 -> 3 -> nil
5 -> 6 -> 4 -> nil
Solution 1: straightforward iteration over linked lists
>Solution result &{7 0x4000118310}
7 -> 0 -> 8 -> nil
Correct result is  
                [7,0,8]
            
TimeLapse 1.426µs
===============
Test count  1 for node {{0 <nil>} {0 <nil>} 
                [0]
            }
0 -> nil
0 -> nil
Solution 1: straightforward iteration over linked lists
>Solution result &{0 <nil>}
0 -> nil
Correct result is  
                [0]
            
TimeLapse 555ns
===============
Test count  2 for node {{9 0x4000118270} {9 0x40001182d0} 
                [8,9,9,9,0,0,0,1]
            }
9 -> 9 -> 9 -> 9 -> 9 -> 9 -> 9 -> nil
9 -> 9 -> 9 -> 9 -> nil
Solution 1: straightforward iteration over linked lists
>Solution result &{8 0x4000118390}
8 -> 9 -> 9 -> 9 -> 0 -> 0 -> 0 -> 1 -> nil
Correct result is  
                [8,9,9,9,0,0,0,1]
            
TimeLapse 1.555µs
===============
TimeLapse Whole Program 819.88µs


*/

//REF
//https://www.geeksforgeeks.org/add-two-numbers-represented-by-linked-list/
//https://leetcode.com/problems/add-two-numbers/description/
