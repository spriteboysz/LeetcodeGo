/**
 * Author: Deean
 * Date: 2022-11-09 00:00
 * FileName: algorithm/P0451. 根据字符出现频率排序.go
 * Description:
 */

package main

import (
	"fmt"
	"sort"
)

func frequencySort(s string) string {
	sByte := []byte(s)
	hash := make(map[byte]int)
	for _, c := range sByte {
		hash[c]++
	}
	sort.Slice(sByte, func(i, j int) bool {
		if hash[sByte[i]] != hash[sByte[j]] {
			return hash[sByte[i]] > hash[sByte[j]]
		} else {
			return sByte[i] < sByte[j]
		}
	})
	return string(sByte)
}

func main() {
	fmt.Println(frequencySort("cccaaa"))
}
