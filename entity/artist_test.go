package entity

import (
	"image"
	_ "image/png"
	"os"
	"testing"
)

func TestDaliArtPng(t *testing.T) {
	// Setting up the test
	filePath := "../resources/test/images/planet_logo.png"
	f, errFileOpen := os.Open(filePath)
	defer f.Close()
	if errFileOpen != nil {
		t.Fail()
	}

	expFilePath := "../resources/test/images/planet_logo.expect.txt"
	expF, errExpFileOpen := os.ReadFile(expFilePath)
	if errExpFileOpen != nil {
		t.Fail()
	}

	image, _, errImageDecode := image.Decode(f)
	if errImageDecode != nil {
		t.Fail()
	}
	// Running the entity
	michelangelo := NewArtist()
	michelangelo.paint(image)
	result := michelangelo.show()
	println(result)
	if result != string(expF) {
		t.Fail()
	}
}
