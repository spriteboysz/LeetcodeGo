/**
 * Author: Deean
 * Date: 2022-11-07 23:58
 * FileName: algorithm/P1816. 截断句子.go
 * Description:
 */

package main

import (
	"fmt"
	"strings"
)

func truncateSentence(s string, k int) string {
	words := strings.Split(s, " ")
	return strings.Join(words[:k], " ")
}

func main() {
	fmt.Println(truncateSentence("What is the solution to this problem", 4))
}
