package main

import (
	"fmt"
	"time"
)

type MyGenericCollection[T any] struct {
	items []T
	prev  *MyGenericCollection[T]
	next  *MyGenericCollection[T]
}

func (mgc *MyGenericCollection[T]) AddAfter(prev *MyGenericCollection[T], value []T) {
	newItem := MyGenericCollection[T]{items: value, prev: prev}
	if value == nil {
		newItem.items = make([]T, 0)
	}
	prev.next = &newItem
}

func (mgc *MyGenericCollection[T]) Append(values []T) {
	last := mgc.Last()
	last.AddAfter(last, values)
}

func (mgc *MyGenericCollection[T]) Delete() {
	if mgc.prev != nil {
		mgc.prev.next = mgc.next
	}
	if mgc.next != nil {
		mgc.next.prev = mgc.prev
	}
}

func (mgc *MyGenericCollection[T]) First() *MyGenericCollection[T] {
	// find first element
	res := mgc
	for res.prev != nil {
		res = res.prev
	}
	return res
}

func (mgc *MyGenericCollection[T]) Last() *MyGenericCollection[T] {
	// find first element
	res := mgc
	for res.next != nil {
		res = res.next
	}
	return res
}

func (mgc *MyGenericCollection[T]) Count() int {
	res := 0
	first := mgc.First()
	for first != nil {
		res++
		first = first.next
	}
	return res
}

func (mgc *MyGenericCollection[T]) CountValues() int {
	res := 0
	first := mgc.First()
	for first != nil {
		res += len(first.items)
		first = first.next
	}
	return res
}

type cacheItem[T any] struct {
	item    T
	expired int64
}

type MyCache[T any] struct {
	items             map[string]cacheItem[T]
	defaultExpiration time.Duration
}

func NewCache[T any](exp time.Duration) *MyCache[T] {
	return &MyCache[T]{items: make(map[string]cacheItem[T]), defaultExpiration: exp}
}

func (mc *MyCache[T]) Set(key string, value T) {
	mc.items[key] = cacheItem[T]{
		item:    value,
		expired: time.Now().Add(mc.defaultExpiration).UnixNano()}
}

func (mc *MyCache[T]) Get(key string) (T, bool) {
	item, found := mc.items[key]
	var empty T
	if !found {
		return empty, false
	}
	if time.Now().UnixNano() > item.expired {
		delete(mc.items, key) // delete expired item
		return empty, false
	}
	return item.item, true
}
func (mc *MyCache[T]) Delete(key string) {
	delete(mc.items, key)
}

func (mc *MyCache[T]) Cleanup() int {
	cur := len(mc.items)
	for s, c := range mc.items {
		if time.Now().UnixNano() > c.expired {
			delete(mc.items, s)
		}
	}
	return cur - len(mc.items)
}

func printAll[T any](in *MyGenericCollection[T]) {
	first := in.First()
	for first != nil {
		fmt.Printf("%v ", first.items)
		first = first.next
	}
	fmt.Println()
}

func main() {
	var mgc MyGenericCollection[int]
	mgc.items = []int{1, 2, 3}
	mgc.Append([]int{4, 5, 6})
	mgc.Append([]int{7, 8, 9})
	mgc.Append(nil)
	mgc.Append(nil)
	printAll(&mgc)
	fmt.Printf("Count: %d, values: %d\n", mgc.Count(), mgc.CountValues())
	deleted := mgc.Last().prev
	deleted.Delete()
	fmt.Printf("Count: %d, values: %d\n", mgc.Count(), mgc.CountValues())
	printAll(&mgc)
	mgc.Last().prev.Delete()
	printAll(&mgc)
	fmt.Printf("Count: %d, values: %d\n", mgc.Count(), mgc.CountValues())
	cache := NewCache[string](time.Millisecond * 5)
	cache.Set("hello", "world")
	cache.Set("hello2", "world2")
	cache.Set("hello3", "world4")
	fmt.Println(cache.Get("hello"))
	time.Sleep(time.Millisecond * 10)
	fmt.Println(cache.Get("hello"))
	fmt.Println(cache.Cleanup())
}
