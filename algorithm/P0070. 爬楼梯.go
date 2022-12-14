/**
 * Author: Deean
 * Date: 2022/12/14 22:35
 * FileName: algorithm/P0070. 爬楼梯.go
 * Description:
 */

package main

import "fmt"

func climbStairs(n int) int {
	p, q, r := 0, 0, 1
	for i := 1; i <= n; i++ {
		p = q
		q = r
		r = p + q
	}
	return r
}

func main() {
	fmt.Println(climbStairs(2))
	fmt.Println(climbStairs(3))
}
