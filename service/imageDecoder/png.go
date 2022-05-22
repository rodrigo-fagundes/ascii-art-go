package imagedecoder

import (
	"bytes"
	"image"
	_ "image/png"
	"mime/multipart"

	log "github.com/sirupsen/logrus"
	"github.com/vincent-petithory/dataurl"
)

type pngDecoder struct {
	IDecoder
}

func NewPngDecoder() pngDecoder {
	return *new(pngDecoder)
}

func (svc pngDecoder) Decode(file multipart.File) (image.Image, error) {
	// FIXME - Errors converting multipart into image.
	imgDecoded, errDecFromURL := dataurl.Decode(file)
	if errDecFromURL != nil {
		log.Error(errDecFromURL, " - Failed decoding image!")
		return nil, errDecFromURL
	}
	img, _, errImageDecode := image.Decode(bytes.NewReader(imgDecoded.Data))
	if errImageDecode != nil {
		log.Error(errImageDecode, " - Failed decoding image!")
		return nil, errImageDecode
	}
	return img, nil
}
