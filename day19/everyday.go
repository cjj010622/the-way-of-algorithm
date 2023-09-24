package day19

/*
请你设计并实现一个满足  LRU (最近最少使用) 缓存 约束的数据结构。
实现 LRUCache 类：
LRUCache(int capacity) 以 正整数 作为容量 capacity 初始化 LRU 缓存
int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
void put(int key, int value) 如果关键字 key 已经存在，则变更其数据值 value ；如果不存在，则向缓存中插入该组 key-value 。如果插入操作导致关键字数量超过 capacity ，则应该 逐出 最久未使用的关键字。
函数 get 和 put 必须以 O(1) 的平均时间复杂度运行。
*/

type node struct {
    key,val int
    prev,next *node
}

type LRUCache struct {
	capacity   int
	cache      map[int]*node
	head, tail *node
}

func Constructor(capacity int) LRUCache {
    head:=new(node)
    tail:=new(node)
    head.next=tail
    tail.prev=head
    return LRUCache{capacity,make(map[int]*node,capacity),head,tail}
}


func (this *LRUCache) Get(key int) int {
    n,ok:=this.cache[key]
    if !ok{
        return -1
    }

    this.moveToFront(n)
    return n.val
}


func (this *LRUCache) Put(key int, value int)  {
    n,ok:=this.cache[key]
    if ok{
        n.val=value
        this.moveToFront(n)
        return
    }

    if len(this.cache)==this.capacity{
        back:=this.tail.prev
        this.remove(back)
        delete(this.cache,back.key)
    }

    n=&node{key:key,val:value}
    this.pushFront(n)
    this.cache[key]=n
}

func (this *LRUCache) moveToFront(n *node) {
    this.remove(n)
    this.pushFront(n)
}

func (this *LRUCache) remove(n *node) {
    n.prev.next=n.next
    n.next.prev=n.prev
    n.prev=nil
    n.next=nil
}

func (this *LRUCache) pushFront(n *node) {
    n.prev=this.head
    n.next=this.head.next
    this.head.next.prev=n
    this.head.next=n
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
