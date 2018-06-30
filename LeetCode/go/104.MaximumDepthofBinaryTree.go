/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
    return helper(0, root, root)
}

func helper(maxDepth int, l *TreeNode, r *TreeNode) int {
    if (l == nil && r == nil) {
        return maxDepth
    } else if (l == nil && r != nil) {
        maxDepth = helper(maxDepth+1, r.Left, r.Right)
    } else if (r == nil && l != nil) {
        maxDepth = helper(maxDepth+1, l.Left, l.Right)
    } else {
        lmaxDepth := helper(maxDepth+1, l.Left, l.Right)
        rmaxDepth := helper(maxDepth+1, r.Left, r.Right)
        if lmaxDepth > rmaxDepth {
            maxDepth = lmaxDepth
        } else {
            maxDepth = rmaxDepth
        }
    }
    
    return maxDepth
}