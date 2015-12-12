package main

import (
	"errors"
	"fmt"
)

// NounService provides operations on strings.
type NounService interface {
	Noun(nounRequest) (string, error)
	Place(string, string) (string, error)
}

type nounService struct{}

// Noun
// --------------------------------------------------
func (nounService) Noun(req nounRequest) (string, error) {

	fmt.Println("Entering Noun Handler") // DEBUGGING
	fmt.Println(req.Noun)                // DEBUGGING

	if req.Noun.Domain == "" {
		return "req.Noun.Domain was empty when Noun was called.", ErrEmpty
	}

	return req.Noun.Domain, nil
}

// Place
// --------------------------------------------------
func (nounService) Place(domain string, category string) (string, error) {
	if domain == "" {
		return "domain was empty when Place was called", ErrEmpty
	}
	return domain + " " + category, nil
}

// ErrEmpty is returned when an input string is empty.
var ErrEmpty = errors.New("empty string")

// ServiceMiddleware is a chainable behavior modifier for NounService.
type ServiceMiddleware func(NounService) NounService
