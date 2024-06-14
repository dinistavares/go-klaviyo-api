package klaviyo

import (
	"fmt"
	"time"
)

// Webhooks service
type WebhooksService struct {
	service
}

type GetWebhooksResponse struct {
	Data     *[]Webhook         `json:"data,omitempty"`
	Links    *GenericLinks      `json:"links,omitempty"`
	Included *[]GenericIncluded `json:"included,omitempty"`
}

type GetWebhookByIDResponse struct {
	Data     *Webhook           `json:"data,omitempty"`
	Included *[]GenericIncluded `json:"included,omitempty"`
}

type CreateWebhookResponse struct {
	Data *Webhook `json:"data,omitempty"`
}

type CreateWebhookCard struct {
	Data *CreateWebhook `json:"data,omitempty"`
}

type CreateWebhook struct {
	ID            string                     `json:"id,omitempty"`
	Type          string                     `json:"type,omitempty"`
	Attributes    *CreateWebhookAttributes   `json:"attributes,omitempty"`
	Relationships *RelationshipWebhookTopics `json:"relationships,omitempty"`
}

type CreateWebhookAttributes struct {
	EndpointURL string `json:"endpoint_url,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	SecretKey   string `json:"secret_key,omitempty"`
}

type WebhookTopics struct {
	Data *[]WebhookTopic `json:"data,omitempty"`
}

type WebhookTopic struct {
	ID   string `json:"id,omitempty"`
	Type string `json:"type,omitempty"`
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
	Enabled     bool       `json:"enabled"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
}

type WebhookRequest struct {
	Meta *WebhookRequestMeta   `json:"meta,omitempty"`
	Data *[]WebhookRequestData `json:"data,omitempty"`
}

type WebhookRequestMeta struct {
	Timestamp        string `json:"timestamp,omitempty"`
	KlaviyoWebhookID string `json:"klaviyo_webhook_id,omitempty"`
	KlaviyoAccountID string `json:"klaviyo_account_id,omitempty"`
	Version          string `json:"version,omitempty"`
}

type WebhookRequestData struct {
	Topic      string                 `json:"topic,omitempty"`
	ExternalID string                 `json:"external_id,omitempty"`
	Payload    *WebhookRequestPayload `json:"payload,omitempty"`
}

type WebhookRequestPayload struct {
	Data *Event `json:"data,omitempty"`
}

type WebhookQueries struct{}

// Query parameters for 'GetWebhooks' method.
type GetWebhooksQueryParams struct {
	QueryValues
}

// Query parameters for 'GetWebhookByID' method.
type GetWebhookByIDQueryParams struct {
	QueryValues
}

// Create Query parameters for webhook routes.
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

// Set webhook fields for for 'GetWebhooks' method.
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

	req, _ := service.newRequest("GET", _url, opts, nil)

	webhooks := new(GetWebhooksResponse)
	response, err := service.client.Do(req, webhooks)

	if err != nil {
		return nil, response, err
	}

	return webhooks, response, nil
}

//  ***********************************************************************************
//  GET WEBHOOK
//  https://developers.klaviyo.com/en/reference/get_webhook_subscription
//  ***********************************************************************************

// Set webhook fields for for 'GetWebhookByID' method.
func (p GetWebhookByIDQueryParams) SetWebhookFields(values []string) {
	fields := queryFields{}
	fields.setWebhookFields(values)

	p.setValues(fields)
}

// Set sort for for 'GetWebhookByID' method.
func (p GetWebhookByIDQueryParams) Include(values []string) {
	p.include(values)
}

// Get webhook. Reference: https://developers.klaviyo.com/en/reference/get_webhook_subscription
func (service *WebhooksService) GetWebhookByID(id string, opts *GetWebhookByIDQueryParams) (*GetWebhookByIDResponse, *Response, error) {
	_url := fmt.Sprintf("%s/webhooks/%s", ApiTypePrivate, id)

	req, _ := service.newRequest("GET", _url, opts, nil)

	webhooks := new(GetWebhookByIDResponse)
	response, err := service.client.Do(req, webhooks)

	if err != nil {
		return nil, response, err
	}

	return webhooks, response, nil
}

