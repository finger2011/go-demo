package main

import (
	"reflect"
	"testing"
)

func createHead(arr []int) *ListNode {
	if len(arr) == 0 {
		var node *ListNode
		return node
	}
	var head = new(ListNode)
	var cur = head
	for i := 0; i < len(arr); i++ {
		node := new(ListNode)
		node.Val = arr[i]
		cur.Next = node
		cur = cur.Next
	}
	return head.Next
}

func Test_splitListToParts(t *testing.T) {
	type args struct {
		head *ListNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want []*ListNode
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				head: createHead([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}),
				k:    3,
			},
			want: []*ListNode{
				createHead([]int{1, 2, 3, 4}),
				createHead([]int{5, 6, 7}),
				createHead([]int{8, 9, 10}),
			},
		},
		{
			name: "test2",
			args: args{
				head: createHead([]int{1, 2, 3}),
				k:    5,
			},
			want: []*ListNode{
				createHead([]int{1}),
				createHead([]int{2}),
				createHead([]int{3}),
				nil,
				nil,
			},
		},
		{
			name: "test3",
			args: args{
				head: createHead([]int{}),
				k:    3,
			},
			want: []*ListNode{nil, nil, nil},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := splitListToParts(tt.args.head, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitListToParts() = %v, want %v", got, tt.want)
			}
		})
	}
}
