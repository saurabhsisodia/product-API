package main
import (
	"fmt"
	"container/heap"
)
type Node struct{
	val []int  // [i,j,sum]
}
type index struct{
	i,j int
}
type Heap []Node
func (h Heap) Len()int{ return len(h) }
func (h Heap) Less(i,j int)bool { return h[i].val[2]<h[j].val[2] }
func (h Heap) Swap(i,j int){ h[i],h[j]=h[j],h[i] }
func (h *Heap) Push(x interface{}){ *h=append(*h,x.(Node)) }
func (h *Heap) Pop()interface{}{
	n:=h.Len()
	val:=(*h)[n-1]
	*h=(*h)[0:n-1]
	return val
}
func main(){
	fmt.Println(kSmallestPairs([]int{1,7,11},[]int{2,4,6},3))
}
// [i,j,sum]
func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	h:=&Heap{Node{[]int{0,0,nums1[0]+nums2[0]}}}
	heap.Init(h)
	n:=1
	visited := make(map[index]bool)
	visited[index{0,0}]=true
	var ans [][]int
	ans=append(ans,[]int{nums1[0],nums2[0]})
	for k>0 && n>0{
		slice := (heap.Pop(h).(Node)).val
		n--
		k--
		ans=append(ans,[]int{nums1[slice[0]],nums2[slice[1]]})
		if slice[0]+1 < len(nums1){
			if _,ok:=visited[index{slice[0]+1,slice[1]}];!ok{
				heap.Push(h,Node{[]int{slice[0]+1,slice[1],nums1[slice[0]+1]+nums2[slice[1]]}})
				n++
				visited[index{slice[0]+1,slice[1]}]=true
			}
		}
		if slice[1]+1 < len(nums2){
			if _,ok:=visited[index{slice[0],slice[1]+1}];!ok{
				heap.Push(h,Node{[]int{slice[0],slice[1]+1,nums1[slice[0]]+nums2[slice[1]+1]}})
				n++
				visited[index{slice[0],slice[1]+1}]=true
			}
		}
	}
	return ans
}