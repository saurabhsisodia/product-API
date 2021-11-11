type Point struct{
	x,y,dist float64
}
type Heap []Point
func (h Heap)Len()int{ return len(h) }
func (h Heap)Less(i,j int)bool { return h[i].dist>h[j].dist }
func (h Heap)Swap(i,j int){ h[i],h[j]=h[j],h[i] }
func (h *Heap)Push(x interface{}){ *h=append(*h,x.(Point)) }
func (h *Heap)Pop()interface{}{
	n := h.Len()
	val := (*h)[n-1]
	*h=(*h)[0:n-1]
	return val
}
func Distance(x,y float64)float64{
	return math.Sqrt(x*x + y*y)
}
func (h Heap)Top()float64{
	return h[0].dist
}
func kClosest(points [][]int, k int) [][]int {
	h := &Heap{}
	heap.Init(h)
	for i:=0;i<len(points);i++{
		x,y := float64(points[i][0]),float64(points[i][1])
		d := Distance(x,y)
		if i<k{
			heap.Push(h,Point{x,y,d})
		}else{
			if d<h.Top(){
				heap.Pop(h)
				heap.Push(h,Point{x,y,d})
			}
		}
	}
    ans := make([][]int,k)
    i:=0
    for _,point := range *h{
        ans[i]=[]int{int(point.x),int(point.y)}
        i++
    }
    return ans
}