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
func makeNounEndpoint(svc NounService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(nounRequest)
		v, err := svc.Noun(req.Noun.Domain)
		if err != nil {
			return nounResponse{v, err.Error()}, nil
		}
		return nounResponse{v, ""}, nil
	}
}

func makePlaceEndpoint(svc NounService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(placeRequest)
		v, err := svc.Place(req.Domain, req.Category)
		if err != nil {
			return placeResponse{v, err.Error()}, nil
		}
		return placeResponse{v, ""}, nil
	}
}

// Decode Requests
// Extracts a user-domain request object from an HTTP request object.
// --------------------------------------------------
func decodeNounRequest(r *http.Request) (interface{}, error) {
	var request nounRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodePlaceRequest(r *http.Request) (interface{}, error) {
	var request placeRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

// Decode Responses
// Extracts a user-domain response object from an HTTP response object.
// --------------------------------------------------
func decodeNounResponse(r *http.Response) (interface{}, error) {
	var response nounResponse
	if err := json.NewDecoder(r.Body).Decode(&response); err != nil {
		return nil, err
	}
	return response, nil
}

func decodePlaceResponse(r *http.Response) (interface{}, error) {
	var response placeResponse
	if err := json.NewDecoder(r.Body).Decode(&response); err != nil {
		return nil, err
	}
	return response, nil
}

// Encode Responses
// Encodes the passed response object to the HTTP response writer.
// --------------------------------------------------
func encodeResponse(w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

// Encode Requests
// Encodes the passed request object into the HTTP request object.
// --------------------------------------------------
func encodeRequest(r *http.Request, request interface{}) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(request); err != nil {
		return err
	}
	r.Body = ioutil.NopCloser(&buf)
	return nil
}

// Define a struct for the request payloads
// --------------------------------------------------
type nounRequest struct {
	Noun struct {
		Category        string    `json:"category"`
		Coordinates     []float64 `json:"coordinates"`
		CountryCode     string    `json:"country-code"`
		Domain          string    `json:"domain"`
		ExtendedAddress string    `json:"extended-address"`
		Image           string    `json:"image"`
		Link            string    `json:"link"`
		Locality        string    `json:"locality"`
		Name            string    `json:"name"`
		PhoneNumber     string    `json:"phone-number"`
		PostalCode      string    `json:"postal-code"`
		Region          string    `json:"region"`
		StreetAddress   string    `json:"street-address"`
		Tags            []string  `json:"tags"`
	} `json:"noun"`
}

type nounResponse struct {
	V   string `json:"v"`
	Err string `json:"err,omitempty"`
}

// --------------------------------------------------

type placeRequest struct {
	Domain   string `json:"domain"`
	Category string `json:"category"`
}

type placeResponse struct {
	V   string `json:"v"`
	Err string `json:"err,omitempty"`
}
