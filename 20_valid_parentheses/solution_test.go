package main

import "testing"

func TestSolution(t *testing.T) {
	if isValid(`()`) != true {
		t.Fail()
	}
	if isValid(`()[]{}`) != true {
		t.Fail()
	}
	if isValid(`(]`) != false {
		t.Fail()
	}
	if isValid(`([])`) != true {
		t.Fail()
	}
	if isValid(`([][{}])`) != true {
		t.Fail()
	}
}

func isValid(s string) bool {
	stack := []rune{}
	for _, c := range s {
		if c == '(' || c == '[' || c == '{' {
			stack = append(stack, c)
		} else {
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			if (c == ')' && top == '(') || (c == ']' && top == '[') || (c == '}' && top == '{') {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}
