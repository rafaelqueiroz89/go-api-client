package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
)

var baseUrl = "http://localhost:8080"

const (
	httpClientTimeout = 15 * time.Second
)

// Client The default HTTP Client
type Client struct {
	HTTPClient *http.Client
	BaseUrl    string
}

// NewClient Client with the BASE_URL set, if running tests it will look for the FORM3 API in the Docker container
//Otherwise it will look for localhost
func (c Client) NewClient() *Client {
	if os.Getenv("BASE_URL") != "" {
		baseUrl = os.Getenv("BASE_URL")
	}
	if c.BaseUrl != "" {
		baseUrl = c.BaseUrl
	}

	return &Client{
		HTTPClient: &http.Client{
			Timeout: httpClientTimeout,
		},
	}
}

// Request HTTP generic request and response
func (c *Client) Request(v interface{}, httpMethod, path string, query map[string]string, data interface{}) (*http.Response, error) {
	path = fmt.Sprintf("%s/%s", baseUrl, path)

	uri, err := url.Parse(path)
	LogError(err)

	body, err := PrepareBody(data)
	LogError(err)

	request, err := http.NewRequest(httpMethod, uri.String(), bytes.NewBuffer(body))
	LogError(err)

	request.Header.Set("Accept", "*/*")
	request.Header.Set("Content-Type", "vnd.api+json")

	AddQueryParams(query, request)

	resp, err := c.HTTPClient.Do(request)
	LogError(err)

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		LogError(err)
	}(resp.Body)
	LogError(err)

	response, err, done := GetResponseData(resp, err)
	if done {
		return response, err
	}

	if v != nil {
		if err = json.NewDecoder(resp.Body).Decode(&v); err != nil {
			return resp, err
		}
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	LogError(err)

	if len(responseBody) > 0 {
		if err := json.Unmarshal(responseBody, &v); err != nil {
			return nil, fmt.Errorf("could not decode response json, %s: ", string(responseBody))
		}
	}

	return resp, err
}

// GetResponseData Get response data and handle http status codes
func GetResponseData(resp *http.Response, err error) (*http.Response, error, bool) {
	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusBadRequest {
		data := Error{}
		if err := json.NewDecoder(resp.Body).Decode(&data); err == nil {
			return resp, errors.New(data.ErrorMessage), true
		}

		return resp, err, true
	}
	return nil, nil, false
}

// AddQueryParams Add params to the query string
func AddQueryParams(query map[string]string, request *http.Request) {
	if len(query) > 0 {
		q := request.URL.Query()
		for k, v := range query {
			q.Add(k, v)
		}
		request.URL.RawQuery = q.Encode()
	}
}

// PrepareBody Decodes the body to display
func PrepareBody(data interface{}) ([]byte, error) {
	bytesJ, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	return bytesJ, nil
}

// LogError Log errors
func LogError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

// Error Struct to unmarshal the HTTP message for error messages
type Error struct {
	ErrorMessage string `json:"error_message"`
}
