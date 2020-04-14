package messages

import "testing"

func TestGreet(t *testing.T) {
	got := Greet("Gopher")
	expect := "Hello, Gopher!\n"

	if got != expect {
		t.Errorf("Did not get expected results. Wanted %q, got: %q\n", expected, got)
	}
}
func TestDepart(t *testing.T) {
	got := Greet("Gopher")
	expect := "Bye, gopher!\n"

	if got != expect {
		t.Errorf("Did not get expected results. Wanted %q, got: %q\n", expected, got)
	}
}
