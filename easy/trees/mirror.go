package trees

import "container/list"

// Challenge:
// Given a binary tree, check whether it is a mirror of itself (ie, symmetric around its center).

// For example, this binary tree [1,2,2,3,4,4,3] is symmetric:

//     1
//    / \
//   2   2
//  / \ / \
// 3  4 4  3
// But the following [1,2,2,null,3,null,3] is not:
//     1
//    / \
//   2   2
//    \   \
//    3    3
// Note:
// Bonus points if you could solve it both recursively and iteratively.
// Solution:
// Run DFS on left and right subtrees in parallel comparing nodes
// Note: every recursive call have to compare opposite branches (remember they should be mirror), i.e. tree1.left vs tree2.right
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return isMirror(root.Left, root.Right)
}

func isMirror(root1, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}

	if root1 != nil && root2 != nil && root1.Val == root2.Val {
		return isMirror(root1.Left, root2.Right) && isMirror(root1.Right, root2.Left)
	}
	return false
}

// Solution:
// Can run BFS through left and right sub-tree in parallel to avoid recursion.
// This way both trees are traversed and nodes are compared
// Note: the BFS should traverse trees in opposite directions to each other, i.e. tree1 first left then right branch and tree2 first right then left branch

func isSymmetric2(root *TreeNode) bool {
	if root == nil {
		return true
	}
	queue1 := list.New()
	queue2 := list.New()

	queue1.PushBack(root.Left)
	queue2.PushBack(root.Right)

	for elem1, elem2 := queue1.Front(), queue2.Front(); elem1 != nil && elem2 != nil; elem1, elem2 = elem1.Next(), elem2.Next() {
		node1 := elem1.Value.(*TreeNode)
		node2 := elem2.Value.(*TreeNode)

		val1, val2 := -1, -1

		if node1 != nil {
			val1 = node1.Val
			queue1.PushBack(node1.Left)
			queue1.PushBack(node1.Right)
		}

		if node2 != nil {
			val2 = node2.Val
			queue2.PushBack(node2.Right)
			queue2.PushBack(node2.Left)
		}

		if val2 != val1 {
			return false
		}
	}

	return true
}
