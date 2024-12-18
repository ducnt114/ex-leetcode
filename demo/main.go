package main

import "fmt"

func main() {
	// Test cases
	testCases := []string{
		//"abcabcbb",
		//"bbbbb",
		//"pwwkew",
		//"",
		//"a",
		"au",
	}

	for _, testCase := range testCases {
		fmt.Printf("Input: %s, Longest Substring Length: %d\n", testCase, lengthOfLongestSubstring(testCase))
	}
}

// lengthOfLongestSubstring finds the length of the longest substring without repeating characters.
func lengthOfLongestSubstring(s string) int {
	// Map to store the last seen index of each character
	charIndex := make(map[rune]int)
	maxLength := 0
	start := 0

	for i, char := range s {
		fmt.Println("*******")
		fmt.Println("i: ", i)
		fmt.Println("char: ", char)
		// If the character is already in the map and is part of the current substring
		if lastIndex, found := charIndex[char]; found && lastIndex >= start {
			// Move the start of the substring right after the last occurrence
			start = lastIndex + 1
		}
		fmt.Println("start: ", start)
		// Update the character's last seen index
		charIndex[char] = i
		// Calculate the maximum length
		currentLength := i - start + 1
		fmt.Println("currentLength: ", currentLength)
		if currentLength > maxLength {
			maxLength = currentLength
		}
		fmt.Println("maxLength: ", maxLength)
		fmt.Println("*******")
	}
	return maxLength
}
