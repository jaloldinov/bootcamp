package main

import "fmt"

type MyHashMap struct {
	data []int
}

func Constructor() MyHashMap {
	return MyHashMap{data: make([]int, 100)}
}

func (this *MyHashMap) Put(key int, value int) {
	this.data[key] = value + 1
}

func (this *MyHashMap) Get(key int) int {
	return this.data[key] - 1
}

func (this *MyHashMap) Remove(key int) {
	this.data[key] = 0
}

func main() {
	hashMap := Constructor()
	hashMap.Put(1, 1)
	hashMap.Put(2, 2)
	fmt.Println(hashMap.Get(1)) // Output: 1
	fmt.Println(hashMap.Get(3)) // Output: -1
	hashMap.Put(2, 1)
	fmt.Println(hashMap.Get(2)) // Output: 1
	hashMap.Remove(2)
	fmt.Println(hashMap.Get(2)) // Output: -1
}