//  ***********************************************************************************
//  CREATE WEBHOOK
//  https://developers.klaviyo.com/en/reference/create_webhook_subscription
//  ***********************************************************************************

// Sets new webhook attributes
func (webhook *CreateWebhookCard) SetWebhookAttributes(attributes *CreateWebhookAttributes) {
	webhook.setCreateWebhookData()

	webhook.Data.Attributes = attributes
}

// Sets new webhook topics
func (webhook *CreateWebhookCard) SetWebhookTopics(topicIDs *[]string) {
	webhook.setCreateWebhookRelationships()

	topics := []WebhookTopic{}

	if topicIDs != nil {
		for _, id := range *topicIDs {
			topic := WebhookTopic{
				Type: "webhook-topic",
				ID:   id,
			}

			topics = append(topics, topic)
		}
	}

	webhook.Data.Relationships.WebhookTopics.Data = &topics
}

// Create webhook. Reference: https://developers.klaviyo.com/en/reference/create_webhook_subscription
func (service *WebhooksService) CreateWebhook(webhook *CreateWebhookCard) (*CreateWebhookResponse, *Response, error) {
	_url := fmt.Sprintf("%s/webhooks", ApiTypePrivate)

	// Ensure type is set to "webhook" if empty
	service.setCreateType(webhook)

	req, _ := service.newRequest("POST", _url, nil, webhook)

	newWebhook := new(CreateWebhookResponse)
	response, err := service.client.Do(req, newWebhook)

	if err != nil {
		return newWebhook, response, err
	}

	return newWebhook, response, nil
}

// Sets CreateEvent.Type to 'event' if it is not set
func (service *WebhooksService) setCreateType(webhook *CreateWebhookCard) {
	if webhook != nil && webhook.Data != nil && webhook.Data.Type == "" {
		webhook.Data.Type = "webhook"
	}
}

// Ensure webhook data pointers are created
func (webhook *CreateWebhookCard) setCreateWebhookData() {
	if webhook.Data == nil {
		webhook.Data = &CreateWebhook{}
	}
}

// Ensure webhook data pointers are created
func (webhook *CreateWebhookCard) setCreateWebhookRelationships() {
	webhook.setCreateWebhookData()

	if webhook.Data.Relationships == nil {
		webhook.Data.Relationships = &RelationshipWebhookTopics{
			WebhookTopics: &WebhookTopics{},
		}
	}
}

//  ***********************************************************************************
//  UPDATE WEBHOOK
//  https://developers.klaviyo.com/en/reference/update_webhook_subscription
//  ***********************************************************************************

// Create webhook. Reference: https://developers.klaviyo.com/en/reference/update_webhook_subscription
func (service *WebhooksService) UpdateWebhook(id string, webhook *CreateWebhookCard) (*CreateWebhookResponse, *Response, error) {
	_url := fmt.Sprintf("%s/webhooks/%s", ApiTypePrivate, id)

	// Ensure type is set to "webhook" if empty
	service.setCreateType(webhook)

	req, _ := service.newRequest("PATCH", _url, nil, webhook)

	newWebhook := new(CreateWebhookResponse)
	response, err := service.client.Do(req, newWebhook)

	if err != nil {
		return newWebhook, response, err
	}

	return newWebhook, response, nil
}

//  ***********************************************************************************
//  DELETE WEBHOOK
//  https://developers.klaviyo.com/en/reference/delete_webhook_subscription
//  ***********************************************************************************

// Create webhook. Reference: https://developers.klaviyo.com/en/reference/delete_webhook_subscription
func (service *WebhooksService) DeleteWebhook(id string) (*Response, error) {
	_url := fmt.Sprintf("%s/webhooks/%s", ApiTypePrivate, id)

	req, _ := service.newRequest("DELETE", _url, nil, nil)

	response, err := service.client.Do(req, nil)

	if err != nil {
		return response, err
	}

	return response, nil
}
