package main

import "fmt"

func main() {
	//strs := []string{"flower", "flow", "flight"}
	//strs := []string{"flower", "flow", "floght"}
	//strs := []string{"ab", "a"}
	strs := []string{"abab", "aba", ""}
	fmt.Println(longestCommonPrefix(strs))
}

func longestCommonPrefix(strs []string) string {
	if len(strs) <= 1 {
		return strs[0]
	}
	firstString := strs[0]
	if len(firstString) == 0 {
		return ""
	}
	arr := make([]int32, 200)
	for i, ch := range firstString {
		arr[i] = ch
	}

	counter := len(firstString) - 1
	for i := 1; i < len(strs); i++ {
		if len(strs[i]) == 0 {
			return ""
		}
		for j, ch := range strs[i] {
			if j > counter {
				break
			}
			if ch != arr[j] {
				counter = j - 1
				break
			} else {
				if j == len(strs[i])-1 {
					counter = j
				}
			}
		}
	}
	if counter < 0 {
		return ""
	}
	res := ""
	for i := 0; i <= counter; i++ {
		res = fmt.Sprintf("%v%c", res, arr[i])
	}
	return res
}
