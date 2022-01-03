package helper

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/getkin/kin-openapi/routers/gorillamux"
)

func SpecTest(t *testing.T, req *http.Request, handler func(http.ResponseWriter, *http.Request)) {
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

	res := httptest.NewRecorder()
	handler(res, req)

	resReadCloser := ioutil.NopCloser(res.Body)
	responseValidationInput := &openapi3filter.ResponseValidationInput{
		RequestValidationInput: requestValidationInput,
		Status:                 res.Code,
		Header:                 res.Header(),
		Body:                   resReadCloser,
	}
	if err := openapi3filter.ValidateResponse(ctx, responseValidationInput); err != nil {
		t.Error(err.Error())
	}
}
