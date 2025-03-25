package rest

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jgill07/pocket-pilot-common-go/pkg/log"
	"go.uber.org/zap"
)

func Logger(c *gin.Context) {
	t := time.Now()
	c.Next()
	latency := time.Since(t)

	status := c.Writer.Status()
	method := c.Request.Method
	path := c.Request.URL.Path

	log.WithFields(
		zap.Int("status", status),
		zap.String("method", method),
		zap.String("path", path),
		zap.Duration("latency", latency),
	).Debug("request")
}
