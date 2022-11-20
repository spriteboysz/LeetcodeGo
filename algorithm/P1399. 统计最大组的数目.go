/**
 * Author: Deean
 * Date: 2022/11/20 22:35
 * FileName: algorithm/P1399. 统计最大组的数目.go
 * Description:
 */

package main

import "fmt"

func countLargestGroup(n int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	hash := map[int]int{}
	maximum := 0
	for i := 1; i <= n; i++ {
		key, i0 := 0, i
		for i0 != 0 {
			key += i0 % 10
			i0 /= 10
		}
		hash[key]++
		maximum = max(maximum, hash[key])
	}
	cnt := 0
	for _, v := range hash {
		if v > maximum {
			maximum = v
		}
		if v == maximum {
			cnt++
		}
	}
	return cnt
}

func main() {
	fmt.Println(countLargestGroup(24))
}
