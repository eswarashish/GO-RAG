package pipe

import "fmt"

func Producer(c chan int){
	for i :=0; i < 5; i++{
		fmt.Println("Value sent",i)
		c <-i 
	}
	close(c)
}