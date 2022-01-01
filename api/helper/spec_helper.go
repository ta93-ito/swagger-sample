package helper

import (
	"context"
	"net/http"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/getkin/kin-openapi/routers/gorillamux"
)

func SpecTest(req *http.Request) error {
	ctx := context.Background()
	loader := &openapi3.Loader{Context: ctx}
	doc, err := loader.LoadFromFile("/Users/takumiito/swagger-sample/openapi.yaml")
	if err != nil {
		return err
	}
	router, err := gorillamux.NewRouter(doc)
	if err != nil {
		return err
	}
	route, pathParams, err := router.FindRoute(req)
	if err != nil {
		return err
	}

	requestValidationInput := &openapi3filter.RequestValidationInput{
		Request:    req,
		PathParams: pathParams,
		Route:      route,
	}
	if err := openapi3filter.ValidateRequest(ctx, requestValidationInput); err != nil {
		return err
	}

	req.URL.Scheme = "http"
	req.URL.Host = "localhost:3000"
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	responseValidationInput := &openapi3filter.ResponseValidationInput{
		RequestValidationInput: requestValidationInput,
		Status:                 res.StatusCode,
		Header:                 res.Header,
		Body:                   res.Body,
	}
	if err := openapi3filter.ValidateResponse(ctx, responseValidationInput); err != nil {
		return err
	}
	return nil
}
