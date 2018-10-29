package gopractice

// https://leetcode.com/problems/flatten-binary-tree-to-linked-list

/**
 * Definition for a binary tree node.
 */
// type TreeNode struct {
//     Val int
//     Left *TreeNode
//     Right *TreeNode
// }
 
func flatten(root *TreeNode)  {
	/**
	 * It's like a pre-order traverse. But if process via pre-order, the left/right
	 * children links would be broken.
	 * So, should process it via reverse of pre-order.
	 */
	
	var pre *TreeNode = nil

	var _helper func(node *TreeNode)
	_helper = func(current *TreeNode) {
		if current == nil {
			return
		}
		_helper(current.Right)
		_helper(current.Left)
		current.Right = pre
		current.Left = nil
		pre = current
	}

	_helper(root)
}
