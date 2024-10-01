package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
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
	// homepage route
	// http.HandleFunc("/", homePage)
	// articles route
	// http.HandleFunc("/articles", getAllArticles)

	// now utilizing Gorilla/Mux Routing
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/all", getAllArticles)

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")
	Articles = []Article{
        {Title: "Hello", Desc: "Article Description", Content: "Article Content"},
        {Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
    }
	handleRequest()
}

// creating a function to retrieve all articles 
func getAllArticles(w http.ResponseWriter, r *http.Request){
	fmt.Println("Endpoint Hit: getAllArticles")
	json.NewEncoder(w).Encode(Articles)
}