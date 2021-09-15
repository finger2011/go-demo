package main

//verify-preorder-sequence-in-binary-search-tree/submissions
func verifyPreorder(preorder []int) bool {
    if len(preorder) < 2 {
        return true
    }
    var left,right []int
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

func verifyPreorder2(preorder []int) bool {
    if len(preorder) < 2 {
        return true
    }
    var stack = []int{}
    var min = ^int(^uint(0)>>1)
    for i := 0; i < len(preorder); i++ {
        if preorder[i] < min {
            return false
        }
		for len(stack) > 0 && preorder[i] > stack[len(stack) - 1] {
			min = stack[len(stack) - 1]
            stack = stack[0:len(stack) - 1]
		}
        stack = append(stack, preorder[i])
    }
    return true
}