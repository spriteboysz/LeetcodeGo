/**
 * Author: Deean
 * Date: 2022/12/9 23:05
 * FileName: algorithm/P0143. 重排链表.go
 * Description:
 */

package main

import (
	"leetcode/lib"
)

func reorderList(head *lib.ListNode) {
	if head == nil {
		return
	}
	cur, nodes := head, []*lib.ListNode{}
	for cur != nil {
		nodes = append(nodes, cur)
		cur = cur.Next
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
}

func main() {
	reorderList(lib.Str2List("[1,2,3,4,5]"))
}
