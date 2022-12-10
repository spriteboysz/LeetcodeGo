/**
 * Author: Deean
 * Date: 2022/12/9 23:34
 * FileName: algorithm/P0086. 分隔链表.go
 * Description:
 */

package main

import (
	"fmt"
	"leetcode/lib"
)

func partition(head *lib.ListNode, x int) *lib.ListNode {
	if head == nil {
		return nil
	}
	lt, gt := []*lib.ListNode{}, []*lib.ListNode{}
	for head != nil {
		if head.Val < x {
			lt = append(lt, head)
		} else {
			gt = append(gt, head)
		}
		head = head.Next
	}
	ordered := append(lt, gt...)
	for i := 0; i < len(ordered)-1; i++ {
		ordered[i].Next = ordered[i+1]
	}
	ordered[len(ordered)-1].Next = nil
	return ordered[0]
}

func main() {
	fmt.Println(lib.List2String(partition(lib.Str2List("[1,4,3,2,5,2]"), 3)))
}
