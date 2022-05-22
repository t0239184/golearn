package middleware

import (
	"bytes"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func ReqResAopLogging() gin.HandlerFunc {
	return func(c *gin.Context) {
		logrus.Info("[ReqResAopLogging] request: ", c.Request.URL.Path, c.Request.Body)

		bodyLogWriter := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = bodyLogWriter

		c.Next()

		statusCode := c.Writer.Status()

		applog := AppLog
		applog.Info("[ReqResAopLogging] Response body: ", statusCode, bodyLogWriter.body.String())

	}
}

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}
