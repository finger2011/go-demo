package main

import "fmt"

func main() {
	des := [][]int{{20, 15, 1}, {20, 17, 0}, {50, 20, 1}, {50, 80, 0}, {80, 19, 1}}
	fmt.Println("root :", createBinaryTree(des).printTree())
	fmt.Println("root2:", createBinaryTree2(des).printTree())
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (node *TreeNode) printTree() []int {
	var ans []int
	nodes := []*TreeNode{node}
	for len(nodes) > 0 {
		tmp := []*TreeNode{}
		for _, node := range nodes {
			ans = append(ans, node.Val)
			if node.Left != nil {
				tmp = append(tmp, node.Left)
			}
			if node.Right != nil {
				tmp = append(tmp, node.Right)
			}
		}
		nodes = tmp
	}
	return ans
}

func createBinaryTree(descriptions [][]int) *TreeNode {
	n := len(descriptions)
	rootMap := make(map[int]int, n)   // 寻找root
	leafMap := make(map[int][]int, n) // 构建树
	for _, des := range descriptions {
		rootMap[des[1]] = des[0]
		if _, exist := leafMap[des[0]]; !exist {
			leafMap[des[0]] = make([]int, 2)
		}
		switch des[2] {
		case 0:
			leafMap[des[0]][0] = des[1]
		case 1:
			leafMap[des[0]][1] = des[1]
		default:
		}
	}
	rootNum := descriptions[0][0]
	for {
		if _, exist := rootMap[rootNum]; exist {
			rootNum = rootMap[rootNum]
		} else {
			break
		}
	}
	var createNode func(int) *TreeNode
	createNode = func(val int) *TreeNode {
		if val == 0 {
			return nil
		}
		node := &TreeNode{
			Val: val,
		}
		if leaf, exist := leafMap[val]; exist {
			node.Left = createNode(leaf[1])
			node.Right = createNode(leaf[0])
		}
		return node
	}

	return createNode(rootNum)
}

// 利用2次异或来判断，在child和paraent中只出现一次的即为根节点的数值
func createBinaryTree2(descriptions [][]int) *TreeNode {
	n := len(descriptions)
	var rootNum int
	nodeMap := make(map[int]*TreeNode, n+1)
	for _, des := range descriptions {
		if _, exist := nodeMap[des[0]]; !exist {
			nodeMap[des[0]] = &TreeNode{Val: des[0]}
			rootNum ^= des[0]
		}
		if _, exist := nodeMap[des[1]]; !exist {
			nodeMap[des[1]] = &TreeNode{Val: des[1]}
			rootNum ^= des[1]
		}
		switch des[2] {
		case 0:
			nodeMap[des[0]].Right = nodeMap[des[1]]
		case 1:
			nodeMap[des[0]].Left = nodeMap[des[1]]
		default:
		}
		rootNum ^= des[1]
	}

	return nodeMap[rootNum]
}
