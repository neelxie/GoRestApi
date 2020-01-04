package main

import {
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
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

# tests POST method
func testPostArticles(w htttp.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "The POST endpoint worked.")
}

func homePage(w htttp.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Neelxie my first endpoint")
}

# the about page
func aboutPage(w htttp.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "This is the about page.")


func handleRequest(){
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.handleFunc("/", homepage)
	myRouter.handleFunc("/about", aboutPage)
	myRouter.handleFunc("/allarticles", allArticles).Methods("GET")
	myRouter.handleFunc("/allarticles", testPostArticles).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main(){
// to run type go run and the file name which is main.go
	handleRequest()
}
