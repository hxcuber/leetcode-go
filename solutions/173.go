package solutions

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	pos  int
	list []int
}

func Constructor(root *TreeNode) BSTIterator {
	return BSTIterator{
		pos:  0,
		list: inOrderTraversal(root),
	}
}

func (this *BSTIterator) Next() int {
	retVal := this.list[this.pos]
	this.pos++
	return retVal
}

func (this *BSTIterator) HasNext() bool {
	return this.pos != len(this.list)
}

func inOrderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	return append(append(inOrderTraversal(root.Left), root.Val), inOrderTraversal(root.Right)...)
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
