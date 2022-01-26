package main

import (
	"container/list"
	"fmt"
)

// https://leetcode-cn.com/problems/lru-cache/

type intMap map[int]*list.Element

func (m intMap) get(key int) *list.Element {
	return m[key]
}

func (m intMap) add(key int, e *list.Element) {
	m[key] = e
}

func (m intMap) len() int {
	return len(m)
}

func (m intMap) delete(key int) {
	delete(m, key)
}

type LRUCache struct {
	capacity int
	dict     intMap
	list     *list.List
}

type kv struct {
	key   int
	value int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		dict:     make(intMap, capacity),
		list:     list.New(),
		capacity: capacity,
	}
}

func (this *LRUCache) len() int {
	dictLen := this.dict.len()
	listLen := this.list.Len()
	if dictLen != listLen {
		err := fmt.Errorf("bug: length mismatch, dict len: %v, list len: %v", dictLen, listLen)
		panic(err)
	}
	return dictLen
}

func (this *LRUCache) lruElement() *list.Element {
	return this.list.Front()
}

func (this *LRUCache) touchElement(e *list.Element) {
	this.list.MoveToBack(e)
}

func (this *LRUCache) deleteElement(e *list.Element) {
	key := this.list.Remove(e).(kv).key
	this.dict.delete(key)
}

func (this *LRUCache) Get(key int) int {
	if e := this.dict.get(key); e != nil {
		this.touchElement(e)
		return e.Value.(kv).value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if this.capacity <= 0 {
		return
	}

	// found then update its value
	if e := this.dict.get(key); e != nil {
		e.Value = kv{key: key, value: value}
		this.touchElement(e)
		return
	}

	// full, remove the LRU element
	if this.len() >= this.capacity {
		e := this.lruElement()
		this.deleteElement(e)
	}

	e := this.list.PushBack(kv{key: key, value: value})
	this.dict.add(key, e)
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
