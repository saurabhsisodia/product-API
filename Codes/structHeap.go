package main

import (
	"fmt"
	"container/heap"
)
type User struct{
	username string
	age int
}
type UserHeap []User

func (h UserHeap) Len()int{ return len(h) }
func (h UserHeap) Less(i,j int)bool { return h[i].age<h[j].age }
func (h UserHeap) Swap(i,j int) { h[i],h[j]=h[j],h[i] }
func (h *UserHeap) Push(x interface{}){ *h=append(*h,x.(User)) }
func (h *UserHeap) Pop()interface{} {
	n := h.Len()
	val := (*h)[n-1]
	*h=(*h)[0:n-1]
	return val 
}
func main(){
	h:=&UserHeap{
		User{"saurabh",10},
		User{"Leo",-10},
		User{"Kate",100},
	}
	heap.Init(h)
	fmt.Println((*h)[0])
	heap.Push(h,User{"Hello",-1})
	fmt.Println(heap.Pop(h))
	fmt.Println(heap.Pop(h))
}