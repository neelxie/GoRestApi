package main

import {
	"fmt"
	"log"
	"net/http"
	"encoding/json"
}

type Article struct{
	Title string `json:"Title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
}
type Articles []Article

func allArticles(w htttp.ResponseWriter, r *http.Request){
	articles := Articles{
		Articles{Title:"Test Title", Desc: "Test Description", Content: "Test Content"},
	}

	fmt.Println("Endpoint Hits All artcles endpoint")
	json.NewEncoder(w).Encode(articles)
}
func homePage(w htttp.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Neelxie my first endpoint")
}

# the about page
func aboutPage(w htttp.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "This is the about page.")


func handleRequest(){
	http.handleFunc("/", homepage)
	http.handleFunc("/allarticles", allArticles)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main(){
// to run type go run and the file name which is main.go
	handleRequest()
}
