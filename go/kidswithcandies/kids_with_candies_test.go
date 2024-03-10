package kidswithcandies

import (
	"slices"
	"testing"
)

func TestKidsWithCandies(t *testing.T) {
	cases := []struct {
		name         string
		candies      []int
		extraCandies int
		want         []bool
	}{
		{
			name:         "Case 1",
			candies:      []int{2, 3, 5, 1, 3},
			extraCandies: 3,
			want:         []bool{true, true, true, false, true},
		},
		{
			name:         "Case 2",
			candies:      []int{4, 2, 1, 1, 2},
			extraCandies: 1,
			want:         []bool{true, false, false, false, false},
		},
		{
			name:         "Case 3",
			candies:      []int{12, 1, 12},
			extraCandies: 10,
			want:         []bool{true, false, true},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := KidsWithCandies(tc.candies, tc.extraCandies)
			if !slices.Equal(got, tc.want) {
				t.Errorf("expected: %v, got: %v", tc.want, got)
			}
		})
	}
}
