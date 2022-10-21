/**
 * Design a data structure that follows the constraints of a <a href="https://en.wikipedia.org/wiki/Cache_replacement_policies#LRU" target="_blank">Least Recently Used (LRU) cache</a>.
 * Implement the LRUCache class:
 *
 * 	LRUCache(int capacity) Initialize the LRU cache with positive size capacity.
 * 	int get(int key) Return the value of the key if the key exists, otherwise return -1.
 * 	void put(int key, int value) Update the value of the key if the key exists. Otherwise, add the key-value pair to the cache. If the number of keys exceeds the capacity from this operation, evict the least recently used key.
 *
 * Follow up:<br />
 * Could you do get and put in O(1) time complexity?
 *
 * Example 1:
 *
 * Input
 * ["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
 * [[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
 * Output
 * [null, null, null, 1, null, -1, null, -1, 3, 4]
 * Explanation
 * LRUCache lRUCache = new LRUCache(2);
 * lRUCache.put(1, 1); // cache is {1=1}
 * lRUCache.put(2, 2); // cache is {1=1, 2=2}
 * lRUCache.get(1);    // return 1
 * lRUCache.put(3, 3); // LRU key was 2, evicts key 2, cache is {1=1, 3=3}
 * lRUCache.get(2);    // returns -1 (not found)
 * lRUCache.put(4, 4); // LRU key was 1, evicts key 1, cache is {4=4, 3=3}
 * lRUCache.get(1);    // return -1 (not found)
 * lRUCache.get(3);    // return 3
 * lRUCache.get(4);    // return 4
 *
 *
 * Constraints:
 *
 * 	1 <= capacity <= 3000
 * 	0 <= key <= 3000
 * 	0 <= value <= 10^4
 * 	At most 3 * 10^4 calls will be made to get and put.
 *
 */

package leetcode

type List struct {
	head *Node
}

type Node struct {
	key, val   int
	next, prev *Node
}

func NewList() *List {
	node := &Node{}
	node.prev = node
	node.next = node
	return &List{
		head: node,
	}
}

func (l *List) add(node *Node) {
	node.next = l.head.next
	node.prev = l.head
	node.next.prev = node
	l.head.next = node
}

func (l *List) remove(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (l *List) tail() *Node {
	return l.head.prev
}

type LRUCache struct {
	list     *List
	dict     map[int]*Node
	capacity int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		list:     NewList(),
		dict:     make(map[int]*Node),
		capacity: capacity,
	}
}

func (cache *LRUCache) Get(key int) int {
	if node, ok := cache.dict[key]; ok {
		cache.list.remove(node)
		cache.list.add(node)
		return node.val
	}
	return -1
}

func (cache *LRUCache) Put(key int, value int) {
	if node, ok := cache.dict[key]; ok {
		node.key, node.val = key, value
		cache.list.remove(node)
		cache.list.add(node)
	} else {
		// evict least recently used key
		if len(cache.dict) == cache.capacity {
			node := cache.list.tail()
			cache.list.remove(node)
			delete(cache.dict, node.key)
		}
		node := &Node{
			key: key,
			val: value,
		}
		cache.list.add(node)
		cache.dict[key] = node
	}
}
