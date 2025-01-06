package main

import (
  "fmt"
)

func main() {
  fmt.Println(isHappy(19))
  fmt.Println(isHappy(2))
}

func isHappy(n int) bool {
  	// Helper function to calculate the sum of the squares of digits
	sumOfSquares := func(num int) int {
		sum := 0
		for num > 0 {
			digit := num % 10
			sum += digit * digit
			num /= 10
		}
		return sum
	}

	seen := make(map[int]bool) // To track numbers we've seen
	for n != 1 && !seen[n] {
		seen[n] = true
		n = sumOfSquares(n)
	}
	return n == 1
}
