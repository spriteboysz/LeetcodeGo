/**
 * Author: Deean
 * Date: 2022-10-21 00:04
 * FileName: algorithm/P1598. 文件夹操作日志搜集器.go
 * Description:
 */

package main

import "fmt"

func minOperations2(logs []string) int {
	depth := 0
	for _, log := range logs {
		if log == "./" {
			continue
		}
		if log != "../" {
			depth++
		} else if depth > 0 {
			depth--
		}
	}
	return depth
}

func main() {
	logs := []string{"d1/", "d2/", "./", "d3/", "../", "d31/"}
	fmt.Println(minOperations2(logs))
}
