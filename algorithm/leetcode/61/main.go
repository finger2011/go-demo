package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	tail := head
	n := 1
	for tail.Next != nil {
		n++
		tail = tail.Next
	}
	if n == 1 {
		return head
	}
	k %= n
	if k == 0 {
		return head
	}
	k = n - k
	var pre *ListNode
	p := head
	for k > 0 {
		pre = p
		p = pre.Next
		k--
	}
	tail.Next = head
	pre.Next = nil
	return p
}
