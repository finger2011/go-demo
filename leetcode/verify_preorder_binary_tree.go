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