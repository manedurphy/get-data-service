package services

type Review struct {
	Rating  int    `json:"rating"`
	Content string `json:"content"`
	Date    string `json:"date"`
	User    User   `json:"user"`
}

type ReviewInfo struct {
	ReviewCount string `json:"reviewCount"`
	Avg         string `json:"avg"`
}

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type ReviewsResponse struct {
	Reviews    []Review   `json:"reviews"`
	ReviewInfo ReviewInfo `json:"reviewInfo"`
}
