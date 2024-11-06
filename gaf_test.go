package gaf

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func NewTestServer(responseHandler http.HandlerFunc) *httptest.Server {
	return httptest.NewServer(responseHandler)
}

func NewTestClient(t *testing.T) *Client {
	responseHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	server := NewTestServer(responseHandler)

	defer server.Close()

	client := NewClient("testApiKey")

	client.BaseURL = server.URL

	if client.apiKey != "testApiKey" {
		t.Errorf("Expected api key to be testApiKey, got %s", client.apiKey)
	}

	return client
}

func NewTestClientWithCustomHandler(t *testing.T, server *httptest.Server) *Client {
	client := NewClient("testApiKey")

	client.BaseURL = server.URL

	if client.apiKey != "testApiKey" {
		t.Errorf("Expected api key to be testApiKey, got %s", client.apiKey)
	}

	return client
}

func TestNewClient(t *testing.T) {
	NewTestClient(t)
}
