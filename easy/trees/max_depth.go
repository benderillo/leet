package trees

import "fmt"

// Challenge:
// Given a binary tree, find its maximum depth.

// The maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.

// Note: A leaf is a node with no children.

// Example:

// Given binary tree [3,9,20,null,null,15,7],

//     3
//    / \
//   9  20
//     /  \
//    15   7
// return its depth = 3.
//
// Solution:
// DFS the tree post-order.
// Each recursion branch, when tracing back, shall return max between depth of its left and right sub-trees
// Time complexity: O(n)

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)

	if leftDepth > rightDepth {
		return leftDepth + 1
	}
	return rightDepth + 1
}

func main1() {
	a := BuildTree([]int{-10, -3, 0, 5, 9})
	fmt.Println(maxDepth(a))
	fmt.Println()
	b := BuildTree([]int{1})
	fmt.Println(maxDepth(b))
	fmt.Println()
	c := BuildTree([]int{2, 1})
	fmt.Println(maxDepth(c))
	fmt.Println()
	d := BuildTree([]int{-1, 1, 2})
	fmt.Println(maxDepth(d))
	fmt.Println()
	e := BuildTree([]int{-22, -20, -10, 10, 11, 12, 13, 14})
	fmt.Println(maxDepth(e))
}
