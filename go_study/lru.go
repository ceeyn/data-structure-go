package main

// 入参、返回值尽可能都用指针
// 结构体里递归调用要用指针
type node struct {
	key  int
	val  int
	next *node
	pre  *node
}

func NewNode(key int) *node {
	n := node{key: key}
	return &n
}

type LRUCache struct {
	dummy    *node
	books    map[int]*node
	capacity int
}

func Constructor(capacity int) LRUCache {
	dummy := NewNode(-1)
	dummy.next = dummy
	dummy.pre = dummy
	lru := LRUCache{
		dummy:    dummy,
		books:    make(map[int]*node),
		capacity: capacity,
	}

	return lru
}

func (this *LRUCache) Get(key int) int {
	n := this.GetNode(key)
	if n == nil {
		return -1
	} else {
		return n.val
	}
}

func (this *LRUCache) Put(key int, value int) {
	n := this.GetNode(key)
	if n == nil {
		n = &node{
			key: key,
			val: value,
		}
		this.PutFront(n)
		this.books[key] = n
		if len(this.books) > this.capacity {
			//这两段代码的不同之处在于它们在删除节点时的处理方式。如果直接删除`this.dummy.pre`节点，`dummy.pre`指向的节点会在调用`this.remove`后发生变化，因此第二次访问`dummy.pre`可能会导致错误。详细分析如下：
			//
			//### 第一段代码
			//```go
			//removedNode := this.dummy.pre
			//this.remove(removedNode)
			//delete(this.books, removedNode.key)
			//```
			//
			//在第一段代码中，`this.dummy.pre`被保存到`removedNode`变量中，然后通过`this.remove(removedNode)`删除该节点，最后使用`removedNode.key`从`books`映射中删除相应的键值对。这种方法确保了在删除操作中不会因为`dummy.pre`的变化而产生不一致性。
			//
			//### 第二段代码
			//```go
			//this.remove(this.dummy.pre)
			//delete(this.books, this.dummy.pre.key)
			//```
			//
			//在第二段代码中，删除节点时直接使用`this.dummy.pre`。在`this.remove(this.dummy.pre)`执行后，`this.dummy.pre`已经被更新，指向新的末尾节点，因此`delete(this.books, this.dummy.pre.key)`可能会尝试访问已被删除的节点，导致错误或不一致性。
			//
			//### 举个例子说明
			//假设链表当前状态如下：
			//```
			//dummy <-> node1 <-> node2 <-> dummy
			//```
			//此时`dummy.pre`指向`node2`。
			//
			//#### 第一段代码执行步骤：
			//1. `removedNode := this.dummy.pre`，此时`removedNode`指向`node2`。
			//2. `this.remove(removedNode)`删除`node2`，更新链表为：
			//```
			//dummy <-> node1 <-> dummy
			//```
			//3. `delete(this.books, removedNode.key)`，从`books`中删除`node2`。
			//
			//#### 第二段代码执行步骤：
			//1. `this.remove(this.dummy.pre)`删除`node2`，更新链表为：
			//```
			//dummy <-> node1 <-> dummy
			//```
			//2. 此时`dummy.pre`已经被更新，指向`node1`，`delete(this.books, this.dummy.pre.key)`试图删除`node1`，而不是`node2`。
			//
			//### 修正后的代码
			//为了确保删除操作的正确性，建议使用第一种方式：
			//
			//```go
			//func (this *LRUCache) Put(key int, value int) {
			//    n := this.GetNode(key)
			//    if n == nil {
			//        n = &node{
			//            key: key,
			//            val: value,
			//        }
			//        this.PutFront(n)
			//        this.books[key] = n
			//        if len(this.books) > this.capacity {
			//            removedNode := this.dummy.pre
			//            this.remove(removedNode)
			//            delete(this.books, removedNode.key)
			//        }
			//    } else {
			//        n.val = value
			//        this.remove(n)
			//        this.PutFront(n)
			//    }
			//}
			//```
			//
			//这样可以确保在删除节点时，`dummy.pre`的变化不会影响删除操作的正确性。
			this.remove(this.dummy.pre)
			delete(this.books, this.dummy.pre.key)
		}
	} else {
		n.val = value
	}

}
func (this *LRUCache) GetNode(key int) *node {
	n, ex := this.books[key]
	if !ex {
		return nil
	} else {
		this.remove(n)
		this.PutFront(n)
		return n
	}
}
func (this *LRUCache) remove(node *node) {
	node.pre.next = node.next
	node.next.pre = node.pre

}
func (this *LRUCache) PutFront(node *node) {
	node.next = this.dummy.next
	this.dummy.next.pre = node
	this.dummy.next = node
	node.pre = this.dummy
}
