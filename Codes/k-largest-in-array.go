package main

import (
	"fmt"
	"container/heap"
)
type Heap []string

// return true if a is smaller than b
func check(a,b string)bool{
	n:=len(a)
	m:=len(b)
	i:=0
    if n<m{
        return true
    }
    if n>m{
        return false
    }
	for i=0;i<n && i<m ;i++{
		if a[i]<b[i]{  // comparision done based on ASCII
			return true
		}
        if a[i]>b[i]{
            return false
        }
	}
	if i>=n && i<m{
		return true
	}
	return false
}
func (h Heap) Len()int{ return len(h) }
func (h Heap) Less(i,j int)bool{ 
	if len(h[i])<len(h[j]){
		return true
	}
	if len(h[i])>len(h[j]){
		return false
	}
	return check(h[i],h[j])
}

func (h Heap) Swap(i,j int){ h[i],h[j]=h[j],h[i] }
func (h *Heap) Push(x interface{}){ *h=append(*h,x.(string)) }
func (h *Heap) Pop() interface{}{
	n:=h.Len()
	val:=(*h)[n-1]
	*h=(*h)[0:n-1]
	return val
}
func main(){
	fmt.Println(kthLargestNumber([]string{"0","0"},2))
}
func kthLargestNumber(nums []string, k int) string {
    
   	h:=&Heap{}
	heap.Init(h)
	for i:=0;i<k;i++{
		heap.Push(h,nums[i])
	}
	for i:=k;i<len(nums);i++{
		
		if !check(nums[i],(*h)[0]){
			heap.Pop(h)
			heap.Push(h,nums[i])
		}
	}
	return (heap.Pop(h)).(string)
}