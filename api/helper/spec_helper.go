package helper

import (
	"context"
	"net/http"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/getkin/kin-openapi/routers/gorillamux"
)

func SpecTest(t *testing.T, req *http.Request) {
	t.Helper()

	ctx := context.Background()
	loader := &openapi3.Loader{Context: ctx}
	doc, err := loader.LoadFromFile("../../openapi.yaml")
	if err != nil {
		t.Error(err.Error())
	}
	router, err := gorillamux.NewRouter(doc)
	if err != nil {
		t.Error(err.Error())
	}
	route, pathParams, err := router.FindRoute(req)
	if err != nil {
		t.Error(err.Error())
	}

	requestValidationInput := &openapi3filter.RequestValidationInput{
		Request:    req,
		PathParams: pathParams,
		Route:      route,
	}
	if err := openapi3filter.ValidateRequest(ctx, requestValidationInput); err != nil {
		t.Error(err.Error())
	}

	req.URL.Scheme = "http"
	req.URL.Host = "localhost:3000"
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Error(err.Error())
	}
	defer res.Body.Close()

	responseValidationInput := &openapi3filter.ResponseValidationInput{
		RequestValidationInput: requestValidationInput,
		Status:                 res.StatusCode,
		Header:                 res.Header,
		Body:                   res.Body,
	}
	if err := openapi3filter.ValidateResponse(ctx, responseValidationInput); err != nil {
		t.Error(err.Error())
	}
}
