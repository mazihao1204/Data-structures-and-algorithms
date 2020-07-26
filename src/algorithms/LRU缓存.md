#### LRU缓存

LRU：最近最少使用缓存机制。支持两个操作

> get(key): 如果key存在于缓存中，则返回对应的值
>
> put(key,value): 如果key存在，则修改该值，如果key不存在，则插入该缓存。当缓存容量达到上限时，应该在写入数据之前，删除最久未使用的数据，为插入数据留出空间



#### 实现

哈希表辅以双向链表实现

+ 双向链表靠近头部的数据是最近使用的，靠近尾部的数据是最久为使用的
+ 哈希表：通过缓存数据的键映射到其在双向链表中的位置



~~~ go
type LRUCache struct {
	size int
	capacity int
	cache map[int]*DlinkNode
	head,tail *DlinkNode
}

type DlinkNode struct {
	key,value int
	prev,next *DlinkNode
}

func initDlinkNode(key, value int) *DlinkNode{
	return &DlinkNode{
		key: key,
		value: value,
	}
}

func Constructor(capacity int) LRUCache{
	l := LRUCache{
		cache: map[int]*DlinkNode{},
		capacity: capacity,
		tail: initDlinkNode(0,0),
		head: initDlinkNode(0,0),
	}
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

func (c *LRUCache) Get(key int) int{
	if _,ok := c.cache[key];!ok{
		return -1
	}
	node := c.cache[key]
	c.moveToHead(node)
	return node.value
}

func (c *LRUCache) Put(key int,value int){
	if _,ok := c.cache[key];!ok{
		node := initDlinkNode(key,value)
		c.cache[key] = node
		c.addToHead(node)
		c.size++
		if c.size > c.capacity{
			removed := c.removeTail()
			delete(c.cache,removed.key)
			c.size--
		}
	}else {
		node := c.cache[key]
		node.value = value
		c.moveToHead(node)
	}
}

func (c *LRUCache) addToHead(node *DlinkNode){
	node.prev = c.head
	node.next = c.head.next
	c.head.next.prev = node
	c.head.next = node
}

func (c *LRUCache) removeNode(node *DlinkNode){
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (c *LRUCache) moveToHead(node *DlinkNode){
	c.removeNode(node)
	c.moveToHead(node)
}

func (c *LRUCache) removeTail() *DlinkNode{
	node := c.tail.prev
	c.removeNode(node)
	return node
}
~~~

