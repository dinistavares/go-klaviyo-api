package klaviyo

import (
	"fmt"
)

// Coupons service
type CouponsService struct {
	service
}

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

type CreateCouponCodeCard struct {
	Data *CreateCouponCode `json:"data,omitempty"`
}

type CreateCouponCode struct {
	Type          string                         `json:"type,omitempty"`
	Attributes    *CreateCouponCodeAttributes    `json:"attributes,omitempty"`
	Relationships *CreateCouponCodeRelationships `json:"relationships,omitempty"`
}

type CreateCouponCodeAttributes struct {
	UniqueCode string `json:"unique_code,omitempty"`
	ExpiresAt  string `json:"expires_at,omitempty"`
}

type CreateCouponCodeRelationships struct {
	Coupon *CreateCouponCodeRelationshipsCoupon `json:"coupon,omitempty"`
}

type CreateCouponCodeRelationshipsCoupon struct {
	Data *Coupon `json:"data,omitempty"`
}

type UpdateCouponCodeCard struct {
	Data *UpdateCouponCode `json:"data,omitempty"`
}

type UpdateCouponCode struct {
	Type       string                      `json:"type,omitempty"`
	ID         string                      `json:"id,omitempty"`
	Attributes *UpdateCouponCodeAttributes `json:"attributes,omitempty"`
}

type UpdateCouponCodeAttributes struct {
	ExpiresAt string `json:"expires_at,omitempty"`
	Status    string `json:"status,omitempty"`
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

	req, _ := service.newRequest("GET", _url, opts, nil)

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

	req, _ := service.newRequest("GET", _url, opts, nil)

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
	coupon.setCouponDataAttributes()

	coupon.Data.Attributes.ExternalID = externalID
}

// Sets new coupon description
func (coupon *CreateCouponCard) SetCouponDescription(description string) {
	coupon.setCouponDataAttributes()

	coupon.Data.Attributes.Description = description
}

// Create coupon. Reference: https://developers.klaviyo.com/en/reference/create_coupon
func (service *CouponsService) CreateCoupon(coupon *CreateCouponCard) (*Response, error) {
	_url := fmt.Sprintf("%s/coupons", ApiTypePrivate)

	// Ensure type is set to "coupon" if empty
	service.setCreateUpdatedCouponType(coupon)

	req, _ := service.newRequest("POST", _url, nil, coupon)

	response, err := service.client.Do(req, nil)

	if err != nil {
		return response, err
	}

	return response, nil
}

// Sets CreateCoupon.Type to 'coupon' if it is not set
func (service *CouponsService) setCreateUpdatedCouponType(coupon *CreateCouponCard) {
	if coupon != nil && coupon.Data != nil && coupon.Data.Type == "" {
		coupon.Data.Type = "coupon"
	}
}

// Ensure coupon data and attribute pointers are created
func (coupon *CreateCouponCard) setCouponDataAttributes() {
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

	req, _ := service.newRequest("GET", _url, opts, nil)

	couponCodes := new(GetCouponCodesResponse)
	response, err := service.client.Do(req, couponCodes)

	if err != nil {
		return nil, response, err
	}

	return couponCodes, response, nil
}

//  ***********************************************************************************
//  CREATE COUPON CODE
//  https://developers.klaviyo.com/en/reference/create_coupon_code
//  ***********************************************************************************

// Sets new coupon unique code
func (coupon *CreateCouponCodeCard) SetCouponUniqueCode(code string) {
	coupon.setCreateCouponCodeDataAttributes()

	coupon.Data.Attributes.UniqueCode = code
}

// Sets new coupon expire
func (coupon *CreateCouponCodeCard) SetCouponExpire(expire string) {
	coupon.setCreateCouponCodeDataAttributes()

	coupon.Data.Attributes.ExpiresAt = expire
}

// Sets new coupon relationship ID
func (coupon *CreateCouponCodeCard) SetCouponRelationshipID(id string) {
	coupon.setCreateCouponCodeDataRelationship()

	coupon.Data.Relationships.Coupon.Data.ID = id
}

