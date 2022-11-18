/**
 * Author: Deean
 * Date: 2022/11/17 23:28
 * FileName: algorithm/P0202. 快乐数.go
 * Description:
 */

package main

import "fmt"

func isHappy(n int) bool {
	process := func(n int) (sum int) {
		for n > 0 {
			sum += (n % 10) * (n % 10)
			n /= 10
		}
		return
	}

	hash := map[int]bool{}
	for n != 1 && !hash[n] {
		n, hash[n] = process(n), true
	}
	return n == 1
}

func main() {
	fmt.Println(isHappy(19))
	fmt.Println(isHappy(2))
}
