package sortedsquares

import (
	"slices"
	"testing"
)

func TestSortedSquares(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "Case 1",
			nums: []int{-4, -1, 0, 3, 10},
			want: []int{0, 1, 9, 16, 100},
		},
		{
			name: "Case 2",
			nums: []int{-7, -3, 2, 3, 11},
			want: []int{4, 9, 9, 49, 121},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := SortedSquares(tc.nums)
			if !slices.Equal(got, tc.want) {
				t.Errorf("want: %v, got: %v", tc.want, tc.nums)
			}
		})
	}
}
