package linkedlist

import (
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func MakeListBackwards(inputs []int) *ListNode {
	var node *ListNode = nil
	for _, v := range inputs {
		tmp := new(ListNode)
		tmp.Val = v
		tmp.Next = node
		node = tmp
	}
	return node
}

func MakeListForward(inputs []int) *ListNode {
	var first *ListNode
	var last *ListNode
	var newNode *ListNode

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

func getValReversly(node *ListNode, str *string) {
	if node.Next != nil {
		getValReversly(node.Next, str)
	}
	*str += strconv.Itoa(node.Val)
}

func (thiz *ListNode) StringRev() string {
	str := ""
	getValReversly(thiz, &str)
	return str
}

func (thiz *ListNode) String() string {
	str := ""
	node := thiz
	for node != nil {
		str += strconv.Itoa(node.Val) + "->"
		node = node.Next
	}
	str += "nil"
	return str
}
