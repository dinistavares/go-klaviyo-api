package klaviyo

import (
	"fmt"
)

// Coupons service
type CouponsService service

type GetCouponsResponse struct {
	Data  *[]Coupon     `json:"data,omitempty"`
	Links *GenericLinks `json:"links,omitempty"`
}

type GetCouponCodesResponse struct {
	Data     *[]CouponCode      `json:"data,omitempty"`
	Links    *GenericLinks      `json:"links,omitempty"`
	Included *[]GenericIncluded `json:"included,omitempty"`
}
type Coupon struct {
	Type       string            `json:"type,omitempty"`
	ID         string            `json:"id,omitempty"`
	Attributes *CouponAttributes `json:"attributes,omitempty"`
	Links      *GenericLinks     `json:"links,omitempty"`
}

type CouponAttributes struct {
	ExternalID  string `json:"external_id,omitempty"`
	Description string `json:"description,omitempty"`
}

type CouponCode struct {
	Type          string                `json:"type,omitempty"`
	ID            string                `json:"id,omitempty"`
	Attributes    *CouponCodeAttributes `json:"attributes,omitempty"`
	Links         *GenericLinks         `json:"links,omitempty"`
	Relationships *Relationships        `json:"relationships,omitempty"`
}

type CouponCodeAttributes struct {
	UniqueCode string `json:"unique_code,omitempty"`
	ExpiresAt  string `json:"expires_at,omitempty"`
	Status     string `json:"status,omitempty"`
}

type CreateCouponCard struct {
	Data *CreateCoupon `json:"data,omitempty"`
}

type CreateCoupon struct {
	Type       string                  `json:"type,omitempty"`
	Attributes *CreateCouponAttributes `json:"attributes,omitempty"`
}

type CreateCouponAttributes struct {
	ExternalID  string `json:"external_id,omitempty"`
	Description string `json:"description,omitempty"`
}

type CouponQueries struct{}

// Query parameters for 'GetCoupons' method.
type GetCouponsQueryParams struct {
	QueryValues
}

// Query parameters for 'GetCouponByID' method.
type GetCouponByIDQueryParams struct {
	QueryValues
}

// Query parameters for 'GetCouponCodes' method.
type GetCouponCodesQueryParams struct {
	QueryValues
}

// Create Query parameters for coupons routes.
func (service *CouponsService) Query() *CouponQueries {
	return &CouponQueries{}
}

//  ***********************************************************************************
//  GET COUPONS
//  https://developers.klaviyo.com/en/reference/get_coupons
//  ***********************************************************************************

// Creates Query parameters for 'GetCoupons'
func (pq CouponQueries) NewGetCoupons() *GetCouponsQueryParams {
	return &GetCouponsQueryParams{
		QueryValues: QueryValues{},
	}
}

// Set coupon fields for for 'GetCoupons' method.
func (p GetCouponsQueryParams) SetCouponFields(values []string) {
	fields := queryFields{}
	fields.setCouponFields(values)

	p.setValues(fields)
}

// Set page cursor for for 'GetCoupons' method.
func (p GetCouponsQueryParams) SetPageCursor(value string) {
	page := queryPage{}
	page.setPageCursor(value)

	p.setValues(page)
}

// Get coupons. Reference: https://developers.klaviyo.com/en/reference/get_coupons
func (service *CouponsService) GetCoupons(opts *GetCouponsQueryParams) (*GetCouponsResponse, *Response, error) {
	_url := fmt.Sprintf("%s/coupons", ApiTypePrivate)

	req, _ := service.client.NewRequest("GET", _url, opts, nil)

	coupons := new(GetCouponsResponse)
	response, err := service.client.Do(req, coupons)

	if err != nil {
		return nil, response, err
	}

	return coupons, response, nil
}

//  ***********************************************************************************
//  GET COUPON BY ID
//  https://developers.klaviyo.com/en/reference/get_coupon
//  ***********************************************************************************

// Creates Query parameters for 'GetCoupons'
func (pq CouponQueries) NewGetCouponByID() *GetCouponByIDQueryParams {
	return &GetCouponByIDQueryParams{
		QueryValues: QueryValues{},
	}
}

// Set coupon fields for for 'GetCouponByID' method.
func (p GetCouponByIDQueryParams) SetCouponFields(values []string) {
	fields := queryFields{}
	fields.setCouponFields(values)

	p.setValues(fields)
}

