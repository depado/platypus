package mocker

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/logrusorgru/aurora"
)

// Responses represents multiple responses.
type Responses []Response

// Pick will pick the appropriate response.
func (rr Responses) Pick(platy string) Response {
	// No desired code/name, return the first response
	if platy == "" {
		return rr[0]
	}

	// Find by name first
	for _, r := range rr {
		if r.Name == platy {
			return r
		}
	}

	// Find by status code next
	o, err := strconv.Atoi(platy)
	if err == nil {
		for _, r := range rr {
			if r.Code == o {
				return r
			}
		}
	}

	return rr[0]
}

// Response represents the response structure for a single endpoint method.
type Response struct {
	Name    string            `yaml:"name"`
	Code    int               `yaml:"code"`
	Echo    bool              `yaml:"echo"`
	Headers map[string]string `yaml:"headers"`
	Preset  string            `yaml:"preset"`
	Body    string            `yaml:"body"`
}

// Info returns a string to print out the information of a response.
func (r Response) Info(prefix string, last bool) string {
	var sb strings.Builder
	s := "├─"
	if last {
		s = "└─"
	}

	sb.WriteString(fmt.Sprintf("%s %s %s", prefix, s, codeToColor(r.Code).String()))
	if r.Name != "" {
		sb.WriteString(aurora.Cyan(" " + r.Name).String())
	}
	if r.Preset != "" {
		switch r.Preset {
		case "json":
			sb.WriteString(" JSON")
		case "text":
			sb.WriteString(" Text")
		default:
			sb.WriteString(" " + r.Preset)
		}
	}
	if r.Echo {
		sb.WriteString(" 🔊")
	}

	return sb.String()
}

func codeToColor(code int) aurora.Value {
	switch {
	case code >= 100 && code < 200:
		return aurora.Cyan(code)
	case code >= 200 && code < 300:
		return aurora.Green(code)
	case code >= 300 && code < 400:
		return aurora.Yellow(code)
	case code >= 400 && code < 500:
		return aurora.Red(code)
	case code >= 500 && code < 600:
		return aurora.BrightRed(code)
	}
	return aurora.White(code)
}
