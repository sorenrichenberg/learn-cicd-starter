package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	t.Fatal("force fail")
	headers := http.Header{}
	headers.Set("Authorization", "ApiKey authHeader")
	got, err := GetAPIKey(headers)
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	want := "authHeader"
	if want != got {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}
