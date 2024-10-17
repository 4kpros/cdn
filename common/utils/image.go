package utils

import (
	"github.com/h2non/bimg"
)

//Resize the image with preserving the aspect ratio
func ResizeImage(inputBuffer []byte, size int, quality int) (outputBuffer []byte, outputSize bimg.ImageSize, err error) {
	outputBuffer, err = bimg.NewImage(inputBuffer).Process(bimg.Options{
		Width:    size,
		Height:   size,
		Quality:  quality,
		Lossless: false,
		Type:     bimg.WEBP,
	})
	outputSize, err = bimg.NewImage(outputBuffer).Size()
	return
}

// Resize the image with preserving the aspect ratio
func ResizeImageDefault(inputBuffer []byte) (outputBuffer []byte, outputSize bimg.ImageSize, err error) {
	outputBuffer, err = bimg.NewImage(inputBuffer).Process(bimg.Options{
		Width:    500,
		Height:   500,
		Quality:  bimg.Quality,
		Lossless: false,
		Type:     bimg.WEBP,
	})
	outputSize, err = bimg.NewImage(outputBuffer).Size()
	outputBuffer = inputBuffer
	return
}
