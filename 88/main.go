package main

import (
  "fmt"
)

func main() {
  {
    nums1 := []int{1,2,3,0,0,0}
    nums2 := []int{2,5,6}
    merge(nums1, 3, nums2, 3)
    fmt.Println(nums1)
  }
  {
    nums1 := []int{1}
    nums2 := []int{}
    merge(nums1, 1, nums2, 0)
    fmt.Println(nums1)
  }
  {
    nums1 := []int{1,2,10,0,0,0}
    nums2 := []int{2,5,6}
    merge(nums1, 3, nums2, 3)
    fmt.Println(nums1)
  }
  {
    nums1 := []int{1,2,4,5,6,0}
    nums2 := []int{3}
    merge(nums1, 5, nums2, 1)
    fmt.Println(nums1)
  }
  {
    nums1 := []int{0}
    nums2 := []int{1}
    merge(nums1, 0, nums2, 1)
    fmt.Println(nums1)
  }
}

func merge(nums1 []int, m int, nums2 []int, n int) {
  var i,j,k int
  k = m+n-1
  i = m-1
  j = n-1
  for k >= 0{
    if i < 0 && j < 0 {
      break
    }
    if i < 0 {
      nums1[k] = nums2[j]
      j--
      k--
      continue
    }
    if j < 0 {
      nums1[k] = nums1[i]
      i--
      k--
      continue
    }
    if nums1[i] > nums2[j] {
      nums1[k] = nums1[i]
      i--
      k--
      continue
    }
    if nums1[i] <= nums2[j] {
      nums1[k] = nums2[j]
      j--
    }
    k--
  }
}
