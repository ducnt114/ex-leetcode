package main

import (
  "fmt"
)

func main() {
  fmt.Println(plusOne([]int{1,2,3}))
}

func plusOne(digits []int) []int {
  needExtend := true
  for _, d := range digits {
    if d < 9 {
      needExtend = false
      break
    }
  }
  var l = len(digits)
  if needExtend {
    l++
  }
  res := make([]int, l)
  var remember = 1
  for i:= len(digits)-1; i>=0; i-- {
    if digits[i]+remember > 9 {
      res[i] = 0
      remember = 1
    } else {
      res[i] = digits[i] + remember
      remember = 0
    }
  }
  if remember > 0 {
    res[0] = 1
  }
  return res
}
