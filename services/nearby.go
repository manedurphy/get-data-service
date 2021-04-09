package services

type NearbyResponse struct {
	NearbyWorkspaces []NearbyWorkspace `json:"nearbyWorkspaces"`
}

type NearbyWorkspace struct {
	WorkspaceId  int    `json:"workspaceId"`
	Name         string `json:"name"`
	Neighborhood string `json:"neighborhood"`
	StreetName   string `json:"streetName"`
	StreetNumber string `json:"streetNumber"`
	Photo        Photo  `json:"photo"`
	Rate         string `json:"rate"`
	Amenities    string `json:"amenities"`
}

type Amenity struct {
	Id          string `json:"_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Img         string `json:"img"`
}

type Photo struct {
	Id          int    `json:"id"`
	WorkspaceId int    `json:"workspaceId"`
	Url         string `json:"url"`
	Description string `json:"description"`
}
