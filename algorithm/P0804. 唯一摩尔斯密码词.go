/**
 * Author: Deean
 * Date: 2022-10-15 22:57
 * FileName: algorithm/P0804. 唯一摩尔斯密码词.go
 * Description:
 */

package main

import "fmt"

func uniqueMorseRepresentations(words []string) int {
	morse := [26]string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
	encoded := map[string]int{}
	for _, word := range words {
		cur := ""
		for _, c := range word {
			cur += morse[c-'a']
		}
		encoded[cur]++
	}
	return len(encoded)
}

func main() {
	words := []string{"gin", "zen", "gig", "msg"}
	fmt.Println(uniqueMorseRepresentations(words))
}
