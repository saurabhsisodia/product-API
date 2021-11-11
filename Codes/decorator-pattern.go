package main
import (
	"fmt"
	"log"
)

func main(){
	//fmt.Println(logging(sum)(1,2))
	//fmt.Println(logging(diff)(1,2))

	// use as many decorator as u want as long as interface if uniform

	fmt.Println(logging(authenticate(sum))(1,2))

}

// another service
// inject extra functionalities without touching core business logic
func logging(f func(int,int)int)func(int,int)int{

	return func(a,b int)int{

		defer func(){
			log.Printf("doing logging")
		}()
		return f(a,b)
	}
}
func authenticate(f func(int,int)int)func(int,int)int{
	return func(a,b int)int{
		defer func(){
			log.Println("Authentication handled here")
		}()
		return f(a,b)
	}
}

// core logic
func sum(a,b int)int{
	return a+b
}
func diff(a,b int)int{
	return a-b
}