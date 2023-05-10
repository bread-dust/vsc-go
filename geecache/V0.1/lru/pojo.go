package lru

// 双向链表节点
type entry struct {
	key   string
	value Value
}


func newEntry(key string, value Value) *entry {
	return &entry{
		key:   key,
		value: value,
	}
}
