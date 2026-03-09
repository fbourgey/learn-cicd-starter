package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey_Success(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "ApiKey abc123")

	got, err := GetAPIKey(headers)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if got != "abc123" {
		t.Fatalf("expected api key %q, got %q", "abc123", got)
	}
}

func TestGetAPIKey_MissingAuthorizationHeader(t *testing.T) {
	headers := http.Header{} // no Authorization header

	got, err := GetAPIKey(headers)
	if got != "" {
		t.Fatalf("expected empty api key, got %q", got)
	}
	if !errors.Is(err, ErrNoAuthHeaderIncluded) {
		t.Fatalf("expected ErrNoAuthHeaderIncluded, got %v", err)
	}
}
