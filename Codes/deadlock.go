package main
import (
	"fmt"
	//"sync"
)
func main(){

	ch := make(chan int)
	ch <-1
	fmt.Println(<-ch)
}