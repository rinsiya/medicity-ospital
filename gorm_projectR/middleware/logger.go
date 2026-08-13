package middleware

import (
	"log/slog"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func RequestLogger(logger *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// Process request
		c.Next()

		// Calculate duration
		duration := time.Since(start)

		// Log request details
		logger.Info("HTTP request",
			"method", c.Request.Method,
			"path", c.Request.URL.Path,
			"status", c.Writer.Status(),
			"duration", duration.String(),
			"ip", c.ClientIP(),
		)
	}
}


func RequestDataLogger(logger *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		reqID := uuid.New().String()

		c.Set("request_id", reqID)

		logger.Info("request started",
			"request_id", reqID,
			"method", c.Request.Method,
			"path", c.Request.URL.Path,
		)

		c.Next()

		logger.Info("request completed",
			"request_id", reqID,
			"status", c.Writer.Status(),
		)
	}
}