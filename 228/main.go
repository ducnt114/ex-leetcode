package main

import "fmt"

func main() {
	fmt.Println(f([]int{0, 1, 2, 4, 5, 7}))
	fmt.Println(f([]int{0, 2, 3, 4, 6, 8, 9}))
	fmt.Println(f([]int{}))
	fmt.Println(f([]int{123}))
}

func f(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}
	res := []string{}
	var start = nums[0]
	var end = nums[0]
	for i, num := range nums {
		if num > end+1 {
			if start == end {
				res = append(res, fmt.Sprintf("%d", start))
			} else {
				res = append(res, fmt.Sprintf("%d->%d", start, end))
			}
			start = num
		}
		end = num
		if i == len(nums)-1 {
			if start == end {
				res = append(res, fmt.Sprintf("%d", start))
			} else {
				res = append(res, fmt.Sprintf("%d->%d", start, end))
			}
		}
	}
	return res
}