// Create coupon code. Reference: https://developers.klaviyo.com/en/reference/create_coupon_code
func (service *CouponsService) CreateCouponCode(couponCode *CreateCouponCodeCard) (*Response, error) {
	_url := fmt.Sprintf("%s/coupon-codes", ApiTypePrivate)

	// Ensure type is set to "coupon-code" if empty
	service.setCreateCouponCodeType(couponCode)

	req, _ := service.newRequest("POST", _url, nil, couponCode)

	response, err := service.client.Do(req, nil)

	if err != nil {
		return response, err
	}

	return response, nil
}

// Sets CreateCouponCode.Type to 'coupon-code' if it is not set
func (service *CouponsService) setCreateCouponCodeType(coupon *CreateCouponCodeCard) {
	if coupon != nil && coupon.Data != nil && coupon.Data.Type == "" {
		coupon.Data.Type = "coupon-code"
	}
}

// Ensure coupon code data and attribute pointers are created
func (coupon *CreateCouponCodeCard) setCreateCouponCodeDataAttributes() {
	if coupon.Data == nil {
		coupon.Data = &CreateCouponCode{}
	}

	if coupon.Data.Attributes == nil {
		coupon.Data.Attributes = &CreateCouponCodeAttributes{}
	}
}

// Ensure coupon code data and relationtionship pointers are created
func (coupon *CreateCouponCodeCard) setCreateCouponCodeDataRelationship() {
	if coupon.Data == nil {
		coupon.Data = &CreateCouponCode{}
	}

	if coupon.Data.Relationships == nil {
		coupon.Data.Relationships = &CreateCouponCodeRelationships{
			Coupon: &CreateCouponCodeRelationshipsCoupon{
				Data: &Coupon{
					Type: "coupon",
				},
			},
		}
	}
}

//  ***********************************************************************************
//  UPDATE COUPON CODE
//  https://developers.klaviyo.com/en/reference/update_coupon_code
//  ***********************************************************************************

// Sets new coupon status
func (coupon *UpdateCouponCodeCard) SetCouponStatus(status string) {
	coupon.setUpdateCouponCodeDataAttributes()

	coupon.Data.Attributes.Status = status
}

// Sets new coupon expire
func (coupon *UpdateCouponCodeCard) SetCouponExpire(expire string) {
	coupon.setUpdateCouponCodeDataAttributes()

	coupon.Data.Attributes.ExpiresAt = expire
}

// Update coupon code. Reference: https://developers.klaviyo.com/en/reference/update_coupon_code
func (service *CouponsService) UpdateCouponCode(id string, couponCode *UpdateCouponCodeCard) (*Response, error) {
	_url := fmt.Sprintf("%s/coupon-codes/%s", ApiTypePrivate, id)

	// Ensure type is set to "coupon-code" if empty
	service.setUpdateCouponCodeType(couponCode)

	// Set coupon ID in body
	couponCode.setCouponID(id)

	req, _ := service.newRequest("PATCH", _url, nil, couponCode)

	response, err := service.client.Do(req, nil)

	if err != nil {
		return response, err
	}

	return response, nil
}

// Sets CreateCouponCode.Type to 'coupon-code' if it is not set
func (service *CouponsService) setUpdateCouponCodeType(coupon *UpdateCouponCodeCard) {
	if coupon != nil && coupon.Data != nil && coupon.Data.Type == "" {
		coupon.Data.Type = "coupon-code"
	}
}

// Ensure coupon code data and attribute pointers are created
func (coupon *UpdateCouponCodeCard) setUpdateCouponCodeDataAttributes() {
	if coupon.Data == nil {
		coupon.Data = &UpdateCouponCode{}
	}

	if coupon.Data.Attributes == nil {
		coupon.Data.Attributes = &UpdateCouponCodeAttributes{}
	}
}

// Sets new coupon unique code
func (coupon *UpdateCouponCodeCard) setCouponID(id string) {
	coupon.setUpdateCouponCodeDataAttributes()

	coupon.Data.ID = id
}