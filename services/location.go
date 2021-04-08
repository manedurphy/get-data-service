package services

type TransitResponse struct {
	NearbyTransitOptions []TransitOption `json:"nearbyTransitOptions"`
}

type TransitOption struct {
	Id          int    `json:"id"`
	AverageCost int    `json:"average_cost"`
	LocationId  int    `json:"location_id"`
	Name        string `json:"name"`
	Type        string `json:"type"`
}
