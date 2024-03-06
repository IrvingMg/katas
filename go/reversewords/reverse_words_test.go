package reversewords

import "testing"

func TestReverseWords(t *testing.T) {
	cases := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "Case 1",
			input: "the sky is blue",
			want:  "blue is sky the",
		},
		{
			name:  "Case 2",
			input: "  hello world  ",
			want:  "world hello",
		},
		{
			name:  "Case 3",
			input: "a good   example",
			want:  "example good a",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := ReverseWords(tc.input)
			if got != tc.want {
				t.Errorf("expected: %v, got: %v", tc.want, got)
			}
		})
	}
}
