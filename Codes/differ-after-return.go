package main
import (
	"fmt"
)

func main(){

	fmt.Println(f("Hello"))
}

func f(s string)(res string){
	defer func(){
		res+=" World"
	}()

	return s + " Hii"
}