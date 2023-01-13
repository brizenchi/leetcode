package sort

/*
	扑克牌摸牌，每多一个数字，就保证一段序列的正确
*/
func insert(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := i; j > 0; j-- {
			if arr[j] < arr[j-1] {
				swap(arr, j, j-1)
			}
		}
	}
	return arr
}
