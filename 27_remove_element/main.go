package main

import "fmt"

func main() {
	arr1 := []int{3, 2, 2, 3}
	fmt.Println(myFn(arr1, 3))
	fmt.Println(arr1)

	fmt.Println("***********")

	arr2 := []int{0, 1, 2, 2, 3, 0, 4, 2}
	fmt.Println(myFn(arr2, 2))
	fmt.Println(arr2)
}

func myFn(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	var i int = 0
	var j int = len(nums) - 1
	for {
		if i >= j {
			break
		}
		if nums[i] != val {
			i++
			continue
		}
		if nums[j] == val {
			j--
			continue
		}
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
	if nums[i] == val {
		return i
	} else {
		return i + 1
	}
}
