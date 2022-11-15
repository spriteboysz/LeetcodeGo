/**
 * Author: Deean
 * Date: 2022/11/15 23:37
 * FileName: algorithm/P0929. 独特的电子邮件地址.go
 * Description:
 */

package main

import (
	"fmt"
	"strings"
)

func numUniqueEmails(emails []string) int {
	hash := map[string]bool{}
	for _, email := range emails {
		names := strings.Split(email, "@")
		local := strings.ReplaceAll(strings.Split(names[0], "+")[0], ".", "")
		hash[local+"@"+names[1]] = true
	}
	return len(hash)
}

func main() {
	fmt.Println(numUniqueEmails([]string{"a@leetcode.com", "b@leetcode.com", "c@leetcode.com"}))
}
