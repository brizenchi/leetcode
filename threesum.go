package main

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	ans := make([][]int, 0)

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		third := len(nums) - 1
		for j := i + 1; j < third; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			for j < third && nums[i]+nums[j]+nums[third] > 0 {
				third--
			}
			if j == third {
				break
			}
			if nums[i]+nums[j]+nums[third] == 0 {
				ans = append(ans, []int{nums[i], nums[j], nums[third]})
			}
		}
	}
	return ans
}
