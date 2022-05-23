package imagedecoder

import (
	"errors"
	"image"

	log "github.com/sirupsen/logrus"
)

type IDecoder interface {
	Decode(content []byte) (image.Image, error)
}

type factory struct{}

func NewFactory() *factory {
	return new(factory)
}

func (fac factory) Build(contentType string) (IDecoder, error) {
	switch contentType {
	case "image/png":
		return NewPngDecoder(), nil
	}
	log.Warnf("Someone tried to send an unsupported file with the type %s", contentType)
	return nil, errors.New("File type not supported")
}
