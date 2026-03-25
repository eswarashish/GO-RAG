package main
import ( "math";"fmt")

type Shape interface{
	Area() float64
}

type Rectangle struct{
	Width float64
	Height float64
}

func (r Rectangle) Area() float64{
	return r.Height * r.Width
}

type Circle struct{
	Radius float64
}

func (c Circle) Area() float64{
	return (math.Pi * c.Radius) * c.Radius
}

func printArea (s Shape){
	fmt.Println(s.Area())
}

func main(){
	// printArea(Circle{Radius: 24.56})
	// printArea(Rectangle{Width: 34.6, Height: 23.54})
	age := 25
	fmt.Println(age)  // Output: 25
	fmt.Println(&age) // Output: 0xc000012088 (The memory address)

	// * operatoor
	var ptr *int = &age
	fmt.Println(*ptr == age)
}

