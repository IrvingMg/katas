package gcdofstrings

import "testing"

func TestGcdOfString(t *testing.T) {
	cases := []struct {
		name string
		str1 string
		str2 string
		want string
	}{
		{
			name: "Case 1",
			str1: "ABCABC",
			str2: "ABC",
			want: "ABC",
		},
		{
			name: "Case 2",
			str1: "ABABAB",
			str2: "ABAB",
			want: "AB",
		},
		{
			name: "Case 3",
			str1: "LEET",
			str2: "CODE",
			want: "",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := GcdOfStrings(tc.str1, tc.str2)
			if got != tc.want {
				t.Errorf("expected: %v, got: %v", tc.want, got)
			}
		})
	}
}
