package main

import (
	"io/ioutil"
	"net/http"
	"testing"
	"time"
	"github.com/mradigen/short/internal/api"
)

func TestApplication(t *testing.T) {
	go func() {
		api.Start()
	}()

	time.Sleep(1 * time.Second) // Wait for the server to start

	baseURL := "http://localhost:8080"

	t.Run("Valid URL Shortening", func(t *testing.T) {
		resp, err := http.Get(baseURL + "/shorten?url=https://aadivishnu.com")
		if err != nil {
			t.Fatalf("failed to send GET request: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			t.Fatalf("expected status 200, got %d", resp.StatusCode)
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Fatalf("failed to read response body: %v", err)
		}

		expectedSubstring := `{"short_url":"`
		if string(body)[:14] != expectedSubstring {
			t.Fatalf("expected response to contain %s, got %s", expectedSubstring, string(body))
		}
	})

	t.Run("Missing URL Parameter", func(t *testing.T) {
		resp, err := http.Get(baseURL + "/shorten")
		if err != nil {
			t.Fatalf("failed to send GET request: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusBadRequest {
			t.Fatalf("expected status 400, got %d", resp.StatusCode)
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Fatalf("failed to read response body: %v", err)
		}

		expectedError := "Missing 'url' query parameter\n"
		if string(body) != expectedError {
			t.Fatalf("expected error %s, got %s", expectedError, string(body))
		}
	})

	t.Run("Invalid HTTP Method", func(t *testing.T) {
		client := &http.Client{}
		req, err := http.NewRequest(http.MethodPost, baseURL+"/shorten?url=https://aadivishnu.com", nil)
		if err != nil {
			t.Fatalf("failed to create POST request: %v", err)
		}

		resp, err := client.Do(req)
		if err != nil {
			t.Fatalf("failed to send POST request: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusMethodNotAllowed {
			t.Fatalf("expected status 405, got %d", resp.StatusCode)
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Fatalf("failed to read response body: %v", err)
		}

		expectedError := "Method Not Allowed\n"
		if string(body) != expectedError {
			t.Fatalf("expected error %s, got %s", expectedError, string(body))
		}
	})
}

