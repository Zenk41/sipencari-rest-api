package discussion_locations

import "github.com/Zenk41/sipencari-rest-api/models"

type DiscussionLocation struct {
	LocationID   uint    `json:"location_id"`
	Lat          float64 `json:"lat"`
	Lng          float64 `json:"lng"`
	LocationName string  `json:"location_name"`
	DiscussionID string  `json:"discussion_id"`
}

func DiscussionLocationResponse(dLocation models.DiscussionLocation) *DiscussionLocation {
	return &DiscussionLocation{
		LocationID:   dLocation.LocationID,
		Lat:          dLocation.Lat,
		Lng:          dLocation.Lng,
		DiscussionID: dLocation.DiscussionID,
		LocationName: dLocation.LocationName,
	}
}

func DiscussionLocationsResponse(dLocations []models.DiscussionLocation) *[]DiscussionLocation {
	var dLocationsResponse []DiscussionLocation
	for _, location := range dLocations {
		response := DiscussionLocation{
			LocationID:   location.LocationID,
			Lat:          location.Lat,
			Lng:          location.Lng,
			DiscussionID: location.DiscussionID,
			LocationName: location.LocationName,
		}
		dLocationsResponse = append(dLocationsResponse, response)
	}
	return &dLocationsResponse
}
