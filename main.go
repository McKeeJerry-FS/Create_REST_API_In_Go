package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Creating an Article struct
type Article struct {
	Id      string `json:"Id"`
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
	myRouter.HandleFunc("/article", createNewArticle).Methods("POST")
	myRouter.HandleFunc("/article/{id}", getArticleById)

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")
	Articles = []Article{
        {Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
        {Id: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
    }
	handleRequest()
}

// creating a function to retrieve all articles 
func getAllArticles(w http.ResponseWriter, r *http.Request){
	fmt.Println("Endpoint Hit: getAllArticles")
	json.NewEncoder(w).Encode(Articles)
}

// creating as function to return just a single article
func getArticleById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	fmt.Fprintf(w, "Key: " + key)
	// loop over all of the articles
	for _, article := range Articles {
		// if the article's id matches the key that was passed in
		if article.Id == key {
			// return the article
			json.NewEncoder(w).Encode(article)
		}
	}
}

// Creating a function to Create a new article
func createNewArticle(w http.ResponseWriter, r *http.Request){
	reqBody, _ := ioutil.ReadAll(r.Body)
	fmt. Fprintf(w, "Successfully created the new article")
	var article Article
	json.Unmarshal(reqBody, &article)

	Articles = append(Articles, article)
	json.NewEncoder(w).Encode(article)
}