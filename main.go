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
	myRouter.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE")
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
	// get the body of the POST request 
	reqBody, _ := ioutil.ReadAll(r.Body)
	fmt. Fprintf(w, "Successfully created the new article")
	var article Article
	// unmarshal it into a new article struct
	json.Unmarshal(reqBody, &article)
	// append the article into the current array of articles
	Articles = append(Articles, article)
	json.NewEncoder(w).Encode(article)
}

//create a function to delete an article 
func deleteArticle(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id := vars["id"]

	for index, article := range Articles{
		if article.Id == id {
			Articles = append(Articles[:index], Articles[index+1:]... )	
			fmt.Println("Article successfully deleted!")	
		}
		fmt.Println("Article was not found")
	}
}