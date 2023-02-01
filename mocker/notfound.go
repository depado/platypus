package mocker

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/logrusorgru/aurora"
)

// NoRouteConf defines the behavior when no route is found in the mock. Defaults
// to a standard 404 response.
type NoRouteConf struct {
	Dump Dump    `yaml:"dump"`
	Echo bool    `yaml:"echo"`
	Code *int    `yaml:"code"`
	Body *string `yaml:"body"`
}

// Handle will add a special route to customize the 404 behavior.
func (nc NoRouteConf) Handle(r *gin.Engine) {

	code := http.StatusNotFound
	if nc.Code != nil {
		code = *nc.Code
	}

	fmt.Printf("\n%s %s", aurora.Underline("No route"), codeToColor(code).String())

	if nc.Echo {
		fmt.Print(" 🔊")
	}
	fmt.Print("\n")

	r.NoRoute(func(c *gin.Context) {
		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return
		}

		nc.Dump.Handle(c.Request, body)

		if nc.Echo {
			for header, values := range c.Request.Header {
				for _, v := range values {
					c.Header(header, v)
				}
			}
			status := http.StatusOK
			if nc.Code != nil {
				status = *nc.Code
			}
			if nc.Body != nil {
				c.String(status, string(body))
			} else {
				c.Status(status)
			}
			return
		}

		if nc.Body == nil {
			c.Status(code)
			return
		}

		c.String(code, *nc.Body)
	})
}
