package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ta93-ito/golang-swagger-sample/api/handler/helper"
)

func TestHealth(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	helper.SpecTest(t, req, Health)
}

// must mock infra
// func TestGetAllTODOs(t *testing.T) {
// 	req := httptest.NewRequest(http.MethodGet, "/todos", nil)
// 	helper.SpecTest(t, req, GETAllTODOs)
// }
