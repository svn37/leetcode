/**
 * Design a simplified version of Twitter where users can post tweets, follow/unfollow another user and is able to see the 10 most recent tweets in the user's news feed. Your design should support the following methods:
 *
 *
 * <ol>
 * postTweet(userId, tweetId): Compose a new tweet.
 * getNewsFeed(userId): Retrieve the 10 most recent tweet ids in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user herself. Tweets must be ordered from most recent to least recent.
 * follow(followerId, followeeId): Follower follows a followee.
 * unfollow(followerId, followeeId): Follower unfollows a followee.
 * </ol>
 *
 *
 * Example:
 *
 * Twitter twitter = new Twitter();
 *
 * // User 1 posts a new tweet (id = 5).
 * twitter.postTweet(1, 5);
 *
 * // User 1's news feed should return a list with 1 tweet id -> [5].
 * twitter.getNewsFeed(1);
 *
 * // User 1 follows user 2.
 * twitter.follow(1, 2);
 *
 * // User 2 posts a new tweet (id = 6).
 * twitter.postTweet(2, 6);
 *
 * // User 1's news feed should return a list with 2 tweet ids -> [6, 5].
 * // Tweet id 6 should precede tweet id 5 because it is posted after tweet id 5.
 * twitter.getNewsFeed(1);
 *
 * // User 1 unfollows user 2.
 * twitter.unfollow(1, 2);
 *
 * // User 1's news feed should return a list with 1 tweet id -> [5],
 * // since user 1 is no longer following user 2.
 * twitter.getNewsFeed(1);
 *
 *
 */

package leetcode

type Heap struct {
	storage []*Tweet
}

func NewHeap(a []*Tweet) *Heap {
	h := &Heap{
		storage: a,
	}
	for i := h.parent(len(h.storage)); i >= 0; i-- {
		h.heapify(i)
	}
	return h
}

func (h *Heap) ExtractMax() *Tweet {
	n := len(h.storage) - 1
	h.storage[0], h.storage[n] = h.storage[n], h.storage[0]
	pop := h.storage[n]
	h.storage = h.storage[:n]
	h.heapify(0)
	return pop
}

func (h *Heap) Push(val *Tweet) {
	h.storage = append(h.storage, val)
	i := len(h.storage) - 1
	parent := h.parent(i)
	for parent != -1 && h.storage[i].timestamp > h.storage[parent].timestamp {
		h.storage[i], h.storage[parent] = h.storage[parent], h.storage[i]
		i = parent
		parent = h.parent(i)
	}
}

func (h *Heap) Empty() bool {
	return len(h.storage) == 0
}

func (h *Heap) parent(i int) int {
	if i == 0 {
		return -1
	}
	if i%2 != 0 {
		return (i - 1) / 2
	}
	return (i - 2) / 2
}

func (h *Heap) leftchild(i int) int {
	return 2*i + 1
}

func (h *Heap) rightchild(i int) int {
	return 2*i + 2
}

func (h *Heap) heapify(i int) {
	l := h.leftchild(i)
	r := h.rightchild(i)
	largest := i

	if l < len(h.storage) && h.storage[l].timestamp > h.storage[largest].timestamp {
		largest = l
	}

	if r < len(h.storage) && h.storage[r].timestamp > h.storage[largest].timestamp {
		largest = r
	}

	if largest != i {
		h.storage[largest], h.storage[i] = h.storage[i], h.storage[largest]
		h.heapify(largest)
	}
}

// singly linked list
type Tweet struct {
	next          *Tweet
	id, timestamp int
}

type Feed struct {
	head *Tweet
}

func (feed *Feed) Add(tweet *Tweet) {
	tweet.next = feed.head
	feed.head = tweet
}

type UserId = int

type User struct {
	feed      *Feed
	followees map[UserId]*Feed
}

func NewUser() *User {
	return &User{
		feed:      &Feed{},
		followees: make(map[UserId]*Feed),
	}
}

func (user *User) Follow(id int, followee *User) {
	user.followees[id] = followee.feed
}

func (user *User) Unfollow(id int) {
	delete(user.followees, id)
}

func (user *User) Post(id, timestamp int) {
	user.feed.Add(&Tweet{
		id:        id,
		timestamp: timestamp,
	})
}

type Twitter struct {
	// unique timestamp
	timestamp int
	users     map[UserId]*User
}

/** Initialize your data structure here. */
func Constructor() Twitter {
	return Twitter{
		users: make(map[UserId]*User),
	}
}

/** Compose a new tweet. */
func (this *Twitter) PostTweet(userId int, tweetId int) {
	user := this.findOrCreateUser(userId)
	this.timestamp++
	user.Post(tweetId, this.timestamp)
}

/** Retrieve the 10 most recent tweet ids in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user herself. Tweets must be ordered from most recent to least recent. */
func (this *Twitter) GetNewsFeed(userId int) []int {
	heap := &Heap{
		storage: make([]*Tweet, 0, 10),
	}

	user := this.findOrCreateUser(userId)
	if user.feed.head != nil {
		heap.Push(user.feed.head)
	}

	for _, feed := range user.followees {
		tweet := feed.head
		if tweet != nil {
			heap.Push(tweet)
		}
	}

	tweets := make([]int, 0, 10)
	for !heap.Empty() && len(tweets) < 10 {
		tweet := heap.ExtractMax()
		if tweet.next != nil {
			heap.Push(tweet.next)
		}
		tweets = append(tweets, tweet.id)
	}
	return tweets
}

/** Follower follows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Follow(followerId int, followeeId int) {
	if followerId == followeeId {
		return
	}

	// not exactly a no-op
	follower := this.findOrCreateUser(followerId)
	followee := this.findOrCreateUser(followeeId)

	follower.Follow(followeeId, followee)
}

/** Follower unfollows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Unfollow(followerId int, followeeId int) {
	follower := this.findOrCreateUser(followerId)

	follower.Unfollow(followeeId)
}

func (this *Twitter) findOrCreateUser(id int) *User {
	if user, ok := this.users[id]; ok {
		return user
	}

	user := NewUser()
	this.users[id] = user
	return user
}
