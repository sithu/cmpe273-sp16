package main

func reverse(input string) string {
	// TODO: implement string reversal. trim all whitespace characters from the input string.
	return ""
}

func Is_palindrome(s string) bool {
	return reverse(reverse(s)) == reverse(s)
}
