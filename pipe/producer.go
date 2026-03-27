package pipe

import "fmt"

func GenerateNumbers(c chan int){
	for i :=range 10{
		fmt.Println("Value sent",i+1)
		c <-i+1
	}
	close(c)
}

func SquareNumbers(in chan int, out chan int){
	for val:=range in{
		fmt.Println("Value recieved",val)
		out <- val*val
	}

	close(out)

}