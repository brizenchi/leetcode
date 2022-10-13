package arr

// https://leetcode.cn/problems/remove-duplicates-from-sorted-array/
func RemoveDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	slow := 0
	for fast := 1; fast < len(nums); fast++ {
		if nums[slow] < nums[fast] {
			nums[slow+1] = nums[fast]
			slow++
		}
	}
	return slow + 1
}
