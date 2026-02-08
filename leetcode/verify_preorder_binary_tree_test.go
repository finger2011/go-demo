package main

import "testing"

func Test_calculate(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test0",
			args: args{n: 0},
			want: 0,
		},
		{
			name: "test1",
			args: args{n: 1},
			want: 1,
		},
		{
			name: "test2",
			args: args{n: 2},
			want: 1,
		},
		{
			name: "test3",
			args: args{n: 3},
			want: 2,
		},
		{
			name: "test4",
			args: args{n: 4},
			want: 3,
		},
		{
			name: "test5",
			args: args{n: 5},
			want: 5,
		},
		{
			name: "test6",
			args: args{n: 6},
			want: 8,
		},
		{
			name: "test7",
			args: args{n: 7},
			want: 13,
		},
		{
			name: "test8",
			args: args{n: 8},
			want: 21,
		},
		{
			name: "test9",
			args: args{n: 9},
			want: 34,
		},
		{
			name: "test10",
			args: args{n: 10},
			want: 55,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculate(tt.args.n); got != tt.want {
				t.Errorf("calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}
