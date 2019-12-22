package main

import {
	"fmt"
	"log"
	"net/http"
}

func homePage(w htttp.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Neelxie my first endpoint")
}