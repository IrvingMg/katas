package parenvalidator_test

import (
	"katas-go/parenvalidator"
	"testing"
)

func TestValidate(t *testing.T) {
	cases := map[string]struct {
		input string
		want  bool
	}{
		"(()(()))": {
			input: "(()(()))",
			want:  true,
		},
		"(())(()": {
			input: "(())(()",
			want:  false,
		},
		"()": {
			input: "()",
			want:  true,
		},
		")(": {
			input: ")(",
			want:  false,
		},
		")": {
			input: ")",
			want:  false,
		},
	}

	for name, data := range cases {
		t.Run(name, func(st *testing.T) {
			got := parenvalidator.Validate(data.input)
			if got != data.want {
				st.Errorf("want: %v, but got: %v", data.want, got)
			}
		})
	}
}

func TestValidateRecursively(t *testing.T) {
	cases := map[string]struct {
		input string
		want  bool
	}{
		"(()(()))": {
			input: "(()(()))",
			want:  true,
		},
		"(())(()": {
			input: "(())(()",
			want:  false,
		},
		"()": {
			input: "()",
			want:  true,
		},
		")(": {
			input: ")(",
			want:  false,
		},
		")": {
			input: ")",
			want:  false,
		},
	}

	for name, data := range cases {
		t.Run(name, func(st *testing.T) {
			got := parenvalidator.ValidateRecursively(data.input)
			if got != data.want {
				st.Errorf("want: %v, but got: %v", data.want, got)
			}
		})
	}
}
