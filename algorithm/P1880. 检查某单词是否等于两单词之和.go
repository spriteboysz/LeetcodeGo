/**
 * Author: Deean
 * Date: 2022-10-17 23:44
 * FileName: algorithm/P1880. 检查某单词是否等于两单词之和.go
 * Description:
 */

package main

import "fmt"

func isSumEqual(firstWord string, secondWord string, targetWord string) bool {
	process := func(s string) (num int) {
		num = 0
		for _, c := range s {
			num = num*10 + int(c-'a')
		}
		return
	}

	return process(firstWord)+process(secondWord) == process(targetWord)
}

func main() {
	fmt.Println(isSumEqual("aaa", "a", "aaaa"))
}
