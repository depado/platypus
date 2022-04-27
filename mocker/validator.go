package mocker

import (
	"fmt"
	"net/http"
)

// RequestValidator wraps a header validator and a body validator.
type RequestValidator struct {
	Headers HeadersValidator `yaml:"headers"`
	Body    BodyValidator    `yaml:"body"`
}

// Handle will handle the validation of an incoming request, given the request
// and the body in byte form.
func (rv RequestValidator) Handle(r *http.Request, body []byte) []string {
	errors := []string{}

	for _, pres := range rv.Headers.Present {
		found := false
		for k := range r.Header {
			if k == pres {
				found = true
			}
		}
		if !found {
			errors = append(errors, fmt.Sprintf("header not present: %s", pres))
		}
	}

	for _, abs := range rv.Headers.Absent {
		for k := range r.Header {
			if k == abs {
				errors = append(errors, fmt.Sprintf("header present: %s", abs))
			}
		}
	}

	return errors
}

// BodyValidator holds the values to validate an incoming request's body.
type BodyValidator struct {
	Contains string `yaml:"contains"`
}

// HeadersValidator holds the values to validate the headers of an incoming
// request.
type HeadersValidator struct {
	Present  []string          `yaml:"present"`
	Absent   []string          `yaml:"absent"`
	Required map[string]string `yaml:"required"`
}
