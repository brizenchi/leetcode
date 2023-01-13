package sort

import (
	"reflect"
	"testing"
)

func Test_pop(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "基础",
			args: args{arr: []int{5, 4, 1, 3, 2}},
			want: []int{1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pop(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func pop2(arr []int) []int {
	return nil
}
