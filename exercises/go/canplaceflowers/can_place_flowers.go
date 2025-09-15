// 605. Can Place Flowers
// You have a long flowerbed in which some of the plots are planted, and some are not.
// However, flowers cannot be planted in adjacent plots.
//
// Given an integer array flowerbed containing 0's and 1's, where 0 means empty and 1
// means not empty, and an integer n, return true if n new flowers can be planted in the
// flowerbed without violating the no-adjacent-flowers rule and false otherwise.
package canplaceflowers

func CanPlaceFlowers(flowerbed []int, n int) bool {
	for i := 0; i < len(flowerbed) && n > 0; i++ {
		if flowerbed[i] == 1 {
			continue
		}

		prevPlotEmpty := i-1 < 0 || flowerbed[i-1] == 0
		nextPlotEmpty := i+1 == len(flowerbed) || flowerbed[i+1] == 0
		if prevPlotEmpty && nextPlotEmpty {
			n--
			flowerbed[i] = 1
		}
	}

	return n == 0
}
