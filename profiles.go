package klaviyo

import (
	"fmt"
	"time"
)

// Profiles service
type ProfilesService struct {
	service
}

type GetProfilesResponse struct {
	Data  *[]Profile    `json:"data,omitempty"`
	Links *GenericLinks `json:"links,omitempty"`
}

type GetProfileByIDResponse struct {
	Data     *Profile           `json:"data,omitempty"`
	Included *[]ProfileIncluded `json:"included,omitempty"`
}

type GetProfileListsResponse struct {
	Data  *[]List       `json:"data,omitempty"`
	Links *GenericLinks `json:"links,omitempty"`
}

type GetProfileSegmentsResponse struct {
	Data  *[]Segment    `json:"data,omitempty"`
	Links *GenericLinks `json:"links,omitempty"`
}

type CreateUpdateProfile struct {
	Data *Profile `json:"data,omitempty"`
}

type Profile struct {
	Type          string             `json:"type,omitempty"`
	ID            string             `json:"id,omitempty"`
	Attributes    *ProfileAttributes `json:"attributes,omitempty"`
	Links         *GenericLinks      `json:"links,omitempty"`
	Relationships *Relationships     `json:"relationships,omitempty"`
	Meta          *ProfileMeta       `json:"meta,omitempty"`
}

type ProfileAttributes struct {
	Email               string                                `json:"email,omitempty"`
	PhoneNumber         string                                `json:"phone_number,omitempty"`
	ExternalID          string                                `json:"external_id,omitempty"`
	FirstName           string                                `json:"first_name,omitempty"`
	LastName            string                                `json:"last_name,omitempty"`
	Organization        string                                `json:"organization,omitempty"`
	Title               string                                `json:"title,omitempty"`
	Image               string                                `json:"image,omitempty"`
	Created             string                                `json:"created,omitempty"`
	Updated             string                                `json:"updated,omitempty"`
	LastEventDate       string                                `json:"last_event_date,omitempty"`
	Properties          interface{}                           `json:"properties,omitempty"`
	Location            *ProfileAttributesLocation            `json:"location,omitempty"`
	Subscriptions       *ProfileAttributesSubscriptions       `json:"subscriptions,omitempty"`
	PredictiveAnalytics *ProfileAttributesPredictiveAnalytics `json:"predictive_analytics,omitempty"`
}

type ProfileAttributesLocation struct {
	Address1  string      `json:"address1,omitempty"`
	Address2  string      `json:"address2,omitempty"`
	City      string      `json:"city,omitempty"`
	Country   string      `json:"country,omitempty"`
	Region    string      `json:"region,omitempty"`
	Zip       string      `json:"zip,omitempty"`
	Timezone  string      `json:"timezone,omitempty"`
	IP        string      `json:"ip,omitempty"`
	Latitude  interface{} `json:"latitude,omitempty"`
	Longitude interface{} `json:"longitude,omitempty"`
}

type ProfileAttributesSubscriptions struct {
	Email *ProfileAttributesSubscriptionsEmail `json:"email,omitempty"`
	Sms   *ProfileAttributesSubscriptionsSms   `json:"sms,omitempty"`
}

type ProfileAttributesSubscriptionsEmail struct {
	Marketing *ProfileAttributesSubscriptionsEmailMarketing `json:"marketing,omitempty"`
}

type ProfileAttributesSubscriptionsSms struct {
	Marketing *ProfileAttributesSubscriptionsSmsMarketing `json:"marketing,omitempty"`
}

type ProfileAttributesSubscriptionsGenericMarketing struct {
	Consent          string     `json:"consent,omitempty"`
	Method           string     `json:"method,omitempty"`
	MethodDetail     string     `json:"method_detail,omitempty"`
	ConsentTimestamp *time.Time `json:"consent_timestamp,omitempty"`
	LastUpdated      *time.Time `json:"last_updated,omitempty"`
}

