/**
 * Author: Deean
 * Date: 2022/12/9 22:43
 * FileName: algorithm/P0117. 填充每个节点的下一个右侧节点指针 II.go
 * Description:
 */

package main

import "fmt"

// Node Definition for a Node.
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	queue := []*Node{root}
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[i]
			if i+1 < size {
				node.Next = queue[i+1]
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		queue = queue[size:]
	}
	return root
}

func main() {
	fmt.Println(connect(nil))
}
