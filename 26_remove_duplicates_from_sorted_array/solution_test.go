package main

import (
  "testing"
)

func TestSolution(t *testing.T) {
  int[] nums = [1,1,2] // Input array
  // int[] nums = [0,0,1,1,1,2,2,3,3,4]; // Input array
  int[] expectedNums = [1,2]; // The expected answer with correct length

  int k = removeDuplicates(nums); // Calls your implementation

  if k != expectedNums.length {
    t.Fail()
  }
  for (int i = 0; i < k; i++) {
    if nums[i] != expectedNums[i] {
      t.Fail()
    }
  }
}

func removeDuplicates(nums []int) int {
    
}
