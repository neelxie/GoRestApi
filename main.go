package main

import {
	"fmt"
	"log"
	"net/http"
	"encoding/json"
}

}
func homePage(w htttp.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Neelxie my first endpoint")
}

# the about page
func aboutPage(w htttp.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "This is the about page.")


func handleRequest(){
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main(){
	handleRequests()
}

// to run type go run and the file name which is main.go
