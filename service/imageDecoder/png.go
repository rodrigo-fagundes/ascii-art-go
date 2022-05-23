package imagedecoder

import (
	"bytes"
	"image"
	_ "image/png"

	log "github.com/sirupsen/logrus"
)

type pngDecoder struct {
	IDecoder
}

func NewPngDecoder() pngDecoder {
	return *new(pngDecoder)
}

func (svc pngDecoder) Decode(content []byte) (image.Image, error) {
	img, _, errImageDecode := image.Decode(bytes.NewReader(content))
	if errImageDecode != nil {
		log.Error(errImageDecode, " - Failed decoding image!")
		return nil, errImageDecode
	}
	return img, nil
}
