/**
 * Author: Deean
 * Date: 2022-10-16 00:03
 * FileName: algorithm/P2325. 解密消息.go
 * Description:
 */

package main

import (
	"fmt"
	"strings"
)

func decodeMessage(key string, message string) string {
	hashMap := make(map[rune]rune)
	for _, c := range key {
		if _, hash := hashMap[c]; c != ' ' && !hash {
			hashMap[c] = rune('a' + len(hashMap))
		}
	}
	var build strings.Builder
	for _, c := range message {
		if c == ' ' {
			build.WriteRune(c)
		} else {
			build.WriteRune(hashMap[c])
		}
	}
	return build.String()
}

func main() {
	ans := decodeMessage("eljuxhpwnyrdgtqkviszcfmabo", "zwx hnfx lqantp mnoeius ycgk vcnjrdb")
	fmt.Println(ans)
}
