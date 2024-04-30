package klaviyo

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

const (
	libraryVersion               = "1.1"
	defaultAuthHeaderName        = "Authorization"
	defaultAuthPrefix            = "Bearer"
	defaultRestEndpointURL       = "https://a.klaviyo.com"
	defaultRestAPIRevision       = "2024-02-15"
	acceptedContentType          = "application/json"
	userAgent                    = "go-klaviyo-api/" + libraryVersion
	clientRequestRetryAttempts   = 2
	clientRequestRetryHoldMillis = 1000
)

var (
	errorDoAllAttemptsExhausted = errors.New("all request attempts were exhausted")
	errorDoAttemptNilRequest    = errors.New("request could not be constructed")

	// API type 'api' for use with private access tokens
	ApiTypePrivate ApiType = "api"

	// API type 'api' for use with public access tokens
	ApiTypePublic ApiType = "client"
)

type ApiType string

type ClientConfig struct {
	HttpClient           *http.Client
	RestEndpointURL      string
	RestEndpointRevision string
}

type auth struct {
	Available   bool
	AccessToken string
	HeaderName  string
	Prefix      string
}

type Client struct {
	config  *ClientConfig
	client  *http.Client
	auth    *auth
	baseURL *url.URL

	Profiles *ProfilesService
}

type service struct {
	client *Client
}

type GenericResponse struct {
	Response *http.Response

	Errors *[]Error `json:"errors,omitempty"`
}

type ErrorSource struct {
	Pointer string `json:"pointer,omitempty"`
}

type Response struct {
	*http.Response
}

type Error struct {
	ID     string       `json:"id,omitempty"`
	Status int          `json:"status,omitempty"`
	Code   string       `json:"code,omitempty"`
	Title  string       `json:"title,omitempty"`
	Detail string       `json:"detail,omitempty"`
	Source *ErrorSource `json:"source,omitempty"`
}

func (response *GenericResponse) Error() string {
	errorString := fmt.Sprintf("%v %v: %d",
	response.Response.Request.Method, response.Response.Request.URL,
	response.Response.StatusCode)

	if response.Errors != nil && len(*response.Errors) > 0 {
		firstError := (*response.Errors)[0]

		errorString = fmt.Sprintf("%s Code: %s Title: %s Detail: %s (ID: %s)",
		 errorString, firstError.Code, firstError.Title, firstError.Detail, firstError.ID)
	}

	return errorString
}

func NewWithConfig(config ClientConfig) *Client {
	if config.HttpClient == nil {
		config.HttpClient = http.DefaultClient
	}

	if config.RestEndpointURL == "" {
		config.RestEndpointURL = defaultRestEndpointURL
	}

	if config.RestEndpointRevision == "" {
		config.RestEndpointRevision = defaultRestAPIRevision
	}

	// Create client
	baseURL, _ := url.Parse(config.RestEndpointURL)

	client := &Client{config: &config, client: config.HttpClient, auth: &auth{}, baseURL: baseURL}

	// Map services
	client.Profiles = &ProfilesService{client: client}

	return client
}

func New() *Client {
	return NewWithConfig(ClientConfig{})
}

func (client *Client) Authenticate(accessToken string) {
	client.auth.HeaderName = defaultAuthHeaderName
	client.auth.Prefix = defaultAuthPrefix
	client.auth.AccessToken = accessToken
	client.auth.Available = true
}

// NewRequest creates an API request
func (client *Client) NewRequest(method, urlStr string, opts *QueryValues, body interface{}) (*http.Request, error) {
	// Append Query Params to URL
	if opts != nil {
		urlStr += opts.encode()
	}

	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	url := client.baseURL.ResolveReference(rel)

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)

		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, url.String(), buf)
	if err != nil {
		return nil, err
	}

	if client.auth.Available {
		req.Header.Add(client.auth.HeaderName, fmt.Sprintf("%s %s", client.auth.Prefix, client.auth.AccessToken))
	}

	req.Header.Add("Accept", acceptedContentType)
	req.Header.Add("Content-type", acceptedContentType)
	req.Header.Add("revision", client.config.RestEndpointRevision)
	req.Header.Add("User-Agent", userAgent)

	return req, nil
}

// Do sends an API request
func (client *Client) Do(req *http.Request, v interface{}) (*Response, error) {
	var lastErr error

	attempts := 0

	for attempts < clientRequestRetryAttempts {
		// Hold before this attempt? (ie. not first attempt)
		if attempts > 0 {
			time.Sleep(clientRequestRetryHoldMillis * time.Millisecond)
		}

		// Dispatch request attempt
		attempts++
		resp, shouldRetry, err := client.doAttempt(req, v)

		// Return response straight away? (we are done)
		if !shouldRetry {
			return resp, err
		}

		// Should retry: store last error (we are not done)
		lastErr = err
	}

	// Set default error? (all attempts failed, but no error is set)
	if lastErr == nil {
		lastErr = errorDoAllAttemptsExhausted
	}

	// All attempts failed, return last attempt error
	return nil, lastErr
}

func (client *Client) doAttempt(req *http.Request, v interface{}) (*Response, bool, error) {
	if req == nil {
		return nil, false, errorDoAttemptNilRequest
	}

	resp, err := client.client.Do(req)

	if checkRequestRetry(resp, err) {
		return nil, true, err
	}

	defer resp.Body.Close()
	
	response := newResponse(resp)

	err = checkResponse(resp)
	if err != nil {
		return response, false, err
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			_, _ = io.Copy(w, resp.Body)
		} else {
			err = json.NewDecoder(resp.Body).Decode(v)
			if err == io.EOF {
				err = nil
			}
		}
	}

	return response, false, err
}

func newResponse(httpResponse *http.Response) *Response {
  response := Response{Response: httpResponse}

  return &response
}

// checkRequestRetry checks if should retry request
func checkRequestRetry(response *http.Response, err error) bool {
	// Low-level error, or response status is a server error? (HTTP 5xx)
	if err != nil || response.StatusCode >= 500 {
		return true
	}

	// No low-level error (should not retry)
	return false
}

// checkResponse checks response for errors
func checkResponse(response *http.Response) error {
	// No error in response? (HTTP 2xx)
	if code := response.StatusCode; 200 <= code && code <= 299 {
		return nil
	}

	// Map response error data (eg. HTTP 4xx)
	errorResponse := &GenericResponse{Response: response}

	data, err := io.ReadAll(response.Body)
	if err == nil && data != nil {
		_ = json.Unmarshal(data, errorResponse)
	}

	return errorResponse
}
