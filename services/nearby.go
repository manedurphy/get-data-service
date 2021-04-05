package services

type NearbyResponse struct {
	NearbyWorkspaces []NearbyWorkspace `json:"nearbyWorkspaces"`
}

type NearbyWorkspace struct {
	WorkspaceId    int       `json:"workspaceId"`
	Name           string    `json:"name"`
	Neighborhood   string    `json:"neighborhood"`
	Photo          Photo     `json:"photo"`
	MembershipRate int       `json:"membership_rate"`
	Amenities      []Amenity `json:"amenities"`
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
