package utils

import (
	"fmt"
	"github.com/h2non/bimg"
)

const defaultImageSize = 1000

func ResizeImage(inputBuffer []byte, width int, height int, quality int) (outputBuffer []byte, outputSize bimg.ImageSize, err error) {
	// Check for supported types
	inputType := bimg.NewImage(inputBuffer).Type()
	if !bimg.IsTypeNameSupported(inputType) {
		outputBuffer = nil
		err = fmt.Errorf("%s %s", "Unsupported image type:", inputType)
		return
	}
	// Check for resizable types
	if inputType == "pdf" || inputType == "svg" || inputType == "magick" || inputType == "heif" || inputType == "avif" {
		outputSize, err = bimg.NewImage(inputBuffer).Size()
		outputBuffer = inputBuffer
		return
	}

	// Apply default type for maximum performance
	outputType := bimg.JPEG
	if inputType == "gif" {
		outputType = bimg.GIF
	}
	if inputType == "png" || inputType == "webp" || inputType == "tiff" {
		outputType = bimg.WEBP
	}

	// Fix width and height
	fixedWidth := width
	fixedHeight := height
	if fixedWidth <= 0 {
		fixedWidth = 1
	}
	if fixedHeight <= 0 {
		fixedHeight = 1
	}
	inputSize, err := bimg.NewImage(inputBuffer).Size()
	if fixedWidth > inputSize.Width {
		fixedWidth = inputSize.Width
	}
	if fixedHeight > inputSize.Height {
		fixedHeight = inputSize.Height
	}

	// Fix quality
	fixedQuality := quality
	if fixedQuality <= 0 {
		fixedQuality = bimg.Quality
	}

	// Apply image processing
	outputBuffer, err = bimg.NewImage(inputBuffer).Resize(fixedWidth, fixedHeight)
	outputBuffer, err = bimg.NewImage(outputBuffer).Process(bimg.Options{
		Quality:  quality,
		Lossless: false,
		Gravity:  bimg.GravityCentre,
		Type:     outputType,
	})
	outputSize, err = bimg.NewImage(outputBuffer).Size()
	return
}
