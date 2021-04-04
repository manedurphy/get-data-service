package services

type NearbyResponse struct {
	Origin           Origin            `json:"origin"`
	NearbyWorkspaces []NearbyWorkspace `json:"nearbyWorkspaces"`
	AllWorkspaceInfo AllWorkspaceInfo  `json:"allWorkspaceInfo"`
	Photos           Photos            `json:"photos"`
}

type Origin struct {
	NearbyWorkspace
	LocationPointer LocationPointer `json:"locationPointer"`
}

type NearbyWorkspace struct {
	Uuid                string `json:"uuid"`
	WorkspaceSlug       string `json:"workspaceSlug"`
	WorkspaceId         int    `json:"workspaceId"`
	RawAddress          string `json:"rawAddress"`
	FormattedAddress    string `json:"formattedAddress"`
	StreetName          string `json:"streetName"`
	StreetNumber        string `json:"streetNumber"`
	Neighborhood        string `json:"neighborhood"`
	City                string `json:"city"`
	State               string `json:"state"`
	Country             string `json:"country"`
	CountryCode         string `json:"countryCode"`
	Zipcode             string `json:"zipcode"`
	LocationPointerUuid string `json:"locationPointerUuid"`
}

type LocationPointer struct {
	Uuid        string `json:"uuid"`
	WorkspaceId string `json:"workspaceId"`
	Geog        Geog
}

type AllWorkspaceInfo struct {
	WorkspaceData            []WorkspaceData          `json:"workspaceData"`
	WorkspaceDescriptionData WorkspaceDescriptionData `json:"workspaceDescriptionData"`
	AmenitiesData            AmenitiesData            `json:"amenitiesData"`
}

type AmenitiesData struct {
	Id        int       `json:"id"`
	Amenities []Amenity `json:"amenities"`
}

type WorkspaceData struct {
	Id             int `json:"id"`
	OfficeCap      int `json:"office_cap"`
	DesksCap       int `json:"desks_cap"`
	MembershipRate int `json:"membership_rate"`
	PassRate       int `json:"pass_rate"`
	RoomRate       int `json:"room_rate"`
}

type WorkspaceDescriptionData struct {
	Id                  int    `josn:"id"`
	Name                string `json:"name"`
	Url                 string `json:"url"`
	DescriptionHeadline string `json:"descriptionHeadline"`
	Description         string `json:"description"`
}

type Amenity struct {
	Id          string `json:"_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Img         string `json:"img"`
}

type Photos struct {
	Id          int    `json:"id"`
	WorkspaceId int    `json:"workspaceId"`
	Url         string `json:"url"`
	Description string `json:"description"`
}

type Geog struct {
	Crs         Crs    `json:"crs"`
	Type        string `json:"type"`
	Coordinates []float64
}

type Crs struct {
	Type       string   `json:"type"`
	Properties Property `json:"properties"`
}

type Property struct {
	Name string `json:"name"`
}
