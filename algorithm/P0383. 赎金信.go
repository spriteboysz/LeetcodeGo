/**
 * Author: Deean
 * Date: 2022/11/15 23:59
 * FileName: algorithm/P0383. 赎金信.go
 * Description:
 */

package main

import "fmt"

func canConstruct(ransomNote string, magazine string) bool {
	alphabet := make([]int, 26)
	for _, c := range ransomNote {
		alphabet[c-'a']--
	}
	for _, c := range magazine {
		alphabet[c-'a']++
	}
	for i := 0; i < 26; i++ {
		if alphabet[i] < 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(canConstruct("aa", "aab"))
}
