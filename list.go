package klaviyo

import (
	"fmt"
	"time"
)

// Lists service
type ListsService service

type GetListsResponse struct {
	Data     *[]List            `json:"data,omitempty"`
	Links    *GenericLinks      `json:"links,omitempty"`
	Included *[]GenericIncluded `json:"included,omitempty"`
}

type List struct {
	Type          string          `json:"type,omitempty"`
	ID            string          `json:"id,omitempty"`
	Attributes    *ListAttributes `json:"attributes,omitempty"`
	Links         *GenericLinks   `json:"links,omitempty"`
	Relationships *Relationships  `json:"relationships,omitempty"`
}

type ListAttributes struct {
	Name         string     `json:"name,omitempty"`
	OptInProcess string     `json:"opt_in_process,omitempty"`
	Created      *time.Time `json:"created,omitempty"`
	Updated      *time.Time `json:"updated,omitempty"`
}

type AddRemoveProfiles struct {
	Data *[]Profile `json:"data,omitempty"`
}

type ListQueries struct{}

// Query parameters for 'GetLists' method.
type GetListsQueryParams struct {
	QueryValues
}

// Create Query parameters for lists routes.
func (service *ListsService) Query() *ListQueries {
	return &ListQueries{}
}

//  ***********************************************************************************
//  GET LISTS
//  https://developers.klaviyo.com/en/reference/get_lists
//  ***********************************************************************************

// Creates Query parameters for 'GetLists'
func (pq ListQueries) NewGetLists() *GetListsQueryParams {
	return &GetListsQueryParams{
		QueryValues: QueryValues{},
	}
}

// Set list fields for for 'GetLists' method.
func (p GetListsQueryParams) SetListFields(values []string) {
	fields := queryFields{}
	fields.setListFields(values)

	p.setValues(fields)
}

// Set tag fields for for 'GetLists' method.
func (p GetListsQueryParams) SetTagFields(values []string) {
	fields := queryFields{}
	fields.setTagFields(values)

	p.setValues(fields)
}

// Set filter for for 'GetLists' method.
func (p GetListsQueryParams) Filter(filter QueryFilter) {
	p.filter(filter)
}

// Set include for for 'GetLists' method.
func (p GetListsQueryParams) Include(values []string) {
	p.include(values)
}

// Set page cursor for for 'GetLists' method.
func (p GetListsQueryParams) SetPageCursor(value string) {
	page := queryPage{}
	page.setPageCursor(value)

	p.setValues(page)
}

// Get lists. Reference: https://developers.klaviyo.com/en/reference/get_lists
func (service *ListsService) GetLists(opts *GetListsQueryParams) (*GetListsResponse, *Response, error) {
	_url := fmt.Sprintf("%s/lists", ApiTypePrivate)

	req, _ := service.client.NewRequest("GET", _url, opts, nil)

	lists := new(GetListsResponse)
	response, err := service.client.Do(req, lists)

	if err != nil {
		return nil, response, err
	}

	return lists, response, nil
}

//  ***********************************************************************************
//  ADD PROFILE TO LIST
//  https://developers.klaviyo.com/en/reference/create_list_relationships
//  ***********************************************************************************

// Add profile to lists. Reference: https://developers.klaviyo.com/en/reference/create_list_relationships
func (service *ListsService) AddProfileToList(listID string, profiles *AddRemoveProfiles) (*Response, error) {
	_url := fmt.Sprintf("%s/lists/%s/relationships/profiles", ApiTypePrivate, listID)

	service.setAddRemoveProfilesType(profiles)

	req, _ := service.client.NewRequest("POST", _url, nil, profiles)

	response, err := service.client.Do(req, nil)

	if err != nil {
		return response, err
	}

	return response, nil
}

//  ***********************************************************************************
//  REMOVE PROFILE FROM LIST
//  https://developers.klaviyo.com/en/reference/delete_list_relationships
//  ***********************************************************************************

// Remove profile from list. Reference: https://developers.klaviyo.com/en/reference/delete_list_relationships
func (service *ListsService) RemoveProfileFromList(listID string, profiles *AddRemoveProfiles) (*Response, error) {
	_url := fmt.Sprintf("%s/lists/%s/relationships/profiles", ApiTypePrivate, listID)

	service.setAddRemoveProfilesType(profiles)

	req, _ := service.client.NewRequest("DELETE", _url, nil, profiles)

	response, err := service.client.Do(req, nil)

	if err != nil {
		return response, err
	}

	return response, nil
}

// Sets Profile.Type to 'profile' if it is not set in any RemoveProfile
func (service *ListsService) setAddRemoveProfilesType(profiles *AddRemoveProfiles) {
	if profiles != nil && profiles.Data != nil {
		for i, profile := range *profiles.Data {
			if profile.Type == "" {
				(*profiles.Data)[i].Type = "profile"
			}
		}
	}
}
