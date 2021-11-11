import (
	"container/heap"
	"sort"
)
type Heap []int
type Node struct{
	friend int
	phase string
	time int
}
type Array []Node
// implementing heap.Interface & sort.Interface
func (h Heap)Len()int{ return len(h) }
func (h Heap)Less(i,j int)bool{ return h[i]<h[j] }
func (h Heap)Swap(i,j int){ h[i],h[j]=h[j],h[i] }
func (h *Heap)Push(x interface{}){ *h=append(*h,x.(int)) }
func (h *Heap)Pop()interface{}{
	n := h.Len()
	val := (*h)[n-1]
	*h=(*h)[0:n-1]
	return val
}
//implement sort.Interface for []Node
func (arr Array)Len()int{ return len(arr) }
func (arr Array)Less(i,j int)bool{
	if arr[i].time<arr[j].time{
		return true
	}
	if arr[i].time==arr[j].time && arr[i].phase=="e"{
		return true
	}
	return false
}
func (arr Array)Swap(i,j int){ arr[i],arr[j]=arr[j],arr[i] }
func smallestChair(times [][]int, targetFriend int) int {
	h:=&Heap{}
	heap.Init(h)
	var arr Array
	for i:=0;i<len(times);i++{
		n1 := Node{friend:i,phase:"s",time:times[i][0]}
		n2 := Node{friend:i,phase:"e",time:times[i][1]}
		arr=append(arr,n1)
		arr=append(arr,n2)
	}
	sort.Sort(arr)
	empty := 0
	//map[friend]seat
	mp := make(map[int]int)
	len := 0
	for _,node := range arr{
		friend,phase,_ :=node.friend,node.phase,node.time
		if phase == "e"{
			heap.Push(h,mp[friend])
			len++
			continue
		}
		if len>0{
			seat := (heap.Pop(h)).(int)  // type assertion
			len--
			mp[friend] = seat
			if friend == targetFriend{
				return seat
			}
			continue
		}
		mp[friend]=empty
        if friend == targetFriend{
            return empty
		}
		empty++
	}
	return empty
}