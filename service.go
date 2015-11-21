package main

import "errors"

// NounService provides operations on strings.
type NounService interface {
	Noun(nounResponse) (string, error)
	Place(string, string) (string, error)
}

type nounService struct{}

// Noun
// --------------------------------------------------
func (nounService) Noun(req nounRequest) (string, error) {
	if req.Noun.Domain == "" {
		return "", ErrEmpty
	}
	return req.Noun.Domain, nil
}

// Place
// --------------------------------------------------
func (nounService) Place(domain string, category string) (string, error) {
	if domain == "" {
		return "", ErrEmpty
	}
	return domain + " " + category, nil
}

// ErrEmpty is returned when an input string is empty.
var ErrEmpty = errors.New("empty string")

// ServiceMiddleware is a chainable behavior modifier for NounService.
type ServiceMiddleware func(NounService) NounService
