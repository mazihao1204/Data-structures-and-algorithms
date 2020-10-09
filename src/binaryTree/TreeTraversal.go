package binaryTree

//二叉树的遍历
//递归解法：
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func TreeTraversal(root *TreeNode) []int{
	res := []int{}
	if root != nil{
		traversal(root,res)
	}
	return res
}

func traversal(root *TreeNode,res []int){
	res = append(res,root.Val)
	traversal(root.Left,res)
	traversal(root.Right,res)
}

//前序遍历
//迭代解法
func preOrderTraversal(root *TreeNode) []int{
	if root == nil{
		return []int{}
	}
	stack,res := []*TreeNode{},[]int{}
	stack = append(stack,root)
	for len(stack) != 0{
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node != nil{
			res = append(res,node.Val)
		}
		if node.Right != nil{
			stack = append(stack,node.Right)
		}
		if node.Left != nil{
			stack = append(stack,node.Left)
		}
	}
	return res
}

//中序遍历
func inorderTraversal(root *TreeNode) []int{

}