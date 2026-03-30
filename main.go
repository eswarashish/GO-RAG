package main

import (
	// "GO-RAG/channels" // import ("GO-RAG/objects")
	// "GO-RAG/pipe"
	// import ("GO-RAG/download")
	// "GO-RAG/datatypes"
	// "GO-RAG/middleware"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"sync"
	"net"
	"log"
	pb "GO-RAG/proto"
	"google.golang.org/grpc"
	"time"
	// "time"
)

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

type Todo struct {
    UserID    int    `json:"userId"` // The tag tells Go which JSON key to map to
    ID        int    `json:"id"`
    Title     string `json:"title"`
    Completed bool   `json:"completed"`
}

// To turn JSON bytes into a Struct:
var myTodo Todo

type Post struct{
	ID int `json:"id"`
	Title string `json:"title"`
}

type server struct{
	pb.UnimplementedCalculatorServer
}

var myPost Post
var wg *sync.WaitGroup


type App struct{
	client *http.Client
}

func (app *App) GetPostHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	if id == "" {
		http.Error(w, "Missing 'id' parameter", http.StatusBadRequest)
		return // MUST return so it stops processing!
	}
	
	url := fmt.Sprintf("https://jsonplaceholder.typicode.com/posts/%s", id)
	
	// Create context
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel() // Safe to defer right away

	// Create Request
	req, _ := http.NewRequestWithContext(ctx, "GET", url, nil)

	// Execute
	client := app.client
	resp, err := client.Do(req)
	
	// THE CRITICAL CHECK
	if err != nil {
		http.Error(w, err.Error(), http.StatusGatewayTimeout)
		return // MUST return here so we don't touch a nil resp!
	}
	
	// Now we know resp is safe
	defer resp.Body.Close()
	
	body, _ := io.ReadAll(resp.Body)
	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}

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

	// scores:= datatypes.Map()
	// val, ok := scores["Hi"]
	// fmt.Print(val,ok)
	// ch := make(chan int)
	// // go channels.Ping(ch)
	// // fmt.Println("waiting for ping")
	// // msg:= <- ch
	// // fmt.Println(msg)
	// go pipe.Producer(ch)

	// for val:= range ch{
	// 	fmt.Println("Value recieved", val)
	// }

	// nums := make(chan int,10)
	// results:= make(chan int,10)

	// go pipe.GenerateNumbers(nums)
	// go pipe.SquareNumbers(nums,results)
	// for val:= range results{
	// 	fmt.Println("Value sqaured is",val)
	// }
	// resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	// if err != nil{
	// 	fmt.Println(err)
	// 	return
	// }
	// defer resp.Body.Close()

	// body,err := io.ReadAll(resp.Body)
	// json.Unmarshal(body, &myTodo)
	// fmt.Println(myTodo)

	// urls:= []string{"https://jsonplaceholder.typicode.com/posts/1","https://jsonplaceholder.typicode.com/posts/2","https://jsonplaceholder.typicode.com/posts/3"}
	
	// for i:= range len(urls){
	// 	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	// 	wg.Add(1)
	// 	go fetchPost(urls[i],&wg,ctx,cancel)
		

	// }
	// wg.Wait()
	
	// // 2. The "Router" - Map the URL to the Handler
	// // http.HandleFunc("/ping", pingHandler)

	// fmt.Println("🚀 Server is running on http://localhost:8080")
	
	// // 3. Start the server (This is an infinite loop that blocks forever)
	// err := http.ListenAndServe(":8080", nil)
	// if err != nil {
	// 	fmt.Println("Server crashed:", err)
	// }

	lis,_:= net.Listen("tcp",":50001")
	s := grpc.NewServer()
	pb.RegisterCalculatorServer(s, &server{})
	log.Println("Server listening on :50001")
	s.Serve(lis)

	}	

func (s *server) Add(ctx context.Context, in *pb.AddRequest) (*pb.AddResponse, error){
	log.Printf("Received: %v + %v", in.A, in.B)
	return  &pb.AddResponse{Result: in.A + in.B},nil
}
	// 1. This is the "Handler" (Like your FastAPI endpoint)
func pingHandler(w http.ResponseWriter, r *http.Request) {
	// Let's create a map to act as our JSON response
	response := map[string]string{
		"status":  "success",
		"message": "Pong! The Go server is alive.",
	}

	// Tell the browser "Hey, I'm sending JSON!"
	w.Header().Set("Content-Type", "application/json")

	// Encode the map into JSON and write it directly to the 'w' (ResponseWriter)
	json.NewEncoder(w).Encode(response)
}
func fetchPost(url string, wg *sync.WaitGroup, ctx context.Context, cancel context.CancelFunc){
	
	defer cancel()
	defer wg.Done()
	req, _ := http.NewRequestWithContext(ctx,"GET",url,nil)
	client:=&http.Client{}
	resp,err:=client.Do(req)
	if err!=nil{
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	body, _:= io.ReadAll(resp.Body)
	// json.Unmarshal(body,&myPost)
	fmt.Println(string(body))
}

