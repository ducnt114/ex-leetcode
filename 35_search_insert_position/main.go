package main

import "fmt"

func main() {
	//fmt.Println([]int{1, 3, 5, 6}[0:4])
	fmt.Println(solution([]int{3}, 2))                     // 0
	fmt.Println(solution([]int{3}, 4))                     // 1
	fmt.Println(solution([]int{1, 3, 5, 6}, 5))            // 2
	fmt.Println(solution([]int{1, 3, 5, 6}, 2))            // 1
	fmt.Println(solution([]int{1, 3, 5, 6}, 7))            // 4
	fmt.Println(solution([]int{1, 3}, 1))                  // 0
	fmt.Println(solution([]int{1, 3}, 3))                  // 1
	fmt.Println(solution([]int{1, 2, 4, 6, 8, 9, 10}, 10)) // 6
}

func solution(nums []int, target int) int {
	var low, high int
	low = 0
	high = len(nums) - 1
	for {
		if low > high {
			break
		}
		mid := (low + high) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return low
}

func solutionTimeExceed(nums []int, target int) int {
	if nums[0] > target {
		return 0
	}
	if nums[len(nums)-1] < target {
		return len(nums)
	}
	if nums[len(nums)-1] == target {
		return len(nums) - 1
	}
	var firstIndex, lastIndex, middleIndex int
	firstIndex = 0
	lastIndex = len(nums)
	middleIndex = (lastIndex - firstIndex) / 2
	for {
		middleIndex = (lastIndex - firstIndex) / 2
		if lastIndex-firstIndex <= 1 {
			break
		}
		if nums[firstIndex:lastIndex][middleIndex] == target {
			break
		}
		if nums[firstIndex:lastIndex][middleIndex] > target {
			lastIndex = middleIndex
		} else {
			firstIndex = middleIndex
		}
	}
	res := firstIndex + middleIndex
	if nums[res] == target {
		return res
	} else if nums[res] > target {
		return res
	} else {
		return res + 1
	}
}
