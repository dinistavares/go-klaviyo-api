package klaviyo

import "time"

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
