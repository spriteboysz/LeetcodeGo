/**
 * Author: Deean
 * Date: 2022/12/16 22:10
 * FileName: algorithm/P0859. 亲密字符串.go
 * Description:
 */

package main

import "fmt"

func buddyStrings(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}

	if s == goal {
		seen := [26]bool{}
		for _, ch := range s {
			if seen[ch-'a'] {
				return true
			}
			seen[ch-'a'] = true
		}
		return false
	}

	first, second := -1, -1
	for i := range s {
		if s[i] != goal[i] {
			if first == -1 {
				first = i
			} else if second == -1 {
				second = i
			} else {
				return false
			}
		}
	}
	return second != -1 && s[first] == goal[second] && s[second] == goal[first]
}

func main() {
	fmt.Println(buddyStrings("ab", "ba"))
}
