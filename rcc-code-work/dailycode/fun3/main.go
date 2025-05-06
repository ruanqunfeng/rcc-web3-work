package main

import (
	"fmt"
	"reflect"
	"sync"
	"unsafe"
)

// hmap是Go map的底层结构(简化版)
type hmap struct {
	B uint8 // log_2 of # of buckets (can hold up to loadFactor * 2^B items)
	// 其他字段省略...
}

// 获取map的桶数量
func getMapBuckets(m interface{}) int {
	// 获取map的reflect.Value
	mv := reflect.ValueOf(m)

	// 获取map的底层hmap结构指针
	mh := (*hmap)(unsafe.Pointer(mv.Pointer()))

	// 返回桶数量
	return 1 << mh.B
}

func getB(m map[int]int) uint8 {
	// 将 map 转为 *hmap 指针，获取内部字段
	return (*hmap)(*(*unsafe.Pointer)(unsafe.Pointer(&m))).B
}

func getBucketNum() {
	m := make(map[int]int)

	for i := 0; i < 20; i++ {
		m[i] = i
		buckets := getMapBuckets(m)
		fmt.Printf("After inserting %d: buckets = %d\n", i, buckets)
	}
}

func getBucketNum2() {
	m := make(map[int]int)
	var lastB uint8 = 0

	for i := 0; i < 20; i++ {
		m[i] = i
		b := getB(m)
		if b != lastB {
			fmt.Printf("插入 %d 个元素后，B = %d，桶数 = %d\n", i+1, b, 1<<b)
			lastB = b
		}
	}
}

func main() {

	getBucketNum2()

	var m sync.Map

	m.Store("key1", "value1")

	v, ok := m.Load("key1")
	if ok {
		fmt.Println(v) // 输出：value1
	}

	// 遍历
	m.Range(func(key, value any) bool {
		fmt.Println(key, value)
		return true
	})
}

type SafeMap[K comparable, V any] struct {
	m  map[K]V
	mu sync.RWMutex
}

// NewSafeMap 返回一个初始化好的线程安全 map
func NewSafeMap[K comparable, V any]() *SafeMap[K, V] {
	return &SafeMap[K, V]{
		m: make(map[K]V),
	}
}

func (s *SafeMap[K, V]) Load(key K) (value V, ok bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	value, ok = s.m[key]
	return
}

func (s *SafeMap[K, V]) Store(key K, value V) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.m[key] = value
}

func (s *SafeMap[K, V]) Delete(key K) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.m, key)
}

func (s *SafeMap[K, V]) Range(f func(key K, value V) bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for k, v := range s.m {
		if !f(k, v) {
			break
		}
	}
}
