package trees

// Challenge:
// Given an array where elements are sorted in ascending order, convert it to a height balanced BST.

// For this problem, a height-balanced binary tree is defined as a binary tree in which the depth of the two subtrees of every node never differ by more than 1.

// Example:

// Given the sorted array: [-10,-3,0,5,9],

// One possible answer is: [0,-3,9,-10,null,5], which represents the following height balanced BST:

//       0
//      / \
//    -3   9
//    /   /
//  -10  5

// Solution:
// Divide and concuer algorithm: Recursively finding a middle of an input array and setting it as a root node,
// Then for the left sub-tree we choose left half of the input array and recursively call the same routine.
// For the right sub-tree we choose the right half of the input array and recursively call the same routine.

// Complexity : O(n) because the algorithm examins each element once only and finding middle element is constant time.
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	root := new(TreeNode)
	centre := len(nums) >> 1
	root.Val = nums[centre]
	root.Left = sortedArrayToBST(nums[:centre])
	root.Right = sortedArrayToBST(nums[centre+1:])

	return root
}
