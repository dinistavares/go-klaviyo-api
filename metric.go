package klaviyo

import (
	"fmt"
	"time"
)

// Metrics service
type MetricsService service

type GetMetricsResponse struct {
	Data     *[]Metric           `json:"data,omitempty"`
	Links    *GenericLinks      `json:"links,omitempty"`
}

type MetricAttributes struct {
	Name        string             `json:"name,omitempty"`
	Created     string             `json:"created,omitempty"`
	Updated     *time.Time         `json:"updated,omitempty"`
	Integration *MetricIntegration `json:"integration,omitempty"`
}

type Metric struct {
	Type       string            `json:"type,omitempty"`
	ID         string            `json:"id,omitempty"`
	Attributes *MetricAttributes `json:"attributes,omitempty"`
	Links      *GenericLinks     `json:"links,omitempty"`
}

type MetricIntegration struct {
	Category *interface{} `json:"category,omitempty"`
	ID       string       `json:"id,omitempty"`
	Key      string       `json:"key,omitempty"`
	Name     string       `json:"name,omitempty"`
	Object   string       `json:"object,omitempty"`
}
type MetricQueries struct{}

// Query parameters for 'GetMetrics' method.
type GetMetricsQueryParams struct {
	QueryValues
}

//  ***********************************************************************************
//  GET METRICS
//  https://developers.klaviyo.com/en/reference/get_metrics
//  ***********************************************************************************

// Creates Query parameters for 'GetMetrics'
func (pq MetricQueries) NewGetMetrics() *GetMetricsQueryParams {
	return &GetMetricsQueryParams{
		QueryValues: QueryValues{},
	}
}

// Set metric fields for for 'GetMetrics' method.
func (p GetMetricsQueryParams) SetMetricFields(values []string) {
	fields := queryFields{}
	fields.setMetricFields(values)

	p.setValues(fields)
}

// Set filter for for 'GetMetrics' method.
func (p GetMetricsQueryParams) Filter(filter QueryFilter) {
	p.filter(filter)
}

// Set page cursor for for 'GetMetrics' method.
func (p GetMetricsQueryParams) SetPageCursor(value string) {
	page := queryPage{}
	page.setPageCursor(value)

	p.setValues(page)
}

// Get metrics. Reference: https://developers.klaviyo.com/en/reference/get_metrics
func (service *MetricsService) GetMetrics(opts *GetMetricsQueryParams) (*GetMetricsResponse, *Response, error) {
	_url := fmt.Sprintf("%s/metrics", ApiTypePrivate)

	req, _ := service.client.NewRequest("GET", _url, opts, nil)

	metrics := new(GetMetricsResponse)
	response, err := service.client.Do(req, metrics)

	if err != nil {
		return nil, response, err
	}

	return metrics, response, nil
}

