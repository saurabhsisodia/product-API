type Node struct{
	profit float64
	a,b int
}
type Heap []Node
func (h Heap) Len()int{ return len(h) }
func (h Heap) Less(i,j int)bool{ return h[i].profit>h[j].profit }
func (h Heap) Swap(i,j int){ h[i],h[j]=h[j],h[i] }
func (h *Heap) Push(x interface{}){ *h=append(*h,x.(Node)) }
func (h *Heap) Pop() interface{}{
	n:=h.Len()
	val:=(*h)[n-1]
	*h=(*h)[0:n-1]
	return val
}
func maxAverageRatio(classes [][]int, extraStudents int) float64 {
	h:=&Heap{}
	heap.Init(h)   // heapify
	var ans float64
	for i:=0;i<len(classes);i++{
		pass,total:=classes[i][0],classes[i][1]
		ans += float64(pass)/float64(total)
		heap.Push(h,Node{profit(pass,total),pass,total})
	}

	for i:=0;i<extraStudents;i++{
		node := (heap.Pop(h)).(Node)
		added_profit,pass,total := node.profit,node.a,node.b
		ans += added_profit
		heap.Push(h,Node{profit(pass+1,total+1),pass+1,total+1})
	}
    return ans/float64(len(classes))
}
func profit(a,b int)float64{
	return float64(a+1)/float64(b+1) - float64(a)/float64(b)
}