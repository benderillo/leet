package trees

import (
	"fmt"
	"math"
)

// Challenge:
// Given a binary tree, determine if it is a valid binary search tree (BST).

// Assume a BST is defined as follows:

// The left subtree of a node contains only nodes with keys less than the node's key.
// The right subtree of a node contains only nodes with keys greater than the node's key.
// Both the left and right subtrees must also be binary search trees.
// Example 1:

// Input:
//     2
//    / \
//   1   3
// Output: true
// Example 2:

//     5
//    / \
//   1   4
//      / \
//     3   6
// Output: false
// Explanation: The input is: [5,1,4,null,null,3,6]. The root node's value
//              is 5 but its right child's value is 4.

// Solution:
// Traverse the tree pre-order passing in Max and Min bounds and compare node values
// Max and min bounds are adjusted depending on whether the recursion takes left or right subtree.
// Time complexity: O(n)
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isValidBSTInternal(root, math.MaxInt64, math.MinInt64)
}

func isValidBSTInternal(root *TreeNode, max, min int) bool {

	if root == nil {
		return true
	}

	if root.Val <= min || root.Val >= max {
		return false
	}
	return isValidBSTInternal(root.Left, root.Val, min) && isValidBSTInternal(root.Right, max, root.Val)
}

func main2() {
	a := sortedArrayToBST([]int{-10, -3, 0, 5, 9})
	fmt.Println(isValidBST(a))
	fmt.Println()
	b := sortedArrayToBST([]int{1})
	fmt.Println(isValidBST(b))
	fmt.Println()
	c := sortedArrayToBST([]int{2, 1})
	fmt.Println(isValidBST(c))
	fmt.Println()
	d := sortedArrayToBST([]int{-1, 1, 2})
	fmt.Println(isValidBST(d))
	fmt.Println()
	e := sortedArrayToBST([]int{-22, -20, -10, 10, 11, 12, 13, 14})
	fmt.Println(isValidBST(e))
}
