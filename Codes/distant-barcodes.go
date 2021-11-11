type Node struct{
	val,freq int
}
type Heap []Node
func (h Heap)Len()int{ return len(h) }
func (h Heap)Less(i,j int)bool { return h[i].freq>=h[j].freq }
func (h Heap)Swap(i,j int){ h[i],h[j]=h[j],h[i] }
func (h *Heap)Push(x interface{}){ *h = append(*h,x.(Node)) }
func (h *Heap)Pop()interface{}{
	n:=h.Len()
	val := (*h)[n-1]
	*h=(*h)[0:n-1]
	return val
}
func (h Heap)Top()Node{
	return h[0]
}
func rearrangeBarcodes(barcodes []int) []int {
	h:=&Heap{}
	heap.Init(h)
	mp := make(map[int]int)
	for _,v:=range barcodes{
		mp[v]++
	}
	for k,v:=range mp{
		heap.Push(h,Node{k,v})
	}
	ans :=[]int{}
	prev:=-1
	for h.Len()>0 || prev==-1{
		if prev==-1 || h.Top().val!=prev{
			node := heap.Pop(h).(Node)
			ans=append(ans,node.val)
			node.freq--
			if node.freq>0{
				heap.Push(h,node)
			}
			prev=node.val
		}else{
            node1:=heap.Pop(h).(Node)
            node:=heap.Pop(h).(Node)
			ans=append(ans,node.val)
			node.freq--
			heap.Push(h,node1)
			if node.freq>0{
				heap.Push(h,node)
			}
			prev=node.val
		}
	}
    return ans
}