package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"golang.org/x/net/context"

	"github.com/go-kit/kit/endpoint"
)

// Make Endpoints
// --------------------------------------------------
func makePlaceEndpoint(svc NounService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(placeRequest)
		v, err := svc.Place(req.S)
		if err != nil {
			return placeResponse{v, err.Error()}, nil
		}
		return placeResponse{v, ""}, nil
	}
}

func makeUppercaseEndpoint(svc NounService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(uppercaseRequest)
		v, err := svc.Uppercase(req.S)
		if err != nil {
			return uppercaseResponse{v, err.Error()}, nil
		}
		return uppercaseResponse{v, ""}, nil
	}
}

// Decode Requests
// --------------------------------------------------
func decodePlaceRequest(r *http.Request) (interface{}, error) {
	var request placeRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeUppercaseRequest(r *http.Request) (interface{}, error) {
	var request uppercaseRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

// Decode Responses
// --------------------------------------------------
func decodePlaceResponse(r *http.Response) (interface{}, error) {
	var response placeResponse
	if err := json.NewDecoder(r.Body).Decode(&response); err != nil {
		return nil, err
	}
	return response, nil
}

func decodeUppercaseResponse(r *http.Response) (interface{}, error) {
	var response uppercaseResponse
	if err := json.NewDecoder(r.Body).Decode(&response); err != nil {
		return nil, err
	}
	return response, nil
}

// Encode Responses
// --------------------------------------------------
func encodeResponse(w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

// Encode Requests
// --------------------------------------------------
func encodeRequest(r *http.Request, request interface{}) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(request); err != nil {
		return err
	}
	r.Body = ioutil.NopCloser(&buf)
	return nil
}

// Grab the item out of the payload
// --------------------------------------------------
type placeRequest struct {
	S string `json:"noun"`
}

type placeResponse struct {
	V   string `json:"v"`
	Err string `json:"err,omitempty"`
}

type uppercaseRequest struct {
	S string `json:"s"`
}

type uppercaseResponse struct {
	V   string `json:"v"`
	Err string `json:"err,omitempty"`
}
