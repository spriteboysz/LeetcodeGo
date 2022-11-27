/**
 * Author: Deean
 * Date: 2022/11/26 22:36
 * FileName: algorithm/P0706. 设计哈希映射.go
 * Description:
 */

package main

import "fmt"

type MyHashMap struct {
	Map [][]int
}

func Constructor0706() MyHashMap {
	return MyHashMap{
		Map: make([][]int, 0),
	}
}

func (this *MyHashMap) Put(key int, value int) {
	for i, v := range this.Map {
		if v[0] == key {
			this.Map[i][1] = value
			return
		}
	}
	this.Map = append(this.Map, []int{key, value})
	return
}

func (this *MyHashMap) Get(key int) int {
	for _, v := range this.Map {
		if v[0] == key {
			return v[1]
		}
	}
	return -1
}

func (this *MyHashMap) Remove(key int) {
	for i, v := range this.Map {
		if v[0] == key {
			this.Map = append(this.Map[0:i], this.Map[i+1:]...)
			return
		}
	}
}

func main() {
	myHashMap := Constructor0706()
	myHashMap.Put(1, 1)           // myHashMap 现在为 [[1,1]]
	myHashMap.Put(2, 2)           // myHashMap 现在为 [[1,1], [2,2]]
	fmt.Println(myHashMap.Get(1)) // 返回 1 ，myHashMap 现在为 [[1,1], [2,2]]
	fmt.Println(myHashMap.Get(3)) // 返回 -1（未找到），myHashMap 现在为 [[1,1], [2,2]]
	myHashMap.Put(2, 1)           // myHashMap 现在为 [[1,1], [2,1]]（更新已有的值）
	fmt.Println(myHashMap.Get(2)) // 返回 1 ，myHashMap 现在为 [[1,1], [2,1]]
	myHashMap.Remove(2)           // 删除键为 2 的数据，myHashMap 现在为 [[1,1]]
	fmt.Println(myHashMap.Get(2)) // 返回 -1（未找到），myHashMap 现在为 [[1,1]]
}
