/**
 * Author: Deean
 * Date: 2022-10-16 20:22
 * FileName: algorithm/P1299. 将每个元素替换为右侧最大元素.go
 * Description:
 */

package main

import "fmt"

func replaceElements(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}
	maximum := arr[len(arr)-1]
	arr[len(arr)-1] = -1
	for i := len(arr) - 2; i >= 0; i-- {
		cur := arr[i]
		arr[i] = maximum
		if cur > maximum {
			maximum = cur
		}
	}
	return arr
}

func main() {
	arr := []int{17, 18, 5, 4, 6, 1}
	fmt.Println(replaceElements(arr))
}
