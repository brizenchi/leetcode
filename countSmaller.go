package main

//func countSmaller(nums []int) []int {
//
//}

func mergeSort3(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	return merge3(mergeSort3(nums[0:len(nums)/2]), mergeSort3(nums[len(nums)/2:]))
}

func merge3(arrA, arrB []int) []int {
	var result []int
	var left, right, i int
	for left < len(arrA) && right < len(arrB) {
		if arrA[left] < arrB[right] {
			result[i] = arrA[left]
			left++
			i++
		} else {
			result[i] = arrB[right]
			right++
			i++
		}
	}
	if left < len(arrA) {
		result = append(result, arrA[left:]...)
	}
	if right < len(arrB) {
		result = append(result, arrB[right:]...)
	}
	return result
}
