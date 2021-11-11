import (
    "container/heap"
)
type Node struct{
    timestamp,price int
}

type Heap []Node
type StockPrice struct {
    mp map[int]int
    maxTimestamp int
    minHeap *Heap
    maxHeap *Heap
}
func Constructor() StockPrice {
    minHeap:=&Heap{}
    maxHeap:=&Heap{}
    heap.Init(minHeap)
    heap.Init(maxHeap)
    return StockPrice{
        mp:make(map[int]int),
        maxTimestamp:0,
        minHeap:minHeap,
        maxHeap:maxHeap,
    }   
}
func (this *StockPrice) Update(timestamp int, price int)  {
    this.mp[timestamp]=price
    this.maxTimestamp=max(this.maxTimestamp,timestamp)

    // store in Heaps
    heap.Push(this.maxHeap,Node{timestamp,-price})
    heap.Push(this.minHeap,Node{timestamp,price})
    
}
func (this *StockPrice) Current() int {
    return this.mp[this.maxTimestamp]   
}
func (this *StockPrice) Maximum() int {
    node:=(heap.Pop(this.maxHeap)).(Node) // type assertion
    time:=node.timestamp
    price:=node.price
    for -price != this.mp[time]{
        node=(heap.Pop(this.maxHeap)).(Node)
        time=node.timestamp
        price=node.price
    }
    heap.Push(this.maxHeap,node)
    return -price
    
}
func (this *StockPrice) Minimum() int {
    node:=(heap.Pop(this.minHeap)).(Node)
    time:=node.timestamp
    price:=node.price
    for price != this.mp[time]{
        node=(heap.Pop(this.minHeap)).(Node)
        time=node.timestamp
        price=node.price
    }
    heap.Push(this.minHeap,node)
    return price
    
}
func max(a,b int)int{
    if a>=b{
        return a
    }
    return b
}
func (h Heap)Len()int{return len(h)}
func (h Heap)Less(i,j int)bool{return h[i].price<h[j].price}
func (h Heap)Swap(i,j int){h[i],h[j]=h[j],h[i]}
func (h *Heap)Push(x interface{}){*h=append(*h,x.(Node))}
func (h *Heap)Pop()interface{}{
    n:=h.Len()
    node:=(*h)[n-1]
    *h=(*h)[0:n-1]
    return node
}
/**
 * Your StockPrice object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Update(timestamp,price);
 * param_2 := obj.Current();
 * param_3 := obj.Maximum();
 * param_4 := obj.Minimum();
 */