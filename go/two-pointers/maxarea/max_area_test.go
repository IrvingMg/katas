package maxarea

import "testing"

func TestMaxArea(t *testing.T) {
	cases := []struct {
		name   string
		height []int
		want   int
	}{
		{
			name:   "Case 1",
			height: []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			want:   49,
		},
		{
			name:   "Case 2",
			height: []int{1, 1},
			want:   1,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := MaxArea(tc.height)
			if got != tc.want {
				t.Errorf("got: %v, want: %v", got, tc.want)
			}
		})
	}
}
