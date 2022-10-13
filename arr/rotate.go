package arr

// 额外数组
//func rotate(nums []int, k int) {
//	if k > len(nums) {
//		k = k % len(nums)
//	}
//	if k == 0 {
//		return
//	}
//	k = len(nums) - k
//	arr := make([]int, len(nums))
//	for i := 0; i < len(nums); i++ {
//		arr[i] = nums[(k+i)%len(nums)]
//	}
//	copy(nums, arr)
//}

// 多次翻转
func rotate(nums []int, k int) {
	if k >= len(nums) {
		k = k % len(nums)
	}
	if k == 0 {
		return
	}

	ro(nums, 0, len(nums)-1)
	ro(nums, 0, k-1)
	ro(nums, k, len(nums)-1)
}
func ro(nums []int, start, end int) {
	for start < end {
		tmp := nums[start]
		nums[start] = nums[end]
		start++
		nums[end] = tmp
		end--
	}
}

// 环形翻转 k=3
1, 2, 3, 4, 5, 6, 7
1, 2, 3, 1, 5, 6, 7
1, 2, 7, 1, 5, 6, 4
1, 2, 7, 1, 5, 3, 4
1, 6, 7, 1, 5, 3, 4
1, 6, 7, 1, 2, 3, 4
5, 6, 7, 1, 2, 3, 4
func rotate(nums []int, k int) {
	for i := 0; i < k; i++ {
		tmp := nums[i+k]
		nums[i+k] = nums[i]
		nums[i] = tmp
	}
}
