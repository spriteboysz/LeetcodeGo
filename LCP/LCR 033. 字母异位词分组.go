/**
 * Author: Deean
 * Date: 2023-10-15 10:12
 * FileName: LCP/LCR 033. 字母异位词分组.go
 * Description:
 */

package main

import (
	"fmt"
	"leetcode/lib"
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	hash := map[string][]string{}
	for _, word := range strs {
		s := []byte(word)
		sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
		hash[string(s)] = append(hash[string(s)], word)
	}

	group := make([][]string, 0, len(hash))
	for _, v := range hash {
		group = append(group, v)
	}
	return group
}

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
	fmt.Println(lib.Str2Array("[1,2,3,4]"))
	fmt.Println(lib.Str2Array2D("[[1,2,3],[4,5,6],[7,8]]"))
}
