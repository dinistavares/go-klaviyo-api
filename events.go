package klaviyo

import (
	"fmt"
)

// Events service
type EventsService service

type GetEventsResponse struct {
	Data     *[]Event           `json:"data,omitempty"`
	Links    *GenericLinks      `json:"links,omitempty"`
	Included *[]GenericIncluded `json:"included,omitempty"`
}

type GetEventByIDResponse struct {
	Data     *Event             `json:"data,omitempty"`
	Included *[]GenericIncluded `json:"included,omitempty"`
}

type CreateEventCard struct {
	Data *CreateEvent `json:"data,omitempty"`
}

type CreateEvent struct {
	Type       string                 `json:"type,omitempty"`
	ID         string                 `json:"id,omitempty"`
	Attributes *CreateEventAttributes `json:"attributes,omitempty"`
}

type CreateEventAttributes struct {
	UniqueID      string                       `json:"unique_id,omitempty"`
	Properties    interface{}                  `json:"properties,omitempty"`
	Value         float64                      `json:"value,omitempty"`
	ValueCurrency string                       `json:"value_currency,omitempty"`
	Time          string                       `json:"time,omitempty"`
	Profile       *CreateUpdateProfile         `json:"profile,omitempty"`
	Metric        *CreateEventAttributesMetric `json:"metric,omitempty"`
}

type CreateEventAttributesMetric struct {
	Data *CreateEventAttributesMetricData `json:"data,omitempty"`
}

type CreateEventAttributesMetricData struct {
	Type       string                                     `json:"type,omitempty"`
	Attributes *CreateEventAttributesMetricDataAttributes `json:"attributes,omitempty"`
}

type CreateEventAttributesMetricDataAttributes struct {
	Name    string `json:"name,omitempty"`
	Service string `json:"service,omitempty"`
}

type Event struct {
	Type          string           `json:"type,omitempty"`
	ID            string           `json:"id,omitempty"`
	Attributes    *EventAttributes `json:"attributes,omitempty"`
	Links         *GenericLinks    `json:"links,omitempty"`
	Relationships *Relationships   `json:"relationships,omitempty"`
}

type EventAttributes struct {
	Timestamp       int         `json:"timestamp,omitempty"`
	EventProperties interface{} `json:"event_properties,omitempty"`
	Datetime        string      `json:"datetime,omitempty"`
	UUID            string      `json:"uuid,omitempty"`
}

type EventQueries struct{}

// Query parameters for 'GetEvents' method.
type GetEventsQueryParams struct {
	QueryValues
}

// Query parameters for 'GetEventByID' method.
type GetEventByIDQueryParams struct {
	QueryValues
}

// Create Query parameters for event routes.
func (service *EventsService) Query() *EventQueries {
	return &EventQueries{}
}

//  ***********************************************************************************
//  GET EVENTS
//  https://developers.klaviyo.com/en/reference/get_events
//  ***********************************************************************************

// Creates Query parameters for 'GetEvents'
func (pq EventQueries) NewGetEvents() *GetEventsQueryParams {
	return &GetEventsQueryParams{
		QueryValues: QueryValues{},
	}
}

// Set profile fields for for 'GetEvents' method.
func (p GetEventsQueryParams) SetProfileFields(values []string) {
	fields := queryFields{}
	fields.setProfileFields(values)

	p.setValues(fields)
}

// Set event fields for for 'GetEvents' method.
func (p GetEventsQueryParams) SetEventFields(values []string) {
	fields := queryFields{}
	fields.setEventFields(values)

	p.setValues(fields)
}

// Set metric fields for for 'GetEvents' method.
func (p GetEventsQueryParams) SetMetricFields(values []string) {
	fields := queryFields{}
	fields.setMetricFields(values)

	p.setValues(fields)
}

// Set filter for for 'GetEvents' method.
func (p GetEventsQueryParams) Filter(filter QueryFilter) {
	p.filter(filter)
}

// Set sort for for 'GetEvents' method.
func (p GetEventsQueryParams) Include(values []string) {
	p.include(values)
}

// Set page cursor for for 'GetEvents' method.
func (p GetEventsQueryParams) SetPageCursor(value string) {
	page := queryPage{}
	page.setPageCursor(value)

	p.setValues(page)
}

// Set sort for for 'GetEvents' method.
func (p GetEventsQueryParams) Sort(value string) {
	p.sort(value)
}

