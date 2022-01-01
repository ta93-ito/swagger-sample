package handler

import (
	"net/http"
	"testing"

	"github.com/ta93-ito/golang-swagger-sample/api/helper"
)

func TestHealth(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/health", nil)
	if err != nil {
		t.Error(err.Error())
	}
	helper.SpecTest(t, req)
}
