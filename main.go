package main

import (
	"bytes"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func init() {
	// Setting up log level for the entire service
	// TODO - In a real service, I'd set the formatter so a log collector agent could consolidate in a cluster-wide aggrregator.
	logLevel := log.ErrorLevel
	if envLogLevel := os.Getenv("LOG_LEVEL"); envLogLevel != "" {
		var errParseLevel error
		logLevel, errParseLevel = log.ParseLevel(envLogLevel)
		if errParseLevel != nil {
			panic("Invalid log level! Fix the configuration.")
		}
	}
	log.SetLevel(logLevel)
}

func main() {
	r := gin.Default()
	// TODO - In a real service, also set up an agent to send telemetry data for profiling and performance metrics
	r.GET("/artify", func(c *gin.Context) {
		file, header, err := c.Request.FormFile("file")
		defer file.Close()
		if err != nil {
			msg := "Failed getting file from request!"
			log.Error(err, msg)
			c.JSON(
				http.StatusBadRequest,
				gin.H{"message": msg},
			)
		}

		buf := bytes.NewBuffer(nil)
		if _, err := io.Copy(buf, file); err != nil {
			msg := "Failed reading image!"
			log.Error(err, msg)
			c.JSON(
				http.StatusBadRequest,
				gin.H{"message": msg},
			)
		}

		decoder := imagedecoder.NewFactory().build(http.DetectContentType(buf.Bytes()))
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
