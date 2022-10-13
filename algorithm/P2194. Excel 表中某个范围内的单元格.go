/**
 * Author: Deean
 * Date: 2022-10-13 23:51
 * FileName: algorithm/P2194. Excel 表中某个范围内的单元格.go
 * Description:
 */

package main

import "fmt"

func cellsInRange(s string) []string {
	cells := make([]string, 0, (s[3]-s[0]+1)*(s[4]-s[1]+1))
	for i := s[0]; i <= s[3]; i++ {
		for j := s[1]; j <= s[4]; j++ {
			cells = append(cells, string(i)+string(j))
		}
	}
	return cells
}

func main() {
	fmt.Println(cellsInRange("A1:F2"))
}
