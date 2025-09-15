package increasingtriplet

import "math"

func IncreasingTriplet(nums []int) bool {
	min, mid := math.MaxInt, math.MaxInt
	for _, n := range nums {
		if n <= min {
			min = n
			continue
		}

		if n <= mid {
			mid = n
			continue
		}

		return true
	}

	return false
}
