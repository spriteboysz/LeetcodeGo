/**
 * Author: Deean
 * Date: 2022-11-09 22:24
 * FileName: algorithm/P0171. Excel 表列序号.go
 * Description:
 */

package main

import "fmt"

func titleToNumber(columnTitle string) int {
	num := 0
	for i, multiple := len(columnTitle)-1, 1; i >= 0; i-- {
		k := columnTitle[i] - 'A' + 1
		num += int(k) * multiple
		multiple *= 26
	}
	return num
}

func main() {
	fmt.Println(titleToNumber("ZY"))
}
