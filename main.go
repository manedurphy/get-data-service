package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/manedurphy/get-data-service/services"
)

type Final struct {
	Reviews              []services.Review          `json:"reviews"`
	ReviewInfo           services.ReviewInfo        `json:"reviewInfo"`
	NearbyWorkspaces     []services.NearbyWorkspace `json:"nearbyWorkspaces"`
	NearbyTransitOptions []services.TransitOption   `json:"nearbyTransitOptions"`
	Photos               []services.Photo2          `json:"photos"`
}

type URL struct {
	path string
	resp string
}

type Body struct {
	content []byte
	resp    string
}

var urls []URL = []URL{
	{path: os.Getenv("REVIEWS_DOMAIN"), resp: "reviews"},
	{path: os.Getenv("NEARBY_DOMAIN"), resp: "nearby"},
	{path: os.Getenv("LOCATION_DOMAIN"), resp: "transit"},
	{path: os.Getenv("PHOTOS_DOMAIN"), resp: "photos"},
}

func handler(w http.ResponseWriter, r *http.Request) {
	bodyCh := make(chan Body)

	id := strings.TrimPrefix(r.URL.Path, "/api/")
	mapResponses := make(map[string]interface{})

	var wg sync.WaitGroup

	var reviews services.ReviewsResponse
	var nearby services.NearbyResponse
	var transit services.TransitResponse
	var photos []services.Photo2

	var final Final

	mapResponses["reviews"] = &reviews
	mapResponses["nearby"] = &nearby
	mapResponses["transit"] = &transit
	mapResponses["photos"] = &photos

	for _, url := range urls {
		wg.Add(1)
		go func(url URL) {
			body, err := final.GetData(url.path + id)

			if err != nil {
				fmt.Println("Error:", err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}

			b := Body{content: body, resp: url.resp}

			bodyCh <- b

			wg.Done()
		}(url)

	}

	go func() {
		wg.Wait()
		close(bodyCh)
	}()

	for b := range bodyCh {
		json.Unmarshal(b.content, mapResponses[b.resp])
	}

	setReviews(&final, &reviews)
	final.NearbyWorkspaces = nearby.NearbyWorkspaces
	final.NearbyTransitOptions = transit.NearbyTransitOptions
	final.Photos = photos

	finalJson, err := json.Marshal(final)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(finalJson)
}

func (f Final) GetData(url string) ([]byte, error) {
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	defer resp.Body.Close()

	body, e := ioutil.ReadAll(resp.Body)

	if e != nil {
		fmt.Println("ERROR", e)
		return nil, e
	}

	return body, nil
}

func setReviews(f *Final, r *services.ReviewsResponse) {
	if r.Reviews != nil {
		f.Reviews = r.Reviews

	} else {
		f.Reviews = []services.Review{}
	}

	f.ReviewInfo = r.ReviewInfo
}

func main() {
	http.HandleFunc("/api/", handler)
	log.Fatal(http.ListenAndServe(":6003", nil))
}
