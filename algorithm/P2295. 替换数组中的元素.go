/**
 * Author: Deean
 * Date: 2022/12/11 22:51
 * FileName: algorithm/P2295. 替换数组中的元素.go
 * Description:
 */

package main

import "fmt"

func arrayChange(nums []int, operations [][]int) []int {
	hash := map[int]int{}
	for i, num := range nums {
		hash[num] = i
	}
	for _, operation := range operations {
		x, y := operation[0], operation[1]
		index := hash[x]
		nums[index] = y
		delete(hash, x)
		hash[y] = index
	}
	return nums
}

func main() {
	fmt.Println(arrayChange([]int{1, 2}, [][]int{{1, 3}, {2, 1}, {3, 2}}))
}
