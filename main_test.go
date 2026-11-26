package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestUserWildcard(t *testing.T) {
	rec := httptest.NewRecorder()
	newMux().ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/users/42", nil))
	if rec.Code != http.StatusOK || !strings.Contains(rec.Body.String(), "`42`") && !strings.Contains(rec.Body.String(), "42") {
		t.Fatalf("unexpected response: %d %s", rec.Code, rec.Body.String())
	}
}

func TestTrailingWildcard(t *testing.T) {
	rec := httptest.NewRecorder()
	newMux().ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/files/a/b.txt", nil))
	if rec.Code != http.StatusOK || !strings.Contains(rec.Body.String(), "a/b.txt") {
		t.Fatalf("unexpected response: %d %s", rec.Code, rec.Body.String())
	}
}

func TestMethodMismatch(t *testing.T) {
	rec := httptest.NewRecorder()
	newMux().ServeHTTP(rec, httptest.NewRequest(http.MethodPost, "/users/42", nil))
	if rec.Code != http.StatusMethodNotAllowed {
		t.Fatalf("expected 405, got %d", rec.Code)
	}
}
