package klaviyo

import (
	"fmt"
	"time"
)

// Events service
type SegmentsService struct {
	service
}

type GetSegmentsResponse struct {
	Data     *[]Segment         `json:"data,omitempty"`
	Links    *GenericLinks      `json:"links,omitempty"`
	Included *[]GenericIncluded `json:"included,omitempty"`
}

type Segment struct {
	Type          string             `json:"type,omitempty"`
	ID            string             `json:"id,omitempty"`
	Attributes    *SegmentAttributes `json:"attributes,omitempty"`
	Links         *GenericLinks      `json:"links,omitempty"`
	Relationships *Relationships     `json:"relationships,omitempty"`
}

type SegmentAttributes struct {
	Name         string    `json:"name,omitempty"`
	Created      time.Time `json:"created,omitempty"`
	Updated      time.Time `json:"updated,omitempty"`
	IsActive     bool      `json:"is_active,omitempty"`
	IsProcessing bool      `json:"is_processing,omitempty"`
	IsStarred    bool      `json:"is_starred,omitempty"`
}

type SegmentQueries struct{}

// Query parameters for 'GetSegments' method.
type GetSegmentQueryParams struct {
	QueryValues
}

// Create Query parameters for segments routes.
func (service *SegmentsService) Query() *SegmentQueries {
	return &SegmentQueries{}
}

//  ***********************************************************************************
//  GET SEGMENTS
//  https://developers.klaviyo.com/en/reference/get_segments
//  ***********************************************************************************

// Creates Query parameters for 'GetSegments'
func (pq SegmentQueries) NewGetSegments() *GetSegmentQueryParams {
	return &GetSegmentQueryParams{
		QueryValues: QueryValues{},
	}
}

// Set segment fields for for 'GetSegments' method.
func (p GetSegmentQueryParams) SetSegmentFields(values []string) {
	fields := queryFields{}
	fields.setSegmentFields(values)

	p.setValues(fields)
}

// Set tag fields for for 'GetSegments' method.
func (p GetSegmentQueryParams) SetTagFields(values []string) {
	fields := queryFields{}
	fields.setTagFields(values)

	p.setValues(fields)
}

// Set filter for for 'GetSegments' method.
func (p GetSegmentQueryParams) Filter(filter QueryFilter) {
	p.filter(filter)
}

// Set include for for 'GetSegments' method.
func (p GetSegmentQueryParams) Include(values []string) {
	p.include(values)
}

// Set page cursor for for 'GetSegments' method.
func (p GetSegmentQueryParams) SetPageCursor(value string) {
	page := queryPage{}
	page.setPageCursor(value)

	p.setValues(page)
}

// Get segments. Reference: https://developers.klaviyo.com/en/reference/get_segments
func (service *SegmentsService) GetSegments(opts *GetSegmentQueryParams) (*GetSegmentsResponse, *Response, error) {
	_url := fmt.Sprintf("%s/segments", ApiTypePrivate)

	req, _ := service.newRequest("GET", _url, opts, nil)

	segments := new(GetSegmentsResponse)
	response, err := service.client.Do(req, segments)

	if err != nil {
		return nil, response, err
	}

	return segments, response, nil
}
