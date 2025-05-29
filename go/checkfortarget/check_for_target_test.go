package checkfortarget

import "testing"

func TestCheckforTarget(t *testing.T) {
	cases := []struct {
		name   string
		nums   []int
		target int
		want   bool
	}{
		{
			name:   "Case 1",
			nums:   []int{1, 2, 4, 6, 8, 9, 14, 15},
			target: 13,
			want:   true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := CheckforTarget(tc.nums, tc.target)
			if got != tc.want {
				t.Errorf("expected: %v, got: %v", tc.want, got)
			}
		})
	}
}
