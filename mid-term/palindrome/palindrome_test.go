package main

import "testing"

func Test_Is_palindrome(t *testing.T) {
	s := "abcc"
	result := Is_palindrome(s)
	if result {
		t.Error("Expected false, got ", result, "for ", s)
	}

	s = "abcba"
	result = Is_palindrome(s)
	if result == false {
		t.Error("Expected true, got ", result, "for ", s)
	}

	s = "12345    43 21"
	result = Is_palindrome(s)
	if result == false {
		t.Error("Expected true, got ", result, "for ", s)
	}
}
