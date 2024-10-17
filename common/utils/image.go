package utils

import (
	"cdn/common/types"
	"github.com/h2non/bimg"
)

const (
	lowSize    = 128
	mediumSize = 250
	largeSize  = 500
	extraLarge = 1000
)

func ResizeImageWithMultipleResolution(buffer []byte) *types.Image {
	var lowImg, mediumImg, largeImg, extraLargeImg []byte
	go resizeImage(buffer, lowSize, lowImg)
	go resizeImage(buffer, mediumSize, mediumImg)
	go resizeImage(buffer, largeSize, largeImg)
	go resizeImage(buffer, extraLarge, extraLargeImg)
	// Await
	return &types.Image{
		Low:        lowImg,
		Medium:     mediumImg,
		Large:      largeImg,
		ExtraLarge: extraLargeImg,
	}
}

// Resize the image with preserving the aspect ratio
func resizeImage(buffer []byte, size int, result []byte) {
	result, err := bimg.NewImage(buffer).Resize(size, size)
	if err == nil {
		_, _ = bimg.NewImage(result).Size()
	}
}
