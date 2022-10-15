/**
 * Author: Deean
 * Date: 2022-10-15 21:23
 * FileName: algorithm/P1281. 整数的各位积和之差.go
 * Description:
 */

package main

import "fmt"

func subtractProductAndSum(n int) int {
	sum, product := 0, 1
	for n > 0 {
		sum += n % 10
		product *= n % 10
		n /= 10
	}
	return product - sum
}

func main() {
	fmt.Println(subtractProductAndSum(4421))
}
