package nrgba

import (
	"fmt"
	"image"
)

// Encode encodes specified data into a NRGBA image of given size
func Encode(data []byte, width, height int) (*image.NRGBA, error) {
	// 4 bytes can be stored in every pixel (NRGBA)
	maxLength := width * height * 4
	if len(data) > maxLength {
		return nil, fmt.Errorf("data does not fit the image size")
	}
	img := image.NewNRGBA(image.Rectangle{image.Point{0, 0}, image.Point{width, height}})
	img.Pix = data
	for i := len(img.Pix); i < maxLength; i++ {
		img.Pix = append(img.Pix, 0)
	}
	return img, nil
}

// Decode decodes a NRGBA image to an original data
func Decode(img *image.NRGBA, trimZeros bool) ([]byte, error) {
	if trimZeros {
		i := len(img.Pix) - 1
		for ; img.Pix[i] == 0 && i >= 0; i-- {
		}
		return img.Pix[:i+1], nil
	}
	return img.Pix, nil
}
