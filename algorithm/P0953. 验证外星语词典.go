/**
 * Author: Deean
 * Date: 2022/12/13 23:08
 * FileName: algorithm/P0953. 验证外星语词典.go
 * Description:
 */

package main

import "fmt"

func isAlienSorted(words []string, order string) bool {
	alphabet := [26]int{}
	for i, c := range order {
		alphabet[c-'a'] = i
	}
	isSorted := func(w1, w2 string) bool {
		i := 0
		for ; i < len(w1) && i < len(w2); i++ {
			ch1, ch2 := w1[i], w2[i]
			if alphabet[ch1-'a'] < alphabet[ch2-'a'] {
				return true
			}
			if alphabet[ch1-'a'] > alphabet[ch2-'a'] {
				return false
			}
		}
		return i == len(w1)
	}

	for i := 0; i < len(words)-1; i++ {
		if !isSorted(words[i], words[i+1]) {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isAlienSorted([]string{"hello", "leetcode"}, "hlabcdefgijkmnopqrstuvwxyz"))
}
