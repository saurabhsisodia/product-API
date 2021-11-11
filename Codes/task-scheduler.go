// implement sort.Interface


type Heap []int
func (h Heap) Len()int{ return len(h) }
func (h Heap) Less(i,j int)bool { return h[i]>h[j] }
func (h Heap) Swap(i,j int){ h[i],h[j] = h[j],h[i] }

// implement heap.Interface

func (h *Heap) Push(x interface{}) { *h = append(*h,x.(int)) }
func (h *Heap) Pop() interface{}{
	n := h.Len()
	val := (*h)[n-1]
	*h = (*h)[0:n-1]
	return val
}
func leastInterval(tasks []byte, n int) int {
	mp := make(map[byte]int)

	for i:=0;i<len(tasks);i++{
		mp[tasks[i]]++
	}
	var arr []int
	for _,v:=range mp{
		arr=append(arr,v)
	}
    h := &Heap{}
    for _,v:=range arr{
        *h=append(*h,v)
    }
	heap.Init(h)
	ans,l := 0,len(arr)
	for l>0{
		interval,k := n+1,0
		tmp := make([]int,min(interval,l))
		for interval > 0 && l>0{
			val := heap.Pop(h).(int)  // type assertion
			l--
			tmp[k]=val
			k++
			ans++
			interval--
		}

		for i:=0;i<k;i++{
			if tmp[i]-1>0{
				heap.Push(h,tmp[i]-1)
				l++
			}
		}
		// if interval is > 0, means we have to leave empty slots
		if l<=0{
			break
		}
		ans += interval

	}
    return ans
}
func min(a,b int)int{
	if a<=b{
		return a
	}
	return b
}