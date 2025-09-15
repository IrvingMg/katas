// Squares of a Sorted Array
//
// Given an integer array nums sorted in non-decreasing order,
// return an array of the squares of each number sorted in non-decreasing order.
//
// Follow up: Squaring each element and sorting the new array is very trivial,
// could you find an O(n) solution using a different approach?
package sortedsquares

func SortedSquares(nums []int) []int {
	i := 0
	j := len(nums) - 1

	res := make([]int, len(nums))
	for k := j; k >= 0; k-- {
		var num int
		if abs(nums[j]) > abs(nums[i]) {
			num = nums[j]
			j--
		} else {
			num = nums[i]
			i++
		}
		res[k] = num * num
	}

	return res
}

func abs(n int) int {
	if n < 0 {
		return -n
	}

	return n
}
