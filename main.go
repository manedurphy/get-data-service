package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"

	"github.com/manedurphy/reviews-proxy-go/services"
)

// type Review struct {
// 	Rating  int    `json:"rating"`
// 	Content string `json:"content"`
// 	Date    string `json:"date"`
// 	User    User   `json:"user"`
// }

// type ReviewInfo struct {
// 	ReviewCount string `json:"reviewCount"`
// 	Avg         string `json:"avg"`
// }

// type User struct {
// 	FirstName string `json:"first_name"`
// 	LastName  string `json:"last_name"`
// }

// type ReviewsResponse struct {
// 	Reviews    []Review   `json:"reviews"`
// 	ReviewInfo ReviewInfo `json:"reviewInfo"`
// }

type Final struct {
	Reviews    []services.Review   `json:"reviews"`
	ReviewInfo services.ReviewInfo `json:"reviewInfo"`
}

var urls []string = []string{"http://server:5002/api/reviews/all/"}

func handler(w http.ResponseWriter, r *http.Request) {
	respCh := make(chan *http.Response)
	id := strings.TrimPrefix(r.URL.Path, "/api/")

	var wg sync.WaitGroup
	var reviews services.ReviewsResponse
	var final Final

	mapResponses := make(map[string]interface{})
	mapResponses["reviews"] = &reviews

	go func() {
		wg.Wait()
		close(respCh)
	}()

	for _, url := range urls {
		wg.Add(1)
		go GetData(url+id, respCh, &wg)
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
	final.ReviewInfo = reviews.ReviewInfo

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
	http.HandleFunc("/api/", handler)
	log.Fatal(http.ListenAndServe(":6003", nil))
}
