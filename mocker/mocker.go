package mocker

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

// MockConf contains all the endpoints and is used for parsing.
type MockConf struct {
	Endpoints []Endpoint   `yaml:"endpoints"`
	NoRoute   *NoRouteConf `yaml:"noroute"`
}

// GenerateRoutes generates the routes and applies them to the r gin.Engine.
func GenerateRoutes(path string, r *gin.Engine) error {
	var err error
	var out []byte

	mc := &MockConf{}
	if out, err = os.ReadFile(path); err != nil {
		return errors.Wrap(err, "open file")
	}
	logrus.Infof("Found mock file %s", path)

	if err = yaml.Unmarshal(out, mc); err != nil {
		return errors.Wrap(err, "unmarshal")
	}
	logrus.Info("Mock file is valid")

	for _, e := range mc.Endpoints {
		e.Compute()
		for _, g := range e.All {
			g.Generate(e.Path, r)
		}
	}

	if mc.NoRoute != nil {
		mc.NoRoute.Handle(r)
	}

	return nil
}