type ProfileAttributesSubscriptionsEmailMarketing struct {
	CanReceiveEmailMarketing bool                                                            `json:"can_receive_email_marketing,omitempty"`
	DoubleOptin              bool                                                            `json:"double_optin,omitempty"`
	CustomMethodDetail       string                                                          `json:"custom_method_detail,omitempty"`
	Suppression              *[]ProfileAttributesSubscriptionsEmailMarketingSuppression      `json:"suppression,omitempty"`
	ListSuppressions         *[]ProfileAttributesSubscriptionsEmailMarketingListSuppressions `json:"list_suppressions,omitempty"`
	ProfileAttributesSubscriptionsGenericMarketing
}

type ProfileAttributesSubscriptionsSmsMarketing struct {
	CanReceiveSmsMarketing bool `json:"can_receive_sms_marketing,omitempty"`
	ProfileAttributesSubscriptionsGenericMarketing
}

type ProfileAttributesSubscriptionsEmailMarketingSuppression struct {
	Reason    string     `json:"reason,omitempty"`
	Timestamp *time.Time `json:"timestamp,omitempty"`
}

type ProfileAttributesSubscriptionsEmailMarketingListSuppressions struct {
	ListID    string     `json:"list_id,omitempty"`
	Reason    string     `json:"reason,omitempty"`
	Timestamp *time.Time `json:"timestamp,omitempty"`
}

type ProfileAttributesPredictiveAnalytics struct {
	HistoricClv              float64 `json:"historic_clv,omitempty"`
	PredictedClv             float64 `json:"predicted_clv,omitempty"`
	TotalClv                 float64 `json:"total_clv,omitempty"`
	HistoricNumberOfOrders   int     `json:"historic_number_of_orders,omitempty"`
	PredictedNumberOfOrders  float64 `json:"predicted_number_of_orders,omitempty"`
	AverageDaysBetweenOrders int     `json:"average_days_between_orders,omitempty"`
	AverageOrderValue        float64 `json:"average_order_value,omitempty"`
	ChurnProbability         float64 `json:"churn_probability,omitempty"`
	ExpectedDateOfNextOrder  string  `json:"expected_date_of_next_order,omitempty"`
}

type ProfileMeta struct {
	PatchProperties *ProfilePatchProperties `json:"patch_properties,omitempty"`
}

type ProfilePatchProperties struct {
	Append   interface{} `json:"append,omitempty"`
	Unappend interface{} `json:"unappend,omitempty"`
	Unset    string      `json:"unset,omitempty"`
}

type ProfileIncludedAttributes struct {
	Name         string `json:"name,omitempty"`
	Created      string `json:"created,omitempty"`
	Updated      string `json:"updated,omitempty"`
	OptInProcess string `json:"opt_in_process,omitempty"`
	IsActive     bool   `json:"is_active,omitempty"`
	IsProcessing bool   `json:"is_processing,omitempty"`
	IsStarred    bool   `json:"is_starred,omitempty"`
}

type ProfileIncluded struct {
	Type       string                     `json:"type,omitempty"`
	ID         string                     `json:"id,omitempty"`
	Attributes *ProfileIncludedAttributes `json:"attributes,omitempty"`
	Links      *GenericLinks              `json:"links,omitempty"`
}

type ProfileQueries struct{}

// Query parameters for 'GetProfiles' method.
type GetProfilesQueryParams struct {
	QueryValues
}

// Query parameters for 'GetProfileByID' method.
type GetProfileByIDQueryParams struct {
	QueryValues
}

type GetProfileListsQueryParams struct {
	QueryValues
}

type GetProfileSegmentsQueryParams struct {
	QueryValues
}

// Create Query parameters for profile routes.
func (service *ProfilesService) Query() *ProfileQueries {
	return &ProfileQueries{}
}

//  ***********************************************************************************
//  GET PROFILES
//  https://developers.klaviyo.com/en/reference/get_profiles
//  ***********************************************************************************

// Creates Query parameters for 'GetProfiles'
func (pq ProfileQueries) NewGetProfiles() *GetProfilesQueryParams {
	return &GetProfilesQueryParams{
		QueryValues: QueryValues{},
	}
}

// Set profile fields for for 'GetProfiles' method.
func (p GetProfilesQueryParams) SetProfileFields(values []string) {
	fields := queryFields{}
	fields.setProfileFields(values)

	p.setValues(fields)
}

// Set profile additional fields for for 'GetProfiles' method.
func (p GetProfilesQueryParams) SetProfileAdditionalFields(values []string) {
	additionalFields := queryAdditionalFields{}
	additionalFields.SetProfileAdditionalFields(values)

	p.setValues(additionalFields)
}

