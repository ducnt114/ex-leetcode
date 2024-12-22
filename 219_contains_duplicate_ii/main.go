package main

import (
  "fmt"
)

func main() {
  fmt.Println(containsNearbyDuplicates([]int{1,2,3,1}, 3))
  fmt.Println(containsNearbyDuplicates([]int{1,0,1,1}, 1))
  fmt.Println(containsNearbyDuplicates([]int{1,2,3,1,2,3}, 2))
}

func containsNearbyDuplicates(nums []int, k int) bool {
  m := make(map[int]int)
  for index, num := range nums {
    if lastIndex, existed := m[num]; existed {
      if index-lastIndex <= k {
        return true
      } else {
        m[num] = index
      }
    } else {
      m[num] = index
    }
  }
  return false
}
