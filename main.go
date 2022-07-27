package main

import (
	"fmt"
	"log"
	"net/http"
)

func main(){
	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintln(w, "Hi")
	})
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintln(w, "Hello")
	})
	fmt.Println("Starting server on 8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}