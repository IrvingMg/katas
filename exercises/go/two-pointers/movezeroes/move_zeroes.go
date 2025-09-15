// 283. Move Zeroes
// Given an integer array nums, move all 0's to the end of it while maintaining the
// relative order of the non-zero elements.
//
// Note that you must do this in-place without making a copy of the array.
package movezeroes

func MoveZeroes(nums []int) {
	for i, j := 0, 0; i < len(nums); i++ {
		if nums[i] == 0 {
			continue
		}

		nums[i], nums[j] = nums[j], nums[i]
		j++
	}
}
