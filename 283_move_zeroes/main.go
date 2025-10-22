package main

import "fmt"

func main() {
	nums1 := []int{0, 1, 0, 3, 12}
	fmt.Println(moveZeroes(nums1))

	nums2 := []int{0}
	fmt.Println(moveZeroes(nums2))

	nums3 := []int{0, 0}
	fmt.Println(moveZeroes(nums3))

	nums4 := []int{0, 0, 1}
	fmt.Println(moveZeroes(nums4))

	nums5 := []int{1, 0, 0}
	fmt.Println(moveZeroes(nums5))

	nums6 := []int{1}
	fmt.Println(moveZeroes(nums6))

	nums7 := []int{1, 2}
	fmt.Println(moveZeroes(nums7))

	nums8 := []int{4, 2, 4, 0, 0, 3, 0, 5, 1, 0}
	fmt.Println(moveZeroes(nums8))
}

func moveZeroes(nums []int) []int {
	zeroIndex := 0
	nonZeroIndex := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			if nums[zeroIndex] != 0 {
				zeroIndex = i
			}
		} else {
			nonZeroIndex = i
		}
		//fmt.Println("zeroIndex:", zeroIndex, "nonZeroIndex:", nonZeroIndex)
		if zeroIndex >= len(nums)-1 {
			break
		}
		if zeroIndex < len(nums)-1 && nonZeroIndex <= len(nums)-1 &&
			zeroIndex < nonZeroIndex && nums[zeroIndex] == 0 && nums[nonZeroIndex] != 0 {
			//fmt.Printf("before swap: %v\n", nums)
			nums[nonZeroIndex], nums[zeroIndex] = nums[zeroIndex], nums[nonZeroIndex] // swap
			//fmt.Printf("after swap: %v\n", nums)
			zeroIndex++
		}
	}
	return nums
}
