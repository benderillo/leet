package linkedlist

import (
	"fmt"
	"math"

	"github.com/leet/easy/linkedlist"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// Challenge:
// Given two sorted lists, functions shall merge them into one list that is sorted
// This function does it with time complexity O(n)
// Space complexity is O(n)
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	l1Node := l1
	l2Node := l2
	var first, last *ListNode

	for l1Node != nil || l2Node != nil {
		cmpVal := math.MaxInt64

		if l2Node != nil {
			cmpVal = l2Node.Val
		}

		for l1Node != nil && l1Node.Val <= cmpVal {
			newNode := new(linkedlist.ListNode)
			newNode.Val = l1Node.Val

			if first == nil {
				first = newNode
				last = newNode
			} else {
				last.Next = newNode
				last = last.Next
			}
			l1Node = l1Node.Next
		}

		cmpVal = math.MaxInt64
		if l1Node != nil {
			cmpVal = l1Node.Val
		}

		for l2Node != nil && l2Node.Val <= cmpVal {
			newNode := new(linkedlist.ListNode)
			newNode.Val = l2Node.Val

			if first == nil {
				first = newNode
				last = newNode
			} else {
				last.Next = newNode
				last = last.Next
			}
			l2Node = l2Node.Next
		}
	}

	return first
}

// This brilliant solution is not mine but I copied it here because I really like it.
// It does in-place merge so the space complexity is O(1)
func mergeTwoListsRecursively(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists(l1, l2.Next)
		return l2
	}
}

func main4() {
	a := linkedlist.MakeListForward([]int{1})
	b := linkedlist.MakeListForward([]int{1})
	fmt.Println(mergeTwoLists(a, b).String())
}
