package gopractice

// https://leetcode.com/problems/lru-cache/

// Tips: use two pseudo nodes for the head and tail of double-linked list.

type VisitNode struct {
	key int
	value int
	pre *VisitNode
	next *VisitNode
}


type LRUCache struct {
	pointers map[int]*VisitNode
	head *VisitNode
	tail *VisitNode
	size int
	cap int
}


func Constructor(capacity int) LRUCache {
	head := &VisitNode{0, 0, nil, nil}
	tail := &VisitNode{0, 0, nil, nil}
	head.next = tail
	tail.pre = head
	return LRUCache{map[int]*VisitNode{}, head, tail, 0, capacity}
}

func (this *LRUCache) pushHead(node *VisitNode) {
	node.next = this.head.next
	node.next.pre = node
	this.head.next = node
	node.pre = this.head
}

func (this *LRUCache) removeNode(node *VisitNode) {
	node.next.pre = node.pre
	node.pre.next = node.next
}

func (this *LRUCache) popTail() *VisitNode{
	node := this.tail.pre
	this.removeNode(node)
	return node
}

func (this *LRUCache) get(key int) (*VisitNode, bool) {
	if node, ok := this.pointers[key]; ok {
		this.removeNode(node)
		this.pushHead(node)
		return node, true
	}
	return nil, false
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.get(key); ok {
		return node.value
	}
	return -1
}


func (this *LRUCache) Put(key int, value int)  {
	if node, ok := this.get(key); !ok {
		if this.size == this.cap {
			evictNode := this.popTail()
			delete(this.pointers, evictNode.key)
			this.size--
		}
		node = &VisitNode{key, value, nil, nil}
		this.pushHead(node)
		this.pointers[key] = node
		this.size++
	} else {
		node.value = value
	}
}
