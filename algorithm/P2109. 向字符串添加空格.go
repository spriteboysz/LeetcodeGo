/**
 * Author: Deean
 * Date: 2022/12/10 16:56
 * FileName: algorithm/P2109. 向字符串添加空格.go
 * Description:
 */

package main

import "fmt"

func addSpaces(s string, spaces []int) string {
	ss, index := []byte{}, 0
	for i, c := range s {
		if index < len(spaces) && i == spaces[index] {
			ss = append(ss, ' ')
			index++
		}
		ss = append(ss, byte(c))
	}
	return string(ss)
}

func main() {
	fmt.Println(addSpaces("icodeinpython", []int{1, 5, 7, 9}))
}
