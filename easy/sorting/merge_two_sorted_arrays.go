package sorting

// Challenge:
// Given two sorted integer arrays nums1 and nums2, merge nums2 into nums1 as one sorted array.

// Note:

// The number of elements initialized in nums1 and nums2 are m and n respectively.
// You may assume that nums1 has enough space (size that is greater or equal to m + n) to hold additional elements from nums2.
// Example:

// Input:
// nums1 = [1,2,3,0,0,0], m = 3
// nums2 = [2,5,6],       n = 3

// Output: [1,2,2,3,5,6]

// Solution:
// Go from the end of nums2, pick one number at a time, and do insertion of it from end to beginning of num1 (like insertion sort)
// Complexity is O(n*m)

// Note: there were solutions in leetcode that would optimise by examining both arrays from the end
// and inserting largest element to the most end of the nums1 array. It would allow for O(n) complexity.
// However, I do not find such solutions correct as they assume that the size of the nums1 is equal but not greater than n+m
func merge(nums1 []int, m int, nums2 []int, n int) {
	// Initialise end as the next position after the end of nums1 array
	end := m - 1
	// pick elements from nums2 from the end one by one
	for srcPos := n - 1; srcPos >= 0; srcPos-- {
		dstPos := end
		// traverse nums1 from the end, shifting all elements by one position until we find match place to insert
		for dstPos >= 0 && nums1[dstPos] > nums2[srcPos] {
			nums1[dstPos+1] = nums1[dstPos]
			dstPos--
		}
		// Insert an element from nums2 array
		nums1[dstPos+1] = nums2[srcPos]
		end++
	}
}
