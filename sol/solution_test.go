package sol

import "testing"

func BenchmarkTest(b *testing.B) {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	for idx := 0; idx < b.N; idx++ {
		lengthOfLIS(nums)
	}
}
func Test_lengthOfLIS(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "nums = [10,9,2,5,3,7,101,18]",
			args: args{nums: []int{10, 9, 2, 5, 3, 7, 101, 18}},
			want: 4,
		},
		{
			name: "nums = [0,1,0,3,2,3]",
			args: args{nums: []int{0, 1, 0, 3, 2, 3}},
			want: 4,
		},
		{
			name: "nums = [7,7,7,7,7,7,7]",
			args: args{nums: []int{7, 7, 7, 7, 7, 7, 7}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLIS(tt.args.nums); got != tt.want {
				t.Errorf("lengthOfLIS() = %v, want %v", got, tt.want)
			}
		})
	}
}
