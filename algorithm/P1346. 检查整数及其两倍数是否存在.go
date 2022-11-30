/**
 * Author: Deean
 * Date: 2022/11/30 23:28
 * FileName: algorithm/P1346. 检查整数及其两倍数是否存在.go
 * Description:
 */

package main

import "fmt"

func checkIfExist(arr []int) bool {
	hash := map[int]int{}
	for i, v := range arr {
		hash[v] = i
	}
	for i := range arr {
		arr[i] *= 2
	}

	for i, v := range arr {
		if k, ok := hash[v]; i != k && ok {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(checkIfExist([]int{3, 1, 7, 11}))
}
