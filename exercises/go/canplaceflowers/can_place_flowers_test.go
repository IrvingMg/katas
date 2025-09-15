package canplaceflowers

import "testing"

func TestCanPlaceFlowers(t *testing.T) {
	cases := []struct {
		name      string
		flowerbed []int
		n         int
		want      bool
	}{
		{
			name:      "Case 1",
			flowerbed: []int{1, 0, 0, 0, 1},
			n:         1,
			want:      true,
		},
		{
			name:      "Case 2",
			flowerbed: []int{1, 0, 0, 0, 1},
			n:         2,
			want:      false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := CanPlaceFlowers(tc.flowerbed, tc.n)
			if got != tc.want {
				t.Errorf("expected: %v, got: %v", tc.want, got)
			}
		})
	}
}
