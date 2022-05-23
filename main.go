package main

import (
	"bytes"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"github.com/rodrigo-fagundes/ascii-art-go/entity"
	imagedecoder "github.com/rodrigo-fagundes/ascii-art-go/service/imageDecoder"
)

var decoderFactory = imagedecoder.NewFactory()
var michelangelo = entity.NewArtist()

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
	r.GET("/probe/liveness", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "I'm alive!"})
	})
	r.POST("/artify", func(c *gin.Context) {
		file, _, err := c.Request.FormFile("file")
		if err != nil {
			msg := "Please, provide an image."
			log.Error(err, msg)
			c.JSON(
				http.StatusBadRequest,
				gin.H{"message": msg},
			)
		}
		defer file.Close()

		buf := bytes.NewBuffer(nil)
		if _, err := io.Copy(buf, file); err != nil {
			msg := "Failed reading image!"
			log.Error(err, msg)
			c.JSON(
				http.StatusBadRequest,
				gin.H{"message": msg},
			)
		}

		decoder, errDecFac := decoderFactory.Build(http.DetectContentType(buf.Bytes()))
		if errDecFac != nil {
			c.JSON(
				http.StatusBadRequest,
				gin.H{"message": errDecFac.Error()},
			)
		}

		img, errDec := decoder.Decode(buf.Bytes())
		if errDec != nil {
			msg := "Failed decoding image!"
			log.Error(errDec, msg)
			c.JSON(
				http.StatusBadRequest,
				gin.H{"message": msg},
			)
		}

		result := michelangelo.Paint(img)
		c.String(http.StatusOK, result)
	})
	r.Run()
}
