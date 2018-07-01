/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
    return helper(root, root)
}

func helper(l *TreeNode, r *TreeNode) bool {
    if (l == nil || r == nil) {
        return l == r
    } else {
        return l.Val == r.Val && helper(l.Left, r.Right) && helper(l.Right, r.Left)
    }
}