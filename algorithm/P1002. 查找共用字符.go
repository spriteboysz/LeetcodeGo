/**
 * Author: Deean
 * Date: 2022-11-09 22:27
 * FileName: algorithm/P1002. 查找共用字符.go
 * Description:
 */

package main

import (
	"fmt"
	"math"
)

func commonChars(words []string) []string {
	alphabet := [26]int{}
	for i := range alphabet {
		alphabet[i] = math.MaxInt64
	}
	for _, word := range words {
		frequency := [26]int{}
		for _, c := range word {
			frequency[c-'a'] += 1
		}
		for i, f := range frequency {
			if f < alphabet[i] {
				alphabet[i] = f
			}
		}
	}
	common := []string{}
	for i := byte(0); i < 26; i++ {
		for j := 0; j < alphabet[i]; j++ {
			common = append(common, string(i+'a'))
		}
	}
	return common
}

func main() {
	fmt.Println(commonChars([]string{"bella", "label", "roller"}))
}
