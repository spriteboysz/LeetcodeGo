/**
 * Author: Deean
 * Date: 2022/11/22 23:07
 * FileName: algorithm/P0744. 寻找比目标字母大的最小字母.go
 * Description:
 */

package main

import "fmt"

func nextGreatestLetter(letters []byte, target byte) byte {
	for _, letter := range letters {
		if letter > target {
			return letter
		}
	}
	return letters[0]
}

func main() {
	fmt.Println(nextGreatestLetter([]byte{'c', 'f', 'j'}, 'd'))
}
