package middlewares

import (
	"time"

	"github.com/ProgrammerPeasant/order-control/utils"
	"github.com/gin-gonic/gin"
)

// PrometheusMiddleware создает middleware для сбора метрик
func PrometheusMiddleware(metrics *utils.Metrics) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		metrics.RequestsInFlight.Inc()

		c.Next()

		metrics.RequestsInFlight.Dec()

		status := c.Writer.Status()

		elapsed := time.Since(start).Seconds()
		metrics.ResponseTime.WithLabelValues(c.Request.Method, c.FullPath()).Observe(elapsed)

		metrics.TotalRequests.WithLabelValues(c.Request.Method, c.FullPath(), string(rune(status))).Inc()

		if status >= 400 {
			errorType := "client_error"
			if status >= 500 {
				errorType = "server_error"
			}
			metrics.RegisterError(errorType, c.Errors.String())
		}
	}
}
