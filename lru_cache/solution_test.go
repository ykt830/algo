package lru_cache

import (
	"fmt"
	"testing"
)

func TestLRUCache_func(t *testing.T) {
	cache := Constructor(2)
	ret := make([]int, 0)
	ret = append(ret, cache.Get(2))
	cache.Put(2, 6)
	ret = append(ret, cache.Get(1))
	cache.Put(1, 5)
	cache.Put(1, 2)
	ret = append(ret, cache.Get(1))
	ret = append(ret, cache.Get(2))
	fmt.Println(ret)
}
