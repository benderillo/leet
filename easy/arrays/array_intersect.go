package arrays

import (
	"fmt"
)

func intersect(nums1 []int, nums2 []int) []int {
    if len(nums1) == 0 || len(nums2) == 0 {
        return []int{}
    }

    if len(nums1) == 1 && len(nums2) == 1 {
        if nums1[0] == nums2[0] {
            return []int{nums1[0]}
        }
    }

    hash := make(map[int]int)
    var toHash []int = nil
    var toLookup []int = nil

    if len(nums1) < len(nums2) {
        toHash = nums1
        toLookup = nums2
    } else {
        toHash = nums2
        toLookup = nums1
    }

    // Fill the hash with occurences of all numbers from the smallest array
    for _, v := range toHash {
        hash[v]++
    }

    result := []int{}
    for _, v := range toLookup {
        if count, ok := hash[v]; ok && count > 0 {
            result = append(result, v)
            hash[v]--
        }
    }
    return result
}

// This method will sort both arrays
// Then it will go as far as shortest of two and:
// - step 1: take next elem in first array and move within second array from second array cur position while its elements smaller than the current
// - step2: if reached the end of second array, finish: there are no intersections
// - step3: if found element that is equal, then write it to result and go to step 1
// - step4: if found element that is larger, then go to step 1.
func intersect_v2() (nums1 []int, nums2 []int) []int {
    if len(nums1) == 0 || len(nums2) == 0 {
        return []int{}
    }

    if len(nums1) == 1 && len(nums2) == 1 {
        if nums1[0] == nums2[0] {
            return []int{nums1[0]}
        }
    }
	return []int {}
}

func main() {
	a := []int{1,2,2,1}
	b := []int{2,2}
	result := intersect(a,b)
	fmt.Println(result)

}
