// Example 3: Given two sorted integer arrays arr1 and arr2,
// return a new array that combines both of them and is also sorted.
package combine

func Combine(arr1, arr2 []int) []int {
	var i, j, k int
	len1, len2 := len(arr1), len(arr2)

	ans := make([]int, len1+len2)
	for i < len1 && j < len2 {
		if arr1[i] < arr2[j] {
			ans[k] = arr1[i]
			i++
		} else {
			ans[k] = arr2[j]
			j++
		}
		k++
	}

	for i < len1 {
		ans[k] = arr1[i]
		i++
		k++
	}

	for j < len2 {
		ans[k] = arr2[j]
		j++
		k++
	}

	return ans
}
