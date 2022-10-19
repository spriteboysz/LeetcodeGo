/**
 * Author: Deean
 * Date: 2022-10-19 22:28
 * FileName: algorithm/P1207. 独一无二的出现次数.go
 * Description:
 */

package main

import "fmt"

func uniqueOccurrences(arr []int) bool {
	hash := map[int]int{}
	for _, num := range arr {
		hash[num]++
	}
	hash2 := map[int]bool{}
	for _, v := range hash {
		if hash2[v] == true {
			return false
		} else {
			hash2[v] = true
		}
	}
	return true
}

func main() {
	arr := []int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0}
	fmt.Println(uniqueOccurrences(arr))
	arr = []int{1, 1, 2, 2}
	fmt.Println(uniqueOccurrences(arr))
}
