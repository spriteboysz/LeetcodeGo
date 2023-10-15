/**
 * Author: Deean
 * Date: 2022-10-11 23:23
 * FileName: lib/leetcode.go
 * Description:
 */

package lib

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
