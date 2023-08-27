package sliding_window_maximum

import (
	"reflect"
	"testing"
)

func Test_maxSlidingWindow2(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test_1",
			args: args{
				nums: []int{9, 10, 9, -7, -4, -8, 2, -6},
				k:    5,
			},
			want: []int{10, 10, 9, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSlidingWindow2(tt.args.nums, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxSlidingWindow2() = %v, want %v", got, tt.want)
			}
		})
	}
}
