package main

import "fmt"

func main() {
	nums := []int{1, 0, 1, 0, 1, 0, 1}
	root := ArrayToTree(nums)
	fmt.Println("sumRootToLeaf(", nums, ") ====> ", sumRootToLeaf(root))

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ArrayToTree 主函数：将数组转换为二叉树
func ArrayToTree(arr []int) *TreeNode {
	if len(arr) == 0 {
		return nil
	}
	// 从索引 0 (根节点) 开始构建
	return buildTreeRecursive(arr, 0)
}

// buildTreeRecursive 递归辅助函数
// 根据 index 2*i+1 和 2*i+2 的规律构建树
func buildTreeRecursive(arr []int, index int) *TreeNode {
	// 边界条件：如果索引超出数组长度，说明该位置没有节点
	if index >= len(arr) {
		return nil
	}

	// 创建当前节点
	node := &TreeNode{Val: arr[index]}

	// 递归构建左子树 (索引 2*i + 1)
	node.Left = buildTreeRecursive(arr, 2*index+1)

	// 递归构建右子树 (索引 2*i + 2)
	node.Right = buildTreeRecursive(arr, 2*index+2)

	return node
}

func sumRootToLeaf(root *TreeNode) int {
	return sumLeaf(root, 0)
}

func sumLeaf(root *TreeNode, routine int) int {
	if root.Left == nil && root.Right == nil {
		return int(routine<<1) + root.Val
	}
	var sum int
	if root.Left != nil {
		sum += sumLeaf(root.Left, int(routine<<1)+root.Val)
	}
	if root.Right != nil {
		sum += sumLeaf(root.Right, int(routine<<1)+root.Val)
	}

	return sum
}
