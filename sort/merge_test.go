package sort

import (
	"reflect"
	"testing"
)

func Test_mergeSort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "基础",
			args: args{arr: []int{5, 4, 1, 3, 2}},
			want: []int{1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeSort(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func mergeSort2(nums, cnt []int) []int {
	if len(nums) < 2 {
		return nums
	}
	res := merge2(mergeSort2(nums[0:len(nums)/2], cnt), mergeSort2(nums[len(nums)/2:], cnt), cnt)
	return res
}

func merge2(arrA, arrB, cnt []int) []int {
	var result []int
	var left, right int
	for left < len(arrA) && right < len(arrB) {
		if arrA[left] < arrB[right] {
			result = append(result, arrA[left])
			left++
		} else {
			result = append(result, arrB[right])
			right++
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
