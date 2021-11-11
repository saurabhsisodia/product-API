type Heap []int
func (h Heap)Len()int { return len(h) }
func (h Heap)Less(i,j int)bool { return h[i]>h[j] }  // for max heap
func (h Heap)Swap(i,j int){ h[i],h[j]=h[j],h[i] }
func (h *Heap)Push(x interface{}){ *h=append(*h,x.(int)) }
func (h *Heap)Pop()interface{}{
	n:=h.Len()
	val :=(*h)[n-1]
	*h=(*h)[0:n-1]
	return val
}
func maximumScore(a int, b int, c int) int {
	h = &Heap{a,b,c}
	heap.Init(h)
	ans := 0

	for true{
		first,second :=heap.Pop(h).(int),heap.Pop(h).(int)
		if first != 0 && second != 0{
			ans ++
			heap.Push(h,first-1)
			heap.Push(h,second-1)
			continue
		}
		break
	}
	return ans
}