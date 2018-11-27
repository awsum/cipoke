package main

import "testing"

func TestMain(t *testing.T) {
	if 2 == 1 {
		t.Fatal(t, "unfortunate")
	}
}
