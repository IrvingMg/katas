package issubsequence

import "testing"

func TestIsSubsequence(t *testing.T) {
	cases := []struct {
		name string
		s    string
		t    string
		want bool
	}{
		{
			name: "Case 1",
			s:    "abc",
			t:    "ahbgdc",
			want: true,
		},
		{
			name: "Case 2",
			s:    "axc",
			t:    "ahbgdc",
			want: false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := IsSubsequence(tc.s, tc.t)
			if got != tc.want {
				t.Errorf("expected: %v, got: %v", tc.want, got)
			}
		})
	}
}
