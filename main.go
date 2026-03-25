package main


import (
	"fmt"
	"errors"
)

// Variables
const nam string = "asdf"
var name int = 9

func main() {
	// name := "asdf"
    // fmt.Printf("Hello, %s Senior Dev. The environment is ready.",name)
	// for i :=0; i<5;i++{
	// 	fmt.Println(i)	}
	// for i := range 7{
	// 	fmt.Println(i)
	// }
	// count:=0
	// for count<10{
	// 	fmt.Println(count)
	// 	count++	}

	ans, err := divide(10,0)

	if err != nil{
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Answer is:",ans)
}

func divide (a int, b int) (int, error){
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a/b, nil
}




