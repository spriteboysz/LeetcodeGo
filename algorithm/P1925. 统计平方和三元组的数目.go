/**
 * Author: Deean
 * Date: 2022-10-21 21:13
 * FileName: algorithm/P1925. 统计平方和三元组的数目.go
 * Description:
 */

package main

import "fmt"

func countTriples(n int) int {
	cnt := 0
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			cur := i*i + j*j
			for k := 1; k <= cur/2 && k <= n; k++ {
				if cur == k*k {
					cnt++
					break
				}
			}
		}
	}
	return cnt
}

func main() {
	fmt.Println(countTriples(10))
}
