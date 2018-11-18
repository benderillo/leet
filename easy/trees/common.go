package trees

import (
	"container/list"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Action func(node *TreeNode)

// Input is flatten or array-like representation of a tree (consider the tree to be full)
// Example input is [3,9,20,null,null,15,7]

func BuildTree(values []int) *TreeNode {

	root := new(TreeNode)
	root.Val = values[0]

	queue := list.New()
	queue.PushBack(root)
	curIndex := 1

	for elem := queue.Front(); elem != nil; elem = elem.Next() {
		node := elem.Value.(*TreeNode)

		val := getElem(values, curIndex)
		if val != -1 {
			node.Left = new(TreeNode)
			node.Left.Val = val
			queue.PushBack(node.Left)
		}
		curIndex++

		val = getElem(values, curIndex)
		if val != -1 {
			node.Right = new(TreeNode)
			node.Right.Val = val
			queue.PushBack(node.Right)
		}
		curIndex++
	}

	return root
}

func getElem(values []int, index int) int {
	if index >= len(values) {
		return -1
	}

	return values[index]
}

func DFSPostOrder(root *TreeNode, action Action) {
	if root.Left != nil {
		DFSPostOrder(root.Left, action)
	}

	if root.Right != nil {
		DFSPostOrder(root.Right, action)
	}

	action(root)
}

func DFSInOrder(root *TreeNode, action Action) {
	if root.Left != nil {
		DFSInOrder(root.Left, action)
	}

	action(root)

	if root.Right != nil {
		DFSInOrder(root.Right, action)
	}
}

func DFSPreOrder(root *TreeNode, action Action) {
	action(root)

	if root.Left != nil {
		DFSPreOrder(root.Left, action)
	}

	if root.Right != nil {
		DFSPreOrder(root.Right, action)
	}
}
