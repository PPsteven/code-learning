package arr_utils

import (
	"reflect"
	"testing"
)

func TestPopArrByIndex(t *testing.T) {
	type args struct {
		arr   interface{}
		index int
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{"", args{&[]int{1, 2, 3}, 1}, 2},
		{"", args{&[]string{"a", "b", "c"}, 1}, "b"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PopArrByIndex(tt.args.arr, tt.args.index); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PopArrByIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}