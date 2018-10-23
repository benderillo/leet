package arrays

import (
	"fmt"
)

func rotate_copy(nums []int, k int) {
	if k == 0 || len(nums) == 1 || len(nums) == 0 || k == len(nums) {
		return
	}

	length := len(nums)

	k %= length

	another := make([]int, len(nums))

	for i, v := range nums {
		another[(i+k)%len(nums)] = v
	}
	for i, v := range another {
		nums[i] = v
	}
}

func rotate_naive(nums []int, k int) {
	if k == 0 || len(nums) == 1 || len(nums) == 0 || k == len(nums) {
		return
	}

	length := len(nums)

	k %= length

	storage := make([]int, k)
	copy(storage[0:], nums[length-k:])
	copy(nums[k:], nums[:])
	copy(nums[0:k], storage[:])
}

func rotate_manual(nums []int, k int) {
	if k == 0 || len(nums) == 1 || len(nums) == 0 || k == len(nums) {
		return
	}

	length := len(nums)

	k %= length

	storage := make([]int, k)

	for i := 0; i < k; i++ {
		storage[i] = nums[length-k+i]
	}

	for i := length - k - 1; i >= 0; i-- {
		nums[i+k] = nums[i]
	}

	for i := 0; i < k; i++ {
		nums[i] = storage[i]
	}
}

func rotate_cyclic(nums []int, k int) {
	if k == 0 || len(nums) == 1 || len(nums) == 0 || k == len(nums) {
		return
	}

	length := len(nums)

	k %= length

	swaps := 0
	for start := 0; swaps < len(nums); start++ {
		cur := start
		oldVal := nums[cur]
		for {
			next := (cur + k) % len(nums)
			tmp := nums[next]
			nums[next] = oldVal
			oldVal = tmp
			swaps++
			cur = next

			if cur == start {
				break
			}
		}
	}
}

func rotate_inverse(nums []int, k int) {
	if k == 0 || len(nums) == 1 || len(nums) == 0 || k == len(nums) {
		return
	}

	length := len(nums)

	k %= length

	// inverse the entire array, now the items that should shift and appear in the beginning are in the beginning
	inverse(nums, 0, len(nums)-1)
	// now we need to inverse the order of those k numbers in the beginning
	inverse(nums, 0, k-1)
	// and now inverse the numbers that come after k
	inverse(nums, k, len(nums)-1)
}

func inverse(nums []int, left, right int) {
	for left < right {
		tmp := nums[left]
		nums[left] = nums[right]
		nums[right] = tmp
		left++
		right--
	}
}

func main() {
	input1 := []int{1, 2, 3, 4, 5, 6}
	input2 := []int{-1, -100, 3, 99}
	fmt.Println(input1)
	rotate_cyclic(input1, 5)
	fmt.Println(input1)
	fmt.Println(input2)
	rotate_cyclic(input2, 2)
	fmt.Println(input2)
}
