package main

import (
	"testing"
)

func TestCheck(t *testing.T) {
	i := 10
	expected := true
	actual := checkeven(i)
	if actual != expected {
		t.Errorf("For %d expected %v but got %v", i, expected, actual)
	}
}
