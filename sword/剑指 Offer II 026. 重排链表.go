/**
 * Author: Deean
 * Date: 2022/12/16 22:33
 * FileName: sword/剑指 Offer II 026. 重排链表.go
 * Description:
 */

package main

import (
	"fmt"
	"leetcode/lib"
)

func reorderList(head *lib.ListNode) {
	if head == nil {
		return
	}
	nodes := []*lib.ListNode{}
	for node := head; node != nil; node = node.Next {
		nodes = append(nodes, node)
	}
	i, j := 0, len(nodes)-1
	for i < j {
		nodes[i].Next = nodes[j]
		i++
		if i == j {
			break
		}
		nodes[j].Next = nodes[i]
		j--
	}
	nodes[i].Next = nil
	fmt.Println(lib.List2String(nodes[0]))
}

func main() {
	reorderList(lib.Str2List("[1,2,3,4]"))
}
