package main

import "testing"

func Add(a, b int) int {
	return a + b
}


func TestAdd(t *testing.T) {
	result := Add(1, 2)
	expected := 3

	if result != expected {
		t.Errorf("Test failed") 
	}

	

	
	result2 := Add(2, 2)
	expected2 := 5
	if result2 != expected2 {
		t.Fatal("Critical failure in test")
	}
}
