package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	t.Run("ResponseCheck", func(t *testing.T) {
		handler := NewHandler(helloHandler)
		server := httptest.NewServer(handler)
		defer server.Close()

		resp, err := http.Get(server.URL)
		if err != nil {
			t.Fatalf("failed to make GET request: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			t.Errorf("expected status OK; got %v", resp.Status)
		}

		expected := "Hello, go"
		actual := make([]byte, len(expected))
		if _, err := resp.Body.Read(actual); err != nil {
			t.Fatalf("failed to read response body: %v", err)
		}
		if string(actual) != expected {
			t.Errorf("expected %q; got %q", expected, string(actual))
		}
	})
}
