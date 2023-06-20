/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
    ret := false
    if p == nil && q == nil {
        return true
    } else if p != nil && q != nil {
        if p.Val == q.Val {
            ret = isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
        }
    }
    return ret
}
