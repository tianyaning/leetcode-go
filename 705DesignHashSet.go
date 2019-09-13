package main

type MyHashSet struct {
	mySet []bool
}

/** Initialize your data structure here. */
func Constructor() MyHashSet {
	mySet := make([]bool, 1000001)
	result := MyHashSet{
		mySet: mySet,
	}
	return result
}

func (this *MyHashSet) Add(key int) {
	this.mySet[key] = true

}

func (this *MyHashSet) Remove(key int) {
	this.mySet[key] = false

}

/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
	return this.mySet[key]

}
