package linkedlist

import (
	"fmt"
)

// Challenge
// Given a linked list, remove the n-th node from the end of list and return its head.
// Example:
// Given linked list: 1->2->3->4->5, and n = 2.

// After removing the second node from the end, the linked list becomes 1->2->3->5.

// Note:
// Given n will always be valid.

// Solution:
// This challenge is easily solved with two running pointers that are n elements apart.
// Also known as turtle and a hare approach.
// When the first pointer reaches the end of the list, the second pointer points to the Nth element from the end of the list

// Time complexity O(n), space complexity O(1)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {

	hare := head
	// Release a hare
	for i := 0; i < n; i++ {
		if hare == nil {
			return nil
		}
		hare = hare.Next
	}

	// Begin the race of tortoise and a hare
	tortoise := head
	prev := tortoise
	for hare != nil {
		hare = hare.Next
		prev = tortoise
		tortoise = tortoise.Next
	}

	if tortoise.Next == nil {
		// Reached the last element
		if tortoise == head {
			// if tortoise hasn't move anywhere, then it was one element list
			return nil
		}

		prev.Next = nil
		return head
	}

	tmp := tortoise.Next
	tortoise.Val = tmp.Val
	tortoise.Next = tmp.Next

	return head
}

func main1() {
	a := MakeListForward([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	fmt.Println(removeNthFromEnd(a, 2).String())
	b := MakeListForward([]int{1, 2, 3, 4, 5})
	fmt.Println(removeNthFromEnd(b, 1).String())
	c := MakeListForward([]int{1})
	fmt.Println(removeNthFromEnd(c, 1).String())
}
