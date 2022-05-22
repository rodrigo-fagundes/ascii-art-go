package imagedecoder

import (
	"errors"
	"image"
	"mime/multipart"
)

type IDecoder interface {
	decode(file multipart.File) (image.Image, error)
}

type factory struct{}

func (fac factory) NewFactory() *factory {
	return new(factory)
}

func (fac factory) build(contentType string) (IDecoder, error) {
	switch contentType {
	case "png":
		return NewPngDecoder(), nil
	}
	return nil, errors.New("File type not supported")
}
