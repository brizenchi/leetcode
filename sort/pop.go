package sort

/*
	比较相邻的元素。如果第一个比第二个大，就交换他们两个。
	每次能找出最大的,
	稳定
*/
func pop(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				swap(arr, j, j+1)
			}
		}
	}
	return arr
}
