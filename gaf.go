package gaf

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"time"
)

const (
	Protocol     = "https"
	BaseDomainV3 = "v3.football.api-sports.io"
)

type Client struct {
	BaseURL    string
	apiKey     string
	HTTPClient *http.Client
}

func NewClient(apiKey string) *Client {
	return &Client{
		BaseURL: fmt.Sprintf("%s://%s", Protocol, BaseDomainV3),
		apiKey:  apiKey,
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}

func (c *Client) sendRequest(req *http.Request, v interface{}) error {
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Accept", "application/json; charset=utf-8")
	req.Header.Set("x-rapidapi-host", BaseDomainV3)
	req.Header.Set("x-rapidapi-key", c.apiKey)

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		return fmt.Errorf("unknown error, status code: %d", res.StatusCode)
	}

	fullResponse := successResponse{
		Response: v,
	}

	if err = json.NewDecoder(res.Body).Decode(&fullResponse); err != nil {
		return err
	}

	if reflect.TypeOf(fullResponse.Errors).Kind().String() == "map" {
		iter := reflect.ValueOf(fullResponse.Errors).MapRange()
		errorMap := make(map[string]string)
		for iter.Next() {
			errorMap[iter.Key().String()] = fmt.Sprintf("%s", iter.Value())
		}

		if len(errorMap) > 0 {
			var errTxt string
			for field, errorMsg := range errorMap {
				errTxt = fmt.Sprintf("%s: %s. ", field, errorMsg)
			}
			return fmt.Errorf(fmt.Sprintf("API error(-s): %s", errTxt))
		}

		return fmt.Errorf("unknown API error")
	}

	return nil
}
