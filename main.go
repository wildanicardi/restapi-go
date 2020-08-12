package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// User is a model for articles
type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Article is a model for articles
type Article struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	User        *User  `json:"user"`
}

// ini User var as a slice
// var articles []User
var articles []Article

// set json resp

func setJSONResp(res http.ResponseWriter /*, httpCode int*/) {
	res.Header().Set("Content-type", "application/json")
	// res.WriteHeader(httpCode)
}

//index articles
func getArticles(res http.ResponseWriter, req *http.Request) {
	setJSONResp(res)
	json.NewEncoder(res).Encode(articles)
}

//Show Detail Article
func getArticle(res http.ResponseWriter, req *http.Request) {
	setJSONResp(res)

	params := mux.Vars(req)
	for _, item := range articles {
		if item.ID == params["id"] {
			json.NewEncoder(res).Encode(item)
			return
		}
	}
	json.NewEncoder(res).Encode(&Article{})
}

//Add Article
func createArticle(res http.ResponseWriter, req *http.Request) {
	var newArticle Article
	json.NewDecoder(req.Body).Decode(&newArticle)
	newArticle.ID = strconv.Itoa(len(articles) + 1)
	newArticle.User.ID = strconv.Itoa(len(articles) + 1)
	articles = append(articles, newArticle)
	json.NewEncoder(res).Encode(newArticle)
}

//Update Article
func updateArticle(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for i, item := range articles {
		if item.ID == params["id"] {
			articles = append(articles[:i], articles[i+1:]...)
			var newArticle Article
			json.NewDecoder(req.Body).Decode(&newArticle)
			newArticle.ID = params["id"]
			articles = append(articles, newArticle)
			json.NewEncoder(res).Encode(newArticle)
		}
	}
}

//Delete Single Article
func deleteArticle(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for i, item := range articles {
		if item.ID == params["id"] {
			articles = append(articles[:i], articles[i+1:]...)
			break
		}
	}
}
func main() {
	//Generata Mock Data

	articles = append(articles,
		Article{
			ID:          "1",
			Title:       "Belajar RESTFULL API GO",
			Description: "Belajar harus sungguh-sungguh",
			User: &User{
				ID:    "1",
				Name:  "Ali",
				Email: "ali@gmail.com",
			}},
		Article{
			ID:          "2",
			Title:       "Belajar Basic GO",
			Description: "Belajar harus sungguh-sungguh",
			User: &User{
				ID:    "2",
				Name:  "AliWildan",
				Email: "aliwildan@gmail.com",
			}},
	)
	// init router
	router := mux.NewRouter()

	//Handle Endpoint Routing
	router.HandleFunc("/api/articles", getArticles).Methods("GET")
	router.HandleFunc("/api/article/{id}", getArticle).Methods("GET")
	router.HandleFunc("/api/articles", createArticle).Methods("POST")
	router.HandleFunc("/api/article/{id}", updateArticle).Methods("POST")
	router.HandleFunc("/api/article/{id}", deleteArticle).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":5000", router))
}
