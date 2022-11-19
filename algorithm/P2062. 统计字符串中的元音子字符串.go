/**
 * Author: Deean
 * Date: 2022/11/19 14:38
 * FileName: algorithm/P2062. 统计字符串中的元音子字符串.go
 * Description:
 */

package main

import "fmt"

func countVowelSubstrings(word string) int {
	cnt, alphabet := 0, [26]int{}
	for i := 0; i < len(word); i++ {
		alphabet = [26]int{}
		for j := i; j < len(word); j++ {
			if word[j] != 'a' && word[j] != 'e' && word[j] != 'i' && word[j] != 'o' && word[j] != 'u' {
				break
			}
			alphabet[word[j]-'a']++
			if alphabet['a'-'a'] > 0 && alphabet['e'-'a'] > 0 && alphabet['i'-'a'] > 0 && alphabet['o'-'a'] > 0 && alphabet['u'-'a'] > 0 {
				cnt++
			}
		}
	}
	return cnt
}

func main() {
	fmt.Println(countVowelSubstrings("cuaieuouac"))
}
