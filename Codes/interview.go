// implement merge sort
package main
import (
	"fmt"
	"strings"
)
func main(){
	var s strings.Builder
	for i:=0;i<5;i++{
		s.Write([]byte("a"))
	}
	fmt.Println(s.String())   // implement stringer
}
/*
func main(){
	square,quit := make(chan int),make(chan bool)
	ans := 0

	go func(){
		for i:=1;i<=5;i++{
			ans+=<-square
		}
		quit<-true
	}()
	findSum(square,quit)
	fmt.Println(ans)
}

func findSum(square chan int,quit chan bool){
		i:=1
		for {
			select{
			case square<-i*i:
				i++
			case <-quit:
				return
			}
		}
}

func MergeSort(arr []int)[]int{
	if len(arr)<=1{
		return arr
	}
	leftDone,rightDone := make(chan bool),make(chan bool)
	mid := len(arr)/2
	left := []int{}
	right :=[]int{}
	go func(){
		left = MergeSort(arr[:mid])
		leftDone<-true
	}()
	go func(){
		right = MergeSort(arr[mid:])
		rightDone<-true
	}()
	<-leftDone
	<-rightDone         // Merge will be called only if left and right both are done
	return Merge(left,right)
}

func Merge(left,right []int)[]int{
	n,m:=len(left),len(right)
	tmp := make([]int,0,n+m)

	for n>0 || m>0{
		if n==0{
			return append(tmp,right...)
		}else if m==0{
			return append(tmp,left...)
		}else if left[0]<right[0]{
			tmp=append(tmp,left[0])
			left=left[1:]
			n--
		}else{
			tmp=append(tmp,right[0])
			right=right[1:]
			m--
		}
	}
	return tmp
}
*/