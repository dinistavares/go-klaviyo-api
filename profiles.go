package klaviyo

import (
	"fmt"
	"time"
)

// Profiles service
type ProfilesService service

type ProfileResponse struct {
	Data  *[]Profile    `json:"data,omitempty"`
	Links *GenericLinks `json:"links,omitempty"`
}

type Profile struct {
	Type          string                `json:"type,omitempty"`
	ID            string                `json:"id,omitempty"`
	Attributes    *ProfileAttributes    `json:"attributes,omitempty"`
	Links         *GenericLinks         `json:"links,omitempty"`
	Relationships *ProfileRelationships `json:"relationships,omitempty"`
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
	Location            *ProfileAttributesLocation            `json:"location,omitempty"`
	Properties          *ProfileAttributesProperties          `json:"properties,omitempty"`
	Subscriptions       *ProfileAttributesSubscriptions       `json:"subscriptions,omitempty"`
	PredictiveAnalytics *ProfileAttributesPredictiveAnalytics `json:"predictive_analytics,omitempty"`
}

type ProfileAttributesLocation struct {
	Address1  string `json:"address1,omitempty"`
	Address2  string `json:"address2,omitempty"`
	City      string `json:"city,omitempty"`
	Country   string `json:"country,omitempty"`
	Latitude  string `json:"latitude,omitempty"`
	Longitude string `json:"longitude,omitempty"`
	Region    string `json:"region,omitempty"`
	Zip       string `json:"zip,omitempty"`
	Timezone  string `json:"timezone,omitempty"`
	IP        string `json:"ip,omitempty"`
}

type ProfileAttributesProperties struct {
	Pseudonym string `json:"pseudonym,omitempty"`
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
	CustomMethodDetail       string                                                          `json:"custom_method_detail,omitempty"`
	DoubleOptin              string                                                          `json:"double_optin,omitempty"`
	Suppression              *[]ProfileAttributesSubscriptionsEmailMarketingSuppression      `json:"suppression,omitempty"`
	ListSuppressions         *[]ProfileAttributesSubscriptionsEmailMarketingListSuppressions `json:"list_suppressions,omitempty"`
	ProfileAttributesSubscriptionsGenericMarketing
}

type ProfileAttributesSubscriptionsSmsMarketing struct {
	CanReceiveSmsMarketing bool `json:"can_receive_sms_marketing,omitempty"`
	ProfileAttributesSubscriptionsGenericMarketing
}

type ProfileAttributesSubscriptionsEmailMarketingSuppression struct {
	Reason    string    `json:"reason,omitempty"`
	Timestamp *time.Time `json:"timestamp,omitempty"`
}

type ProfileAttributesSubscriptionsEmailMarketingListSuppressions struct {
	ListID    string    `json:"list_id,omitempty"`
	Reason    string    `json:"reason,omitempty"`
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

type RelationshipLists struct {
	Links *GenericLinks `json:"links,omitempty"`
}

type RelationshipSegments struct {
	Links *GenericLinks `json:"links,omitempty"`
}

type ProfileRelationships struct {
	Lists    *RelationshipLists    `json:"lists,omitempty"`
	Segments *RelationshipSegments `json:"segments,omitempty"`
}

type ProfileQueries struct {}

// Query parameters for 'GetProfiles' method.
type GetProfilesQueryParams struct {
	QueryValues
}

// Query parameters for 'GetProfileByID' method.
type GetProfileByIDQueryParams struct {
	QueryValues
}

func (service *ProfilesService) Query() *ProfileQueries {
	return &ProfileQueries{}
}

//  ***********************************************************************************
//  GET PROFILES (https://developers.klaviyo.com/en/reference/get_profiles)
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
func (service *ProfilesService) GetProfiles(opts *GetProfilesQueryParams) (*ProfileResponse, *Response, error) {
	_url := fmt.Sprintf("%s/profiles", ApiTypePrivate)

	req, _ := service.client.NewRequest("GET", _url, opts, nil)

	profiles := new(ProfileResponse)
	response, err := service.client.Do(req, profiles)

	if err != nil {
		return nil, response, err
	}

	return profiles, response, nil
}

//  ***********************************************************************************
//  GET PROFILES BY ID (https://developers.klaviyo.com/en/reference/get_profile)
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

// Set sort for for 'GetProfiles' method.
func (p GetProfileByIDQueryParams) Include(values []string) {
	p.include(values)
}

// Get Profiles. Reference: https://developers.klaviyo.com/en/reference/get_profile
func (service *ProfilesService) GetProfileByID(id string, opts *GetProfileByIDQueryParams) (*ProfileResponse, *Response, error) {
	_url := fmt.Sprintf("%s/profiles/%s", ApiTypePrivate, id)

	req, _ := service.client.NewRequest("GET", _url, opts, nil)

	profiles := new(ProfileResponse)
	response, err := service.client.Do(req, profiles)

	if err != nil {
		return nil, response, err
	}

	return profiles, response, nil
}
