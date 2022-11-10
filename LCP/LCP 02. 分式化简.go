/**
 * Author: Deean
 * Date: 2022-11-10 22:50
 * FileName: LCP/LCP 02. 分式化简.go
 * Description:
 */

package main

import "fmt"

func fraction(cont []int) []int {
	size := len(cont)
	fz, fm := 1, cont[size-1]
	for i := size - 2; i >= 0; i-- {
		fz += cont[i] * fm
		fz, fm = fm, fz
	}
	return []int{fm, fz}
}

func main() {
	fmt.Println(fraction([]int{3, 2, 0, 2}))
}
