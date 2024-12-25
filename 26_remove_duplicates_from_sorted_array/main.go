package main

import (
  "fmt"
)

func main() {
  x1 := []int{1,1,2}
  fmt.Println(removeDuplicates(x1))
  fmt.Println(x1)

  x2 := []int{22}
  fmt.Println(removeDuplicates(x2))
  fmt.Println(x2)

  x3 := []int{0,0,1,1,1,2,2,3,3,4}
  fmt.Println(removeDuplicates(x3))
  fmt.Println(x3)
}

func removeDuplicates(nums []int) int {
  min := nums[0]
  lastIndex := 1
  for i:=1; i<len(nums); i++ {
    if nums[i] > min {
      nums[lastIndex] = nums[i]
      min = nums[lastIndex]
      lastIndex++
    }
  }
  return lastIndex
}
