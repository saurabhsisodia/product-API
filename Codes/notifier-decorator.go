package main

import (
	"fmt"
	"log"
)
func main(){
	//fmt.Println(SlackService(EmailService(SMSService(sum)))(10,10))
	fmt.Println(
		SlackService(
			EmailService(
				SMSService(
					sum,
				),
			),
		)(10,10),
	)
}


//core logic
func sum(a,b int)int{
	return a+b
}

//add SMS service
func SMSService(f func(int,int)int)func(int,int)int{
	return func(a,b int)int{
		defer func(){
			log.Println("sending SMS")
		}()
		return f(a,b)
	}
}

// add SlackNotification Service
func SlackService(f func(int,int)int)func(int,int)int{
	return func(a,b int)int{
		defer func(){
			log.Println("sending slack message")
		}()
		return f(a,b)
	}
}

// add EmailService
func EmailService(f func(int,int)int)func(int,int)int{
	return func(a,b int)int{
		defer func(){
			log.Println("sending email msg")
		}()
		return f(a,b)
	}
}