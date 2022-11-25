/**
 * Author: Deean
 * Date: 2022/11/25 21:51
 * FileName: algorithm/P0278. 第一个错误的版本.go
 * Description:
 */

package main

import "fmt"

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func isBadVersion(version int) bool {
	if version >= 3 {
		return true
	}
	return false
}

func firstBadVersion(n int) int {
	left, right := 1, n
	for left <= right {
		mid := left + (right-left)/2
		if isBadVersion(mid) {
			right = mid - 1
			continue
		}
		left = mid + 1
	}
	return left
}

func main() {
	fmt.Println(firstBadVersion(5))
}
