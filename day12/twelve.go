package day12

import "math/rand"

type RandomizedSet struct {
	Hash  map[int]int
	Array []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		Hash:  map[int]int{},
		Array: []int{},
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.Hash[val]; ok {
		return false
	}
	this.Hash[val] = len(this.Array)
	this.Array = append(this.Array, val)
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	index, ok := this.Hash[val]
	if !ok {
		return false
	}
	l := len(this.Array) - 1
	this.Hash[this.Array[l]] = index
	this.Array[index] = this.Array[l]
	delete(this.Hash, val)
	this.Array = this.Array[:l]
	return true
}

func (this *RandomizedSet) GetRandom() int {
	return this.Array[rand.Intn(len(this.Array))]
}
