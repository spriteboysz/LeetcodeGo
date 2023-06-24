/**
 * Author: Deean
 * Date: 2023-06-24 20:53
 * FileName: algorithm/P2678. 老人的数目.go
 * Description:
 */

package main

import (
	"fmt"
)

func countSeniors(details []string) int {
	cnt := 0
	for _, detail := range details {
		if detail[11]&15*10+detail[12]&15 > 60 {
			cnt += 1
		}
	}
	return cnt
}

func main() {
	fmt.Println(countSeniors([]string{"7868190130M7522", "5303914400F9211", "9273338290F4010"}))
}
