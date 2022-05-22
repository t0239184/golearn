package middleware

import (
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

var AppLog *logrus.Logger
var WebLog *logrus.Logger

func Setup() {
	initAppLog()
	initAccessLog()
}

func initAppLog() {
	logFileName := "app.log"
	AppLog = initFileLog(logFileName)
}

func initAccessLog() {
	logFileName := "access.log"
	WebLog = initFileLog(logFileName)
}

func initFileLog(logFileName string) *logrus.Logger {
	log := logrus.New()
	log.Formatter = &logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	}
	logPath := "logs/"
	logName := logPath + logFileName
	var f *os.File
	var err error
	if _, err := os.Stat(logName); os.IsNotExist(err) {
		err = os.Mkdir(logPath, os.ModePerm)
		if err != nil {
			fmt.Println("mkdir failed!")
		}
		f, err = os.Create(logName)
	} else {
		f, err = os.OpenFile(logName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	}

	if err != nil {
		fmt.Println("open log file failed")
	}

	log.Out = f
	log.Level = logrus.InfoLevel
	
	return log
}

func LoggerToFile() gin.HandlerFunc {

	return func(c *gin.Context) {
		requestId := setupRequestId(c)

		startTime := time.Now()
		c.Next()
		endTime := time.Now()
		latencyTime := fmt.Sprintf("%6v", endTime.Sub(startTime))
		reqMethod := c.Request.Method
		reqUri := c.Request.RequestURI
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		AppLog.WithFields(logrus.Fields{
			"requestId":   requestId,
			"http_status": statusCode,
			"total_time":  latencyTime,
			"ip":          clientIP,
			"method":      reqMethod,
			"uri":         reqUri,
		}).Info("access")
	}
}

func setupRequestId(c *gin.Context) string {
	appLog := AppLog
	appLog.Info("[Trace] Generate UUID for Request.")
	requestId := c.Request.Header.Get("X-Request-Id")
	if requestId == "" {
		requestId = uuid.New().String()
		c.Request.Header.Set("X-Request-Id", requestId)
	}

	c.Header("X-Request-Id", requestId)
	return requestId
}
