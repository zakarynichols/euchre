package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	int1 := 1
	int2 := 2
	want := 3
	sum := add(int1, int2)
	if sum != want {
		t.Fatalf(`add(1, 2) = %d. want sum to equal %d.`, sum, want)
	}
}
