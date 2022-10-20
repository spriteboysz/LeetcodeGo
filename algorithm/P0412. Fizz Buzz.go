/**
 * Author: Deean
 * Date: 2022-10-21 00:00
 * FileName: algorithm/P0412. Fizz Buzz.go
 * Description:
 */

package main

import (
	"fmt"
	"strconv"
)

func fizzBuzz(n int) []string {
	var zz []string
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			zz = append(zz, "FizzBuzz")
		} else if i%3 == 0 {
			zz = append(zz, "Fizz")
		} else if i%5 == 0 {
			zz = append(zz, "Buzz")
		} else {
			zz = append(zz, strconv.Itoa(i))
		}
	}
	return zz
}

func main() {
	fmt.Println(fizzBuzz(15))
}
