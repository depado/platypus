package mocker

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/logrusorgru/aurora"
)

// Dump is a type alias
type Dump []string

// Contains will return whether or not a specific string is present in the
// slice or not.
func (d Dump) Contains(s string) bool {
	for _, v := range d {
		if v == s {
			return true
		}
	}
	return false
}

// Handle will handle request dumping if any.
func (d Dump) Handle(r *http.Request, body []byte) {
	fmt.Println("\n--------------------------------------------------")
	if d.Contains("host") {
		fmt.Printf("%s %s\n\n", aurora.Blue("Host:"), r.Host)
	}

	if d.Contains("headers") {
		for k, v := range r.Header {
			fmt.Printf("%s %s\n", aurora.BrightBlue(k+":"), strings.Join(v, ","))
		}
		fmt.Println()
	}

	if d.Contains("body") && len(body) > 0 {
		fmt.Printf("%s\n\n", string(body))
	}
}
