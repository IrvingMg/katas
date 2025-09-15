package mergealternately

import "testing"

func TestMergeAlternately(t *testing.T) {
	cases := []struct {
		name  string
		word1 string
		word2 string
		want  string
	}{
		{
			name:  "Case 1",
			word1: "abc",
			word2: "pqr",
			want:  "apbqcr",
		},
		{
			name:  "Case 2",
			word1: "ab",
			word2: "pqrs",
			want:  "apbqrs",
		},
		{
			name:  "Case 3",
			word1: "abcd",
			word2: "pq",
			want:  "apbqcd",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := MergeAlternately(tc.word1, tc.word2)
			if got != tc.want {
				t.Errorf("expected: %v, got: %v", tc.want, got)
			}
		})
	}
}
