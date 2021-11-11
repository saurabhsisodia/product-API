type Heap []int
func (h Heap)Len()int{ return len(h) }
func (h Heap)Less(i,j int)bool{ return h[i]>=h[j] }
func (h Heap)Swap(i,j int){ h[i],h[j]=h[j],h[i] }
func (h *Heap)Push(x interface{}){ *h=append(*h,x.(int))}
func (h *Heap)Pop()interface{}{
	n := h.Len()
	val := (*h)[n-1]
	(*h) = (*h)[0:n-1]
	return val 
}
func kthSmallest(matrix [][]int, k int) int {
	h := &Heap{}
	heap.Init(h)
	l := 0
	for i:=0;i<len(matrix);i++{
		for j:=0;j<len(matrix[0]);j++{
			if l==k{
				if matrix[i][j]<=(*h)[0]{
					heap.Pop(h)
					heap.Push(h,matrix[i][j])
				}
			}else{
				heap.Push(h,matrix[i][j])
                l++
			}
		}
	}
	return (*h)[0]
}
