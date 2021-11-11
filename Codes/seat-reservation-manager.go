type Heap []int

//implement heap.Interface

func (h Heap)Len()int{ return len(h) }
func (h Heap)Less(i,j int)bool{ return h[i]<h[j] }
func (h Heap)Swap(i,j int){ h[i],h[j]=h[j],h[i] }
func (h *Heap)Push(x interface{}){ *h=append(*h,x.(int)) }
func (h *Heap)Pop()interface{}{
    n:=h.Len()
    val:=(*h)[n-1]
    *h=(*h)[0:n-1]
    return val
}
type SeatManager struct {
    h *Heap
    curr int
    length int
}
func Constructor(n int) SeatManager {
    h:=&Heap{}
    heap.Init(h)
    return SeatManager{h:h,curr:1,length:0}
}
func (this *SeatManager) Reserve() int {
    if this.length==0{
        this.curr++
        return this.curr-1
    }
    this.length--
    return heap.Pop(this.h).(int)
}
func (this *SeatManager) Unreserve(seatNumber int)  {
    heap.Push(this.h,seatNumber)
    this.length++
}


/**
 * Your SeatManager object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Reserve();
 * obj.Unreserve(seatNumber);
 */