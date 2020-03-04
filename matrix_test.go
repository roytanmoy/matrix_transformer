package main

import "testing"

func TestEcho(t *testing.T) {
	records := [][]string{
		{"1", "2", "3"},
		{"4", "5", "6"},
		{"7", "8", "9"},
	}
	expected := "1,2,3\n4,5,6\n7,8,9"
	result := Echo(records)
	if result != expected {
		t.Errorf("Echo was incorrect, got: %s, expected: %s.", result, expected)
	}

	records = [][]string{
		{"1", "2", "3"},
		{"4", "5", "6"},
		{"X", "Y", "Z"},
	}
	expected = "1,2,3\n4,5,6\nX,Y,Z"
	result = Echo(records)
	if result != expected {
		t.Errorf("Echo was incorrect, got: %s, expected: %s.", result, expected)
	}
}

func TestInvert(t *testing.T) {
	records := [][]string{
		{"1", "2", "3"},
		{"4", "5", "6"},
		{"7", "8", "9"},
	}
	expected := "1,4,7\n2,5,8\n3,6,9"
	result := Invert(records)
	if result != expected {
		t.Errorf("Invert was incorrect, got: %s, expected: %s.", result, expected)
	}

	records = [][]string{
		{"1", "2", "3"},
		{"4", "5", "6"},
		{"X", "Y", "Z"},
	}
	expected = "1,4,X\n2,5,Y\n3,6,Z"
	result = Invert(records)
	if result != expected {
		t.Errorf("Invert was incorrect, got: %s, expected: %s.", result, expected)
	}
}
