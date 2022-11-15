/**
 * Author: Deean
 * Date: 2022/11/15 23:28
 * FileName: interview/面试题 08.03. 魔术索引.go
 * Description:
 */

package main

import "fmt"

func findMagicIndex(nums []int) int {
	for i, num := range nums {
		if i == num {
			return i
		}
	}
	return -1
}
func main() {
	fmt.Println(findMagicIndex([]int{1, 1, 1}))
}
