package reversearray_test

import (
	"katas-go/reversearray"
	"reflect"
	"testing"
)

func TestReverseString(t *testing.T) {
	cases := map[string]struct {
		input []string
		want  []string
	}{
		"hello": {
			input: []string{"hello"},
			want:  []string{"olleh"},
		},
		"hello,world": {
			input: []string{"hello", "world"},
			want:  []string{"olleh", "dlrow"},
		},
		"alpha,beta,charlie,delta,foxtrot": {
			input: []string{"alpha", "beta", "charlie", "delta", "foxtrot"},
			want:  []string{"ahpla", "ateb", "eilrahc", "atled", "tortxof"},
		},
	}

	for name, data := range cases {
		t.Run(name, func(st *testing.T) {
			got := reversearray.ReverseMultiString(data.input)
			if !reflect.DeepEqual(got, data.want) {
				st.Errorf("want: %+v, but got: %+v", data.want, got)
			}
		})
	}
}
