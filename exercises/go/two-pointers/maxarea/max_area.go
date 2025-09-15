// 11. Container With Most Water
// You are given an integer array height of length n. There are n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).
// Find two lines that together with the x-axis form a container, such that the container contains the most water.
// Return the maximum amount of water a container can store.
package maxarea

func MaxArea(height []int) int {
	i, j := 0, len(height)-1

	var maxArea int
	for i < j {
		currHeight := min(height[i], height[j])
		currArea := currHeight * (j - i)

		maxArea = max(maxArea, currArea)

		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}

	return maxArea
}
