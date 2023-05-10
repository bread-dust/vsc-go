package lru

import "container/list"

// 缓存结构体
type cache struct {
	maxBytes   int64                         //最大可以缓存多少字节
	curBytes   int64                         // 当前缓存的字节数
	linkedList *list.List                    //双向链表
	cacheMap   map[string]*list.Element      //字典，key是字符串，value是双向链表中对应节点的指针
	OnEvicted  func(key string, value Value) //某条记录被移除时的回调函数
}

func New(maxBytes int64, onEvicted func(string, Value)) *cache {
	if maxBytes <= 0 {
		panic("maxBytes of cache must be lager than 0")
	}

	return &cache{
		maxBytes:   maxBytes,
		linkedList: list.New(),
		cacheMap:   make(map[string]*list.Element),
		OnEvicted:  onEvicted,
	}
}

// Get 从缓存中获取key对应的值
func (c *cache) Get(key string) (value Value, ok bool) {
	if ele, ok := c.cacheMap[key]; ok {
		c.linkedList.MoveToFront(ele) // 将该节点移动到链表头部
		kv := ele.Value.(*entry)      // 类型断言
		return kv.value, true
	}
	

// Add 添加一个缓存
func (c *cache) Add(key string, value Value) {
	// 判断是更新还是移除
	if ele, ok := c.cacheMap[key]; ok { // 更新
		c.linkedList.MoveToFront(ele) // 移动到链表头部
		kv := ele.Value.(*entry) 	// 类型断言
		c.curBytes += int64(value.Len()) - int64(kv.value.Len()) // 更新缓存大小
		kv.value = value // 更新值
	} else {
		// 添加
		ele := c.linkedList.PushFront(&entry{key, value}) // 在链表头部添加一个节点
		c.cacheMap[key] = ele 						   // 在字典中添加一个key-value
		c.curBytes += int64(len(key)) + int64(value.Len()) // 更新缓存大小
	}

	// 判断是否需要移除 maxBytes >= curBytes
	for c.maxBytes != 0 && c.curBytes > c.maxBytes {
		c.RemoveOldest()
	}

}

// RemoveOldest 移除最近最少访问的节点
func (c *cache) RemoveOldest() {
	list := c.linkedList

	// 获取最后一个节点
	revElem := list.Back()
	if revElem ==nil{ // 链表为空
		return
	}
		
	// 从链表中删除该节点
	list.Remove(revElem)
	e:=revElem.Value.(*entry) // 类型断言
	delete(c.cacheMap,e.key) // 从字典中删除该节点
	c.curBytes -= int64(len(e.key)) + int64(e.value.Len()) // 更新缓存大小

	// 驱逐元素之后的回调函数
	if c.OnEvicted != nil {
		c.OnEvicted(e.key, e.value)
	}
}


