/**
 * Author: Deean
 * Date: 2022-10-21 22:25
 * FileName: algorithm/P0509. 斐波那契数.go
 * Description:
 */

package main

import "fmt"

func fib(n int) int {
	if n <= 1 {
		return n
	}
	p, q, r := 0, 0, 1
	for i := 2; i <= n; i++ {
		p = q
		q = r
		r = p + q
	}
	return r
}

func main() {
	fmt.Println(fib(2))
	fmt.Println(fib(3))
	fmt.Println(fib(4))
}
