package klaviyo

import "fmt"

type GetProfileRelationShipsGenericResponse struct {
	Data *[]GenericRelationshipData `json:"data,omitempty"`
}

//  ***********************************************************************************
//  GET PROFILE RELATIONSHIPS LISTS
//  https://developers.klaviyo.com/en/reference/get_profile_relationships_lists
//  ***********************************************************************************

// Get Profiles. Reference: https://developers.klaviyo.com/en/reference/get_profile_relationships_lists
func (service *ProfilesService) GetProfileRelationshipsLists(id string) (*GetProfileRelationShipsGenericResponse, *Response, error) {
	_url := fmt.Sprintf("%s/profiles/%s/relationships/lists", ApiTypePrivate, id)

	req, _ := service.client.NewRequest("GET", _url, nil, nil)

	lists := new(GetProfileRelationShipsGenericResponse)
	response, err := service.client.Do(req, lists)

	if err != nil {
		return nil, response, err
	}

	return lists, response, nil
}

//  ***********************************************************************************
//  GET PROFILE RELATIONSHIPS SEGMENTS
//  https://developers.klaviyo.com/en/reference/get_profile_relationships_segments
//  ***********************************************************************************

// Get Profiles. Reference: https://developers.klaviyo.com/en/reference/get_profile_relationships_segments
func (service *ProfilesService) GetProfileRelationshipsSegments(id string) (*GetProfileRelationShipsGenericResponse, *Response, error) {
	_url := fmt.Sprintf("%s/profiles/%s/relationships/segments", ApiTypePrivate, id)

	req, _ := service.client.NewRequest("GET", _url, nil, nil)

	segments := new(GetProfileRelationShipsGenericResponse)
	response, err := service.client.Do(req, segments)

	if err != nil {
		return nil, response, err
	}

	return segments, response, nil
}