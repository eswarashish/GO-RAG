package channels

import "fmt"

func Ping(c chan string){
	fmt.Println("Pinging")
	msg:= "pong"
	c <- msg
}