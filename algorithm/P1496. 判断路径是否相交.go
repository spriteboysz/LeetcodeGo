/**
 * Author: Deean
 * Date: 2022/12/14 23:19
 * FileName: algorithm/P1496. 判断路径是否相交.go
 * Description:
 */

package main

import "fmt"

func isPathCrossing(path string) bool {
	hash := map[string]bool{}
	x, y := 0, 0
	hash[fmt.Sprintf("%d#%d", x, y)] = true
	for i := 0; i < len(path); i++ {
		switch path[i] {
		case 'N':
			y += 1
		case 'S':
			y -= 1
		case 'W':
			x -= 1
		case 'E':
			x += 1
		}
		position := fmt.Sprintf("%d#%d", x, y)
		if hash[position] {
			return true
		}
		hash[position] = true
	}
	return false
}

func main() {
	fmt.Println(isPathCrossing("NESWW"))
}
