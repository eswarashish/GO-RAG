package main
import ( "math";"fmt";"sync")
// import ("GO-RAG/objects")
// import ("GO-RAG/download")
import ("GO-RAG/datatypes")

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

var wg sync.WaitGroup
func main(){
	// printArea(Circle{Radius: 24.56})
	// printArea(Rectangle{Width: 34.6, Height: 23.54})
	// age := 25
	// fmt.Println(age)  // Output: 25
	// fmt.Println(&age) // Output: 0xc000012088 (The memory address)

	// // * operatoor
	// var ptr *int = &age
	// // fmt.Println(*ptr == age)
	// account := objects.Account{Owner: "Ashish"}
	// personal := objects.PersonalAccount{Account: account }
	// objects.PerfomTransaction(&personal,64.5,true)
	// fmt.Println(personal.GetBalance())
	
	// files := []string{"tags.json", "profile.jpg", "video.mp4", "report.pdf", "config.yaml"}
	// i := 0
	// for i<len(files){
	// 	wg.Add(1)//this only represents the number of routines going to get spawned or called now
	// 	go download.DownloadFile(files[i],&wg)
	// 	i++
	
		
	// }
	// wg.Wait()
	// fmt.Println("All files downloaded successfully")
	// names := datatypes.Names()
	// fmt.Print(names[0:1])
	// numbers := datatypes.Numbers()
	// results := make(map[string]int)
	// high, low := 0,0
	// for _,val := range numbers{
		
	// 	if val>25{
	// 		high++
	// 		key := fmt.Sprintf("High-%d", high)
	// 		results[key] = val
	// 	}else if val<=25{
	// 		low++
	// 		key := fmt.Sprintf("Low-%d", low)
	// 		results[key] = val
	// 	}
	// }
	// fmt.Print(results)

	scores:= datatypes.Map()
	val, ok := scores["Hi"]
	fmt.Print(val,ok)
	}	

