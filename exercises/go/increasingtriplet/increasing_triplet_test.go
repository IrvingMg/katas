package increasingtriplet

import (
	"testing"
)

func TestIncreasingTriplet(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		want bool
	}{
		{
			name: "Case 1",
			nums: []int{1, 2, 3, 4, 5},
			want: true,
		},
		{
			name: "Case 2",
			nums: []int{5, 4, 3, 2, 1},
			want: false,
		},
		{
			name: "Case 3",
			nums: []int{2, 1, 5, 0, 4, 6},
			want: true,
		},
		{
			name: "Case 4",
			nums: []int{20, 100, 10, 12, 5, 13},
			want: true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := IncreasingTriplet(tc.nums)
			if got != tc.want {
				t.Errorf("expected: %v, got: %v", tc.want, got)
			}
		})
	}
}
