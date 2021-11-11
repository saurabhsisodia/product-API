package main
import (
	"container/heap"
)
type Heap []int
func (h Heap)Len()int{ return len(h) }
func (h Heap)Less(i,j int)bool{ return h[i]<h[j] }
func (h Heap)Swap(i,j int){ h[i],h[j]=h[j],h[i] }
func (h *Heap)Push(x interface{}){ *h=append(*h,x.(int)) }
func (h *Heap)Pop()interface{}{
	n:=h.Len()
	val:=(*h)[n-1]
	*h=(*h)[0:n-1]
	return val
}
func minStoneSum(piles []int, k int) int{
	h:=&Heap{}
	heap.Init(h)
	for _,v:=range piles{
		heap.Push(h,-v)
	}

	for k>0{
		val:=(heap.Pop(h)).(int)
		val=val - val/2
		heap.Push(h,-val)
	}
	ans := 0
	for i := 0;i<h.Len();i++{
		ans += (*h)[i] 
	}
	return -ans
}