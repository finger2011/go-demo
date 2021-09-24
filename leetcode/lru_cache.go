package main

import (
	"fmt"
)

//https://leetcode-cn.com/problems/lru-cache/
//leetcode 146
func testLRUCache() {
	lRUCache := Constructor(2)
	lRUCache.Put(1, 1) // 缓存是 {1=1}
	// lRUCache.printValues()
	lRUCache.Put(2, 2) // 缓存是 {1=1, 2=2}
	// lRUCache.printValues()
	lRUCache.Get(1) // 返回 1
	// lRUCache.printValues()
	lRUCache.Put(3, 3) // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
	// lRUCache.printValues()
	a := lRUCache.Get(2) // 返回 -1 (未找到)
	fmt.Printf("get 2 result:%v\n", a)
	// lRUCache.printValues()
	lRUCache.Put(4, 4) // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
	// lRUCache.printValues()
	a = lRUCache.Get(1) // 返回 -1 (未找到)
	fmt.Printf("get 1 result:%v\n", a)
	lRUCache.Get(3) // 返回 3
	// lRUCache.printValues()
	lRUCache.Get(4) // 返回 4
	// lRUCache.printValues()
}

//LRUCache lru cache struct
type LRUCache struct {
	capacity int
	length   int
	htable   map[ValueInt]*ListNodeStruct
	head     *ListNodeStruct
	tail     *ListNodeStruct
}

//ListNodeStruct list node
type ListNodeStruct struct {
	Key  ValueInt
	Val  ValueInt
	Next *ListNodeStruct
	Prev *ListNodeStruct
}

//ValueInt int value
type ValueInt int

//Constructor create lrucache
func Constructor(capacity int) LRUCache {
	var c = LRUCache{
		capacity: capacity,
		htable:   make(map[ValueInt]*ListNodeStruct, capacity),
		head:     new(ListNodeStruct),
		tail:     new(ListNodeStruct),
	}
	c.head.Next = c.tail
	c.tail.Prev = c.head
	return c
}

//Get get cache by key
func (c *LRUCache) Get(key int) int {
	if c.htable[ValueInt(key)] == nil {
		return -1
	}
	c.activeKey(c.htable[ValueInt(key)], false)
	return int(c.htable[ValueInt(key)].Val)
}

//Put put cache key
func (c *LRUCache) Put(key int, value int) {
	if c.htable[ValueInt(key)] != nil {
		c.htable[ValueInt(key)].Val = ValueInt(value)
		c.activeKey(c.htable[ValueInt(key)], false)
	} else {
		var node = new(ListNodeStruct)
		node.Val = ValueInt(value)
		node.Key = ValueInt(key)
		c.htable[ValueInt(key)] = node
		c.activeKey(node, true)
		c.length++
	}
}

func (c *LRUCache) printValues() {
	cur := c.head.Next
	fmt.Println("print cache:")
	for {
		fmt.Printf("key:%v ====> value:%v\n", cur.Key, cur.Val)
		cur = cur.Next
		if cur == nil || cur.Next == nil {
			break
		}
	}
}

func (c *LRUCache) activeKey(node *ListNodeStruct, insert bool) error {
	if !insert {
		var prev = node.Prev
		var next = node.Next
		prev.Next = next
		next.Prev = prev
	} else {
		if c.length == c.capacity && c.length > 0 {
			c.htable[c.head.Next.Key] = nil
			c.head.Next.Next.Prev = c.head
			c.head.Next = c.head.Next.Next
			c.length--
		}
	}
	var prev = c.tail.Prev
	prev.Next = node
	node.Next = c.tail
	node.Prev = prev
	c.tail.Prev = node
	return nil
}
