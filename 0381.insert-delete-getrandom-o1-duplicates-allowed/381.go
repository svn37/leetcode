/**
 * Design a data structure that supports all following operations in average O(1) time.
 * Note: Duplicate elements are allowed.
 *
 * <ol>
 * insert(val): Inserts an item val to the collection.
 * remove(val): Removes an item val from the collection if present.
 * getRandom: Returns a random element from current collection of elements. The probability of each element being returned is linearly related to the number of same value the collection contains.
 * </ol>
 *
 *
 * Example:
 *
 * // Init an empty collection.
 * RandomizedCollection collection = new RandomizedCollection();
 *
 * // Inserts 1 to the collection. Returns true as the collection did not contain 1.
 * collection.insert(1);
 *
 * // Inserts another 1 to the collection. Returns false as the collection contained 1. Collection now contains [1,1].
 * collection.insert(1);
 *
 * // Inserts 2 to the collection, returns true. Collection now contains [1,1,2].
 * collection.insert(2);
 *
 * // getRandom should return 1 with the probability 2/3, and returns 2 with the probability 1/3.
 * collection.getRandom();
 *
 * // Removes 1 from the collection, returns true. Collection now contains [1,2].
 * collection.remove(1);
 *
 * // getRandom should return 1 and 2 both equally likely.
 * collection.getRandom();
 *
 *
 */

package leetcode

import "math/rand"

type RandomizedCollection struct {
	set  map[int]map[int]struct{}
	list []int
}

/** Initialize your data structure here. */
func Constructor() RandomizedCollection {
	return RandomizedCollection{
		set: make(map[int]map[int]struct{}),
	}
}

/** Inserts a value to the collection. Returns true if the collection did not already contain the specified element. */
func (this *RandomizedCollection) Insert(val int) bool {
	idx := len(this.list)
	this.list = append(this.list, val)

	indexes, ok := this.set[val]
	if !ok {
		indexes = make(map[int]struct{})
		this.set[val] = indexes
	}
	indexes[idx] = struct{}{}
	return !ok
}

/** Removes a value from the collection. Returns true if the collection contained the specified element. */
func (this *RandomizedCollection) Remove(val int) bool {
	indexes, ok := this.set[val]
	if !ok {
		return false
	}

	var idx int
	for i := range indexes {
		idx = i
		break
	}
	delete(indexes, idx)

	if len(indexes) == 0 {
		delete(this.set, val)
	}

	if idx != len(this.list)-1 {
		this.list[idx], this.list[len(this.list)-1] = this.list[len(this.list)-1], this.list[idx]
		// change the index of the swapped element
		indexes = this.set[this.list[idx]]
		delete(indexes, len(this.list)-1)
		indexes[idx] = struct{}{}
	}

	this.list = this.list[:len(this.list)-1]

	return true
}

/** Get a random element from the collection. */
func (this *RandomizedCollection) GetRandom() int {
	random := rand.Intn(len(this.list))
	return this.list[random]
}
