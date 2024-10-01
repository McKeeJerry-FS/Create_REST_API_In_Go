package main

import (
	"fmt"
	"log"
	"net/http"
)

// Creating an Article struct
type Article struct {
	Title 	string `json:"Title"`
	Desc  	string `json:"Description"`
	Content string `json:"Content"`
}

// Declaring an array of Article
var Articles []Article


func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequest(){
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	Articles = []Article{
        {Title: "Hello", Desc: "Article Description", Content: "Article Content"},
        {Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
    }
	handleRequest()
}