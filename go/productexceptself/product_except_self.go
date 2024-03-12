// 238. Product of Array Except Self
// Given an integer array nums, return an array answer such that answer[i] is equal
// to the product of all the elements of nums except nums[i].
//
// The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.
//
// You must write an algorithm that runs in O(n) time and without using the division operation.
package productexceptself

func ProductExceptSelf(nums []int) []int {
	size := len(nums)

	left := make([]int, size)
	left[0] = 1
	for i := 1; i < size; i++ {
		left[i] = left[i-1] * nums[i-1]
	}

	right := make([]int, size)
	right[size-1] = 1
	for i := size - 2; i >= 0; i-- {
		right[i] = right[i+1] * nums[i+1]
	}

	answers := make([]int, size)
	for i := 0; i < size; i++ {
		answers[i] = left[i] * right[i]
	}

	return answers
}