// Get coupon by ID. Reference: https://developers.klaviyo.com/en/reference/get_coupon
func (service *CouponsService) GetCouponByID(id string, opts *GetCouponsQueryParams) (*GetCouponsResponse, *Response, error) {
	_url := fmt.Sprintf("%s/coupons/%s", ApiTypePrivate, id)

	req, _ := service.client.NewRequest("GET", _url, opts, nil)

	coupons := new(GetCouponsResponse)
	response, err := service.client.Do(req, coupons)

	if err != nil {
		return nil, response, err
	}

	return coupons, response, nil
}

//  ***********************************************************************************
//  CREATE COUPON
//  https://developers.klaviyo.com/en/reference/create_coupon
//  ***********************************************************************************

// Sets new coupon external ID
func (coupon *CreateCouponCard) SetCouponExternalID(externalID string) {
	coupon.setEventDataAttributes()

	coupon.Data.Attributes.ExternalID = externalID
}

// Sets new coupon description
func (coupon *CreateCouponCard) SetCouponDescription(description string) {
	coupon.setEventDataAttributes()

	coupon.Data.Attributes.Description = description
}

// Create coupon. Reference: https://developers.klaviyo.com/en/reference/create_coupon
func (service *CouponsService) CreateCoupon(coupon *CreateCouponCard) (*Response, error) {
	_url := fmt.Sprintf("%s/coupons", ApiTypePrivate)

	// Ensure type is set to "coupon" if empty
	service.setCreateUpdatedType(coupon)

	req, _ := service.client.NewRequest("POST", _url, nil, coupon)

	response, err := service.client.Do(req, nil)

	if err != nil {
		return response, err
	}

	return response, nil
}

// Sets CreateCoupon.Type to 'coupon' if it is not set
func (service *CouponsService) setCreateUpdatedType(coupon *CreateCouponCard) {
	if coupon != nil && coupon.Data != nil && coupon.Data.Type == "" {
		coupon.Data.Type = "coupon"
	}
}

// Ensure coupon data and attribute pointers are created
func (coupon *CreateCouponCard) setEventDataAttributes() {
	if coupon.Data == nil {
		coupon.Data = &CreateCoupon{}
	}

	if coupon.Data.Attributes == nil {
		coupon.Data.Attributes = &CreateCouponAttributes{}
	}
}

//  ***********************************************************************************
//  GET COUPON CODES
//  https://developers.klaviyo.com/en/reference/get_coupon_codes
//  ***********************************************************************************

// Creates Query parameters for 'GetCouponCodes'
func (pq CouponQueries) NewGetCouponCodes() *GetCouponCodesQueryParams {
	return &GetCouponCodesQueryParams{
		QueryValues: QueryValues{},
	}
}

// Set coupon codes fields for for 'GetCouponCodes' method.
func (p GetCouponCodesQueryParams) SetCouponCodesFields(values []string) {
	fields := queryFields{}
	fields.setCouponCodeFields(values)

	p.setValues(fields)
}

// Set coupon fields for for 'GetCouponCodes' method.
func (p GetCouponCodesQueryParams) SetCouponFields(values []string) {
	fields := queryFields{}
	fields.setCouponFields(values)

	p.setValues(fields)
}

// Set filter for for 'GetCouponCodes' method.
func (p GetCouponCodesQueryParams) Filter(filter QueryFilter) {
	p.filter(filter)
}

// Set include for for 'GetCouponCodes' method.
func (p GetCouponCodesQueryParams) Include(values []string) {
	p.include(values)
}

// Set page cursor for for 'GetCouponCodes' method.
func (p GetCouponCodesQueryParams) SetPageCursor(value string) {
	page := queryPage{}
	page.setPageCursor(value)

	p.setValues(page)
}

// Get coupons codes. Reference: https://developers.klaviyo.com/en/reference/get_coupon_codes
func (service *CouponsService) GetCouponCodes(opts *GetCouponCodesQueryParams) (*GetCouponCodesResponse, *Response, error) {
	_url := fmt.Sprintf("%s/coupon-codes", ApiTypePrivate)

	req, _ := service.client.NewRequest("GET", _url, opts, nil)

	couponCodes := new(GetCouponCodesResponse)
	response, err := service.client.Do(req, couponCodes)

	if err != nil {
		return nil, response, err
	}

	return couponCodes, response, nil
}
