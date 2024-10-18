package utils

import (
	"fmt"
	"github.com/h2non/bimg"
)

const defaultImageSize = 1000

func ResizeImage(inputBuffer []byte, width int, height int, quality int, compression int, crop bool) (outputBuffer []byte, outputSize bimg.ImageSize, err error) {
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
		fixedWidth = defaultImageSize
	}
	if fixedHeight <= 0 {
		fixedHeight = defaultImageSize
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

	// Fix compression
	fixedCompression := compression
	if fixedCompression <= 0 {
		fixedCompression = 0
	}

	// Apply image processing
	outputBuffer, err = bimg.NewImage(inputBuffer).Resize(fixedWidth, fixedHeight)
	outputBuffer, err = bimg.NewImage(outputBuffer).Process(bimg.Options{
		Quality:     quality,
		Compression: compression,
		Crop:        crop,
		Lossless:    false,
		Gravity:     bimg.GravityCentre,
		Type:        outputType,
	})
	outputSize, err = bimg.NewImage(outputBuffer).Size()
	return
}

func ResizeImageDefault(inputBuffer []byte) (outputBuffer []byte, outputSize bimg.ImageSize, err error) {
	outputBuffer, outputSize, err = ResizeImage(inputBuffer, defaultImageSize, defaultImageSize, 100, 0, false)
	return
}
