package sort

/*
	选择排序, 每次找出最大或最小
*/

func selection(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		min := i
		for j := i; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		swap(arr, min, i)
	}
	return arr
}
