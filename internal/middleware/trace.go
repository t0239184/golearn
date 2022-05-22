package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
)

func Trace() gin.HandlerFunc {
	return func(c *gin.Context) {
		appLog := AppLog
		start := time.Now()
		c.Next()
		end := time.Now()
		latency := end.Sub(start)
		appLog.Info("[Trace] Cost time: ", latency)
	}
}
