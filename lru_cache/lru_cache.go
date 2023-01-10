// https://leetcode.com/problems/lru-cache/
/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
package lrucache

type LRUCache struct {
	head      *Node
	tail      *Node
	nodeByKey map[int]*Node
	capacity  int
}

type Node struct {
	key   int
	value int
	prev  *Node
	next  *Node
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		nodeByKey: make(map[int]*Node, capacity),
		capacity:  capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.nodeByKey[key]
	if ok {
		this.Remove(node)
		this.Add(node)
		return node.value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	node, ok := this.nodeByKey[key]
	if ok {
		node.value = value
		this.Remove(node)
		this.Add(node)
		return
	}

	if len(this.nodeByKey) == this.capacity {
		delete(this.nodeByKey, this.tail.key)
		this.Remove(this.tail)
	}

	node = &Node{key: key, value: value}
	this.nodeByKey[key] = node
	this.Add(node)

}

// adds node to the head of the list
func (this *LRUCache) Add(node *Node) {
	node.prev = nil
	node.next = this.head
	if this.head != nil {
		this.head.prev = node
	}
	this.head = node
	if this.tail == nil {
		this.tail = node
	}
}

func (this *LRUCache) Remove(node *Node) {
	// link node.next
	if node != this.head {
		// if node is not head, it will have a prev
		node.prev.next = node.next
	} else {
		this.head = node.next
	}

	// link node.prev
	if node != this.tail {
		// if node is not tail, it will have a next
		node.next.prev = node.prev
	} else {
		this.tail = node.prev
	}
}
