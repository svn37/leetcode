/**
 * Implement a data structure supporting the following operations:
 *
 *
 * <ol>
 * Inc(Key) - Inserts a new key <Key> with value 1. Or increments an existing key by 1. Key is guaranteed to be a non-empty string.
 * Dec(Key) - If Key's value is 1, remove it from the data structure. Otherwise decrements an existing key by 1. If the key does not exist, this function does nothing. Key is guaranteed to be a non-empty string.
 * GetMaxKey() - Returns one of the keys with maximal value. If no element exists, return an empty string "".
 * GetMinKey() - Returns one of the keys with minimal value. If no element exists, return an empty string "".
 * </ol>
 *
 *
 *
 * Challenge: Perform all these in O(1) time complexity.
 *
 */

package leetcode

type List struct {
	head, tail *Bucket
}

type Bucket struct {
	key        int
	next, prev *Bucket
	set        map[string]struct{}
}

func NewBucket(key int) *Bucket {
	return &Bucket{
		key: key,
		set: make(map[string]struct{}),
	}
}

func (l *List) append(b *Bucket) *Bucket {
	key := 1
	if b != nil {
		key = b.key + 1
	}

	bucket := NewBucket(key)
	bucket.prev = b

	if b == nil {
		bucket.next = l.head
		l.head = bucket
	} else {
		bucket.next = b.next
		if b.next != nil {
			b.next.prev = bucket
		}
		b.next = bucket
	}

	if l.tail == b {
		l.tail = bucket
	}
	return bucket
}

func (l *List) prepend(b *Bucket) *Bucket {
	bucket := NewBucket(b.key - 1)
	bucket.next = b
	bucket.prev = b.prev
	b.prev = bucket
	if l.head == b {
		l.head = bucket
	}
	return bucket
}

func (l *List) insert(b *Bucket, key string) {
	b.set[key] = struct{}{}
}

func (l *List) delete(b *Bucket, key string) {
	if b == nil {
		return
	}

	delete(b.set, key)

	if len(b.set) == 0 {
		if l.head == b {
			l.head = b.next
		}

		if l.tail == b && b.prev == nil {
			l.tail = l.head
		} else if l.tail == b && b.prev != nil {
			l.tail = b.prev
		}

		if b.prev != nil {
			b.prev.next = b.next
		}

		if b.next != nil {
			b.next.prev = b.prev
		}
	}
}

func (b *Bucket) size() int {
	return len(b.set)
}

func (b *Bucket) randomValue() string {
	for key := range b.set {
		return key
	}
	return ""
}

type AllOne struct {
	list           *List
	keyCountMap    map[string]int
	countBucketMap map[int]*Bucket
}

/** Initialize your data structure here. */
func Constructor() AllOne {
	return AllOne{
		list:           &List{},
		keyCountMap:    make(map[string]int),
		countBucketMap: make(map[int]*Bucket),
	}
}

/** Inserts a new key <Key> with value 1. Or increments an existing key by 1. */
func (this *AllOne) Inc(key string) {
	count := this.keyCountMap[key]
	this.keyCountMap[key]++

	bucket := this.countBucketMap[count]

	nextBucket, ok := this.countBucketMap[count+1]
	if !ok {
		nextBucket = this.list.append(bucket)
		this.countBucketMap[count+1] = nextBucket
	}
	this.list.insert(nextBucket, key)

	if bucket != nil && bucket.size() == 1 {
		delete(this.countBucketMap, count)
	}
	this.list.delete(bucket, key)
}

/** Decrements an existing key by 1. If Key's value is 1, remove it from the data structure. */
func (this *AllOne) Dec(key string) {
	count := this.keyCountMap[key]
	if count < 1 {
		return
	}

	if count == 1 {
		delete(this.keyCountMap, key)
	} else {
		this.keyCountMap[key]--
	}

	bucket := this.countBucketMap[count]

	if count > 1 {
		prevBucket, ok := this.countBucketMap[count-1]
		if !ok {
			prevBucket = this.list.prepend(bucket)
			this.countBucketMap[count-1] = prevBucket
		}
		this.list.insert(prevBucket, key)
	}

	if bucket.size() == 1 {
		delete(this.countBucketMap, count)
	}

	this.list.delete(bucket, key)
}

/** Returns one of the keys with Minimal value. */
func (this *AllOne) GetMinKey() string {
	bucket := this.list.head
	if bucket == nil {
		return ""
	}
	return bucket.randomValue()
}

/** Returns one of the keys with maximal value. */
func (this *AllOne) GetMaxKey() string {
	bucket := this.list.tail
	if bucket == nil {
		return ""
	}
	return bucket.randomValue()
}
