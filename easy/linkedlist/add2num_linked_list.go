package linkedlist

import (
	"fmt"
)

// leetcode add two numbers (linked list) challenge
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	num1 := l1
	num2 := l2

	var first *ListNode = nil
	var last *ListNode = nil
	var newNode *ListNode = nil
	var carry = 0
	for (num1 != nil) || (num2 != nil) {
		var val1 = 0
		if num1 != nil {
			val1 = num1.Val
		}
		var val2 = 0
		if num2 != nil {
			val2 = num2.Val
		}

		sum := carry + val1 + val2
		carry = sum / 10

		newNode = new(ListNode)
		newNode.Val = sum % 10

		if first == nil {
			first = newNode
			last = first
		} else {
			last.Next = newNode
			last = newNode
		}

		if num1 != nil {
			num1 = num1.Next
		}

		if num2 != nil {
			num2 = num2.Next
		}
	}
	if carry > 0 {
		last.Next = new(ListNode)
		last.Next.Val = carry
	}

	return first
}

func main() {
	a := MakeListForward([]int{9, 8, 8, 1})
	b := MakeListBackwards([]int{4, 6, 5})
	fmt.Println("a: " + a.StringRev())
	fmt.Println("b: " + b.StringRev())
	fmt.Println("c: " + addTwoNumbers(a, b).StringRev())
}
