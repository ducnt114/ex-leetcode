package main

import (
  "fmt"
)

func main() {
  fmt.Println(generate(5))
  fmt.Println(generate(1))
}

func generate(numRows int) [][]int {
  res := make([][]int, 0)
  for i:=0; i < numRows; i++ {
    if i == 0 {
      res = append(res, []int{1})
      continue
    }
    item := make([]int, i+1)
    for index, _ := range item {
      if index == 0 || index == i {
        item[index] = 1
      } else {
        item[index] = res[i-1][index-1] + res[i-1][index]
      }
    }
    res = append(res, item)
  }
  return res
}
