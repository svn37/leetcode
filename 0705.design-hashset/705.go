/**
 * Design a HashSet without using any built-in hash table libraries.
 *
 * To be specific, your design should include these functions:
 *
 *
 * 	add(value): Insert a value into the HashSet.
 * 	contains(value) : Return whether the value exists in the HashSet or not.
 * 	remove(value): Remove a value in the HashSet. If the value does not exist in the HashSet, do nothing.
 *
 *
 * <br />
 * Example:
 *
 *
 * MyHashSet hashSet = new MyHashSet();
 * hashSet.add(1);
 * hashSet.add(2);
 * hashSet.contains(1);    // returns true
 * hashSet.contains(3);    // returns false (not found)
 * hashSet.add(2);
 * hashSet.contains(2);    // returns true
 * hashSet.remove(2);
 * hashSet.contains(2);    // returns false (already removed)
 *
 *
 * <br />
 * Note:
 *
 *
 * 	All values will be in the range of [0, 1000000].
 * 	The number of operations will be in the range of [1, 10000].
 * 	Please do not use the built-in HashSet library.
 *
 *
 */

package leetcode

type MyHashSet struct {
	arr [1000000]bool
}

/** Initialize your data structure here. */
func Constructor() MyHashSet {
	arr := [1000000]bool{}
	for i := range arr {
		arr[i] = false
	}
	return MyHashSet{arr}
}

func (this *MyHashSet) Add(key int) {
	this.arr[key] = true
}

func (this *MyHashSet) Remove(key int) {
	this.arr[key] = false
}

/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
	return this.arr[key]
}
