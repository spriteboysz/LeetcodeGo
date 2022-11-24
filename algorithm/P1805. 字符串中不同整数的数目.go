/**
 * Author: Deean
 * Date: 2022/11/23 23:01
 * FileName: algorithm/P1805. 字符串中不同整数的数目.go
 * Description:
 */

package main

import (
	"fmt"
	"regexp"
	"strings"
)

func numDifferentIntegers(word string) int {
	reg := regexp.MustCompile("[0-9]+")
	arr := reg.FindAllString(word, -1)
	hash := make(map[string]bool, 0)
	for _, v := range arr {
		v = strings.TrimLeft(v, "0")
		hash[v] = true
	}
	return len(hash)
}

func main() {
	fmt.Println(numDifferentIntegers("a123bc34d8ef34"))
}
