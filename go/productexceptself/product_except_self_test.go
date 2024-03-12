package productexceptself

import (
	"slices"
	"testing"
)

func TestProductExceptSelf(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "Case 1",
			nums: []int{1, 2, 3, 4},
			want: []int{24, 12, 8, 6},
		},
		{
			name: "Case 2",
			nums: []int{-1, 1, 0, -3, 3},
			want: []int{0, 0, 9, 0, 0},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := ProductExceptSelf(tc.nums)
			if !slices.Equal(got, tc.want) {
				t.Errorf("expected: %v, got: %v", tc.want, got)
			}
		})
	}
}
