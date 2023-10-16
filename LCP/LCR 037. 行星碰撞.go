/**
 * Author: Deean
 * Date: 2023-10-15 16:07
 * FileName: LCP/LCR 037. 行星碰撞.go
 * Description:
 */

package main

import "fmt"

func asteroidCollision(asteroids []int) []int {
	stack := []int{}
	for _, aster := range asteroids {
		alive := true
		for alive && aster < 0 && len(stack) > 0 && stack[len(stack)-1] > 0 {
			alive = stack[len(stack)-1] < -aster // aster 是否存在
			if stack[len(stack)-1] <= -aster {   // 栈顶小行星爆炸
				stack = stack[:len(stack)-1]
			}
		}
		if alive {
			stack = append(stack, aster)
		}
	}
	return stack
}

func main() {
	fmt.Println(asteroidCollision([]int{5, 10, -5}))
}
