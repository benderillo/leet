package linkedlist

import (
	"fmt"
)

// Challenge:
// Given a singly linked list, determine if it is a palindrome.
// Example 1:
// Input: 1->2
// Output: false
// Example 2:
// Input: 1->2->2->1
// Output: true

// Solution:
// Need to compare first half of the list with the second half
// If they match then the list contains palindrome.
// To do so, we shall do the following:
// 1. Find the middle of the list (will use tortoise and hare approach)
// 2. Reverse one half of the list (in place)
// 3. Go node by node for both parts of the list and compare
// 4. Reverse back the part that was reversed (to preserve original state of input list)
// Time complexity is O(n) and space complexity is O(1)

func isPalindrome(head *ListNode) bool {
	result := true
	tortoise := head
	hare := head

	if head == nil || head.Next == nil {
		return result
	}

	// Let's find the middle element
	for hare.Next != nil && hare.Next.Next != nil {
		tortoise = tortoise.Next
		hare = hare.Next.Next
	}

	// Inverse the list from the middle to the tail
	tortoise.Next = inverseList(tortoise.Next)
	// remember the head of inversed list
	backup := tortoise

	// shift one step to get to the start of inverted list
	tortoise = tortoise.Next
	for tortoise != nil {
		if head.Val != tortoise.Val {
			result = false
			break
		}
		head = head.Next
		tortoise = tortoise.Next
	}

	// restore original order of input list
	backup.Next = inverseList(backup.Next)

	return result
}

func inverseList(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		tmp := head.Next
		head.Next = prev
		prev = head
		head = tmp
	}
	return prev
}

func main9() {
	a := MakeListForward([]int{1, 2, 3, 4})
	fmt.Println(a.String(), isPalindrome(a))
	b := MakeListForward([]int{1})
	fmt.Println(b.String(), isPalindrome(b))
	c := MakeListForward([]int{1, 1})
	fmt.Println(c.String(), isPalindrome(c))
	d := MakeListForward([]int{1, 2, 1})
	fmt.Println(d.String(), isPalindrome(d))
}
