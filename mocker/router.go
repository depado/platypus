package mocker

import (
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

// GenerateRoutes generates the routes and applies them to the r gin.Engine
func GenerateRoutes(path string, r *gin.Engine) error {
	var err error
	var out []byte

	all := &All{}
	if out, err = ioutil.ReadFile(path); err != nil {
		return errors.Wrap(err, "open file")
	}

	if err = yaml.Unmarshal(out, all); err != nil {
		return errors.Wrap(err, "unmarshal")
	}

	for _, e := range all.Endpoints {
		e.Compute()
		for _, g := range e.All {
			g.Generate(e.Path, r)
		}
	}
	return nil
}
