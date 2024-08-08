package solutions

func balanceBST(root *TreeNode) *TreeNode {
	sorted := inOrderTraversal(root)
	b := (*TreeNode)(nil)

	b = insertSpecial(b, sorted)
	return b
}

func insert(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{
			Val:   val,
			Left:  nil,
			Right: nil,
		}
	}

	if val < root.Val {
		root.Left = insert(root.Left, val)
	} else {
		root.Right = insert(root.Right, val)
	}
	return root
}

func insertSpecial(root *TreeNode, sorted []int) *TreeNode {
	if sorted == nil || len(sorted) == 0 {
		return root
	}

	half := len(sorted) / 2

	root = insert(root, sorted[half])

	if half != 0 {
		root = insertSpecial(root, sorted[:half])
		root = insertSpecial(root, sorted[half+1:])
	}
	return root
}
