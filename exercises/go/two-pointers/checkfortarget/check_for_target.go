// Example 2: Given a sorted array of unique integers and a target integer,
// return true if there exists a pair of numbers that sum to target, false otherwise.
// This problem is similar to Two Sum. (In Two Sum, the input is not sorted).
//
// For example, given nums = [1, 2, 4, 6, 8, 9, 14, 15] and target = 13, return true because 4 + 9 = 13.
package checkfortarget

func CheckforTarget(nums []int, target int) bool {
	i := 0
	j := len(nums) - 1
	for i < j {
		sum := nums[i] + nums[j]
		if sum == target {
			return true
		}

		if sum > target {
			j--
		} else {
			i++
		}
	}

	return false
}
