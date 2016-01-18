package util

import "testing"

func TestReverseHelloWorld(t *testing.T) {

	input := "Hello World !"
	expected := "! dlroW olleH"
	r := Reverse(input)
	if r != expected {
		t.Errorf("Reserve failed: %s instead of %s for input %s", r, expected, input)
	}
}
