package main

// verify-preorder-sequence-in-binary-search-tree/submissions
func verifyPreorder(preorder []int) bool {
	if len(preorder) < 2 {
		return true
	}
	var left, right []int
	for i := 1; i < len(preorder); i++ {
		if preorder[i] < preorder[0] {
			left = append(left, preorder[i])
		} else {
			right = preorder[i:]
			break
		}
	}
	for i := 1; i < len(right); i++ {
		if right[i] < preorder[0] {
			return false
		}
	}
	var check bool
	if len(left) > 1 {
		check = verifyPreorder(left)
		if check == false {
			return false
		}
	}
	if len(right) > 1 {
		check = verifyPreorder(right)
		if check == false {
			return false
		}
	}
	return true
}

// 局部递减，整体递增
// 5 2 1 3 8 7 9
// 5 2 1递减，向下向左遍历
// 5 8 9 递增，向下向右遍历
func verifyPreorder2(preorder []int) bool {
	if len(preorder) < 2 {
		return true
	}
	var stack = []int{}
	var min = ^int(^uint(0) >> 1)
	for i := 0; i < len(preorder); i++ {
		if preorder[i] < min {
			return false
		}
		for len(stack) > 0 && preorder[i] > stack[len(stack)-1] {
			min = stack[len(stack)-1]
			stack = stack[0 : len(stack)-1]
		}
		stack = append(stack, preorder[i])
	}
	return true
}

// calculate returns the n-th Fibonacci number.
// This is a simple recursive implementation where
// calculate(0) = 0, calculate(1) = 1, and for n > 1,
// calculate(n) = calculate(n-1) + calculate(n-2).
func calculate(n int) int {
	if n <= 1 {
		return n
	}
	return calculate(n-1) + calculate(n-2)
}
