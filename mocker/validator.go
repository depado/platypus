package mocker

import (
	"fmt"
	"net/http"
	"strings"
)

// RequestValidator wraps a header validator and a body validator.
type RequestValidator struct {
	Headers HeadersValidator `yaml:"headers"`
	Body    BodyValidator    `yaml:"body"`
}

// Handle will handle the validation of an incoming request, given the request
// and the body in byte form.
func (rv RequestValidator) Handle(r *http.Request, body []byte) []error {
	errors := rv.Headers.Validate(r)
	if err := rv.Body.Validate(body); err != nil {
		errors = append(errors, err)
	}

	return errors
}

// BodyValidator holds the values to validate an incoming request's body.
type BodyValidator struct {
	Contains string `yaml:"contains"`
}

// Validate will validate that the incoming body matches the expectations of
// the body validator.
func (bv BodyValidator) Validate(body []byte) error {
	if bv.Contains != "" {
		if !strings.Contains(string(body), bv.Contains) {
			return fmt.Errorf("body doesn't contain '%s'", bv.Contains)
		}
	}
	return nil
}

// HeadersValidator holds the values to validate the headers of an incoming
// request.
type HeadersValidator struct {
	Present []string          `yaml:"present"`
	Absent  []string          `yaml:"absent"`
	Match   map[string]string `yaml:"match"`
}

// Validate will validate that the incoming HTTP request validates the various
// rules.
func (hv HeadersValidator) Validate(r *http.Request) []error {
	errors := []error{}

	for _, pres := range hv.Present {
		found := false
		for k := range r.Header {
			if k == pres {
				found = true
			}
		}
		if !found {
			errors = append(errors, fmt.Errorf("header not present: %s", pres))
		}
	}

	for _, abs := range hv.Absent {
		for k := range r.Header {
			if k == abs {
				errors = append(errors, fmt.Errorf("header present: %s", abs))
			}
		}
	}

	for reqkey, reqval := range hv.Match {
		found := false
		for hk, hv := range r.Header {
			if reqkey == hk {
				for _, v := range hv {
					if reqval == v {
						found = true
						break
					}
				}
			}
		}
		if !found {
			errors = append(errors, fmt.Errorf("incorrect header value or missing header: %s:%s", reqkey, reqval))
		}
	}

	return errors
}
