package integers

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	got := Add(1, 2)
	want := 3

	if got != want {
		t.Errorf("got: '%d', want: '%d'", got, want)
	}
}

func ExampleAdd() {
	got := Add(2, 3)
	fmt.Println(got)
	// Output: 5
}
