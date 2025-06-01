package reversestring

import "testing"

func TestReverseString(t *testing.T) {
	cases := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "hello",
			s:    "hello",
			want: "olleh",
		}, {
			name: "Hannah",
			s:    "Hannah",
			want: "hannaH",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := []byte(tc.s)
			ReverseStringInPlace(got)
			if string(got) != tc.want {
				t.Errorf("want: %v, got: %s", tc.want, got)
			}
		})
	}
}