// Set page cursor for for 'GetProfiles' method.
func (p GetProfilesQueryParams) SetPageCursor(value string) {
	page := queryPage{}
	page.setPageCursor(value)

	p.setValues(page)
}

// Set page size for for 'GetProfiles' method.
func (p GetProfilesQueryParams) SetPageSize(value int) {
	page := queryPage{}
	page.setPageSize(value)

	p.setValues(page)
}

// Set sort for for 'GetProfiles' method.
func (p GetProfilesQueryParams) Sort(value string) {
	p.sort(value)
}

// Set filter for for 'GetProfiles' method.
func (p GetProfilesQueryParams) Filter(filter QueryFilter) {
	p.filter(filter)
}

// Get Profiles. Reference: https://developers.klaviyo.com/en/reference/get_profiles
func (service *ProfilesService) GetProfiles(opts *GetProfilesQueryParams) (*GetProfilesResponse, *Response, error) {
	_url := fmt.Sprintf("%s/profiles", ApiTypePrivate)

	req, _ := service.newRequest("GET", _url, opts, nil)

	profiles := new(GetProfilesResponse)
	response, err := service.client.Do(req, profiles)

	if err != nil {
		return nil, response, err
	}

	return profiles, response, nil
}

//  ***********************************************************************************
//  GET PROFILES BY ID
//  https://developers.klaviyo.com/en/reference/get_profile
//  ***********************************************************************************

// Creates Query parameters for 'GetProfileByID'
func (pq ProfileQueries) NewGetProfileByID() *GetProfileByIDQueryParams {
	return &GetProfileByIDQueryParams{
		QueryValues: QueryValues{},
	}
}

// Set profile fields for for 'GetProfileByID' method.
func (p GetProfileByIDQueryParams) SetProfileFields(values []string) {
	fields := queryFields{}
	fields.setProfileFields(values)

	p.setValues(fields)
}

// Set list fields for for 'GetProfileByID' method.
func (p GetProfileByIDQueryParams) SetListFields(values []string) {
	fields := queryFields{}
	fields.setListFields(values)

	p.setValues(fields)
}

// Set segment fields for for 'GetProfileByID' method.
func (p GetProfileByIDQueryParams) SetSegmentFields(values []string) {
	fields := queryFields{}
	fields.setSegmentFields(values)

	p.setValues(fields)
}

// Set include for for 'GetProfiles' method.
func (p GetProfileByIDQueryParams) Include(values []string) {
	p.include(values)
}

// Get Profiles. Reference: https://developers.klaviyo.com/en/reference/get_profile
func (service *ProfilesService) GetProfileByID(id string, opts *GetProfileByIDQueryParams) (*GetProfileByIDResponse, *Response, error) {
	_url := fmt.Sprintf("%s/profiles/%s", ApiTypePrivate, id)

	req, _ := service.newRequest("GET", _url, opts, nil)

	profile := new(GetProfileByIDResponse)
	response, err := service.client.Do(req, profile)

	if err != nil {
		return nil, response, err
	}

	return profile, response, nil
}

//  ***********************************************************************************
//  CREATE PROFILE
//  https://developers.klaviyo.com/en/reference/create_profile
//  ***********************************************************************************

// Create a new profile. Reference: https://developers.klaviyo.com/en/reference/create_profile
func (service *ProfilesService) CreateProfile(profile *CreateUpdateProfile) (*CreateUpdateProfile, *Response, error) {
	_url := fmt.Sprintf("%s/profiles", ApiTypePrivate)

	// Ensure type is set to "profile" if empty
	service.setCreateUpdatedType(profile)

	req, _ := service.newRequest("POST", _url, nil, profile)

	newProfile := new(CreateUpdateProfile)
	response, err := service.client.Do(req, newProfile)

	if err != nil {
		return nil, response, err
	}

	return newProfile, response, nil
}

//  ***********************************************************************************
//  UPDATE PROFILE
//  https://developers.klaviyo.com/en/reference/update_profile
//  ***********************************************************************************

