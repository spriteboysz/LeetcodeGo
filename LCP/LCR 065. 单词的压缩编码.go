/**
 * Author: Deean
 * Date: 2023-10-15 22:46
 * FileName: LCP/LCR 065. 单词的压缩编码.go
 * Description:
 */

package main

import "fmt"

func minimumLengthEncoding(words []string) int {
	hash := map[string]bool{}
	for _, word := range words {
		hash[word] = true
	}
	for _, word := range words {
		for i := 1; i < len(word); i++ {
			if hash[word[i:]] {
				delete(hash, word[i:])
			}
		}
	}
	sum := 0
	for word := range hash {
		sum += len(word) + 1
	}
	return sum
}

func main() {
	fmt.Println(minimumLengthEncoding([]string{"time", "me", "bell"}))
}
