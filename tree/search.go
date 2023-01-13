package tree

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 深度遍历
// 先序
func preFirstDFS(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Println(root.Val)
	preFirstDFS(root.Left)
	preFirstDFS(root.Right)
}

// 非递归
func preFirstDSF2(root *TreeNode) {

}

// 中序
func midFirstDFS(root *TreeNode) {
	if root == nil {
		return
	}

	midFirstDFS(root.Left)
	fmt.Println(root.Val)
	midFirstDFS(root.Right)
}

// 后序
func lastFirstDFS(root *TreeNode) {
	if root == nil {
		return
	}
	lastFirstDFS(root.Left)
	lastFirstDFS(root.Right)

	fmt.Println(root.Val)
}

func bfs(root *TreeNode) {

}
