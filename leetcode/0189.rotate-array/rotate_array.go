package leetcode

func Rotate(nums []int, k int) {
	for i := 0; i < k; i++ {
		t := nums[len(nums)-1]
		for j := len(nums) - 1; j > 0; j-- {
			nums[j] = nums[j-1]
		}
		nums[0] = t
	}
}
