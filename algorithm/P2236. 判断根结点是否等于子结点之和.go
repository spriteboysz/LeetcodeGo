/**
 * Author: Deean
 * Date: 2022-10-12 23:25
 * FileName: algorithm/P2236. 判断根结点是否等于子结点之和.go
 * Description:
 */

package main

import "fmt"

func checkTree(root *TreeNode) bool {
	return root.Val == root.Left.Val+root.Right.Val
}

func main() {
	fmt.Println()
}
