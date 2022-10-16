/**
 * Author: Deean
 * Date: 2022-10-16 10:35
 * FileName: algorithm/P1832. 判断句子是否为全字母句.go
 * Description:
 */

package main

import "fmt"

func checkIfPangram(sentence string) bool {
	var alphabet [26]bool
	for _, c := range sentence {
		alphabet[c-'a'] = true
	}
	for _, v := range alphabet {
		if v == false {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(checkIfPangram("thequickbrownfoxjumpsoverthelazydog"))
	fmt.Println(checkIfPangram("leetcode"))
}
