package main

import "fmt"

// 键值对
type pair struct {
	key int
	val string
}

// 基于数组实现哈希表
type arrayHashMap struct {
	buckets []*pair
}

// 初始化哈希表
func initArrayHashMap() *arrayHashMap {
	// 初始化数组，包含 100 个桶
	return &arrayHashMap{
		buckets: make([]*pair, 100),
	}
}

// 哈希函数
func (s *arrayHashMap) hashFunc(key int) int {
	return key % 100
}

// 添加
func (s *arrayHashMap) put(key int, val string) {
	pair := &pair{
		key: key,
		val: val,
	}
	index := s.hashFunc(key)
	s.buckets[index] = pair
}

// 删除
func (s *arrayHashMap) remove(key int) {
	s.buckets[s.hashFunc(key)] = nil
}

/* 获取所有键对 */
func (s *arrayHashMap) pairSet() []*pair {
	var pairs []*pair
	for _, pair := range s.buckets {
		if pair != nil {
			pairs = append(pairs, pair)
		}
	}
	return pairs
}

/* 获取所有键 */
func (s *arrayHashMap) keySet() []int {
	var keys []int
	for _, pair := range s.buckets {
		if pair != nil {
			keys = append(keys, pair.key)
		}
	}
	return keys
}

/* 获取所有值 */
func (s *arrayHashMap) valueSet() []string {
	var values []string
	for _, pair := range s.buckets {
		if pair != nil {
			values = append(values, pair.val)
		}
	}
	return values
}

/* 打印哈希表 */
func (s *arrayHashMap) print() {
	for _, pair := range s.buckets {
		if pair != nil {
			fmt.Println(pair.key, "->", pair.val)
		}
	}
}

func main() {
	hashTable := initArrayHashMap()

	hashTable.put(12836, "a")
	hashTable.put(15937, "b")
	hashTable.put(16750, "c")

}
