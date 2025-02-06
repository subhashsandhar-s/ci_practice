package main

import "testing"

func TestAdd(t *testing.T) {
	actual := Add(2, 3)
	expected := 10
	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}