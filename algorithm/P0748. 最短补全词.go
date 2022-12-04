/**
 * Author: Deean
 * Date: 2022/12/4 16:44
 * FileName: algorithm/P0748. 最短补全词.go
 * Description:
 */

package main

import "fmt"

func shortestCompletingWord(licensePlate string, words []string) string {
	cache := map[rune]int{}
	for _, c := range licensePlate {
		if c >= 'a' && c <= 'z' {
			cache[c-'a']++
		}
		if c >= 'A' && c <= 'Z' {
			cache[c-'A']++
		}
	}
	shortest := ""
	for _, word := range words {
		temp := map[rune]int{}
		for _, c := range word {
			temp[c-'a']++
		}
		flag := true
		for k, v := range cache {
			if v > temp[k] {
				flag = false
				break
			}
		}
		if flag && (len(word) < len(shortest) || len(shortest) == 0) {
			shortest = word
		}
	}
	return shortest
}

func main() {
	fmt.Println(shortestCompletingWord("1s3 PSt", []string{"step", "steps", "stripe", "stepple"}))
}
