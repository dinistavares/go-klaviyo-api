package klaviyo

import (
	"fmt"
	"time"
)

// Webhooks service
type WebhooksService service

type GetWebhooksResponse struct {
	Data     *[]Webhook         `json:"data,omitempty"`
	Links    *GenericLinks      `json:"links,omitempty"`
	Included *[]GenericIncluded `json:"included,omitempty"`
}

type Webhook struct {
	Type          string                     `json:"type,omitempty"`
	ID            string                     `json:"id,omitempty"`
	Attributes    *WebhookAttributes         `json:"attributes,omitempty"`
	Links         *GenericLinks              `json:"links,omitempty"`
	Relationships *RelationshipWebhookTopics `json:"relationships,omitempty"`
}

type WebhookAttributes struct {
	Name        string     `json:"name,omitempty"`
	Description string     `json:"description,omitempty"`
	EndpointURL string     `json:"endpoint_url,omitempty"`
	Enabled     bool       `json:"enabled,omitempty"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
}

type WebhookQueries struct{}

// Query parameters for 'GetWebhooks' method.
type GetWebhooksQueryParams struct {
	QueryValues
}

// Create Query parameters for event routes.
func (service *WebhooksService) Query() *WebhookQueries {
	return &WebhookQueries{}
}

//  ***********************************************************************************
//  GET WEBHOOKS
//  https://developers.klaviyo.com/en/reference/get_webhook_subscriptions
//  ***********************************************************************************

// Creates Query parameters for 'GetWebhooks'
func (pq WebhookQueries) NewGetWebhooks() *GetWebhooksQueryParams {
	return &GetWebhooksQueryParams{
		QueryValues: QueryValues{},
	}
}

// Set event fields for for 'GetWebhooks' method.
func (p GetWebhooksQueryParams) SetWebhookFields(values []string) {
	fields := queryFields{}
	fields.setWebhookFields(values)

	p.setValues(fields)
}

// Set sort for for 'GetWebhooks' method.
func (p GetWebhooksQueryParams) Include(values []string) {
	p.include(values)
}

// Get webhooks. Reference: https://developers.klaviyo.com/en/reference/get_webhook_subscriptions
func (service *WebhooksService) GetWebhooks(opts *GetWebhooksQueryParams) (*GetWebhooksResponse, *Response, error) {
	_url := fmt.Sprintf("%s/webhooks", ApiTypePrivate)

	req, _ := service.client.NewRequest("GET", _url, opts, nil)

	webhooks := new(GetWebhooksResponse)
	response, err := service.client.Do(req, webhooks)

	if err != nil {
		return nil, response, err
	}

	return webhooks, response, nil
}
