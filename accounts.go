package klaviyo

import (
	"fmt"
)

// Accounts service
type AccountsService service

type GetAccountsResponse struct {
	Data  *[]Account    `json:"data,omitempty"`
	Links *GenericLinks `json:"links,omitempty"`
}

type GetAccountByIDResponse struct {
	Data  *Account      `json:"data,omitempty"`
	Links *GenericLinks `json:"links,omitempty"`
}

type Account struct {
	Type       string             `json:"type"`
	ID         string             `json:"id"`
	Attributes *AccountAttributes `json:"attributes"`
	Links      *GenericLinks      `json:"links"`
}

type AccountAttributes struct {
	TestAccount        bool                       `json:"test_account"`
	Industry           string                     `json:"industry"`
	Timezone           string                     `json:"timezone"`
	PreferredCurrency  string                     `json:"preferred_currency"`
	PublicAPIKey       string                     `json:"public_api_key"`
	Locale             string                     `json:"locale"`
	ContactInformation *AccountContactInformation `json:"contact_information"`
}

type AccountContactInformation struct {
	DefaultSenderName  string                `json:"default_sender_name"`
	DefaultSenderEmail string                `json:"default_sender_email"`
	WebsiteURL         string                `json:"website_url"`
	OrganizationName   string                `json:"organization_name"`
	StreetAddress      *AccountStreetAddress `json:"street_address"`
}

type AccountStreetAddress struct {
	Address1 string `json:"address1"`
	Address2 string `json:"address2"`
	City     string `json:"city"`
	Region   string `json:"region"`
	Country  string `json:"country"`
	Zip      string `json:"zip"`
}

type AccountQueries struct{}

// Query parameters for 'GetAccounts' method.
type GetAccountsQueryParams struct {
	QueryValues
}

// Query parameters for 'GetAccountByID' method.
type GetAccountByIDQueryParams struct {
	QueryValues
}

// Query parameters for 'GetAccountCodes' method.
type GetAccountCodesQueryParams struct {
	QueryValues
}

// Create Query parameters for accounts routes.
func (service *AccountsService) Query() *AccountQueries {
	return &AccountQueries{}
}

//  ***********************************************************************************
//  GET ACCOUNTS
//  https://developers.klaviyo.com/en/reference/get_accounts
//  ***********************************************************************************

// Creates Query parameters for 'GetAccounts'
func (pq AccountQueries) NewGetAccounts() *GetAccountsQueryParams {
	return &GetAccountsQueryParams{
		QueryValues: QueryValues{},
	}
}

// Set account fields for for 'GetAccounts' method.
func (p GetAccountsQueryParams) SetAccountFields(values []string) {
	fields := queryFields{}
	fields.setAccountFields(values)

	p.setValues(fields)
}

// Get accounts. Reference: https://developers.klaviyo.com/en/reference/get_accounts
func (service *AccountsService) GetAccounts(opts *GetAccountsQueryParams) (*GetAccountsResponse, *Response, error) {
	_url := fmt.Sprintf("%s/accounts", ApiTypePrivate)

	req, _ := service.client.NewRequest("GET", _url, opts, nil)

	accounts := new(GetAccountsResponse)
	response, err := service.client.Do(req, accounts)

	if err != nil {
		return nil, response, err
	}

	return accounts, response, nil
}

//  ***********************************************************************************
//  GET ACCOUNT BY ID
//  https://developers.klaviyo.com/en/reference/get_account
//  ***********************************************************************************

// Creates Query parameters for 'GetAccountByID'
func (pq AccountQueries) NewGetAccountByID() *GetAccountByIDQueryParams {
	return &GetAccountByIDQueryParams{
		QueryValues: QueryValues{},
	}
}

// Set account fields for for 'GetAccountByID' method.
func (p GetAccountByIDQueryParams) SetAccountFields(values []string) {
	fields := queryFields{}
	fields.setAccountFields(values)

	p.setValues(fields)
}

// Get account by ID. Reference: https://developers.klaviyo.com/en/reference/get_account
func (service *AccountsService) GetAccountByID(id string, opts *GetAccountsQueryParams) (*GetAccountByIDResponse, *Response, error) {
	_url := fmt.Sprintf("%s/accounts/%s", ApiTypePrivate, id)

	req, _ := service.client.NewRequest("GET", _url, opts, nil)

	accounts := new(GetAccountByIDResponse)
	response, err := service.client.Do(req, accounts)

	if err != nil {
		return nil, response, err
	}

	return accounts, response, nil
}
