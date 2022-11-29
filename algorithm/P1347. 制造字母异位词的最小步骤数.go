/**
 * Author: Deean
 * Date: 2022/11/29 22:58
 * FileName: algorithm/P1347. 制造字母异位词的最小步骤数.go
 * Description:
 */

package main

import "fmt"

func minSteps(s string, t string) int {
	alphabet := [26]int{}
	for _, c := range s {
		alphabet[c-'a']++
	}
	fmt.Println(alphabet)
	for _, c := range t {
		alphabet[c-'a']--
	}
	fmt.Println(alphabet)
	sum := 0
	for _, num := range alphabet {
		if num < 0 {
			num = -num
		}
		sum += num
	}
	return sum
}

func main() {
	fmt.Println(minSteps("leetcode", "practice"))
}
