package main
import (
	"fmt"
	"container/heap"
)
type IntHeap []int
// implement heap.Interface
/*
	type Interface interface{  //heap
		sort.Interface 
		Push(x interface{})
		Pop() interface{}
	}
	type Interface interface  //sort{
		Len()int
		Less(i,j int)bool
		Swap(i,j int)
	}
*/
func (h IntHeap) Len()int{ return len(h) }
func (h IntHeap) Less(i,j int)bool{ return h[i]<h[j]}
func (h IntHeap) Swap (i,j int) { h[i],h[j]=h[j],h[i]}
// Push and pop are changing Len and cap so, underlying array can change, so use Pointer receiver
func (h *IntHeap) Push(x interface{}){
	*h=append(*h,x.(int))  // type assertion to fetch Dynamic value, if dynamic type is int
}
func (h *IntHeap) Pop() interface{}{
	n:=h.Len()
	val:=(*h)[n-1]
	*h=(*h)[0:n-1]
	return val
}
func main(){
	h:=&IntHeap{1,2,3,4,5,0,-2,10,-5,100}
	// initialize heap in O(n) where n = h.Len()
	heap.Init(h)
	fmt.Println((*h)[0])
	for i:=0;i<5;i++{
		heap.Push(h,i)
	}
	heap.Push(h,10)

	fmt.Println(heap.Pop(h))

}