type Node struct{
	str byte
	count int
}
type Heap []Node
func (h Heap)Len()int{ return len(h) }
func (h Heap)Less(i,j int)bool{ return h[i].count>=h[j].count }
func (h Heap)Swap(i,j int){ h[i],h[j]=h[j],h[i] }
func (h *Heap)Push(x interface{}){ *h = append(*h,x.(Node)) }
func (h *Heap)Pop()interface{}{
	n := h.Len()
	val := (*h)[n-1]
	*h = (*h)[0:n-1]
	return val
}
func (h Heap)Top()Node{
	return h[0]
}
func reorganizeString(s string) string {
	l := 0  // heap Length
	n := len(s)

	mp := make(map[byte]int)
	for i:=0;i<n;i++{
		mp[s[i]]++
	}
	h := &Heap{}
	heap.Init(h)
	for k,v := range mp{
		node := Node{k,v}
		heap.Push(h,node)
		l++
	}
	ans := make([]byte,n)
	i:=0
	var prev byte = 0
	for l>0{
		if i==0 || prev != h.Top().str {
			node := heap.Pop(h).(Node)   // type assertion
			ans[i]=node.str
			node.count--
			l--
			i++
			if node.count>0{
				heap.Push(h,node)
				l++
			}
			prev = node.str
		}else{
			node1 := heap.Pop(h).(Node)
			l--
			if l==0{
				return ""
			}
			node := heap.Pop(h).(Node)
			ans[i]=node.str
			node.count--
			i++
			l--
			heap.Push(h,node1)
            l++
			if node.count>0{
				heap.Push(h,node)
				l++
			}
            prev = node.str
		}
	}
	return string(ans)
}