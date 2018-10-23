package linkedlist

import (
	"fmt"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

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

func (thiz *ListNode) String() string {
	str := ""
	getValReversly(thiz, &str)
	return str
}

func getValReversly(node *ListNode, str *string) {
	if node.Next != nil {
		getValReversly(node.Next, str)
	}
	*str += strconv.Itoa(node.Val)
}

func makeListForward(inputs []int) *ListNode {
	var first *ListNode = nil
	var last *ListNode = nil
	var newNode *ListNode = nil

	for _, v := range inputs {
		newNode = new(ListNode)
		newNode.Val = v

		if first == nil {
			first = newNode
			last = newNode
		} else {
			last.Next = newNode
			last = newNode
		}

	}
	return first
}

func makeListBackwards(inputs []int) *ListNode {
	var node *ListNode = nil
	for _, v := range inputs {
		tmp := new(ListNode)
		tmp.Val = v
		tmp.Next = node
		node = tmp
	}
	return node
}

func main() {
	a := makeListForward([]int{9, 8, 8, 1})
	b := makeListBackwards([]int{4, 6, 5})
	fmt.Println("a: " + a.String())
	fmt.Println("b: " + b.String())
	fmt.Println("c: " + addTwoNumbers(a, b).String())
}
