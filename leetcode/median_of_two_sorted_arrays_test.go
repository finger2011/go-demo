package main

import "testing"

func Test_findMedianSortedArrays1(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				nums1: []int{1, 3},
				nums2: []int{2, 4},
			},
			want: float64(2.5),
		},
		{
			name: "test2",
			args: args{
				nums1: []int{1, 3},
				nums2: []int{2},
			},
			want: float64(2),
		},
		{
			name: "test3",
			args: args{
				nums1: []int{},
				nums2: []int{2, 3},
			},
			want: float64(2.5),
		},
		{
			name: "test4",
			args: args{
				nums1: []int{2},
				nums2: []int{1, 3, 4},
			},
			want: float64(2.5),
		},
		{
			name: "test5",
			args: args{
				nums1: []int{1, 2},
				nums2: []int{-1, 3},
			},
			want: float64(1.5),
		},
		{
			name: "test2",
			args: args{
				nums1: []int{2},
				nums2: []int{1, 3, 4, 5},
			},
			want: float64(3),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMedianSortedArrays1(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("findMedianSortedArrays1() = %v, want %v", got, tt.want)
			}
		})
	}
}
