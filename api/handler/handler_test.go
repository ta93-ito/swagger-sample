package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealth(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "localhost:3000/health", nil)
	res := httptest.NewRecorder()

	Health(res, req)
	if res.Code != http.StatusOK {
		t.Errorf("want 200, but got %d", res.Code)
	}
	wantBody := "server is healthy!"
	if str := res.Body.String(); str != wantBody {
		t.Errorf("want is %s, but got %s", wantBody, str)
	}
}
