package main
import (
	"fmt"
)

func main(){

	a:=[]int{1,2,3,4,5,6,7}
	fmt.Printf("%p\n",&a)
	f(&a)
	fmt.Println(a)

}

func f(a *[]int){
	//change th Len and may be Cap\
	// as well as pointer to underlying array
	*a=append(*a,10)  
	*a=append(*a,10)
	*a=append(*a,10)
	*a=append(*a,10)
	*a=append(*a,10)
	fmt.Printf("%p\n",a)
	fmt.Println(*a)
}