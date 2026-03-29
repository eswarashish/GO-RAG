package main

import (
	"log"
	"net/http"
)

func main(){
	log.Println("Server starting")
	if err:= http.ListenAndServe(":8080",nil); err!=nil{
		log.Fatal("failed to start server:", err)
	}


}