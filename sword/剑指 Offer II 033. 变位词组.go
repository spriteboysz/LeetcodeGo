/**
 * Author: Deean
 * Date: 2022/11/29 23:08
 * FileName: sword/剑指 Offer II 033. 变位词组.go
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
