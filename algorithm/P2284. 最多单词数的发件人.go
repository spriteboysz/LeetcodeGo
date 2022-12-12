/**
 * Author: Deean
 * Date: 2022/12/11 23:08
 * FileName: algorithm/P2284. 最多单词数的发件人.go
 * Description:
 */

package main

import (
	"fmt"
	"strings"
)

func largestWordCount(messages []string, senders []string) string {
	hash := map[string]int{}
	for i, sender := range senders {
		value := len(strings.Split(messages[i], " "))
		hash[sender] += value
	}
	target := ""
	for k, v := range hash {
		if v > hash[target] || v == hash[target] && k > target {
			target = k
		}
	}
	return target
}

func main() {
	fmt.Println(largestWordCount([]string{"Hello userTwooo", "Hi userThree", "Wonderful day Alice", "Nice day userThree"},
		[]string{"Alice", "userTwo", "userThree", "Alice"}))
}
