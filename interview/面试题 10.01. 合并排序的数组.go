/**
 * Author: Deean
 * Date: 2022/11/20 22:11
 * FileName: interview/面试题 10.01. 合并排序的数组.go
 * Description:
 */

package main

import (
	"fmt"
	"sort"
)

func merge(A []int, m int, B []int, n int) {
	for i := m; i < m+n; i++ {
		A[i] = B[i-m]
	}
	sort.Ints(A)
	fmt.Println(A)
}

func main() {
	merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)
}
