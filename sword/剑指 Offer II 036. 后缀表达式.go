/**
 * Author: Deean
 * Date: 2022/12/12 23:18
 * FileName: sword/剑指 Offer II 036. 后缀表达式.go
 * Description:
 */

package main

import (
	"fmt"
	"strconv"
)

func evalRPN(tokens []string) int {
	stack := []int{}
	for _, token := range tokens {
		val, err := strconv.Atoi(token)
		if err == nil {
			stack = append(stack, val)
		} else {
			num1, num2 := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			switch token {
			case "+":
				stack = append(stack, num1+num2)
			case "-":
				stack = append(stack, num1-num2)
			case "*":
				stack = append(stack, num1*num2)
			default:
				stack = append(stack, num1/num2)
			}
		}
	}
	return stack[0]
}

func main() {
	fmt.Println(evalRPN([]string{"4", "13", "5", "/", "+"}))
}
