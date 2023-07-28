package search_a_2d_matrix_ii

import "testing"

func Test_searchMatrix(t *testing.T) {
	type args struct {
		matrix [][]int
		target int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test_1",
			args: args{
				matrix: [][]int{
					{1, 3, 5, 7, 9},
					{2, 4, 6, 8, 10},
					{11, 13, 15, 17, 19},
					{12, 14, 16, 18, 20},
					{21, 22, 23, 24, 25},
				},
				target: 13,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchMatrix(tt.args.matrix, tt.args.target); got != tt.want {
				t.Errorf("searchMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
