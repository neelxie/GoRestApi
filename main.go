package main

import {
	"fmt"
	"log"
	"net/http"
}

func homePage(w htttp.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Neelxie my first endpoint")
}

func handleRequest(){
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main(){
	handleRequests()
}

// to run type go run and the file name which is main.go