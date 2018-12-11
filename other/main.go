package main

import "fmt"

func main() {
	//构建树,原始数据
	root := &TreeNode{
		Val:  1,
		Left: nil,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
	}
	fmt.Println(inorderTraversal(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)      //value
	stack := make([]*TreeNode, 0) //栈，按先序遍历完的节点，存储root节点与左节点
	p := root                     //根
	//循环代替递归
	for p != nil || len(stack) > 0 { //root不为nil或stack值
		for p != nil { //
			stack = append(stack, p) //
			p = p.Left               //取出root的Left节点
		}

		p = stack[len(stack)-1]         //取得末尾节点
		stack = stack[0 : len(stack)-1] //更新栈
		result = append(result, p.Val)  //取出当前节点值,存到数组
		p = p.Right                     //取出root的right节点
	}
	return result
}
