package linkedlist

import (
	"fmt"
)

// Challenge:
// Reverse a singly linked list.
// Example:
// Input: 1->2->3->4->5->NULL
// Output: 5->4->3->2->1->NULL

// Solution:
// Two approaches, one goes through the list iteratively and swaps Next pointer
// Second approach does it recursively, it goes through the list till the end going deeper in recursion,
// then swaps the pointers on the way back

// Time complexity of both is O(n)
func reverseListIteratively(head *ListNode) *ListNode {

	currentNode := head
	var prevNode *ListNode

	for currentNode != nil {
		nextNode := currentNode.Next
		currentNode.Next = prevNode
		prevNode = currentNode
		currentNode = nextNode
	}

	return prevNode
}

func reverseListRecursively(head *ListNode) *ListNode {
	return reverse(nil, head)
}

func reverse(prev, current *ListNode) *ListNode {

	if current == nil {
		return prev
	}

	next := current.Next
	current.Next = prev

	return reverse(current, next)
}

func main3() {

	a := MakeListForward([]int{1})
	fmt.Println(reverseListRecursively(a))
	b := MakeListForward([]int{1, 2})
	fmt.Println(reverseListRecursively(b))
	c := MakeListForward([]int{1, 2, 3, 4})
	fmt.Println(reverseListRecursively(c))
}
