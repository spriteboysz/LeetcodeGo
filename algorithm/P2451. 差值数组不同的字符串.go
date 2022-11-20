/**
 * Author: Deean
 * Date: 2022/11/20 16:46
 * FileName: algorithm/P2451. 差值数组不同的字符串.go
 * Description:
 */

package main

import "fmt"

func oddString(words []string) string {
	hash := map[string][]string{}
	for _, word := range words {
		cur := make([]byte, len(words[0])-1)
		for i := 0; i < len(word)-1; i++ {
			cur[i] = word[i] - word[i+1]
		}
		t := string(cur)
		hash[t] = append(hash[t], word)
	}
	for _, v := range hash {
		if len(v) == 1 {
			return v[0]
		}
	}
	return ""
}

func main() {
	fmt.Println(oddString([]string{"aaa", "bob", "ccc", "ddd"}))
}
