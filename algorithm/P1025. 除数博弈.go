/**
 * Author: Deean
 * Date: 2022-11-09 23:45
 * FileName: algorithm/P1025. 除数博弈.go
 * Description:
 */

package main

import "fmt"

func divisorGame(n int) bool {
	return n%2 == 0
}

func main() {
	fmt.Println(divisorGame(2))
	fmt.Println(divisorGame(3))
}
