/**
 * Author: Deean
 * Date: 2022-10-16 23:24
 * FileName: algorithm/P1812. 判断国际象棋棋盘中一个格子的颜色.go
 * Description:
 */

package main

import "fmt"

func squareIsWhite(coordinates string) bool {
	return (coordinates[0]+coordinates[1])%2 == 1
}

func main() {
	fmt.Println(squareIsWhite("h3"))
	fmt.Println(squareIsWhite("c7"))
}
