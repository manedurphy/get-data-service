package services

type PhotosResponse struct {
	Photos []Photo2 `json:"photos"`
}

type Photo2 struct {
	WorkspaceId string `json:"workspaceid"`
	StayDetails string `json:"staydetails"`
	Url         string `json:"url"`
	Description string `json:"description"`
}
