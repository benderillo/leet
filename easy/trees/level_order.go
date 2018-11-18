package trees

import "container/list"

// Challenge:
// Given a binary tree, return the level order traversal of its nodes' values. (ie, from left to right, level by level).

// For example:
// Given binary tree [3,9,20,null,null,15,7],
//     3
//    / \
//   9  20
//     /  \
//    15   7
// return its level order traversal as:
// [
//   [3],
//   [9,20],
//   [15,7]
// ]

// Solution:
// Run BFS but insert a special node-marker upon completion of each layer to separate the layers line by line
// This solution was inspired by "counting" approach here: https://www.geeksforgeeks.org/print-level-order-traversal-line-line/
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	results := [][]int{}

	queue := list.New()

	queue.PushBack(root)
	// Nil node will be a marker that tells us where the level ends
	queue.PushBack(nil)

	currentLevel := []int{}
	for elem := queue.Front(); elem != nil; elem = elem.Next() {
		node, ok := elem.Value.(*TreeNode)
		if !ok {
			if len(currentLevel) == 0 {
				// No more nodes left
				break
			}
			// Found marker, then need to push new marker to the queue to signal end of current level
			queue.PushBack(nil)
			// Add current level to the list and create another empty level to fill up
			results = append(results, currentLevel)
			currentLevel = []int{}
			continue
		}

		currentLevel = append(currentLevel, node.Val)
		if node.Left != nil {
			queue.PushBack(node.Left)
		}
		if node.Right != nil {
			queue.PushBack(node.Right)
		}
	}

	return results
}
