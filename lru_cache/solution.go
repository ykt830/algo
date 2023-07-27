package lru_cache

/*
请你设计并实现一个满足  LRU (最近最少使用) 缓存 约束的数据结构。
实现 LRUCache 类：
LRUCache(int capacity) 以 正整数 作为容量 capacity 初始化 LRU 缓存
int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
void put(int key, int value) 如果关键字 key 已经存在，则变更其数据值 value ；如果不存在，则向缓存中插入该组 key-value 。如果插入操作导致关键字数量超过 capacity ，则应该 逐出 最久未使用的关键字。
函数 get 和 put 必须以 O(1) 的平均时间复杂度运行。

示例：

输入
["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
[[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
输出
[null, null, null, 1, null, -1, null, -1, 3, 4]

解释
LRUCache lRUCache = new LRUCache(2);
lRUCache.put(1, 1); // 缓存是 {1=1}
lRUCache.put(2, 2); // 缓存是 {1=1, 2=2}
lRUCache.get(1);    // 返回 1
lRUCache.put(3, 3); // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
lRUCache.get(2);    // 返回 -1 (未找到)
lRUCache.put(4, 4); // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
lRUCache.get(1);    // 返回 -1 (未找到)
lRUCache.get(3);    // 返回 3
lRUCache.get(4);    // 返回 4

提示：
1 <= capacity <= 3000
0 <= key <= 10000
0 <= value <= 105
最多调用 2 * 105 次 get 和 put
*/

type LRUCache struct {
	cap  int
	hash map[int]*Node
	head *Node
	tail *Node
}

type Node struct {
	prv  *Node
	next *Node
	key  int
	val  int
}

func InsertNode(prv *Node, key, val int) *Node {
	nxt := prv.next
	if nxt == nil {
		newNode := Node{
			key:  key,
			val:  val,
			prv:  prv,
			next: nil,
		}
		prv.next = &newNode
		return &newNode
	}

	newNode := Node{
		key:  key,
		val:  val,
		prv:  prv,
		next: nxt,
	}
	prv.next = &newNode
	nxt.prv = &newNode

	return &newNode
}

func RemoveNode(node *Node) {
	prv, nxt := node.prv, node.next
	if prv != nil {
		prv.next = nxt
	}
	if nxt != nil {
		nxt.prv = prv
	}
}

func Constructor(capacity int) LRUCache {
	c := LRUCache{
		cap:  capacity,
		hash: make(map[int]*Node),
		head: &Node{},
		tail: &Node{},
	}

	c.head.next, c.tail.prv = c.tail, c.head

	return c
}

func (cache *LRUCache) Get(key int) int {
	node, ok := cache.hash[key]
	if !ok {
		return -1
	}

	RemoveNode(node)
	cache.hash[key] = InsertNode(cache.head, key, node.val)

	return cache.hash[key].val
}

func (cache *LRUCache) Put(key int, value int) {
	if fill := len(cache.hash) >= cache.cap; cache.hash[key] == nil && fill {
		delete(cache.hash, cache.tail.prv.key)
		RemoveNode(cache.tail.prv)
	}

	if node, ok := cache.hash[key]; ok {
		delete(cache.hash, key)
		RemoveNode(node)
	}

	cache.hash[key] = InsertNode(cache.head, key, value)
}
