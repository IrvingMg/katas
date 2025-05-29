package combine

import (
	"slices"
	"testing"
)

func TestCombine(t *testing.T) {
	cases := []struct {
		name string
		arr1 []int
		arr2 []int
		want []int
	}{
		{
			name: "Case 1",
			arr1: []int{1, 4, 7, 20},
			arr2: []int{3, 5, 6},
			want: []int{1, 3, 4, 5, 6, 7, 20},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := Combine(tc.arr1, tc.arr2)
			if !slices.Equal(got, tc.want) {
				t.Errorf("expected: %v, got: %v", tc.want, got)
			}
		})
	}
}
