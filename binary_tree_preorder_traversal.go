package gopractice

// https://leetcode.com/problems/binary-tree-preorder-traversal


// Definition for a binary tree node.
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	result := []int{}
	var _helper func(*TreeNode)
	_helper = func(root *TreeNode) {
		if root == nil {
			return
		}
		result = append(result, root.Val)
		_helper(root.Left)
		_helper(root.Right)
	}

	_helper(root)
	return result
}

func preorderTraversal2(root *TreeNode) []int {
	
	result := []int{}
	if root == nil {
		return result
	}
	stack := []*TreeNode{}
	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			result = append(result, root.Val)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		root = root.Right
	}
	return result
}