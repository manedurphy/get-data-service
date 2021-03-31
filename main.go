package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

type Review struct {
	Rating  int    `json:"rating"`
	Content string `json:"content"`
	Date    string `json:"date"`
	User    User   `json:"user"`
}

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type ReviewsResponse struct {
	Reviews []Review
}

type Post struct {
	UserId int    `json:"userId"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

type Final struct {
	Reviews []Review `json:"reviews"`
	Posts   []Post   `json:"posts"`
}

var urls []string = []string{"http://localhost:5002/api/reviews/all/1", "https://jsonplaceholder.typicode.com/posts"}

func handler(w http.ResponseWriter, r *http.Request) {
	respCh := make(chan *http.Response)
	var wg sync.WaitGroup

	var reviews ReviewsResponse
	var posts []Post
	var final Final

	mapResponses := make(map[string]interface{})

	mapResponses["reviews"] = &reviews
	mapResponses["posts"] = &posts

	go func() {
		wg.Wait()
		close(respCh)
	}()

	for _, url := range urls {
		wg.Add(1)
		go GetData(url, respCh, &wg)
	}

	for resp := range mapResponses {
		current := <-respCh

		c, _ := ioutil.ReadAll(current.Body)
		err := json.Unmarshal(c, mapResponses[resp])

		if err != nil {
			fmt.Println(err)
		}
	}

	final.Reviews = reviews.Reviews
	final.Posts = posts

	finalJson, err := json.Marshal(final)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(finalJson)
}

func GetData(url string, respCh chan<- *http.Response, wg *sync.WaitGroup) {
	defer wg.Done()
	resp, _ := http.Get(url)
	respCh <- resp
}

func main() {
	http.HandleFunc("/api", handler)
	log.Fatal(http.ListenAndServe(":6003", nil))
}
