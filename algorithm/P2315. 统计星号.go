/**
 * Author: Deean
 * Date: 2022-10-14 22:26
 * FileName: algorithm/P2315. 统计星号.go
 * Description:
 */

package main

import (
	"fmt"
	"strings"
)

func countAsterisks(s string) int {
	ans := 0
	sp := strings.Split(s, "|")
	for i := 0; i < len(sp); i += 2 {
		ans += strings.Count(sp[i], "*")
	}
	return ans
}

func main() {
	fmt.Println(countAsterisks("yo|uar|e**|b|e***au|tifu|l"))
}
