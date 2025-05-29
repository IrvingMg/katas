package palindrome

import "testing"

func TestPalindrome(t *testing.T) {
	cases := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "Case 1",
			s:    "abcdcba",
			want: true,
		},
		{
			name: "Case 2",
			s:    "racecar",
			want: true,
		},
		{
			name: "Case 3",
			s:    "palindrome",
			want: false,
		},
		{
			name: "Case 4",
			s:    "aceba",
			want: false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := IsPalindrome(tc.s)
			if got != tc.want {
				t.Errorf("expected: %v, got: %v", tc.want, got)
			}
		})
	}
}
