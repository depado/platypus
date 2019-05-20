package infra

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/sirupsen/logrus"
)

// NewCorsConfig generates a new cors config with the given parameters
func NewCorsConfig(enable, all bool, origins, methods, headers, expose []string) *cors.Config {
	if !enable {
		return nil
	}
	c := &cors.Config{
		AllowCredentials: true,
		MaxAge:           50 * time.Second,
		AllowMethods:     methods,
		AllowHeaders:     headers,
		ExposeHeaders:    expose,
	}

	switch {
	case len(origins) > 0:
		c.AllowOrigins = origins
	case all:
		c.AllowAllOrigins = true
	default:
		logrus.WithField("error", "allow all origin disabled but no allowed origin provided").Fatal("Couldn't configure CORS")
	}
	logrus.WithFields(logrus.Fields{
		"AllowMethods":    methods,
		"AllowHeaders":    headers,
		"ExposeHeaders":   expose,
		"AllowOrigins":    origins,
		"AllowAllOrigins": all,
	}).Debug("CORS configuration")
	return c
}
