package main

import (
	"errors"
	"strings"
)

// NounService provides operations on strings.
type NounService interface {
	Place(string) (string, error)
	Uppercase(string) (string, error)
}

type nounService struct{}

// Place
// --------------------------------------------------
func (nounService) Place(payload string) (string, error) {
	if payload == "" {
		return "", ErrEmpty
	}
	return payload, nil
}

// Uppercase
// --------------------------------------------------
func (nounService) Uppercase(s string) (string, error) {
	if s == "" {
		return "", ErrEmpty
	}
	return strings.ToUpper(s), nil
}

// ErrEmpty is returned when an input string is empty.
var ErrEmpty = errors.New("empty string")

// ServiceMiddleware is a chainable behavior modifier for NounService.
type ServiceMiddleware func(NounService) NounService
