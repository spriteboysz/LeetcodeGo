/**
 * Author: Deean
 * Date: 2022/11/25 22:36
 * FileName: sword/剑指 Offer 10- I. 斐波那契数列.go
 * Description:
 */

package main

import "fmt"

func fib(n int) int {
	const mod int = 1e9 + 7
	if n <= 1 {
		return n
	}
	p, q, r := 0, 0, 1
	for i := 2; i <= n; i++ {
		p = q
		q = r
		r = (p + q) % mod
	}
	return r
}

func main() {
	fmt.Println(fib(2))
	fmt.Println(fib(5))
}
