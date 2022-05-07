package mocker

import (
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// EndpointMethod represents a single method associated to a parent endpoint.
type EndpointMethod struct {
	Echo      bool             `yaml:"echo"`
	Dump      Dump             `yaml:"dump"`
	Validate  RequestValidator `yaml:"validate"`
	Responses Responses        `yaml:"responses"`
}

// Info returns the string representing the information.
func (e EndpointMethod) Info(last bool) string {
	var sb strings.Builder
	pref := "\n│ "
	if last {
		pref = "\n  "
	}

	for i, r := range e.Responses {
		sb.WriteString(r.Info(pref, i == len(e.Responses)-1))
	}

	return sb.String()
}

// ToHandler generates a handler to apply on the router.
func (e EndpointMethod) ToHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		// Extract the body
		body, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return
		}

		// Dump if required
		e.Dump.Handle(c.Request, body)

		// Validate headers and body if any validation is required
		if err := e.Validate.Handle(c.Request, body); len(err) != 0 {
			var errs []string
			for _, ie := range err {
				errs = append(errs, ie.Error())
			}
			c.JSON(http.StatusBadRequest, gin.H{"error": "request validation failed", "reasons": errs})
			return
		}

		// If no response is defined, return a 200 without a body
		if e.Responses == nil {
			c.Status(http.StatusOK)
			return
		}

		// Pick the desired response
		desired := c.Query("platy")
		r := e.Responses.Pick(desired)

		// Echo mode
		if r.Echo {
			for header, values := range c.Request.Header {
				for _, v := range values {
					c.Header(header, v)
				}
			}
			status := http.StatusOK
			if r.Code != 0 {
				status = r.Code
			}
			if body != nil {
				c.String(status, string(body))
			} else {
				c.Status(status)
			}
			return
		}

		switch r.Preset {
		case "json":
			c.Header("Content-Type", "application/json; charset=utf-8")
		}

		for k, v := range r.Headers {
			c.Header(k, v)
		}

		if r.Body != "" {
			c.String(r.Code, r.Body)
		} else {
			c.Status(r.Code)
		}
	}
}
