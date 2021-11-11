package main
import (
	"fmt"
	"sync"
	"time"
)
type Node struct{
	mu sync.Mutex
	val int
}
func main(){
	var wg sync.WaitGroup

	f := func(node1,node2 *Node){
		defer wg.Done()

		node1.mu.Lock()
		defer node1.mu.Unlock()

		time.Sleep(2*time.Second)

		node2.mu.Lock()
		defer node2.mu.Unlock()
		fmt.Println(node1.val+node2.val)
	}
	wg.Add(2)

	var node1,node2 Node
	go f(&node1,&node2)
	go f(&node2,&node1)
	wg.Wait()
}