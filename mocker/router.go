package mocker

import (
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
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
	logrus.Infof("Found mock file %s", path)

	if err = yaml.Unmarshal(out, all); err != nil {
		return errors.Wrap(err, "unmarshal")
	}
	logrus.Info("Mock file is valid")

	for _, e := range all.Endpoints {
		e.Compute()
		for _, g := range e.All {
			g.Generate(e.Path, r)
		}
	}
	return nil
}
