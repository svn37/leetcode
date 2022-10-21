/**
 * Create a timebased key-value store class TimeMap, that supports two operations.
 *
 * 1. set(string key, string value, int timestamp)
 *
 *
 * 	Stores the key and value, along with the given timestamp.
 *
 *
 * 2. get(string key, int timestamp)
 *
 *
 * 	Returns a value such that set(key, value, timestamp_prev) was called previously, with timestamp_prev <= timestamp.
 * 	If there are multiple such values, it returns the one with the largest timestamp_prev.
 * 	If there are no values, it returns the empty string ("").
 *
 *
 *
 *
 * <div>
 * Example 1:
 *
 *
 * Input: inputs = <span id="example-input-1-1">["TimeMap","set","get","get","set","get","get"]</span>, inputs = <span id="example-input-1-2">[[],["foo","bar",1],["foo",1],["foo",3],["foo","bar2",4],["foo",4],["foo",5]]</span>
 * Output: <span id="example-output-1">[null,null,"bar","bar",null,"bar2","bar2"]</span>
 * Explanation: <span id="example-output-1">
 * TimeMap kv;
 * kv.set("foo", "bar", 1); // store the key "foo" and value "bar" along with timestamp = 1
 * kv.get("foo", 1);  // output "bar"
 * kv.get("foo", 3); // output "bar" since there is no value corresponding to foo at timestamp 3 and timestamp 2, then the only value is at timestamp 1 ie "bar"
 * kv.set("foo", "bar2", 4);
 * kv.get("foo", 4); // output "bar2"
 * kv.get("foo", 5); //output "bar2"
 * </span>
 *
 *
 * <div>
 * Example 2:
 *
 *
 * Input: inputs = <span id="example-input-2-1">["TimeMap","set","set","get","get","get","get","get"]</span>, inputs = <span id="example-input-2-2">[[],["love","high",10],["love","low",20],["love",5],["love",10],["love",15],["love",20],["love",25]]</span>
 * Output: <span id="example-output-2">[null,null,null,"","high","high","low","low"]</span>
 *
 * </div>
 * </div>
 *
 *
 *
 * Note:
 *
 * <ol>
 * 	All key/value strings are lowercase.
 * 	All key/value strings have length in the range [1, 100]
 * 	The timestamps for all TimeMap.set operations are strictly increasing.
 * 	1 <= timestamp <= 10^7
 * 	TimeMap.set and TimeMap.get functions will be called a total of 120000 times (combined) per test case.
 * </ol>
 *
 */

package leetcode

type Value struct {
	timestamp int
	value     string
}

type TimeMap struct {
	m map[string][]Value
}

func Constructor() TimeMap {
	return TimeMap{
		m: make(map[string][]Value),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.m[key] = append(this.m[key], Value{timestamp, value})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	return bsearch(this.m[key], timestamp)
}

func bsearch(arr []Value, target int) string {
	lo, hi := 0, len(arr)
	for lo < hi {
		mid := (lo + hi) / 2
		if arr[mid].timestamp <= target {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	if lo == 0 {
		return ""
	}
	return arr[lo-1].value
}
