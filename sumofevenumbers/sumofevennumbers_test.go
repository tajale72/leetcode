package main

import (
	"testing"
)

func TestAddSumofEvenNumbers(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	d := SumOfEvenNumbers(arr)

	if d == 0 {
		t.Errorf("Expected 20 but got %v", d)
	}

	if d != 20 {
		t.Errorf("Expected 20 but got %v", d)
	}

}
