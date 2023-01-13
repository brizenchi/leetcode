package sort

import (
	"reflect"
	"testing"
)

func Test_heapSort(t *testing.T) {
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
			args: args{arr: []int{55, 94, 87, 1, 4, 32, 11, 77, 39, 42, 64, 53, 70, 12, 9}},
			want: []int{1, 4, 9, 11, 12, 32, 39, 42, 53, 55, 64, 70, 77, 87, 94},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := heapSort(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("heapSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
