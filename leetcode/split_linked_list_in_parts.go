package main

import (
	"fmt"
)

//https://leetcode-cn.com/problems/split-linked-list-in-parts/
//leetcode 725

func testSplitList() {
	var head = new(ListNode)
	var k = 3
	var arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// var k = 5
	// var arr = []int{1, 2, 3}
	// var k = 3
	// var arr = []int{}
	var cur = head
	for i := 0; i < len(arr); i++ {
		node := new(ListNode)
		node.Val = arr[i]
		cur.Next = node
		cur = cur.Next
	}
	var result = splitListToParts2(head.Next, k)
	fmt.Printf("result length:%d\n", len(result))
	for i := 0; i < len(result); i++ {
		fmt.Printf("==========>node:%d\n", i)
		var node = result[i]
		for node != nil {
			fmt.Printf("node value:%d\n", node.Val)
			node = node.Next
		}
	}
}

//ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func splitListToParts2(head *ListNode, k int) []*ListNode {
	var result []*ListNode
	if head == nil {
		for i := 0; i < k; i++ {
			result = append(result, nil)
		}
		return result
	}
	var length int
	var cur = head
	for cur != nil {
		length++
		cur = cur.Next
	}
	cur = head
	prev := head
	step := length / k
	mod := length % k
	for cur != nil {
		result = append(result, cur)
		for i := 0; i < step; i++ {
			prev = cur
			cur = cur.Next
		}
		if mod > 0 {
			mod--
			prev = cur
			cur = cur.Next
		}
		prev.Next = nil
	}
	if k > len(result) {
		length := k - len(result)
		for i := 0; i < length; i++ {
			result = append(result, nil)
		}
	}
	return result
}

func splitListToParts(head *ListNode, k int) []*ListNode {
	var result []*ListNode
	if head == nil {
		for i := 0; i < k; i++ {
			result = append(result, nil)
		}
		return result
	}
	var length, mod int
	var cur *ListNode
	cur = head
	for cur != nil {
		for i := 0; i < k; i++ {
			cur = cur.Next
			if cur == nil {
				mod = i + 1
				break
			}
		}
		if cur == nil {
			break
		}
		length++
	}
	cur = head
	for cur != nil {
		var node = new(ListNode)
		nodeHead := node

		if length > 0 {
			for i := 0; i < length; i++ {
				node.Next = cur
				node = node.Next
				cur = cur.Next
			}
		}

		if mod > 0 {
			mod--
			node.Next = cur
			node = node.Next
			cur = cur.Next
		}
		node.Next = nil
		result = append(result, nodeHead.Next)
	}

	if k > len(result) {
		length := k - len(result)
		for i := 0; i < length; i++ {
			result = append(result, nil)
		}
	}
	return result
}
