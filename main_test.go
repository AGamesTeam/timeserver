package main

import (
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
)

var timePattern = regexp.MustCompile(`^\d{2}:\d{2}:\d{2}$`)

func TestTimeHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	timeHandler(w, req)

	resp := w.Result()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200, got %d", resp.StatusCode)
	}

	contentType := resp.Header.Get("Content-Type")
	if contentType != "text/plain" {
		t.Errorf("expected Content-Type text/plain, got %s", contentType)
	}

	body := w.Body.String()
	if !timePattern.MatchString(body) {
		t.Errorf("expected HH:MM:SS format, got %q", body)
	}
}
