package main

import (
	"fmt"
	"github.com/gorhill/cronexpr"
	"math"
	"time"
)

func main() {
	nextTime := cronexpr.MustParse("00 46 15,16 * * ? *").Next(time.Now()).Unix()
	fmt.Println(nextTime)
	//lc := Constructor(2)
	//lc.Put(1, 1)
	//lc.Put(2, 2)
	//fmt.Println(lc.Get(1))
	//lc.Put(3, 3)
	//fmt.Println(lc.Get(2))
	//lc.Put(4, 4)
	//fmt.Println(lc.Get(1))
	//fmt.Println(lc.Get(3))
	//fmt.Println(lc.Get(4))
	//lengthOfLongestSubstring("aab")
	//head := &ListNode{
	//	Val: 1,
	//}
	//head2 := &ListNode{
	//	Val: 1,
	//}
	//node := head
	//node2 := head2
	//v1 := []int{1, 2, 4}
	//v2 := []int{1, 3, 4}
	//
	//for k, v := range v1 {
	//	node.Val = v
	//	if k != len(v1)-1 {
	//		node.Next = &ListNode{}
	//	}
	//	node = node.Next
	//}
	//for k, v := range v2 {
	//	node2.Val = v
	//	if k != len(v1)-1 {
	//		node2.Next = &ListNode{}
	//	}
	//	node2 = node2.Next
	//}
	//mergeTwoLists(head, head2)
	//removeNthFromEnd(head)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if maxDepth(root.Left) >= maxDepth(root.Right) {
		return 1 + maxDepth(root.Left)
	} else {
		return 1 + maxDepth(root.Right)
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	//head -> 1 -> 2 -> 3

	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}

	p := head
	head = head.Next
	p.Next = nil

	for head.Next != nil {
		tmp := head
		head = head.Next
		tmp.Next = p
		p = tmp
	}
	head.Next = p

	return head
}
func maxProfit(prices []int) int {
	profit := 0
	buyPrice := math.MaxInt
	for i := 0; i < len(prices); i++ {
		if buyPrice > prices[i] {
			buyPrice = prices[i]
		} else if prices[i]-buyPrice > profit {
			profit = prices[i] - buyPrice
		}
	}
	return profit
}

func reverse(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	tail := head
	head = head.Next
	tail.Next = nil
	for head != nil {
		fast := head.Next
		head.Next = tail
		tail = head
		head = fast
	}

	return tail
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	head := &ListNode{}
	current := head
	for list1 != nil && list2 != nil {
		if list1.Val > list2.Val {
			current.Next = list2
			list2 = list2.Next
			current = current.Next
		} else {
			current.Next = list1
			list1 = list1.Next
			current = current.Next
		}
	}
	if list1 == nil {
		current.Next = list2
	}
	if list2 == nil {
		current.Next = list1
	}
	return head.Next
}
func removeNthFromEnd(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	tail := head
	head = head.Next
	tail.Next = nil
	for head != nil {
		fast := head.Next
		head.Next = tail
		tail = head
		head = fast
	}
	return tail
}

/*
Given an integer array nums and an integer k, return the kth largest element in the array.

Note that it is the kth largest element in the sorted order, not the kth distinct element.

You must solve it in O(n) time complexity.

Example 1:

Input: nums = [3,2,1,5,6,4], k = 2
Output: 5
Example 2:

Input: nums = [3,2,3,1,2,4,5,5,6], k = 4
Output: 4


Constraints:

1 <= k <= nums.length <= 105
-104 <= nums[i] <= 104
*/
func findKthLargest(nums []int, k int) int {
	return 0
}

func lengthOfLongestSubstring(s string) int {
	maxLength := 0
	repeat := map[byte]bool{}
	for i := 0; i < len(s); i++ {
		length := 0
		for index := i; index < len(s); index++ {
			if _, ok := repeat[s[index]]; ok {
				if length > maxLength {
					maxLength = length
				}
				length = 0
				repeat = map[byte]bool{}
				break
			} else {
				repeat[s[index]] = true
				length++
				if index == len(s)-1 {
					if length > maxLength {
						maxLength = length
					}
					return maxLength
				}
			}

		}
	}
	return maxLength
}

/*
Input
["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
[[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
Output
[null, null, null, 1, null, -1, null, -1, 3, 4]

Explanation
LRUCache lRUCache = new LRUCache(2);
lRUCache.put(1, 1); // cache is {1=1}
lRUCache.put(2, 2); // cache is {1=1, 2=2}
lRUCache.get(1);    // return 1
lRUCache.put(3, 3); // LRU key was 2, evicts key 2, cache is {1=1, 3=3}
lRUCache.get(2);    // returns -1 (not found)
lRUCache.put(4, 4); // LRU key was 1, evicts key 1, cache is {4=4, 3=3}
lRUCache.get(1);    // return -1 (not found)
lRUCache.get(3);    // return 3
lRUCache.get(4);    // return 4
*/

// LRUCache(int capacity) Initialize the LRU cache with positive size capacity.
type LRUCache struct {
	size       int
	capacity   int
	cache      map[int]*DLinkedNode
	head, tail *DLinkedNode
}

type DLinkedNode struct {
	key, value int
	prev, next *DLinkedNode
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		size:     2,
		cache:    map[int]*DLinkedNode{},
		head:     initDLinkedNode(0, 0),
		tail:     initDLinkedNode(0, 0),
		capacity: capacity,
	}
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}
func initDLinkedNode(key, value int) *DLinkedNode {
	return &DLinkedNode{
		key:   key,
		value: value,
	}
}

// int get(int key) Return the value of the key if the key exists, otherwise return -1.
func (this *LRUCache) Get(key int) int {
	if _, ok := this.cache[key]; ok {
		moveToHead(this, key)
		return this.cache[key].value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	// 存在, 更新value, 调整位置,
	// 不存在, 检查是否超出容量
	//  		超了: 删除最后一个, 加到最新一个
	// 			没超: 修改size, 加到最新一个
	if _, ok := this.cache[key]; ok {
		this.cache[key].value = value
		moveToHead(this, key)
		return
	}
	p := initDLinkedNode(key, value)
	this.cache[key] = p
	if this.size == this.capacity {
		delete(this.cache, this.tail.key)
		p.next = this.head
		this.head.prev = p
		this.head = p
		this.tail = this.tail.prev
		this.tail.next = nil
		return
	}

	this.size++
	this.cache[key] = p
	p.next = this.head
	this.head.prev = p
	this.head = p
}
func moveToHead(this *LRUCache, key int) {
	if this.cache[key].prev != nil {
		this.cache[key].prev.next = this.cache[key].next
	}
	if this.cache[key].next != nil {
		this.cache[key].next.prev = this.cache[key].prev
	}

	this.cache[key].next = this.head
	this.cache[key].prev = nil

	this.head.prev = this.cache[key]

	this.head = this.cache[key]
}
