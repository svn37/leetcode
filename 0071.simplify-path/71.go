/**
 * Given an absolute path for a file (Unix-style), simplify it. Or in other words, convert it to the canonical path.
 * In a UNIX-style file system, a period '.' refers to the current directory. Furthermore, a double period '..' moves the directory up a level.
 * Note that the returned canonical path must always begin with a slash '/', and there must be only a single slash '/' between two directory names. The last directory name (if it exists) must not end with a trailing '/'. Also, the canonical path must be the shortest string representing the absolute path.
 *
 * Example 1:
 *
 * Input: path = "/home/"
 * Output: "/home"
 * Explanation: Note that there is no trailing slash after the last directory name.
 *
 * Example 2:
 *
 * Input: path = "/../"
 * Output: "/"
 * Explanation: Going one level up from the root directory is a no-op, as the root level is the highest level you can go.
 *
 * Example 3:
 *
 * Input: path = "/home//foo/"
 * Output: "/home/foo"
 * Explanation: In the canonical path, multiple consecutive slashes are replaced by a single one.
 *
 * Example 4:
 *
 * Input: path = "/a/./b/../../c/"
 * Output: "/c"
 *
 *
 * Constraints:
 *
 * 	1 <= path.length <= 3000
 * 	path consists of English letters, digits, period '.', slash '/' or '_'.
 * 	path is a valid Unix path.
 *
 */

package leetcode

import "bytes"

type Deque struct {
	storage []string
}

func (d *Deque) Push(str string) {
	d.storage = append(d.storage, str)
}

func (d *Deque) Pop() string {
	if len(d.storage) == 0 {
		return ""
	}
	str := d.storage[len(d.storage)-1]
	d.storage = d.storage[:len(d.storage)-1]
	return str
}

func (d *Deque) Poll() string {
	if len(d.storage) == 0 {
		return ""
	}
	str := d.storage[0]
	d.storage = d.storage[1:]
	return str
}

func (d *Deque) Empty() bool {
	return len(d.storage) == 0
}

func absolutePath(deque *Deque) string {
	var path bytes.Buffer

	for !deque.Empty() {
		path.WriteRune('/')
		path.WriteString(deque.Poll())
	}

	if path.Len() == 0 {
		return "/"
	}
	return path.String()
}

func simplifyPath(path string) string {
	deque := &Deque{}

	path += "/"
	cur := ""
	for _, char := range path {
		if char == '/' {
			// go one directory up
			if cur == ".." {
				deque.Pop()
				// push path if it has changed
			} else if len(cur) > 0 && cur != "." {
				deque.Push(cur)
			}
			cur = ""
		} else {
			cur += string(char)
		}
	}

	return absolutePath(deque)
}
