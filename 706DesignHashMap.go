package main

type MyHashMap struct {
	mySet []int
}

/** Initialize your data structure here. */
func Constructor() MyHashMap {
	mySet := make([]int, 1000001)
	result := MyHashMap{
		mySet: mySet,
	}
	return result

}

/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int) {
	this.mySet[key] = value + 1
}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	return this.mySet[key] - 1
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int) {
	this.mySet[key] = 0
}