// Get events. Reference: https://developers.klaviyo.com/en/reference/get_events
func (service *EventsService) GetEvents(opts *GetEventsQueryParams) (*GetEventsResponse, *Response, error) {
	_url := fmt.Sprintf("%s/events", ApiTypePrivate)

	req, _ := service.client.NewRequest("GET", _url, opts, nil)

	events := new(GetEventsResponse)
	response, err := service.client.Do(req, events)

	if err != nil {
		return nil, response, err
	}

	return events, response, nil
}

//  ***********************************************************************************
//  GET EVENT
//  https://developers.klaviyo.com/en/reference/get_event
//  ***********************************************************************************

// Creates Query parameters for 'NewGetEventByID'
func (pq EventQueries) NewGetEventByID() *GetEventByIDQueryParams {
	return &GetEventByIDQueryParams{
		QueryValues: QueryValues{},
	}
}

// Set event fields for for 'GetEventByID' method.
func (p GetEventByIDQueryParams) SetEventFields(values []string) {
	fields := queryFields{}
	fields.setEventFields(values)

	p.setValues(fields)
}

// Set metric fields for for 'GetEventByID' method.
func (p GetEventByIDQueryParams) SetMetricFields(values []string) {
	fields := queryFields{}
	fields.setMetricFields(values)

	p.setValues(fields)
}

// Set profile fields for for 'GetEventByID' method.
func (p GetEventByIDQueryParams) SetProfileFields(values []string) {
	fields := queryFields{}
	fields.setProfileFields(values)

	p.setValues(fields)
}

// Set sort for for 'GetEventByID' method.
func (p GetEventByIDQueryParams) Include(values []string) {
	p.include(values)
}

// Get event. Reference: https://developers.klaviyo.com/en/reference/get_event
func (service *EventsService) GetEventByID(id string, opts *GetEventByIDQueryParams) (*GetEventByIDResponse, *Response, error) {
	_url := fmt.Sprintf("%s/events/%s", ApiTypePrivate, id)

	req, _ := service.client.NewRequest("GET", _url, opts, nil)

	events := new(GetEventByIDResponse)
	response, err := service.client.Do(req, events)

	if err != nil {
		return nil, response, err
	}

	return events, response, nil
}

//  ***********************************************************************************
//  CREATE EVENTS
//  https://developers.klaviyo.com/en/reference/create_events
//  ***********************************************************************************

// Sets new event metric
func (event *CreateEventCard) SetEventMetric(name string, service string) {
	event.setEventDataAttributes()

	event.Data.Attributes.Metric = &CreateEventAttributesMetric{
		Data: &CreateEventAttributesMetricData{
			Type: "metric",
			Attributes: &CreateEventAttributesMetricDataAttributes{
				Name:    name,
				Service: service,
			},
		},
	}
}

// Sets new event profile
func (event *CreateEventCard) SetEventProfile(profile *Profile) {
	event.setEventDataAttributes()

	if profile != nil && profile.Type == "" {
		profile.Type = "profile"
	}

	event.Data.Attributes.Profile = &CreateUpdateProfile{
		Data: profile,
	}
}

// Sets new event properties
func (event *CreateEventCard) SetEventProperties(properties interface{}) {
	event.setEventDataAttributes()

	event.Data.Attributes.Properties = properties
}

// Create event. Reference: https://developers.klaviyo.com/en/reference/create_events
func (service *EventsService) CreateEvent(event *CreateEventCard) (*Response, error) {
	_url := fmt.Sprintf("%s/events", ApiTypePrivate)

	// Ensure type is set to "event" if empty
	service.setCreateUpdatedType(event)

	req, _ := service.client.NewRequest("POST", _url, nil, event)

	events := new(GetEventsResponse)
	response, err := service.client.Do(req, events)

	if err != nil {
		return response, err
	}

	return response, nil
}

// Sets CreateEvent.Type to 'event' if it is not set
func (service *EventsService) setCreateUpdatedType(event *CreateEventCard) {
	if event != nil && event.Data != nil && event.Data.Type == "" {
		event.Data.Type = "event"
	}
}

// Ensure event data and attribute pointers are created
func (event *CreateEventCard) setEventDataAttributes() {
	if event.Data == nil {
		event.Data = &CreateEvent{}
	}

	if event.Data.Attributes == nil {
		event.Data.Attributes = &CreateEventAttributes{}
	}
}
