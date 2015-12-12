package main

import (
	"bytes"
	"encoding/json"
	"fmt"
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
		v, err := svc.Noun(req)

		fmt.Println("Making Noun Endpoint") // DEBUGGING
		fmt.Println(req.Noun)               // DEBUGGING

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

		fmt.Println("Making Place Endpoint") // DEBUGGING
		fmt.Println(req)                     // DEBUGGING

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

	fmt.Println("Decoding Noun Request") // DEBUGGING
	fmt.Println(request)                 // DEBUGGING

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
		Category        string    `json:"category" xml:"category"`
		Coordinates     []float64 `json:"coordinates" xml:"coordinates"`
		CountryCode     string    `json:"country-code" xml:"country-code"`
		Domain          string    `json:"domain" xml:"domain"`
		ExtendedAddress string    `json:"extended-address" xml:"extended-address"`
		Image           string    `json:"image" xml:"image"`
		Link            string    `json:"link" xml:"link"`
		Locality        string    `json:"locality" xml:"locality"`
		Name            string    `json:"name" xml:"name"`
		PhoneNumber     string    `json:"phone-number" xml:"phone-number"`
		PostalCode      string    `json:"postal-code" xml:"postal-code"`
		Region          string    `json:"region" xml:"region"`
		StreetAddress   string    `json:"street-address" xml:"street-address"`
		Tags            []string  `json:"tags" xml:"tags"`
	} `json:"noun" xml:"noun"`
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
