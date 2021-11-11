type Node struct{
	count,expiry int
}
type Heap []Node
func (h Heap)Len()int { return len(h) }
func (h Heap)Less(i,j int)bool{ return h[i].expiry<h[j].expiry }
func (h Heap)Swap(i,j int){ h[i],h[j]=h[j],h[i] }
func (h *Heap)Push(x interface{}){ *h=append(*h,x.(Node))}
func (h *Heap)Pop()interface{}{
	n := h.Len()
	val :=(*h)[n-1]
	*h = (*h)[0:n-1]
	return val
}
func eatenApples(apples []int, days []int) int {

	h :=&Heap{}
	heap.Init(h)
	l := 0
	ans := 0
    n := len(apples)
    for i:=0;i<n || l>0;i++{
		if i<n{
			heap.Push(h,Node{apples[i],i+days[i]})
			l++
		}
        for l>0 && ((*h)[0].expiry<=i || (*h)[0].count==0){
			heap.Pop(h)
			l--
		}
        if l>0{
            ans++
            (*h)[0].count--
            if (*h)[0].count<=0{
                heap.Pop(h)
                l--
            }
        }
	}
	return ans
}