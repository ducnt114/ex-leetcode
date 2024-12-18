package main

import "fmt"

func main() {
	// Test cases
	testCases := []struct {
		s1 string
		s2 string
	}{
		{"sadbutsad", "sad"},
		{"asadbutsad", "sad"},
		{"leetcode", "leeto"},
		{"a", "a"},
		{"au", "a"},
	}

	for _, testCase := range testCases {
		fmt.Println(solution(testCase.s1, testCase.s2))
	}
}

func solution(haystack string, needle string) int {
	for i := 0; i <= len(haystack)-len(needle); i++ {
		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}
	return -1
}
