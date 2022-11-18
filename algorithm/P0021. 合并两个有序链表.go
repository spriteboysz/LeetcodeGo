/**
 * Author: Deean
 * Date: 2022/11/18 23:12
 * FileName: algorithm/P0021. 合并两个有序链表.go
 * Description:
 */

package main

import (
	"fmt"
	"leetcode/lib"
)

func mergeTwoLists(list1 *lib.ListNode, list2 *lib.ListNode) *lib.ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list1.Val > list2.Val {
		list1, list2 = list2, list1
	}
	list1.Next = mergeTwoLists(list1.Next, list2)
	return list1
}
func main() {
	fmt.Println(lib.List2String(mergeTwoLists(lib.Str2List("[1,2,3]"), lib.Str2List("[1,3,4]"))))
}
