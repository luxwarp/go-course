package main

import "testing"

func TestGreeter(t *testing.T) {
	if greeter("mikael") != "Hello mikael" {
		t.Error("Expected Hello mikael")
	}
}
