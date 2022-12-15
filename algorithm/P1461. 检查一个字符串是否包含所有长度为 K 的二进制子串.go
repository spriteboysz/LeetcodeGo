/**
 * Author: Deean
 * Date: 2022/12/15 23:09
 * FileName: algorithm/P1461. 检查一个字符串是否包含所有长度为 K 的二进制子串.go
 * Description:
 */

package main

import "fmt"

func hasAllCodes(s string, k int) bool {
	hash := map[string]bool{}
	for i := 0; i <= len(s)-k; i++ {
		hash[s[i:i+k]] = true
	}
	return len(hash) == 1<<k
}

func main() {
	fmt.Println(hasAllCodes("00110110", 2))
}
