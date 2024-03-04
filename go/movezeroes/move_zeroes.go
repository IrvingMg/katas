// 283. Move Zeroes
// Given an integer array nums, move all 0's to the end of it while maintaining the
// relative order of the non-zero elements.
//
// Note that you must do this in-place without making a copy of the array.
package movezeroes

func MoveZeroes(nums []int) {
	var zeroIdx int
	for i := 1; i < len(nums); i++ {
		if nums[zeroIdx] != 0 {
			zeroIdx++
			continue
		}

		if nums[i] != 0 {
			nums[zeroIdx], nums[i] = nums[i], 0
			zeroIdx++
		}
	}
}
