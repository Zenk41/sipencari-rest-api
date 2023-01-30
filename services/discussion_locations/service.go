package discussion_locations

import (
	payload "github.com/Zenk41/sipencari-rest-api/dto/payload/discussion_locations"
	response "github.com/Zenk41/sipencari-rest-api/dto/response/discussion_locations"

	repository "github.com/Zenk41/sipencari-rest-api/repositories/discussion_locations"
)

type DisLocationService interface {
	GetByID(DisLocationID string) (response.DiscussionLocation, error)
	GetAll() ([]response.DiscussionLocation, error)
	Update(payload payload.DiscussionLocation, DisLocationID string, LocationName string) (response.DiscussionLocation, error)
	GetByDiscussionID(dicussionID string) (response.DiscussionLocation, error)
	UpdateByDiscussionID(payload payload.DiscussionLocation, dicussionID string, LocationName string) (response.DiscussionLocation, error)
}

type disLocationService struct {
	repository repository.DisLocationRepository
}

func NewDisLocationService(repository repository.DisLocationRepository) DisLocationService {
	return &disLocationService{repository: repository}
}

func (dls *disLocationService) GetByID(DisLocationID string) (response.DiscussionLocation, error) {
	location, err := dls.repository.GetByID(DisLocationID)
	if err != nil {
		return response.DiscussionLocation{}, err
	}
	return *response.DiscussionLocationResponse(location), nil
}
func (dls *disLocationService) GetAll() ([]response.DiscussionLocation, error) {
	locations, err := dls.repository.GetAll()
	if err != nil {
		return []response.DiscussionLocation{}, err
	}
	return *response.DiscussionLocationsResponse(locations), nil
}

func (dls *disLocationService) Update(payload payload.DiscussionLocation, DisLocationID string, LocationName string) (response.DiscussionLocation, error) {
	location, err := dls.repository.GetByID(DisLocationID)
	if err != nil {
		return response.DiscussionLocation{}, err
	}
	location.Lat = payload.Lat
	location.Lng = payload.Lng
	location.LocationName = LocationName

	updatedLocation, err := dls.repository.Update(location)
	if err != nil {
		return response.DiscussionLocation{}, err
	}
	return *response.DiscussionLocationResponse(updatedLocation), err

}

func (dls *disLocationService) GetByDiscussionID(dicussionID string) (response.DiscussionLocation, error) {
	location, err := dls.repository.GetByDiscussionID(dicussionID)
	if err != nil {
		return response.DiscussionLocation{}, err
	}
	return *response.DiscussionLocationResponse(location), nil
}

func (dls *disLocationService) UpdateByDiscussionID(payload payload.DiscussionLocation, dicussionID string, LocationName string) (response.DiscussionLocation, error) {
	location, err := dls.repository.GetByDiscussionID(dicussionID)
	if err != nil {
		return response.DiscussionLocation{}, err
	}
	location.Lat = payload.Lat
	location.Lng = payload.Lng
	location.LocationName = LocationName

	updatedLocation, err := dls.repository.Update(location)
	if err != nil {
		return response.DiscussionLocation{}, err
	}
	return *response.DiscussionLocationResponse(updatedLocation), err
}
