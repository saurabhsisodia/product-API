package main
import (
	"fmt"
	_"time"
	"sync"
)
func main(){
	var data int
	var memoryAccess sync.Mutex
	go func(){
		memoryAccess.Lock()
		data++   // write
		memoryAccess.Unlock()
	}()

	//time.Sleep(1*time.Second)   //sleep for 1 second, u just increases the inefficiency in ur program
	if data==0{   // read
		fmt.Println("data = ",data)
	}
}