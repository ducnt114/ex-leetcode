package main

import "fmt"

func main() {
	testCases := make(map[string]int)
	testCases["abcabcbb"] = 3
	testCases["abcdabcbb"] = 4
	testCases["bbbbb"] = 1
	testCases["pwwkew"] = 3
	testCases[""] = 0
	testCases["a"] = 1
	testCases["tmmzuxt"] = 5

	for k, v := range testCases {
		if lengthOfLongestSubstring(k) != v {
			fmt.Printf("[WRONG] input: %v, expected: %d, actual: %d\n", k, v, lengthOfLongestSubstring(k))
		} else {
			fmt.Printf("[CORRECT] input: %v, expected: %d, actual: %d\n", k, v, lengthOfLongestSubstring(k))
		}
	}
}

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	m := make(map[uint8]int)
	var left, right int
	var maxLen int
	for right < len(s) {
		index, existed := m[s[right]]
		if !existed || index < left {
			m[s[right]] = right
			right++
			currentLength := right - left
			if currentLength > maxLen {
				maxLen = currentLength
			}
		} else {
			left = m[s[right]] + 1
			m[s[right]] = right
			right++
		}
	}
	return maxLen
}
