package sort

import (
	"reflect"
	"testing"
)

func Test_insert(t *testing.T) {
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
			if got := insert2(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insert() = %v, want %v", got, tt.want)
			}
		})
	}
}
func insert2(arr []int) []int {

	return arr
}
