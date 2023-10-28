package streamgenerator_test

import (
	"katas-go/streamgenerator"
	"testing"
)

func TestGenerator(t *testing.T) {
	// given
	input := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h'}
	want := input

	// when
	got := streamgenerator.MakeGenerator(input)

	// then
	for _, wantVal := range want {
		gotVal := <-got

		if wantVal != gotVal {
			t.Errorf("want: %v, but got: %v", wantVal, gotVal)
		}
	}

}
