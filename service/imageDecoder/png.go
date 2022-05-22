package imagedecoder

import (
	"image"
	_ "image/png"
	"mime/multipart"
)

type pngDecoder struct {
	IDecoder
}

func NewPngDecoder() pngDecoder {
	return *new(pngDecoder)
}

func (svc pngDecoder) decode(file multipart.File) (image.Image, error) {
	image, _, errImageDecode := image.Decode(file)
	if errImageDecode != nil {
		return nil, errImageDecode
	}
	return image, nil
}
