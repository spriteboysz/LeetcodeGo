/**
 * Author: Deean
 * Date: 2022/12/4 17:11
 * FileName: algorithm/P1694. 重新格式化电话号码.go
 * Description:
 */

package main

import (
	"fmt"
	"strings"
)

func reformatNumber(number string) string {
	s := strings.ReplaceAll(number, " ", "")
	s = strings.ReplaceAll(s, "-", "")
	reformat, i := []string{}, 0
	for ; i+4 < len(s); i += 3 {
		reformat = append(reformat, s[i:i+3])
	}
	s = s[i:]
	if len(s) < 4 {
		reformat = append(reformat, s)
	} else {
		reformat = append(reformat, s[:2], s[2:])
	}
	return strings.Join(reformat, "-")
}

func main() {
	fmt.Println(reformatNumber("1-23-45 6"))
}
