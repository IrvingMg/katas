package reversevowels

import "testing"

func TestReverseVowels(t *testing.T) {
	cases := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "Case 1",
			s:    "hello",
			want: "holle",
		},
		{
			name: "Case 2",
			s:    "leetcode",
			want: "leotcede",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := ReverseVowels(tc.s)
			if got != tc.want {
				t.Errorf("expected: %v, got: %v", tc.want, got)
			}
		})
	}
}
