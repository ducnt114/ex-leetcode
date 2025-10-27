type NumArray struct {
    nums []int
}


func Constructor(nums []int) NumArray {
    calculatedNums := make([]int, len(nums))
    for i:=0;i<len(nums);i++ {
        if i == 0 {
            calculatedNums[i] = nums[i]
        } else {
            calculatedNums[i] = calculatedNums[i-1] + nums[i]
        }
    }
    res := NumArray {
        nums: calculatedNums,
    }
    return res
}


func (this *NumArray) SumRange(left int, right int) int {
    if left == 0 {
        return this.nums[right]
    }
    return this.nums[right] - this.nums[left-1]
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
