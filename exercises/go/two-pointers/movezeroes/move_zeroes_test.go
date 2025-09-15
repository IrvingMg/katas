package movezeroes

import (
	"slices"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "Case 1",
			nums: []int{0, 1, 0, 3, 12},
			want: []int{1, 3, 12, 0, 0},
		},
		{
			name: "Case 2",
			nums: []int{0},
			want: []int{0},
		},
		{
			name: "Case 3",
			nums: []int{1, 0, 1},
			want: []int{1, 1, 0},
		},
		{
			name: "Case 4",
			nums: []int{1, 2},
			want: []int{1, 2},
		},
		{
			name: "Case 5",
			nums: []int{4, 2, 4, 0, 0, 3, 0, 5, 1, 0},
			want: []int{4, 2, 4, 3, 5, 1, 0, 0, 0, 0},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			MoveZeroes(tc.nums)
			if !slices.Equal(tc.nums, tc.want) {
				t.Errorf("expected: %v, got: %v", tc.want, tc.nums)
			}
		})
	}
}
