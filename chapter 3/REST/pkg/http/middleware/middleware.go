package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type Config func(hcc *healthCheckConfig)

type healthCheckConfig struct {
	service      string
	db           *sqlx.DB
	withDatabase bool
}

func defaultConfig(service string) *healthCheckConfig {
	var config = healthCheckConfig{
		service:      service,
		withDatabase: false,
	}

	return &config
}

func HealthCheck(svc string, r *gin.Engine, ops ...Config) {
	config := defaultConfig(svc)

	for _, v := range ops {
		v(config)
	}

	r.GET(fmt.Sprintf("%s/health-check", svc), config.handler)
}

func (hcc *healthCheckConfig) handler(c *gin.Context) {
	if hcc.withDatabase {
		err := hcc.db.Ping()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"status":  "postgres ping failed",
				"message": err.Error(),
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "service online",
	})
}
