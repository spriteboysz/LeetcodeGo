/**
 * Author: Deean
 * Date: 2022/11/18 23:21
 * FileName: algorithm/P2269. 找到一个数字的 K 美丽值.go
 * Description:
 */

package main

import (
	"fmt"
	"strconv"
)

func divisorSubstrings(num int, k int) int {
	cnt, s := 0, strconv.Itoa(num)
	for i := 0; i+k <= len(s); i++ {
		n, _ := strconv.Atoi(s[i : i+k])
		if n != 0 && num%n == 0 {
			cnt++
		}
	}
	return cnt
}

func main() {
	fmt.Println(divisorSubstrings(240, 2))
	fmt.Println(divisorSubstrings(430043, 2))
}
