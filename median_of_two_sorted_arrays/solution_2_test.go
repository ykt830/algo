package median_of_two_sorted_arrays

import "testing"

func Test_findMedianSortedArrays2(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "test_1",
			args: args{
				nums1: []int{},
				nums2: []int{2, 3},
			},
			want: 2.5,
		},
		{
			name: "test_2",
			args: args{
				nums1: []int{1, 2},
				nums2: []int{3, 4},
			},
			want: 2.5,
		},
		{
			name: "test_3",
			args: args{
				nums1: []int{1, 3},
				nums2: []int{2},
			},
			want: 2,
		},
		{
			name: "test_4",
			args: args{
				nums1: []int{1, 3},
				nums2: []int{},
			},
			want: 2,
		},
		{
			name: "test_5",
			args: args{
				nums1: []int{1},
				nums2: []int{},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMedianSortedArrays2(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("findMedianSortedArrays2() = %v, want %v", got, tt.want)
			}
		})
	}
}
