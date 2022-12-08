/**
 * Author: Deean
 * Date: 2022/12/7 23:24
 * FileName: algorithm/P0049. 字母异位词分组.go
 * Description:
 */

package main

import "fmt"

func groupAnagrams(strs []string) [][]string {
	hash := map[[26]int][]string{}
	for _, str := range strs {
		key := [26]int{}
		for _, c := range str {
			key[c-'a']++
		}
		value := hash[key]
		value = append(value, str)
		hash[key] = value
	}
	anagrams := [][]string{}
	for _, v := range hash {
		anagrams = append(anagrams, v)
	}
	return anagrams
}

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}
