package bracketsvalidator_test

import (
	"katas-go/bracketsvalidator"
	"testing"
)

func TestValidate(t *testing.T) {
	cases := map[string]struct {
		input string
		want  bool
	}{
		"()": {
			input: "()",
			want:  true,
		},
		"[()]": {
			input: "[()]",
			want:  true,
		},
		"{[()]}": {
			input: "{[()]}",
			want:  true,
		},
		"{{[]()}}}}": {
			input: "{{[]()}}}}",
			want:  false,
		},
		"{[(])}": {
			input: "{[(])}",
			want:  false,
		},
	}

	for name, data := range cases {
		t.Run(name, func(st *testing.T) {
			got := bracketsvalidator.Validate(data.input)
			if got != data.want {
				st.Errorf("want: %v, but got: %v", data.want, got)
			}
		})
	}
}