// Create a new profile. Reference: https://developers.klaviyo.com/en/reference/update_profile
func (service *ProfilesService) UpdateProfile(id string, profile *CreateUpdateProfile) (*CreateUpdateProfile, *Response, error) {
	_url := fmt.Sprintf("%s/profiles/%s", ApiTypePrivate, id)

	// Ensure type is set to "profile" if empty
	service.setCreateUpdatedType(profile)

	req, _ := service.newRequest("PATCH", _url, nil, profile)

	newProfile := new(CreateUpdateProfile)
	response, err := service.client.Do(req, newProfile)

	if err != nil {
		return nil, response, err
	}

	return newProfile, response, nil
}

//  ***********************************************************************************
//  CREATE OR UPDATE PROFILE
// 	https://developers.klaviyo.com/en/reference/create_or_update_profile
//  ***********************************************************************************

// Create a new profile. Reference: https://developers.klaviyo.com/en/reference/create_or_update_profile
func (service *ProfilesService) CreateOrUpdateProfile(profile *CreateUpdateProfile) (*CreateUpdateProfile, *Response, error) {
	_url := fmt.Sprintf("%s/profile-import", ApiTypePrivate)

	// Ensure type is set to "profile" if empty
	service.setCreateUpdatedType(profile)

	req, _ := service.newRequest("POST", _url, nil, profile)

	newProfile := new(CreateUpdateProfile)
	response, err := service.client.Do(req, newProfile)

	if err != nil {
		return nil, response, err
	}

	return newProfile, response, nil
}

//  ***********************************************************************************
//  GET PROFILE LISTS
// 	https://developers.klaviyo.com/en/reference/get_profile_lists
//  ***********************************************************************************

// Creates Query parameters for 'GetProfileLists'
func (pq ProfileQueries) NewGetProfileLists() *GetProfileListsQueryParams {
	return &GetProfileListsQueryParams{
		QueryValues: QueryValues{},
	}
}

// Set list fields for for 'GetProfileLists' method.
func (p GetProfileListsQueryParams) SetListFields(values []string) {
	fields := queryFields{}
	fields.setListFields(values)

	p.setValues(fields)
}

// Get Profiles. Reference: https://developers.klaviyo.com/en/reference/get_profile_lists
func (service *ProfilesService) GetProfileLists(id string, opts *GetProfileListsQueryParams) (*GetProfileListsResponse, *Response, error) {
	_url := fmt.Sprintf("%s/profiles/%s/lists", ApiTypePrivate, id)

	req, _ := service.newRequest("GET", _url, opts, nil)

	lists := new(GetProfileListsResponse)
	response, err := service.client.Do(req, lists)

	if err != nil {
		return nil, response, err
	}

	return lists, response, nil
}

//  ***********************************************************************************
//  GET PROFILE SEGMENTS
//  https://developers.klaviyo.com/en/reference/get_profile_segments
//  ***********************************************************************************

// Creates Query parameters for 'GetProfileSegments'
func (pq ProfileQueries) NewGetProfileSegments() *GetProfileSegmentsQueryParams {
	return &GetProfileSegmentsQueryParams{
		QueryValues: QueryValues{},
	}
}

// Set list fields for for 'GetProfileSegments' method.
func (p GetProfileSegmentsQueryParams) SetListFields(values []string) {
	fields := queryFields{}
	fields.setSegmentFields(values)

	p.setValues(fields)
}

// Get Profiles. Reference: https://developers.klaviyo.com/en/reference/get_profile_segments
func (service *ProfilesService) GetProfileSegments(id string, opts *GetProfileSegmentsQueryParams) (*GetProfileSegmentsResponse, *Response, error) {
	_url := fmt.Sprintf("%s/profiles/%s/segments", ApiTypePrivate, id)

	req, _ := service.newRequest("GET", _url, opts, nil)

	lists := new(GetProfileSegmentsResponse)
	response, err := service.client.Do(req, lists)

	if err != nil {
		return nil, response, err
	}

	return lists, response, nil
}

// Sets CreateUpdateProfile.Type to 'profile' if it is not set
func (service *ProfilesService) setCreateUpdatedType(profile *CreateUpdateProfile) {
	if profile != nil && profile.Data != nil && profile.Data.Type == "" {
		profile.Data.Type = "profile"
	}
}
