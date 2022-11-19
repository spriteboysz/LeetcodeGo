/**
 * Author: Deean
 * Date: 2022/11/19 22:35
 * FileName: algorithm/P1071. 字符串的最大公因子.go
 * Description:
 */

package main

import "fmt"

func gcdOfStrings(str1 string, str2 string) string {
	var gcd func(a, b int) int
	gcd = func(a, b int) int {
		if b == 0 {
			return a
		}
		return gcd(b, a%b)
	}
	if str1+str2 != str2+str1 {
		return ""
	}
	if len(str1) >= len(str2) {
		return str1[:gcd(len(str1), len(str2))]
	}
	return str2[:gcd(len(str2), len(str1))]
}

func main() {
	fmt.Println(gcdOfStrings("ABABAB", "ABAB"))
}
