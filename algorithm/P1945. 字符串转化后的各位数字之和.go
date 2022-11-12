/**
 * Author: Deean
 * Date: 2022/11/12 19:43
 * FileName: algorithm/P1945. 字符串转化后的各位数字之和.go
 * Description:
 */

package main

import "fmt"

func getLucky(s string, k int) int {
	res := 0
	for _, c := range s {
		a := int(c - 'a' + 1)
		if a < 10 {
			res += a
		} else {
			res += a/10 + a%10
		}
	}
	k -= 1
	for k > 0 {
		tmp := 0
		for res > 0 {
			tmp += res % 10
			res = res / 10
		}
		res = tmp
		k--
	}
	return res
}

func main() {
	fmt.Println(getLucky("leetcode", 2))
}
